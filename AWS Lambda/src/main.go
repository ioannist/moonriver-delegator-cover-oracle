package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	inactivitycover "stakebaby.com/moonriver-delegator-cover-oracle/model/inactivitycover"
	oraclemaster "stakebaby.com/moonriver-delegator-cover-oracle/model/oraclemaster"
	"stakebaby.com/moonriver-delegator-cover-oracle/services/exporter"
	"stakebaby.com/moonriver-delegator-cover-oracle/util"
)

var (
	ctx         = context.Background()
	url         = os.Getenv("ETH_URL")
	client, err = ethclient.DialContext(ctx, url)
)
var ORACLE_MEMBER_PKEY = os.Getenv("ORACLE_MEMBER_PKEY")
var ORACLE_MEMBER = common.HexToAddress(os.Getenv("ORACLE_MEMBER"))
var ORACLE_MASTER = common.HexToAddress(os.Getenv("ORACLE_MASTER"))
var INACTIVITY_COVER = common.HexToAddress(os.Getenv("INACTIVITY_COVER"))
var MAX_DELEGATORS_IN_REPORT int
var rpcUrl = os.Getenv("RPC_URL")
var oracleAddress common.Address
var privateKey *ecdsa.PrivateKey
var GAS_LIMIT int
var roundCompleted = false
var members = []common.Address{}
var zeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")

type Member struct {
	IsMember                 bool     `json:"isMember"`
	Active                   bool     `json:"active"`
	Deposit                  *big.Int `json:"deposit"`
	MaxCoveredDelegation     *big.Int `json:"maxCoveredDelegation"`
	DelegatorsReportedInEra  *big.Int `json:"delegatorsReportedInEra"`
	LastPushedEra            *big.Int `json:"lastPushedEra"`
	NoZeroPtsCoverAfterEra   *big.Int `json:"noZeroPtsCoverAfterEra"`
	NoActiveSetCoverAfterEra *big.Int `json:"noActiveSetCoverAfterEra"`
	WentInactiveEra          *big.Int `json:"wentInactiveEra"`
}

const (
	sleepTime = 333
)

func init() {
	GAS_LIMIT_i, _ := strconv.Atoi(os.Getenv("GAS_LIMIT"))
	GAS_LIMIT = util.MaxInt(GAS_LIMIT_i, 10000000)
	MAX_DELEGATORS_IN_REPORT, _ = strconv.Atoi(os.Getenv("MAX_DELEGATORS_IN_REPORT"))
	fmt.Printf("Loaded members: %v\n", members)
}

