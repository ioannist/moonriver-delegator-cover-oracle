package exporter

import (
	_ "embed"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ws "github.com/gorilla/websocket"
	"github.com/itering/subscan/util"
	"github.com/itering/substrate-api-rpc"
	"github.com/prometheus/client_golang/prometheus"
	oraclemaster "stakebaby.com/moonriver-delegator-cover-oracle/model/oraclemaster"
)

type Exporter struct {
	endpoint           string
	registry           *prometheus.Registry
	metricDescriptions map[string]*prometheus.Desc
}

type ActiveEra struct {
	Current uint64 `json:"current"`
	First   uint64 `json:"first"`
	Length  uint64 `json:"length"`
}

type Delegation struct {
	Owner  string `json:"owner"`
	Amount string `json:"amount"`
}

type Delegations struct {
	Total       string       `json:"total"`
	Delegations []Delegation `json:"delegations"`
}

type CollatorInfo struct {
	Address                       string
	Bond                          string         `json:"bond"`
	DelegationCount               uint64         `json:"delegation_count"`
	TotalCounted                  string         `json:"total_counted"`
	LowestTopDelegationAmount     string         `json:"lowest_top_delegation_amount"`
	HighestBottomDelegationAmount string         `json:"highest_bottom_delegation_amount"`
	LowestBottomDelegationAmount  string         `json:"lowest_bottom_delegation_amount"`
	TopCapacity                   string         `json:"top_capacity"`
	BottomCapacity                string         `json:"bottom_capacity"`
	Status                        map[string]any `json:"status"`
}

type PoolCandidate struct {
	Owner  string `json:"owner"`
	Amount string `json:"amount"`
}

const (
	sleepTime      = 333
	metadataSpecID = 1
)

//go:embed moonbeam-types.json
var customTypes []byte

