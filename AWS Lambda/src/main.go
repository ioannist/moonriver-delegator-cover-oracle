package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	inactivitycover "stakebaby.com/moonriver-delegator-cover-oracle/model/inactivitycover"
	oraclemaster "stakebaby.com/moonriver-delegator-cover-oracle/model/oraclemaster"
	"stakebaby.com/moonriver-delegator-cover-oracle/services/exporter"
	"stakebaby.com/moonriver-delegator-cover-oracle/util"
)

/**
TODO
* fetch constants from contract
* exclude revokes and decreases
* go back in time if no delegators
**/

var client *ethclient.Client
var err error
var url = os.Getenv("ETH_URL")
var ORACLE_MEMBER_PKEY = os.Getenv("ORACLE_MEMBER_PKEY")
var ORACLE_MEMBER = common.HexToAddress(os.Getenv("ORACLE_MEMBER"))
var ORACLE_MASTER = common.HexToAddress(os.Getenv("ORACLE_MASTER"))
var INACTIVITY_COVER = common.HexToAddress(os.Getenv("INACTIVITY_COVER"))
var MAX_DELEGATORS_IN_REPORT int
var rpcUrl = os.Getenv("RPC_URL")
var oracleAddress common.Address
var privateKey *ecdsa.PrivateKey
var GAS_LIMIT int
var members = []common.Address{}
var zeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
var prevTxHash common.Hash = common.BigToHash(common.Big0)

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
	WentActiveEra            *big.Int `json:"wentActiveEra"`
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

func handler(ctx context.Context, req events.CloudWatchEvent) error {

	fmt.Printf("\n%+v\n", req)
	client, err = ethclient.DialContext(ctx, url)

	fmt.Println("Set up oracle key and address")
	err = setupKeys()
	if err != nil {
		return err
	}
	fmt.Printf("Submitting from %v\n", oracleAddress)

	fmt.Println("Get oracle config from contract")
	maxDelegatorsInReport, _, err := getOracleConfig()
	if maxDelegatorsInReport != 0 {
		MAX_DELEGATORS_IN_REPORT = int(maxDelegatorsInReport)
	}

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

	fmt.Println("Load active members")
	memberCount, err := getMemberCount()
	if err != nil {
		return err
	}
	members, err = getActiveMemberAddresses(memberCount, eraFromChain.Current-1)
	if err != nil {
		return err
	}
	fmt.Printf("Loaded members: %v\n", members)

	fmt.Println("Check that oracle has enough balance to execute txs")
	err = validateBalance(ctx, oracleAddress)
	if err != nil {
		return err
	}

	// Prepare oracle data structure
	od := oraclemaster.TypesOracleData{}

	// To get the data for round-1 (completed round), we must query on any block of the current round (we choose the first block)
	// This will ensure that we are not quering a incomplete round, and we are not quering from a very remote block (no data)
	err = queryOracleData(&od, int(eraFromChain.First), int(eraFromChain.Current-1), members, client, oracleAddress)
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

	if collatorsWithZeroPoints == 0 && !isSpecialPeriod() {
		return fmt.Errorf("No collator has zero points in this round. Will not file report to save gas.\n")
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
			// fmt.Printf("member:\n%+v\n", member)
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
		return nil
	}
	od.Collators = filteredCollatorData

	fmt.Println("Sign and send data to Oracle contract")
	pushAttempts := 0
	for {
		_, err = pushData(ctx, od.Round.Uint64(), eraNonceFromContract, &od, int64(pushAttempts))
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "already known") {
				if pushAttempts > 5 {
					return fmt.Errorf("Max attempts reached")
				}
				fmt.Println("Already known! will try again...")
				pushAttempts++
				continue
			} else {
				return err
			}
		}
		break

	}
	fmt.Println("Finished")
	return nil
}