func handler(_ context.Context, req events.CloudWatchEvent) error {
	fmt.Printf("\n%+v\n", req)

	fmt.Println("Set up oracle key and address")
	err = setupKeys()
	if err != nil {
		return err
	}
	fmt.Printf("Submitting from %v\n", oracleAddress)

	fmt.Println("Query current era data from contract (1st)")
	eraFromContract, _, _, _, err := isReportedLastEra(oracleAddress)
	if err != nil {
		return err
	}
	fmt.Println("Query chain data (substrate call) for previous round")
	eraFromChain, err := queryChainEra()
	if err != nil {
		return err
	}
	if eraFromChain.Current <= eraFromContract {
		return fmt.Errorf("Invalid era on contract or chain: %v vs %v", eraFromChain, eraFromContract)
	}
	if eraFromChain.Current-1 == eraFromContract && roundCompleted {
		fmt.Println("Nothing to do")
		return nil
	}
	roundCompleted = false

	fmt.Println("Load active members")
	memberCount, err := getMemberCount()
	if err != nil {
		return err
	}
	members, err = getActiveMemberAddresses(memberCount, eraFromContract)
	if err != nil {
		return err
	}
	fmt.Printf("Loaded members: %v\n", members)

	fmt.Println("Check that oracle has enough balance to execute txs")
	err = validateBalance(oracleAddress)
	if err != nil {
		return err
	}

	// Prepare oracle data structure
	od := oraclemaster.TypesOracleData{}

	// To get the data for round-1 (completed round), we must query on any block of the current round (we choose the first block)
	// This will ensure that we are not quering a incomplete round, and we are not quering from a very remote block (no data)
	err = queryOracleData(&od, int(eraFromChain.First), int(eraFromChain.Current-1), members)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", od)

	collatorsWithZeroPoints := 0
	for _, col := range od.Collators {
		if col.Points.Cmp(big.NewInt(0)) == 0 {
			collatorsWithZeroPoints++
		}
	}
	fmt.Printf("Detected %v collators with zero points in this round\n", collatorsWithZeroPoints)
	if collatorsWithZeroPoints > len(od.Collators)*2/3 {
		return fmt.Errorf("Too many collators have 0 points in this round; network/oracle failure? will not report.\n")
	}

	fmt.Println("Query current era data from contract")
	eraFromContract, eraNonceFromContract, firstEraNonceFromContract, reportedByMe, err := isReportedLastEra(oracleAddress)
	fmt.Printf("eraFromContract %v\neraNonceFromContract %v\nfirstEraNonceFromContract %v\nreportedByMe %v\n", eraFromContract, eraNonceFromContract, firstEraNonceFromContract, reportedByMe)
	if err != nil {
		return err
	}

	if eraFromChain.Current-1 == eraFromContract && reportedByMe {
		fmt.Println("Already reported the last completed era")
		return nil
	}

	fmt.Println("Break era into era nonces")
	// the order of reporting collators is decided
	filteredCollatorData := []oraclemaster.TypesCollatorData{}
	indexToReport := eraNonceFromContract - firstEraNonceFromContract
	if indexToReport < 0 {
		return fmt.Errorf("Invalid eraNonceFromContract %v or firstEraNonce %v\n", eraNonceFromContract, firstEraNonceFromContract)
	}

	od.Finalize = true // default is to send all data for the collator/s
	if eraFromChain.Current-1 > eraFromContract || indexToReport == 0 {
		// first, report all non-zero-points collators
		//if eraFromChain.Current-1 > eraFromContract {
		// the first caller for the new round will run clearReporting to start fresh, which will also increment era nonce
		// eraNonceFromContract++
		//}
		fmt.Println("Reporting non-zero-point collators")
		for _, col := range od.Collators {
			if col.Points.Cmp(big.NewInt(0)) == 1 {
				filteredCollatorData = append(filteredCollatorData, col)
			}
		}
	} else {
		// then, report any zero-point or non-active collators, one at a time
		collatorWithZeroPointsIndex := uint64(1)
		for _, col := range od.Collators {
			member, err := getMember(col.CollatorAccount)
			fmt.Printf("zero-point member:\n%+v\n", member)
			if col.Points.Cmp(big.NewInt(0)) == 0 || !col.Active {
				// if the lastPushedEraNonce is the previous nonce, we processed
				if collatorWithZeroPointsIndex == indexToReport {
					fmt.Printf("Reporting zero-point collator %v with index %v\n", col.CollatorAccount, collatorWithZeroPointsIndex)
					// split the report of each collator, to multiple reports, each one with a max of 150 delegators
					if err != nil {
						return err
					}
					fmt.Printf("Sending delegators from index %v\n", member.DelegatorsReportedInEra.Uint64())
					col.TopActiveDelegations = col.TopActiveDelegations[member.DelegatorsReportedInEra.Uint64():] // exclude previously submitted delegators
					// if we cannot fit all remaining delegators in the report, then mark finalize as false
					if len(col.TopActiveDelegations) > MAX_DELEGATORS_IN_REPORT {
						fmt.Printf("Cannot finalize, will send %v delegators\n", MAX_DELEGATORS_IN_REPORT)
						col.TopActiveDelegations = col.TopActiveDelegations[0:MAX_DELEGATORS_IN_REPORT]
						od.Finalize = false
					}
					filteredCollatorData = append(filteredCollatorData, col)
					break
				}
				collatorWithZeroPointsIndex++
			}
		}
	}
	if len(filteredCollatorData) == 0 {
		fmt.Println("No collators to report")
		roundCompleted = true
		return nil
	}
	od.Collators = filteredCollatorData

	fmt.Println("Sign and send data to Oracle contract")
	err = pushData(od.Round.Uint64(), eraNonceFromContract, &od, 0)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "already known") {
			err = pushData(od.Round.Uint64(), eraNonceFromContract, &od, 1)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	fmt.Println("Finished")
	return nil
}

func queryOracleData(od *oraclemaster.TypesOracleData, blockNum int, round int, members []common.Address) error {
	var err error
	xp, err := exporter.NewExporter(rpcUrl, "")
	if err != nil {
		return err
	}
	err = xp.GetOracleData(od, blockNum, round, members)
	return err
}

func queryChainEra() (exporter.ActiveEra, error) {
	var err error
	xp, err := exporter.NewExporter(rpcUrl, "")
	if err != nil {
		return exporter.ActiveEra{}, err
	}
	return xp.GetActiveEra()
}

func pushData(eraID uint64, eraNonce uint64, data *oraclemaster.TypesOracleData, noncePlus int64) error {
	nonce, err := client.PendingNonceAt(context.Background(), oracleAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("gasPrice %v, nonce %v\n", gasPrice, nonce)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce) + noncePlus)
	auth.Value = big.NewInt(0)        // in wei
	auth.GasLimit = uint64(GAS_LIMIT) // in units
	auth.GasPrice = gasPrice          //big.NewInt(1)

	oracleMaster, err := oraclemaster.NewMoonriverDelegatorCoverOracle(ORACLE_MASTER, client)
	if err != nil {
		return err
	}
	_, err = oracleMaster.ReportPara(
		auth,
		oracleAddress,
		big.NewInt(int64(eraID)),
		big.NewInt(int64(eraNonce)),
		*data,
	)
	// fmt.Printf("Submitted tx hash: %v\n", tx.Hash())
	return err
}

