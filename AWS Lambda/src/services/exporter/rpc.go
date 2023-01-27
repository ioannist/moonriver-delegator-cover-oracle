package exporter

import (
	"fmt"
	"math/rand"
	"strings"

	ws "github.com/gorilla/websocket"
	"github.com/itering/substrate-api-rpc/metadata"
	"github.com/itering/substrate-api-rpc/rpc"
	"github.com/itering/substrate-api-rpc/storage"
	"github.com/itering/substrate-api-rpc/storageKey"
	"github.com/sirupsen/logrus"
)

type EraRewardPoints struct {
	Total       uint32 `json:"total"`
	Individuals []struct {
		AccountId   string `json:"col1"`
		RewardPoint uint32 `json:"col2"`
	} `json:"individual"`
}

const (
	wsBlockHash = iota + 1
	wsBlock
	wsEvent
	wsSpec
)
const nilStorage = storage.StateStorage("null")

func prepareMetadata(conn *ws.Conn) error {
	response, err := sendWsRequest(conn, rpc.StateGetMetadata(0))
	if err != nil {
		return err
	}

	encoded := response.Result.(string)

	metadata.Latest(&metadata.RuntimeRaw{
		Spec: metadataSpecID,
		Raw:  strings.TrimPrefix(encoded, "0x"),
	})

	return nil
}

func sendWsRequest(conn *ws.Conn, data []byte) (*rpc.JsonRpcResult, error) {
	logrus.Tracef("RPC call: %s", data)
	v := &rpc.JsonRpcResult{}

	if err := conn.WriteMessage(ws.TextMessage, data); err != nil {
		return nil, fmt.Errorf("conn.WriteMessage: %w", err)
	}

	if err := conn.ReadJSON(v); err != nil {
		return nil, fmt.Errorf("conn.ReadJSON: %w", err)
	}

	logrus.Tracef("RPC raw result: %+v", v)

	if v.Error != nil {
		return nil, fmt.Errorf("RPC error: %s", v.Error.Message)
	}

	return v, nil
}

func getBlockHash(conn *ws.Conn, block int) (string, error) {

	rpcRequest := rpc.ChainGetBlockHash(wsBlockHash, block)
	rpcResponse, err := sendWsRequest(conn, rpcRequest)
	if err != nil {
		return "", err
	}

	if rpcResponse.Result == nil {
		return "", nil
	}

	encoded, ok := rpcResponse.Result.(string)
	if !ok {
		return "", fmt.Errorf("unable to get block hash\n")
	}

	if encoded == "" {
		return "", fmt.Errorf("no hash found\n")
	}

	return string(encoded), nil
}

func readStorage(conn *ws.Conn, storageSection string, storageMethod string, blockHash string, args ...string) (storage.StateStorage, error) {
	
	storageKey := storageKey.EncodeStorageKey(storageSection, storageMethod, args...)
	logrus.Tracef("Encoded storage key: %s", storageKey)

	rpcRequest := rpc.StateGetStorage(rand.Intn(10000), "0x"+storageKey.EncodeKey, blockHash)
	rpcResponse, err := sendWsRequest(conn, rpcRequest)
	if err != nil {
		return nilStorage, err
	}

	if rpcResponse.Result == nil {
		return nilStorage, nil
	}

	encoded, ok := rpcResponse.Result.(string)
	if !ok {
		return nilStorage, fmt.Errorf("unable to parse storage %s.%s, raw result: %+v", storageSection, storageMethod, rpcResponse.Result)
	}

	if encoded == "" {
		return nilStorage, nil
	}

	storage, err := storage.Decode(encoded, storageKey.ScaleType, nil)
	logrus.Tracef("Decoded storage: %v", storage)

	return storage, err
}
