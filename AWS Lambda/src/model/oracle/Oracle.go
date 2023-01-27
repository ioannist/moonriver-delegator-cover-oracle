// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package moonriver_delegator_cover_oracle

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TypesCollatorData is an auto generated low-level Go binding around an user-defined struct.
type TypesCollatorData struct {
	CollatorAccount      common.Address
	Points               *big.Int
	Active               bool
	Bond                 *big.Int
	DelegationsTotal     *big.Int
	TopActiveDelegations []TypesDelegationsData
}

// TypesDelegationsData is an auto generated low-level Go binding around an user-defined struct.
type TypesDelegationsData struct {
	OwnerAccount common.Address
	Amount       *big.Int
}

// TypesOracleData is an auto generated low-level Go binding around an user-defined struct.
type TypesOracleData struct {
	TotalStaked   *big.Int
	TotalSelected *big.Int
	OrbitersCount *big.Int
	Round         *big.Int
	BlockHash     [32]byte
	BlockNumber   *big.Int
	Awarded       *big.Int
	Collators     []TypesCollatorData
	Finalize      bool
}

// MoonriverDelegatorCoverOracleMetaData contains all meta data concerning the MoonriverDelegatorCoverOracle contract.
var MoonriverDelegatorCoverOracleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"eraId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"eraNonce\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"ReportSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"ReportingCleared\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ORACLE_MASTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PUSHABLES\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eraNonce\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"name\":\"erasToBlockHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracleMaster\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_pushable\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quorum\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_eraId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_eraNonce\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"totalSelected\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"orbitersCount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"blockNumber\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"awarded\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collatorAccount\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"points\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"bond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegationsTotal\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ownerAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.DelegationsData[]\",\"name\":\"topActiveDelegations\",\"type\":\"tuple[]\"}],\"internalType\":\"structTypes.CollatorData[]\",\"name\":\"collators\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"finalize\",\"type\":\"bool\"}],\"internalType\":\"structTypes.OracleData\",\"name\":\"_staking\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_oracleCollator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"veto\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"vetoDisabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"newEra\",\"type\":\"bool\"}],\"name\":\"reportPara\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_quorum\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"_eraId\",\"type\":\"uint128\"}],\"name\":\"softenQuorum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_pushable\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_toAdd\",\"type\":\"bool\"}],\"name\":\"addRemovePushable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"isReported\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearReporting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_eraId\",\"type\":\"uint128\"}],\"name\":\"getReportBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MoonriverDelegatorCoverOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use MoonriverDelegatorCoverOracleMetaData.ABI instead.
var MoonriverDelegatorCoverOracleABI = MoonriverDelegatorCoverOracleMetaData.ABI

// MoonriverDelegatorCoverOracle is an auto generated Go binding around an Ethereum contract.
type MoonriverDelegatorCoverOracle struct {
	MoonriverDelegatorCoverOracleCaller     // Read-only binding to the contract
	MoonriverDelegatorCoverOracleTransactor // Write-only binding to the contract
	MoonriverDelegatorCoverOracleFilterer   // Log filterer for contract events
}

// MoonriverDelegatorCoverOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type MoonriverDelegatorCoverOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MoonriverDelegatorCoverOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MoonriverDelegatorCoverOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MoonriverDelegatorCoverOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MoonriverDelegatorCoverOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MoonriverDelegatorCoverOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MoonriverDelegatorCoverOracleSession struct {
	Contract     *MoonriverDelegatorCoverOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MoonriverDelegatorCoverOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MoonriverDelegatorCoverOracleCallerSession struct {
	Contract *MoonriverDelegatorCoverOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// MoonriverDelegatorCoverOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MoonriverDelegatorCoverOracleTransactorSession struct {
	Contract     *MoonriverDelegatorCoverOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// MoonriverDelegatorCoverOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type MoonriverDelegatorCoverOracleRaw struct {
	Contract *MoonriverDelegatorCoverOracle // Generic contract binding to access the raw methods on
}

// MoonriverDelegatorCoverOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MoonriverDelegatorCoverOracleCallerRaw struct {
	Contract *MoonriverDelegatorCoverOracleCaller // Generic read-only contract binding to access the raw methods on
}

// MoonriverDelegatorCoverOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MoonriverDelegatorCoverOracleTransactorRaw struct {
	Contract *MoonriverDelegatorCoverOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMoonriverDelegatorCoverOracle creates a new instance of MoonriverDelegatorCoverOracle, bound to a specific deployed contract.
func NewMoonriverDelegatorCoverOracle(address common.Address, backend bind.ContractBackend) (*MoonriverDelegatorCoverOracle, error) {
	contract, err := bindMoonriverDelegatorCoverOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracle{MoonriverDelegatorCoverOracleCaller: MoonriverDelegatorCoverOracleCaller{contract: contract}, MoonriverDelegatorCoverOracleTransactor: MoonriverDelegatorCoverOracleTransactor{contract: contract}, MoonriverDelegatorCoverOracleFilterer: MoonriverDelegatorCoverOracleFilterer{contract: contract}}, nil
}

// NewMoonriverDelegatorCoverOracleCaller creates a new read-only instance of MoonriverDelegatorCoverOracle, bound to a specific deployed contract.
func NewMoonriverDelegatorCoverOracleCaller(address common.Address, caller bind.ContractCaller) (*MoonriverDelegatorCoverOracleCaller, error) {
	contract, err := bindMoonriverDelegatorCoverOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleCaller{contract: contract}, nil
}

// NewMoonriverDelegatorCoverOracleTransactor creates a new write-only instance of MoonriverDelegatorCoverOracle, bound to a specific deployed contract.
func NewMoonriverDelegatorCoverOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*MoonriverDelegatorCoverOracleTransactor, error) {
	contract, err := bindMoonriverDelegatorCoverOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleTransactor{contract: contract}, nil
}

// NewMoonriverDelegatorCoverOracleFilterer creates a new log filterer instance of MoonriverDelegatorCoverOracle, bound to a specific deployed contract.
func NewMoonriverDelegatorCoverOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*MoonriverDelegatorCoverOracleFilterer, error) {
	contract, err := bindMoonriverDelegatorCoverOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleFilterer{contract: contract}, nil
}

// bindMoonriverDelegatorCoverOracle binds a generic wrapper to an already deployed contract.
func bindMoonriverDelegatorCoverOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MoonriverDelegatorCoverOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MoonriverDelegatorCoverOracle.Contract.MoonriverDelegatorCoverOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MoonriverDelegatorCoverOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MoonriverDelegatorCoverOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MoonriverDelegatorCoverOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.contract.Transact(opts, method, params...)
}

// ORACLEMASTER is a free data retrieval call binding the contract method 0x3fd2c16a.
//
// Solidity: function ORACLE_MASTER() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) ORACLEMASTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "ORACLE_MASTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ORACLEMASTER is a free data retrieval call binding the contract method 0x3fd2c16a.
//
// Solidity: function ORACLE_MASTER() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ORACLEMASTER() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ORACLEMASTER(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// ORACLEMASTER is a free data retrieval call binding the contract method 0x3fd2c16a.
//
// Solidity: function ORACLE_MASTER() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) ORACLEMASTER() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ORACLEMASTER(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// PUSHABLES is a free data retrieval call binding the contract method 0x402e4749.
//
// Solidity: function PUSHABLES(uint256 ) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) PUSHABLES(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "PUSHABLES", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PUSHABLES is a free data retrieval call binding the contract method 0x402e4749.
//
// Solidity: function PUSHABLES(uint256 ) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) PUSHABLES(arg0 *big.Int) (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.PUSHABLES(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// PUSHABLES is a free data retrieval call binding the contract method 0x402e4749.
//
// Solidity: function PUSHABLES(uint256 ) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) PUSHABLES(arg0 *big.Int) (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.PUSHABLES(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// EraNonce is a free data retrieval call binding the contract method 0xf53d2ead.
//
// Solidity: function eraNonce() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) EraNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "eraNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EraNonce is a free data retrieval call binding the contract method 0xf53d2ead.
//
// Solidity: function eraNonce() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) EraNonce() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.EraNonce(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// EraNonce is a free data retrieval call binding the contract method 0xf53d2ead.
//
// Solidity: function eraNonce() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) EraNonce() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.EraNonce(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// ErasToBlockHashes is a free data retrieval call binding the contract method 0xadec8b73.
//
// Solidity: function erasToBlockHashes(uint128 ) view returns(bytes32)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) ErasToBlockHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "erasToBlockHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ErasToBlockHashes is a free data retrieval call binding the contract method 0xadec8b73.
//
// Solidity: function erasToBlockHashes(uint128 ) view returns(bytes32)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ErasToBlockHashes(arg0 *big.Int) ([32]byte, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ErasToBlockHashes(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// ErasToBlockHashes is a free data retrieval call binding the contract method 0xadec8b73.
//
// Solidity: function erasToBlockHashes(uint128 ) view returns(bytes32)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) ErasToBlockHashes(arg0 *big.Int) ([32]byte, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ErasToBlockHashes(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// GetReportBlockHash is a free data retrieval call binding the contract method 0x0e53fb48.
//
// Solidity: function getReportBlockHash(uint128 _eraId) view returns(bytes32)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) GetReportBlockHash(opts *bind.CallOpts, _eraId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "getReportBlockHash", _eraId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetReportBlockHash is a free data retrieval call binding the contract method 0x0e53fb48.
//
// Solidity: function getReportBlockHash(uint128 _eraId) view returns(bytes32)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) GetReportBlockHash(_eraId *big.Int) ([32]byte, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetReportBlockHash(&_MoonriverDelegatorCoverOracle.CallOpts, _eraId)
}