func isReportedLastEra(oracleAddress common.Address) (uint64, uint64, uint64, bool, error) {
	oracleMaster, err := oraclemaster.NewMoonriverDelegatorCoverOracle(ORACLE_MASTER, client)
	if err != nil {
		return 0, 0, 0, false, err
	}
	/*eraId, err := oracleMaster.EraId(&bind.CallOpts{
		From: oracleAddress,
	})*/
	lastReported, err := oracleMaster.IsReportedLastEra(&bind.CallOpts{
		From: oracleAddress,
	}, oracleAddress)
	fmt.Printf("Era from contract is %v\n", lastReported.LastEra)
	return lastReported.LastEra.Uint64(), lastReported.LastEraNonce.Uint64(), lastReported.LastFirstEraNonce.Uint64(), lastReported.Reported, nil
}

func getMember(collator common.Address) (Member, error) {
	inactivtyCover, err := inactivitycover.NewMoonriverDelegatorCoverOracle(INACTIVITY_COVER, client)
	if err != nil {
		return Member{}, err
	}
	isMember, active, deposit, maxCoveredDelegation, delegatorsReportedInEra, lastPushedEra, noZeroPtsCoverAfterEraZero, noActiveSetCoverAfterEra, wentInactiveEra, err := inactivtyCover.GetMember(&bind.CallOpts{
		From: oracleAddress,
	}, collator)
	member := Member{
		IsMember:                 isMember,
		Active:                   active,
		Deposit:                  deposit,
		MaxCoveredDelegation:     maxCoveredDelegation,
		DelegatorsReportedInEra:  delegatorsReportedInEra,
		LastPushedEra:            lastPushedEra,
		NoZeroPtsCoverAfterEra:   noZeroPtsCoverAfterEraZero,
		NoActiveSetCoverAfterEra: noActiveSetCoverAfterEra,
		WentInactiveEra:          wentInactiveEra,
	}
	return member, nil
}

func getMemberAddress(index int) (common.Address, error) {
	inactivtyCover, err := inactivitycover.NewMoonriverDelegatorCoverOracle(INACTIVITY_COVER, client)
	memberAddress, err := inactivtyCover.MemberAddresses(&bind.CallOpts{
		From: oracleAddress,
	}, big.NewInt(int64(index)))
	return memberAddress, err
}

func getActiveMemberAddresses(count uint64, eraId uint64) ([]common.Address, error) {
	members := []common.Address{}
	inactivtyCover, err := inactivitycover.NewMoonriverDelegatorCoverOracle(INACTIVITY_COVER, client)
	if err != nil {
		return nil, err
	}
	for index := uint64(0); index < count; index++ {
		memberAddress, err := inactivtyCover.MemberAddresses(&bind.CallOpts{
			From: oracleAddress,
		}, big.NewInt(int64(index)))
		if err != nil {
			return members, err
		}
		member, err := getMember(memberAddress)
		if err != nil {
			return members, err
		}
		// if the member is active, or it just went inactive (in the currently reported era)
		if member.Active || member.WentInactiveEra.Uint64() == eraId {
			members = append(members, memberAddress)
			fmt.Printf("member loaded:\n%+v\n", member)
		}
	}
	return members, nil
}

func getMemberCount() (uint64, error) {
	inactivtyCover, err := inactivitycover.NewMoonriverDelegatorCoverOracle(INACTIVITY_COVER, client)
	if err != nil {
		return 0, err
	}
	count, err := inactivtyCover.GetMembersCount(&bind.CallOpts{
		From: oracleAddress,
	})
	if err != nil {
		return 0, err
	}
	return count.Uint64(), nil
}

func getFirstEraNonce() (uint64, error) {
	oracleMaster, err := oraclemaster.NewMoonriverDelegatorCoverOracle(ORACLE_MASTER, client)
	if err != nil {
		return 0, err
	}
	firstEraNonce, err := oracleMaster.FirstEraNonce(&bind.CallOpts{
		From: oracleAddress,
	})
	fmt.Printf("firstEraNonce from contract is %v\n", firstEraNonce.Uint64())
	return firstEraNonce.Uint64(), nil
}

func validateBalance(oracleAddress common.Address) error {
	balance, err := client.BalanceAt(ctx, oracleAddress, nil)
	if err != nil {
		return err
	}
	minimumBalance, _ := new(big.Int).SetString("500000000000000000", 10) // in Wei
	if balance.Cmp(minimumBalance) == -1 {
		return fmt.Errorf("Manager balance below threshold")
	}
	fmt.Printf("Balance is %v\n", balance.String())
	return nil
}

func setupKeys() error {
	privateKey, err = crypto.HexToECDSA(ORACLE_MEMBER_PKEY)
	if err != nil {
		return err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return err
	}
	oracleAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