func queryOracleData(od *oraclemaster.TypesOracleData, blockNum int, round int, members []common.Address, client *ethclient.Client, oracleAddress common.Address) error {
	var err error
	xp, err := exporter.NewExporter(rpcUrl, "")
	if err != nil {
		return err
	}
	err = xp.GetOracleData(od, blockNum, round, members, client, oracleAddress)
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

func pushData(ctx context.Context, eraID uint64, eraNonce uint64, data *oraclemaster.TypesOracleData, noncePlus int64) (bool, error) {

	// the following check is useful in periods of extreme network load
	// we stop submitting txs until the previous tx is resolved
	if prevTxHash.Big().Cmp(common.Big0) != 0 {
		_, isPending, err := client.TransactionByHash(ctx, prevTxHash)

		receipt, err := client.TransactionReceipt(context.Background(), prevTxHash)
		if receipt.Status == types.ReceiptStatusFailed {
			fmt.Println("Previous transaction was reverted.")
		} else {
			fmt.Println("Previous transaction was successful.")
		}
		if err != nil {
			fmt.Printf("%+v\n", err)
		}

		if err != nil && !strings.Contains(strings.ToLower(err.Error()), "not found") {
			return false, err
		}
		if isPending {
			fmt.Printf("The previous tx is still pending; will wait\n")
			return true, nil
		}
	}

	nonce, err := client.PendingNonceAt(ctx, oracleAddress)
	if err != nil {
		return false, err
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return false, err
	}
	fmt.Printf("gasPrice %v, nonce %v\n", gasPrice, nonce)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce) + noncePlus)
	auth.Value = big.NewInt(0)                                 // in wei
	auth.GasLimit = uint64(GAS_LIMIT)                          // in units
	auth.GasPrice = big.NewInt(0).Mul(big.NewInt(4), gasPrice) //big.NewInt(1)
	auth.GasPrice.Div(auth.GasPrice, big.NewInt(2))

	fmt.Printf("data %+v", data)

	oracleMaster, err := oraclemaster.NewMoonriverDelegatorCoverOracle(ORACLE_MASTER, client)
	if err != nil {
		return false, err
	}
	tx, err := oracleMaster.ReportPara(
		auth,
		oracleAddress,
		big.NewInt(int64(eraID)),
		big.NewInt(int64(eraNonce)),
		*data,
	)
	if err != nil {
		return false, err
	}

	prevTxHash = tx.Hash()
	fmt.Printf("Submitted tx hash: %v, with nonce %v\n", tx.Hash(), tx.Nonce())
	return true, nil
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
	isMember, active, deposit, maxCoveredDelegation, delegatorsReportedInEra, lastPushedEra, noZeroPtsCoverAfterEraZero, noActiveSetCoverAfterEra, wentInactiveEra, wentActiveEra, err := inactivtyCover.GetMember(&bind.CallOpts{
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
		WentActiveEra:            wentActiveEra,
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
		fmt.Printf("%+v\n", member)
		// if the member is active (but it was not activated in the round being currently reported),
		// or if it just went inactive (in the currently reported era)
		if (member.Active && member.WentActiveEra.Uint64() < eraId) || member.WentInactiveEra.Uint64() == eraId {
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

func getOracleConfig() (uint64, uint64, error) {
	oracleMaster, err := oraclemaster.NewMoonriverDelegatorCoverOracle(ORACLE_MASTER, client)
	if err != nil {
		return 0, 0, err
	}
	maxDelegatorsInReport, runPeriodInSeconds, err := oracleMaster.GetOracleBinaryConfig(&bind.CallOpts{
		From: oracleAddress,
	})
	return maxDelegatorsInReport.Uint64(), runPeriodInSeconds.Uint64(), nil
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

func validateBalance(ctx context.Context, oracleAddress common.Address) error {
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

func isSpecialPeriod() bool {
	currentTime := time.Now()
	day := currentTime.YearDay() // gets the day of the year, from 1 to 366
	hour := currentTime.Hour()   // gets the hour of the day, from 0 to 23

	return day%3 == 0 && hour < 6
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