// GetReportBlockHash is a free data retrieval call binding the contract method 0x0e53fb48.
//
// Solidity: function getReportBlockHash(uint128 _eraId) view returns(bytes32)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) GetReportBlockHash(_eraId *big.Int) ([32]byte, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetReportBlockHash(&_MoonriverDelegatorCoverOracle.CallOpts, _eraId)
}

// IsReported is a free data retrieval call binding the contract method 0x89b80361.
//
// Solidity: function isReported(uint256 _index) view returns(uint128, bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) IsReported(opts *bind.CallOpts, _index *big.Int) (*big.Int, bool, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "isReported", _index)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// IsReported is a free data retrieval call binding the contract method 0x89b80361.
//
// Solidity: function isReported(uint256 _index) view returns(uint128, bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) IsReported(_index *big.Int) (*big.Int, bool, error) {
	return _MoonriverDelegatorCoverOracle.Contract.IsReported(&_MoonriverDelegatorCoverOracle.CallOpts, _index)
}

// IsReported is a free data retrieval call binding the contract method 0x89b80361.
//
// Solidity: function isReported(uint256 _index) view returns(uint128, bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) IsReported(_index *big.Int) (*big.Int, bool, error) {
	return _MoonriverDelegatorCoverOracle.Contract.IsReported(&_MoonriverDelegatorCoverOracle.CallOpts, _index)
}

// AddRemovePushable is a paid mutator transaction binding the contract method 0x30206bcc.
//
// Solidity: function addRemovePushable(address _pushable, bool _toAdd) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) AddRemovePushable(opts *bind.TransactOpts, _pushable common.Address, _toAdd bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "addRemovePushable", _pushable, _toAdd)
}

// AddRemovePushable is a paid mutator transaction binding the contract method 0x30206bcc.
//
// Solidity: function addRemovePushable(address _pushable, bool _toAdd) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) AddRemovePushable(_pushable common.Address, _toAdd bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.AddRemovePushable(&_MoonriverDelegatorCoverOracle.TransactOpts, _pushable, _toAdd)
}

// AddRemovePushable is a paid mutator transaction binding the contract method 0x30206bcc.
//
// Solidity: function addRemovePushable(address _pushable, bool _toAdd) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) AddRemovePushable(_pushable common.Address, _toAdd bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.AddRemovePushable(&_MoonriverDelegatorCoverOracle.TransactOpts, _pushable, _toAdd)
}

// ClearReporting is a paid mutator transaction binding the contract method 0xb75b518c.
//
// Solidity: function clearReporting() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) ClearReporting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "clearReporting")
}