func NewExporter(endpoint string, customTypesFilePath string) (*Exporter, error) {
	conn, _, err := ws.DefaultDialer.Dial(endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	if err := prepareMetadata(conn); err != nil {
		return nil, err
	}
	substrate.RegCustomTypes(customTypes)
	e := &Exporter{
		endpoint: endpoint,
		registry: prometheus.NewRegistry(),
	}
	return e, nil
}

func (e *Exporter) GetActiveEra() (ActiveEra, error) {
	conn, _, err := ws.DefaultDialer.Dial(e.endpoint, nil)
	if err != nil {
		return ActiveEra{}, err
	}
	defer conn.Close()

	var activeEra ActiveEra
	if storage, err := readStorage(conn, "parachainStaking", "round", ""); err != nil {
		return ActiveEra{}, err
	} else if err = json.Unmarshal([]byte(storage), &activeEra); err != nil {
		return ActiveEra{}, fmt.Errorf("storage staking.activeEra invalid: %w", err)
	} else {
		fmt.Printf("active_era_index %v\n", float64(activeEra.Current))
	}
	return activeEra, nil
}

func (e *Exporter) GetOracleData(od *oraclemaster.TypesOracleData, blockNumber int, round int, members []common.Address) error {
	conn, _, err := ws.DefaultDialer.Dial(e.endpoint, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	blockHash, err := getBlockHash(conn, blockNumber)
	if err != nil {
		return err
	}
	copy(od.BlockHash[:], blockHash)
	od.Collators = []oraclemaster.TypesCollatorData{}
	od.BlockNumber = big.NewInt(int64(blockNumber))
	od.Round = big.NewInt(int64(round))

	fmt.Printf("Collecting data for blockHash: %v\n", blockHash)
	ps := "parachainStaking"

	fmt.Println("Get total staked")
	if data, err := readStorage(conn, ps, "total", blockHash); err != nil {
		return err
	} else {
		n := stringToBigInt(data.ToString())
		od.TotalStaked = n
	}
	time.Sleep(sleepTime * time.Millisecond)

	fmt.Println("Get total selected candidates")
	if data, err := readStorage(conn, ps, "totalSelected", blockHash); err != nil {
		return err
	} else {
		n := stringToBigInt(data.ToString())
		od.TotalSelected = n
	}
	time.Sleep(sleepTime * time.Millisecond)

	fmt.Println("Get count of orbiters")
	if data, err := readStorage(conn, "moonbeamOrbiters", "counterForCollatorsPool", blockHash); err != nil {
		return err
	} else {
		n := stringToBigInt(data.ToString())
		od.OrbitersCount = n
	}
	time.Sleep(sleepTime * time.Millisecond)

	fmt.Println("Get total awarded points for this round")
	encodedEra := make([]byte, 4)
	binary.LittleEndian.PutUint32(encodedEra, uint32(od.Round.Uint64()))
	if data, err := readStorage(conn, ps, "points", blockHash, util.BytesToHex(encodedEra)); err != nil {
		return err
	} else {
		n := stringToBigInt(data.ToString())
		od.Awarded = n
	}
	fmt.Printf("points: %v\n", od.Awarded)
	time.Sleep(sleepTime * time.Millisecond)

	/*var collators []PoolCandidate
	if storage, err := readStorage(conn, ps, "candidatePool", blockHash); err != nil {
		return err
	} else if err = json.Unmarshal([]byte(storage), &collators); err != nil {
		return fmt.Errorf("storage session.validators invalid: %w", err)
	}*/

	time.Sleep(sleepTime * time.Millisecond)
	collatorInfos := []CollatorInfo{}

	for _, memberAddress := range members {
		var collatorInfo = CollatorInfo{
			Address: memberAddress.Hex(),
		}
		if storage, err := readStorage(conn, ps, "candidateInfo", blockHash, memberAddress.Hex()); err != nil {
			fmt.Printf("Could not get data for collator: %v\n", memberAddress.Hex())
		} else if err = json.Unmarshal([]byte(storage), &collatorInfo); err != nil {
			fmt.Printf("storage parachainStaking.collatorInfo invalid: %v\n", err)
		} else {
			// fmt.Printf("%+v\n", collatorInfo)
			collatorInfos = append(collatorInfos, collatorInfo)
		}
	}

	for _, collatorInfo := range collatorInfos {
		var collatorData oraclemaster.TypesCollatorData
		if data, err := readStorage(conn, ps, "awardedPts", blockHash, util.BytesToHex(encodedEra), collatorInfo.Address); err != nil {
			return err
		} else {

			time.Sleep(sleepTime * time.Millisecond)
			collatorData.Points = stringToBigInt(data.ToString())
			_, isActive := collatorInfo.Status["Active"]
			collatorData.Active = isActive
			collatorData.Bond = stringToBigInt(collatorInfo.Bond)
			collatorData.CollatorAccount = common.HexToAddress(collatorInfo.Address)
			collatorData.TopActiveDelegations = []oraclemaster.TypesDelegationsData{}
			collatorData.DelegationsTotal = new(big.Int).Sub(stringToBigInt(collatorInfo.TotalCounted), stringToBigInt(collatorInfo.Bond))

			if collatorData.Points.Cmp(big.NewInt(0)) == 0 {
				var topDelegations Delegations
				if storage, err := readStorage(conn, ps, "topDelegations", blockHash, collatorInfo.Address); err != nil {
					return err
				} else if err = json.Unmarshal([]byte(storage), &topDelegations); err != nil {
					return fmt.Errorf("storage session.validators invalid: %w", err)
				} else {
					for _, delegation := range topDelegations.Delegations {
						deleg := oraclemaster.TypesDelegationsData{
							OwnerAccount: common.HexToAddress(delegation.Owner),
							Amount:       stringToBigInt(delegation.Amount),
						}
						collatorData.TopActiveDelegations = append(collatorData.TopActiveDelegations, deleg)
					}
				}
				time.Sleep(sleepTime * time.Millisecond)
			}

		}
		od.Collators = append(od.Collators, collatorData)
		// fmt.Printf("%+v\n", collatorData)
	}

	return nil
}

func stringToBigInt(s string) *big.Int {
	b := new(big.Int)
	b, ok := b.SetString(s, 10)
	if !ok {
		b = big.NewInt(0)
	}
	return b
}
