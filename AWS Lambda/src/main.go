package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	oraclemaster "stakebaby.com/moonriver-delegator-cover-oracle/model/oraclemaster"
	"stakebaby.com/moonriver-delegator-cover-oracle/services/exporter"
	"stakebaby.com/moonriver-delegator-cover-oracle/util"
)

var (
	ctx         = context.Background()
	url         = os.Getenv("ETH_URL")
	client, err = ethclient.DialContext(ctx, url)
)
var ORACLE_MEMBER_ADDRESS = os.Getenv("ORACLE_MEMBER")
var ORACLE_MEMBER = common.HexToAddress(os.Getenv("ORACLE_MEMBER"))
var ORACLE_MASTER = common.HexToAddress(os.Getenv("ORACLE_MASTER"))
var rpcUrl = os.Getenv("RPC_URL")
var managerAddress common.Address
var privateKey *ecdsa.PrivateKey
var GAS_LIMIT int

const (
	sleepTime = 333
)

func init() {
	GAS_LIMIT_i, _ := strconv.Atoi(os.Getenv("GAS_LIMIT"))
	GAS_LIMIT = util.MaxInt(GAS_LIMIT_i, 10000000)
}

func handler(_ context.Context, req events.CloudWatchEvent) error {
	fmt.Printf("\n%+v\n", req)

	fmt.Println("Set up manager key and address")
	err = setupKeys()
	if err != nil {
		return err
	}

	fmt.Println("Check that manager has enough balance to execute txs")
	err = validateBalance(managerAddress)
	if err != nil {
		return err
	}

	fmt.Println("Query chain data (substrate call) for previous round")
	eraFromChain, err := queryChainEra()
	if err != nil {
		return err
	}

	// Prepare oracle data structure
	od := oraclemaster.TypesOracleData{}

	// To get the data for round-1 (completed round), we must query on any block of the current round (we choose the first block)
	// This will ensure that we are not quering a incomplete round, and we are not quering from a very remote block (no data)
	err = queryOracleData(&od, int(eraFromChain.First), int(eraFromChain.Current-1))
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
	eraFromContract, eraNonceFromContract, firstEraNonceFromContract, reportedByMe, err := isReportedLastEra(managerAddress)
	if err != nil {
		return err
	}

	if eraFromContract > eraFromChain.Current {
		return fmt.Errorf("Invalid era on contract or chain: %v vs %v", eraFromChain, eraFromContract)
	}
	if eraFromChain.Current-1 == eraFromContract && reportedByMe {
		fmt.Println("Already reported the last completed era")
		return nil
	}

	fmt.Println("Filter out report for this era nonce (one collator per report for gas tx safety)")
	// the order of reporting collators is decided
	indexToReport := eraNonceFromContract - firstEraNonceFromContract;
	if indexToReport < 1 {
		return fmt.Errorf("Invalid eraNonceFromContract %v or firstEraNonce %v", eraNonceFromContract, firstEraNonceFromContract)
	}
	filteredCollatorData := []oraclemaster.TypesCollatorData{}
	
	if indexToReport == 0 {
		// first, report all non-zero-points collators
		fmt.Println("Reporting non-zero-point collators")
		for _, col := range od.Collators {
			if col.Points.Cmp(big.NewInt(0)) == 1 {
				filteredCollatorData = append(filteredCollatorData, col)
			}
		}		
	} else {
		// then, report any zero-point collators, one at a time
		collatorWithZeroPointsIndex := uint64(1)
		for _, col := range od.Collators {
			if col.Points.Cmp(big.NewInt(0)) == 0 {
				if collatorWithZeroPointsIndex == indexToReport {
					fmt.Printf("Reporting zero-point collator %v with index %v\n", col.CollatorAccount, collatorWithZeroPointsIndex)
					filteredCollatorData = append(filteredCollatorData, col)
					break;
				}
				collatorWithZeroPointsIndex++;
			}
		}
	}
	if len(filteredCollatorData) == 0 {
		fmt.Println("No collators to report")
		return nil
	}
	od.Collators = filteredCollatorData

	fmt.Println("Sign and send data to Oracle contract")
	err = pushData(od.Round.Uint64(), eraNonceFromContract, &od)
	if err != nil {
		return err
	}

	fmt.Println("Finished")
	return nil
}

func queryOracleData(od *oraclemaster.TypesOracleData, blockNum int, round int) error {
	var err error
	xp, err := exporter.NewExporter(rpcUrl, "")
	if err != nil {
		return err
	}
	err = xp.GetOracleData(od, blockNum, round)
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

func pushData(eraID uint64, eraNonce uint64, data *oraclemaster.TypesOracleData) error {
	nonce, err := client.PendingNonceAt(context.Background(), managerAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)        // in wei
	auth.GasLimit = uint64(GAS_LIMIT) // in units
	auth.GasPrice = gasPrice          //big.NewInt(1)

	oracleMaster, err := oraclemaster.NewMoonriverDelegatorCoverOracle(ORACLE_MASTER, client)
	if err != nil {
		return err
	}
	_, err = oracleMaster.ReportPara(
		auth,
		ORACLE_MEMBER,
		big.NewInt(int64(eraID)),
		big.NewInt(int64(eraNonce)),
		*data,
	)
	// fmt.Printf("Submitted tx hash: %v\n", tx.Hash())
	return err
}

func isReportedLastEra(managerAddress common.Address) (uint64, uint64, uint64, bool, error) {
	oracleMaster, err := oraclemaster.NewMoonriverDelegatorCoverOracle(ORACLE_MASTER, client)
	if err != nil {
		return 0, 0, 0, false, err
	}
	/*eraId, err := oracleMaster.EraId(&bind.CallOpts{
		From: managerAddress,
	})*/
	lastReported, err := oracleMaster.IsReportedLastEra(&bind.CallOpts{
		From: managerAddress,
	}, managerAddress)
	fmt.Printf("Era from contract is %v\n", lastReported.LastEra)
	return lastReported.LastEra.Uint64(), lastReported.LastEraNonce.Uint64(), lastReported.LastFirstEraNonce.Uint64(), lastReported.Reported, nil
}

func getFirstEraNonce() (uint64, error) {
	oracleMaster, err := oraclemaster.NewMoonriverDelegatorCoverOracle(ORACLE_MASTER, client)
	if err != nil {
		return 0, err
	}
	firstEraNonce, err := oracleMaster.FirstEraNonce(&bind.CallOpts{
		From: managerAddress,
	})
	fmt.Printf("firstEraNonce from contract is %v\n", firstEraNonce.Uint64())
	return firstEraNonce.Uint64(), nil
}

func validateBalance(managerAddress common.Address) error {
	balance, err := client.BalanceAt(ctx, managerAddress, nil)
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
	privateKey, err = crypto.HexToECDSA(ORACLE_MEMBER_ADDRESS)
	if err != nil {
		return err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return err
	}
	managerAddress = crypto.PubkeyToAddress(*publicKeyECDSA)
	return nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