// ClearReporting is a paid mutator transaction binding the contract method 0xb75b518c.
//
// Solidity: function clearReporting() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ClearReporting() (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ClearReporting(&_MoonriverDelegatorCoverOracle.TransactOpts)
}

// ClearReporting is a paid mutator transaction binding the contract method 0xb75b518c.
//
// Solidity: function clearReporting() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) ClearReporting() (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ClearReporting(&_MoonriverDelegatorCoverOracle.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _oracleMaster, address _pushable) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) Initialize(opts *bind.TransactOpts, _oracleMaster common.Address, _pushable common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "initialize", _oracleMaster, _pushable)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _oracleMaster, address _pushable) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Initialize(_oracleMaster common.Address, _pushable common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Initialize(&_MoonriverDelegatorCoverOracle.TransactOpts, _oracleMaster, _pushable)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _oracleMaster, address _pushable) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) Initialize(_oracleMaster common.Address, _pushable common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Initialize(&_MoonriverDelegatorCoverOracle.TransactOpts, _oracleMaster, _pushable)
}

// ReportPara is a paid mutator transaction binding the contract method 0x889cc58b.
//
// Solidity: function reportPara(uint256 _index, uint256 _quorum, uint128 _eraId, uint128 _eraNonce, (uint256,uint128,uint128,uint128,bytes32,uint128,uint128,(address,uint128,bool,uint256,uint256,(address,uint256)[])[],bool) _staking, address _oracleCollator, bool veto, bool vetoDisabled, bool newEra) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) ReportPara(opts *bind.TransactOpts, _index *big.Int, _quorum *big.Int, _eraId *big.Int, _eraNonce *big.Int, _staking TypesOracleData, _oracleCollator common.Address, veto bool, vetoDisabled bool, newEra bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "reportPara", _index, _quorum, _eraId, _eraNonce, _staking, _oracleCollator, veto, vetoDisabled, newEra)
}

// ReportPara is a paid mutator transaction binding the contract method 0x889cc58b.
//
// Solidity: function reportPara(uint256 _index, uint256 _quorum, uint128 _eraId, uint128 _eraNonce, (uint256,uint128,uint128,uint128,bytes32,uint128,uint128,(address,uint128,bool,uint256,uint256,(address,uint256)[])[],bool) _staking, address _oracleCollator, bool veto, bool vetoDisabled, bool newEra) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ReportPara(_index *big.Int, _quorum *big.Int, _eraId *big.Int, _eraNonce *big.Int, _staking TypesOracleData, _oracleCollator common.Address, veto bool, vetoDisabled bool, newEra bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ReportPara(&_MoonriverDelegatorCoverOracle.TransactOpts, _index, _quorum, _eraId, _eraNonce, _staking, _oracleCollator, veto, vetoDisabled, newEra)
}

// ReportPara is a paid mutator transaction binding the contract method 0x889cc58b.
//
// Solidity: function reportPara(uint256 _index, uint256 _quorum, uint128 _eraId, uint128 _eraNonce, (uint256,uint128,uint128,uint128,bytes32,uint128,uint128,(address,uint128,bool,uint256,uint256,(address,uint256)[])[],bool) _staking, address _oracleCollator, bool veto, bool vetoDisabled, bool newEra) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) ReportPara(_index *big.Int, _quorum *big.Int, _eraId *big.Int, _eraNonce *big.Int, _staking TypesOracleData, _oracleCollator common.Address, veto bool, vetoDisabled bool, newEra bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ReportPara(&_MoonriverDelegatorCoverOracle.TransactOpts, _index, _quorum, _eraId, _eraNonce, _staking, _oracleCollator, veto, vetoDisabled, newEra)
}

// SoftenQuorum is a paid mutator transaction binding the contract method 0xf461e54f.
//
// Solidity: function softenQuorum(uint8 _quorum, uint128 _eraId) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) SoftenQuorum(opts *bind.TransactOpts, _quorum uint8, _eraId *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "softenQuorum", _quorum, _eraId)
}

// SoftenQuorum is a paid mutator transaction binding the contract method 0xf461e54f.
//
// Solidity: function softenQuorum(uint8 _quorum, uint128 _eraId) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) SoftenQuorum(_quorum uint8, _eraId *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SoftenQuorum(&_MoonriverDelegatorCoverOracle.TransactOpts, _quorum, _eraId)
}

// SoftenQuorum is a paid mutator transaction binding the contract method 0xf461e54f.
//
// Solidity: function softenQuorum(uint8 _quorum, uint128 _eraId) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) SoftenQuorum(_quorum uint8, _eraId *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SoftenQuorum(&_MoonriverDelegatorCoverOracle.TransactOpts, _quorum, _eraId)
}

// MoonriverDelegatorCoverOracleReportSubmittedIterator is returned from FilterReportSubmitted and is used to iterate over the raw logs and unpacked data for ReportSubmitted events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleReportSubmittedIterator struct {
	Event *MoonriverDelegatorCoverOracleReportSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MoonriverDelegatorCoverOracleReportSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleReportSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MoonriverDelegatorCoverOracleReportSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MoonriverDelegatorCoverOracleReportSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleReportSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleReportSubmitted represents a ReportSubmitted event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleReportSubmitted struct {
	EraId    *big.Int
	EraNonce *big.Int
	Oracle   common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReportSubmitted is a free log retrieval operation binding the contract event 0xbbf23e6bd8173d3423fb94a419f7193574bdca2f2b640364201f708b6a13cbe8.
//
// Solidity: event ReportSubmitted(uint128 eraId, uint128 eraNonce, address oracle)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterReportSubmitted(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleReportSubmittedIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "ReportSubmitted")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleReportSubmittedIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "ReportSubmitted", logs: logs, sub: sub}, nil
}

// WatchReportSubmitted is a free log subscription operation binding the contract event 0xbbf23e6bd8173d3423fb94a419f7193574bdca2f2b640364201f708b6a13cbe8.
//
// Solidity: event ReportSubmitted(uint128 eraId, uint128 eraNonce, address oracle)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchReportSubmitted(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleReportSubmitted) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "ReportSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleReportSubmitted)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "ReportSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReportSubmitted is a log parse operation binding the contract event 0xbbf23e6bd8173d3423fb94a419f7193574bdca2f2b640364201f708b6a13cbe8.
//
// Solidity: event ReportSubmitted(uint128 eraId, uint128 eraNonce, address oracle)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseReportSubmitted(log types.Log) (*MoonriverDelegatorCoverOracleReportSubmitted, error) {
	event := new(MoonriverDelegatorCoverOracleReportSubmitted)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "ReportSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleReportingClearedIterator is returned from FilterReportingCleared and is used to iterate over the raw logs and unpacked data for ReportingCleared events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleReportingClearedIterator struct {
	Event *MoonriverDelegatorCoverOracleReportingCleared // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MoonriverDelegatorCoverOracleReportingClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleReportingCleared)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MoonriverDelegatorCoverOracleReportingCleared)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MoonriverDelegatorCoverOracleReportingClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleReportingClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleReportingCleared represents a ReportingCleared event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleReportingCleared struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterReportingCleared is a free log retrieval operation binding the contract event 0x131fc6443e37553a3c958042fd059569756675c2d859c083e17ec4928c3c70fb.
//
// Solidity: event ReportingCleared()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterReportingCleared(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleReportingClearedIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "ReportingCleared")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleReportingClearedIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "ReportingCleared", logs: logs, sub: sub}, nil
}

// WatchReportingCleared is a free log subscription operation binding the contract event 0x131fc6443e37553a3c958042fd059569756675c2d859c083e17ec4928c3c70fb.
//
// Solidity: event ReportingCleared()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchReportingCleared(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleReportingCleared) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "ReportingCleared")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleReportingCleared)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "ReportingCleared", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReportingCleared is a log parse operation binding the contract event 0x131fc6443e37553a3c958042fd059569756675c2d859c083e17ec4928c3c70fb.
//
// Solidity: event ReportingCleared()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseReportingCleared(log types.Log) (*MoonriverDelegatorCoverOracleReportingCleared, error) {
	event := new(MoonriverDelegatorCoverOracleReportingCleared)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "ReportingCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
