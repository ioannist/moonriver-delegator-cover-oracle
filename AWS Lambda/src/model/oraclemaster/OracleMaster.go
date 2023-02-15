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
	Bond                 *big.Int
	DelegationsTotal     *big.Int
	Points               *big.Int
	CollatorAccount      common.Address
	Active               bool
	TopActiveDelegations []TypesDelegationsData
}

// TypesDelegationsData is an auto generated low-level Go binding around an user-defined struct.
type TypesDelegationsData struct {
	Amount       *big.Int
	OwnerAccount common.Address
}

// TypesOracleData is an auto generated low-level Go binding around an user-defined struct.
type TypesOracleData struct {
	TotalStaked   *big.Int
	TotalSelected *big.Int
	OrbitersCount *big.Int
	Round         *big.Int
	BlockNumber   *big.Int
	BlockHash     [32]byte
	Awarded       *big.Int
	Finalize      bool
	Collators     []TypesCollatorData
}

// MoonriverDelegatorCoverOracleMetaData contains all meta data concerning the MoonriverDelegatorCoverOracle contract.
var MoonriverDelegatorCoverOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"MemberAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"MemberRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"QUORUM\",\"type\":\"uint8\"}],\"name\":\"QuorumChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SudoRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUTH_MANAGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"INACTIVITY_COVER\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"MAX_MEMBERS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"ORACLE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"QUORUM\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"collatorsToOracles\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"eraId\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"firstEraNonce\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"lastEraVetoOracleVoted\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"members\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"oracleBinaryConfig\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"MAX_DELEGATORS_IN_REPORT\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"RUN_PERIOD_IN_SECONDS\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"oraclePointBitmaps\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"oraclesToCollators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"internalType\":\"contractIProxy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"staking\",\"outputs\":[{\"internalType\":\"contractParachainStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"sudo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"vetoOracleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auth_manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_inactivity_cover\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_quorum\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"maxDelegatorsInReport\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"runPeriodInSeconds\",\"type\":\"uint128\"}],\"name\":\"setOracleBinaryConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_quorum\",\"type\":\"uint8\"}],\"name\":\"setQuorum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracleMember\",\"type\":\"address\"}],\"name\":\"addOracleMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracleMember\",\"type\":\"address\"}],\"name\":\"removeOracleMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_pushable\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_toAdd\",\"type\":\"bool\"}],\"name\":\"addRemovePushable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clearReporting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_itsProxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_someCollator\",\"type\":\"address\"}],\"name\":\"removeSudo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vetoOracleMember\",\"type\":\"address\"}],\"name\":\"setVetoOracleMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collator\",\"type\":\"address\"}],\"name\":\"registerAsOracleMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracleMember\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collator\",\"type\":\"address\"}],\"name\":\"unregisterOracleMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collator\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_eraId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_eraNonce\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"totalSelected\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"orbitersCount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"blockNumber\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"awarded\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"finalize\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"bond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegationsTotal\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"points\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collatorAccount\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ownerAccount\",\"type\":\"address\"}],\"internalType\":\"structTypes.DelegationsData[]\",\"name\":\"topActiveDelegations\",\"type\":\"tuple[]\"}],\"internalType\":\"structTypes.CollatorData[]\",\"name\":\"collators\",\"type\":\"tuple[]\"}],\"internalType\":\"structTypes.OracleData\",\"name\":\"_report\",\"type\":\"tuple\"}],\"name\":\"reportPara\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracleMember\",\"type\":\"address\"}],\"name\":\"isReportedLastEra\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"lastEra\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"lastEraNonce\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"lastFirstEraNonce\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"reported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getOracleBinaryConfig\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracleMember\",\"type\":\"address\"}],\"name\":\"getOraclePointBitmap\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getMembersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"_getEra\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]",
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

// AUTHMANAGER is a free data retrieval call binding the contract method 0xa1e206c7.
//
// Solidity: function AUTH_MANAGER() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) AUTHMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "AUTH_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AUTHMANAGER is a free data retrieval call binding the contract method 0xa1e206c7.
//
// Solidity: function AUTH_MANAGER() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) AUTHMANAGER() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.AUTHMANAGER(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// AUTHMANAGER is a free data retrieval call binding the contract method 0xa1e206c7.
//
// Solidity: function AUTH_MANAGER() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) AUTHMANAGER() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.AUTHMANAGER(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// INACTIVITYCOVER is a free data retrieval call binding the contract method 0xc4f305d9.
//
// Solidity: function INACTIVITY_COVER() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) INACTIVITYCOVER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "INACTIVITY_COVER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// INACTIVITYCOVER is a free data retrieval call binding the contract method 0xc4f305d9.
//
// Solidity: function INACTIVITY_COVER() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) INACTIVITYCOVER() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.INACTIVITYCOVER(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// INACTIVITYCOVER is a free data retrieval call binding the contract method 0xc4f305d9.
//
// Solidity: function INACTIVITY_COVER() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) INACTIVITYCOVER() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.INACTIVITYCOVER(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// MAXMEMBERS is a free data retrieval call binding the contract method 0xea0e35b1.
//
// Solidity: function MAX_MEMBERS() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) MAXMEMBERS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "MAX_MEMBERS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXMEMBERS is a free data retrieval call binding the contract method 0xea0e35b1.
//
// Solidity: function MAX_MEMBERS() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) MAXMEMBERS() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MAXMEMBERS(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// MAXMEMBERS is a free data retrieval call binding the contract method 0xea0e35b1.
//
// Solidity: function MAX_MEMBERS() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) MAXMEMBERS() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MAXMEMBERS(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) ORACLE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "ORACLE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ORACLE() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ORACLE(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) ORACLE() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ORACLE(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// QUORUM is a free data retrieval call binding the contract method 0x2e80d9b6.
//
// Solidity: function QUORUM() view returns(uint8)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) QUORUM(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "QUORUM")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// QUORUM is a free data retrieval call binding the contract method 0x2e80d9b6.
//
// Solidity: function QUORUM() view returns(uint8)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) QUORUM() (uint8, error) {
	return _MoonriverDelegatorCoverOracle.Contract.QUORUM(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// QUORUM is a free data retrieval call binding the contract method 0x2e80d9b6.
//
// Solidity: function QUORUM() view returns(uint8)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) QUORUM() (uint8, error) {
	return _MoonriverDelegatorCoverOracle.Contract.QUORUM(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// GetEra is a free data retrieval call binding the contract method 0xef763b22.
//
// Solidity: function _getEra() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) GetEra(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "_getEra")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEra is a free data retrieval call binding the contract method 0xef763b22.
//
// Solidity: function _getEra() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) GetEra() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetEra(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// GetEra is a free data retrieval call binding the contract method 0xef763b22.
//
// Solidity: function _getEra() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) GetEra() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetEra(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// CollatorsToOracles is a free data retrieval call binding the contract method 0x729086fe.
//
// Solidity: function collatorsToOracles(address ) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) CollatorsToOracles(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "collatorsToOracles", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CollatorsToOracles is a free data retrieval call binding the contract method 0x729086fe.
//
// Solidity: function collatorsToOracles(address ) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) CollatorsToOracles(arg0 common.Address) (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.CollatorsToOracles(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// CollatorsToOracles is a free data retrieval call binding the contract method 0x729086fe.
//
// Solidity: function collatorsToOracles(address ) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) CollatorsToOracles(arg0 common.Address) (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.CollatorsToOracles(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// EraId is a free data retrieval call binding the contract method 0x3f109d23.
//
// Solidity: function eraId() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) EraId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "eraId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EraId is a free data retrieval call binding the contract method 0x3f109d23.
//
// Solidity: function eraId() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) EraId() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.EraId(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// EraId is a free data retrieval call binding the contract method 0x3f109d23.
//
// Solidity: function eraId() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) EraId() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.EraId(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// FirstEraNonce is a free data retrieval call binding the contract method 0xfb6dd3f0.
//
// Solidity: function firstEraNonce() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) FirstEraNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "firstEraNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstEraNonce is a free data retrieval call binding the contract method 0xfb6dd3f0.
//
// Solidity: function firstEraNonce() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) FirstEraNonce() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.FirstEraNonce(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// FirstEraNonce is a free data retrieval call binding the contract method 0xfb6dd3f0.
//
// Solidity: function firstEraNonce() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) FirstEraNonce() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.FirstEraNonce(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// GetMembersCount is a free data retrieval call binding the contract method 0x09772f8f.
//
// Solidity: function getMembersCount() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) GetMembersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "getMembersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMembersCount is a free data retrieval call binding the contract method 0x09772f8f.
//
// Solidity: function getMembersCount() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) GetMembersCount() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetMembersCount(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// GetMembersCount is a free data retrieval call binding the contract method 0x09772f8f.
//
// Solidity: function getMembersCount() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) GetMembersCount() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetMembersCount(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) GetOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "getOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) GetOracle() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetOracle(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) GetOracle() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetOracle(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// GetOracleBinaryConfig is a free data retrieval call binding the contract method 0xe9c53598.
//
// Solidity: function getOracleBinaryConfig() view returns(uint128, uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) GetOracleBinaryConfig(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "getOracleBinaryConfig")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetOracleBinaryConfig is a free data retrieval call binding the contract method 0xe9c53598.
//
// Solidity: function getOracleBinaryConfig() view returns(uint128, uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) GetOracleBinaryConfig() (*big.Int, *big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetOracleBinaryConfig(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// GetOracleBinaryConfig is a free data retrieval call binding the contract method 0xe9c53598.
//
// Solidity: function getOracleBinaryConfig() view returns(uint128, uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) GetOracleBinaryConfig() (*big.Int, *big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetOracleBinaryConfig(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// GetOraclePointBitmap is a free data retrieval call binding the contract method 0xedfcaad4.
//
// Solidity: function getOraclePointBitmap(address _oracleMember) view returns(uint32)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) GetOraclePointBitmap(opts *bind.CallOpts, _oracleMember common.Address) (uint32, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "getOraclePointBitmap", _oracleMember)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetOraclePointBitmap is a free data retrieval call binding the contract method 0xedfcaad4.
//
// Solidity: function getOraclePointBitmap(address _oracleMember) view returns(uint32)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) GetOraclePointBitmap(_oracleMember common.Address) (uint32, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetOraclePointBitmap(&_MoonriverDelegatorCoverOracle.CallOpts, _oracleMember)
}

// GetOraclePointBitmap is a free data retrieval call binding the contract method 0xedfcaad4.
//
// Solidity: function getOraclePointBitmap(address _oracleMember) view returns(uint32)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) GetOraclePointBitmap(_oracleMember common.Address) (uint32, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetOraclePointBitmap(&_MoonriverDelegatorCoverOracle.CallOpts, _oracleMember)
}

// IsReportedLastEra is a free data retrieval call binding the contract method 0x8fbad6b4.
//
// Solidity: function isReportedLastEra(address _oracleMember) view returns(uint128 lastEra, uint128 lastEraNonce, uint128 lastFirstEraNonce, bool reported)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) IsReportedLastEra(opts *bind.CallOpts, _oracleMember common.Address) (struct {
	LastEra           *big.Int
	LastEraNonce      *big.Int
	LastFirstEraNonce *big.Int
	Reported          bool
}, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "isReportedLastEra", _oracleMember)

	outstruct := new(struct {
		LastEra           *big.Int
		LastEraNonce      *big.Int
		LastFirstEraNonce *big.Int
		Reported          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastEra = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastEraNonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastFirstEraNonce = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Reported = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// IsReportedLastEra is a free data retrieval call binding the contract method 0x8fbad6b4.
//
// Solidity: function isReportedLastEra(address _oracleMember) view returns(uint128 lastEra, uint128 lastEraNonce, uint128 lastFirstEraNonce, bool reported)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) IsReportedLastEra(_oracleMember common.Address) (struct {
	LastEra           *big.Int
	LastEraNonce      *big.Int
	LastFirstEraNonce *big.Int
	Reported          bool
}, error) {
	return _MoonriverDelegatorCoverOracle.Contract.IsReportedLastEra(&_MoonriverDelegatorCoverOracle.CallOpts, _oracleMember)
}

// IsReportedLastEra is a free data retrieval call binding the contract method 0x8fbad6b4.
//
// Solidity: function isReportedLastEra(address _oracleMember) view returns(uint128 lastEra, uint128 lastEraNonce, uint128 lastFirstEraNonce, bool reported)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) IsReportedLastEra(_oracleMember common.Address) (struct {
	LastEra           *big.Int
	LastEraNonce      *big.Int
	LastFirstEraNonce *big.Int
	Reported          bool
}, error) {
	return _MoonriverDelegatorCoverOracle.Contract.IsReportedLastEra(&_MoonriverDelegatorCoverOracle.CallOpts, _oracleMember)
}

// LastEraVetoOracleVoted is a free data retrieval call binding the contract method 0x7bfa0000.
//
// Solidity: function lastEraVetoOracleVoted() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) LastEraVetoOracleVoted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "lastEraVetoOracleVoted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastEraVetoOracleVoted is a free data retrieval call binding the contract method 0x7bfa0000.
//
// Solidity: function lastEraVetoOracleVoted() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) LastEraVetoOracleVoted() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.LastEraVetoOracleVoted(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// LastEraVetoOracleVoted is a free data retrieval call binding the contract method 0x7bfa0000.
//
// Solidity: function lastEraVetoOracleVoted() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) LastEraVetoOracleVoted() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.LastEraVetoOracleVoted(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members(uint256 ) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) Members(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "members", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members(uint256 ) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Members(arg0 *big.Int) (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Members(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x5daf08ca.
//
// Solidity: function members(uint256 ) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) Members(arg0 *big.Int) (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Members(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// OracleBinaryConfig is a free data retrieval call binding the contract method 0x7bfd54bb.
//
// Solidity: function oracleBinaryConfig() view returns(uint128 MAX_DELEGATORS_IN_REPORT, uint128 RUN_PERIOD_IN_SECONDS)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) OracleBinaryConfig(opts *bind.CallOpts) (struct {
	MAXDELEGATORSINREPORT *big.Int
	RUNPERIODINSECONDS    *big.Int
}, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "oracleBinaryConfig")

	outstruct := new(struct {
		MAXDELEGATORSINREPORT *big.Int
		RUNPERIODINSECONDS    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MAXDELEGATORSINREPORT = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RUNPERIODINSECONDS = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OracleBinaryConfig is a free data retrieval call binding the contract method 0x7bfd54bb.
//
// Solidity: function oracleBinaryConfig() view returns(uint128 MAX_DELEGATORS_IN_REPORT, uint128 RUN_PERIOD_IN_SECONDS)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) OracleBinaryConfig() (struct {
	MAXDELEGATORSINREPORT *big.Int
	RUNPERIODINSECONDS    *big.Int
}, error) {
	return _MoonriverDelegatorCoverOracle.Contract.OracleBinaryConfig(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// OracleBinaryConfig is a free data retrieval call binding the contract method 0x7bfd54bb.
//
// Solidity: function oracleBinaryConfig() view returns(uint128 MAX_DELEGATORS_IN_REPORT, uint128 RUN_PERIOD_IN_SECONDS)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) OracleBinaryConfig() (struct {
	MAXDELEGATORSINREPORT *big.Int
	RUNPERIODINSECONDS    *big.Int
}, error) {
	return _MoonriverDelegatorCoverOracle.Contract.OracleBinaryConfig(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// OraclePointBitmaps is a free data retrieval call binding the contract method 0xb98165f4.
//
// Solidity: function oraclePointBitmaps(address ) view returns(uint32)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) OraclePointBitmaps(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "oraclePointBitmaps", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// OraclePointBitmaps is a free data retrieval call binding the contract method 0xb98165f4.
//
// Solidity: function oraclePointBitmaps(address ) view returns(uint32)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) OraclePointBitmaps(arg0 common.Address) (uint32, error) {
	return _MoonriverDelegatorCoverOracle.Contract.OraclePointBitmaps(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// OraclePointBitmaps is a free data retrieval call binding the contract method 0xb98165f4.
//
// Solidity: function oraclePointBitmaps(address ) view returns(uint32)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) OraclePointBitmaps(arg0 common.Address) (uint32, error) {
	return _MoonriverDelegatorCoverOracle.Contract.OraclePointBitmaps(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// OraclesToCollators is a free data retrieval call binding the contract method 0xb857998a.
//
// Solidity: function oraclesToCollators(address ) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) OraclesToCollators(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "oraclesToCollators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OraclesToCollators is a free data retrieval call binding the contract method 0xb857998a.
//
// Solidity: function oraclesToCollators(address ) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) OraclesToCollators(arg0 common.Address) (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.OraclesToCollators(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// OraclesToCollators is a free data retrieval call binding the contract method 0xb857998a.
//
// Solidity: function oraclesToCollators(address ) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) OraclesToCollators(arg0 common.Address) (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.OraclesToCollators(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Paused() (bool, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Paused(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) Paused() (bool, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Paused(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "proxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Proxy() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Proxy(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) Proxy() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Proxy(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) Staking(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "staking")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Staking() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Staking(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// Staking is a free data retrieval call binding the contract method 0x4cf088d9.
//
// Solidity: function staking() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) Staking() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Staking(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// Sudo is a free data retrieval call binding the contract method 0xe77c48e9.
//
// Solidity: function sudo() view returns(bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) Sudo(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "sudo")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Sudo is a free data retrieval call binding the contract method 0xe77c48e9.
//
// Solidity: function sudo() view returns(bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Sudo() (bool, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Sudo(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// Sudo is a free data retrieval call binding the contract method 0xe77c48e9.
//
// Solidity: function sudo() view returns(bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) Sudo() (bool, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Sudo(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// VetoOracleMember is a free data retrieval call binding the contract method 0x30a8e602.
//
// Solidity: function vetoOracleMember() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) VetoOracleMember(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "vetoOracleMember")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VetoOracleMember is a free data retrieval call binding the contract method 0x30a8e602.
//
// Solidity: function vetoOracleMember() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) VetoOracleMember() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.VetoOracleMember(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// VetoOracleMember is a free data retrieval call binding the contract method 0x30a8e602.
//
// Solidity: function vetoOracleMember() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) VetoOracleMember() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.VetoOracleMember(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// AddOracleMember is a paid mutator transaction binding the contract method 0xc68bf1a0.
//
// Solidity: function addOracleMember(address _collator, address _oracleMember) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) AddOracleMember(opts *bind.TransactOpts, _collator common.Address, _oracleMember common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "addOracleMember", _collator, _oracleMember)
}

// AddOracleMember is a paid mutator transaction binding the contract method 0xc68bf1a0.
//
// Solidity: function addOracleMember(address _collator, address _oracleMember) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) AddOracleMember(_collator common.Address, _oracleMember common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.AddOracleMember(&_MoonriverDelegatorCoverOracle.TransactOpts, _collator, _oracleMember)
}

// AddOracleMember is a paid mutator transaction binding the contract method 0xc68bf1a0.
//
// Solidity: function addOracleMember(address _collator, address _oracleMember) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) AddOracleMember(_collator common.Address, _oracleMember common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.AddOracleMember(&_MoonriverDelegatorCoverOracle.TransactOpts, _collator, _oracleMember)
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

// Initialize is a paid mutator transaction binding the contract method 0x3073cecf.
//
// Solidity: function initialize(address _auth_manager, address _oracle, address _inactivity_cover, uint8 _quorum) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) Initialize(opts *bind.TransactOpts, _auth_manager common.Address, _oracle common.Address, _inactivity_cover common.Address, _quorum uint8) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "initialize", _auth_manager, _oracle, _inactivity_cover, _quorum)
}

// Initialize is a paid mutator transaction binding the contract method 0x3073cecf.
//
// Solidity: function initialize(address _auth_manager, address _oracle, address _inactivity_cover, uint8 _quorum) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Initialize(_auth_manager common.Address, _oracle common.Address, _inactivity_cover common.Address, _quorum uint8) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Initialize(&_MoonriverDelegatorCoverOracle.TransactOpts, _auth_manager, _oracle, _inactivity_cover, _quorum)
}

// Initialize is a paid mutator transaction binding the contract method 0x3073cecf.
//
// Solidity: function initialize(address _auth_manager, address _oracle, address _inactivity_cover, uint8 _quorum) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) Initialize(_auth_manager common.Address, _oracle common.Address, _inactivity_cover common.Address, _quorum uint8) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Initialize(&_MoonriverDelegatorCoverOracle.TransactOpts, _auth_manager, _oracle, _inactivity_cover, _quorum)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Pause() (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Pause(&_MoonriverDelegatorCoverOracle.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) Pause() (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Pause(&_MoonriverDelegatorCoverOracle.TransactOpts)
}

// RegisterAsOracleMember is a paid mutator transaction binding the contract method 0xb984f923.
//
// Solidity: function registerAsOracleMember(address _collator) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) RegisterAsOracleMember(opts *bind.TransactOpts, _collator common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "registerAsOracleMember", _collator)
}

// RegisterAsOracleMember is a paid mutator transaction binding the contract method 0xb984f923.
//
// Solidity: function registerAsOracleMember(address _collator) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) RegisterAsOracleMember(_collator common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.RegisterAsOracleMember(&_MoonriverDelegatorCoverOracle.TransactOpts, _collator)
}

// RegisterAsOracleMember is a paid mutator transaction binding the contract method 0xb984f923.
//
// Solidity: function registerAsOracleMember(address _collator) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) RegisterAsOracleMember(_collator common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.RegisterAsOracleMember(&_MoonriverDelegatorCoverOracle.TransactOpts, _collator)
}

// RemoveOracleMember is a paid mutator transaction binding the contract method 0x85cea53f.
//
// Solidity: function removeOracleMember(address _collator, address _oracleMember) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) RemoveOracleMember(opts *bind.TransactOpts, _collator common.Address, _oracleMember common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "removeOracleMember", _collator, _oracleMember)
}

// RemoveOracleMember is a paid mutator transaction binding the contract method 0x85cea53f.
//
// Solidity: function removeOracleMember(address _collator, address _oracleMember) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) RemoveOracleMember(_collator common.Address, _oracleMember common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.RemoveOracleMember(&_MoonriverDelegatorCoverOracle.TransactOpts, _collator, _oracleMember)
}

// RemoveOracleMember is a paid mutator transaction binding the contract method 0x85cea53f.
//
// Solidity: function removeOracleMember(address _collator, address _oracleMember) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) RemoveOracleMember(_collator common.Address, _oracleMember common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.RemoveOracleMember(&_MoonriverDelegatorCoverOracle.TransactOpts, _collator, _oracleMember)
}

// RemoveSudo is a paid mutator transaction binding the contract method 0x075efca9.
//
// Solidity: function removeSudo(uint256 code, address _itsProxy, address _someCollator) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) RemoveSudo(opts *bind.TransactOpts, code *big.Int, _itsProxy common.Address, _someCollator common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "removeSudo", code, _itsProxy, _someCollator)
}

// RemoveSudo is a paid mutator transaction binding the contract method 0x075efca9.
//
// Solidity: function removeSudo(uint256 code, address _itsProxy, address _someCollator) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) RemoveSudo(code *big.Int, _itsProxy common.Address, _someCollator common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.RemoveSudo(&_MoonriverDelegatorCoverOracle.TransactOpts, code, _itsProxy, _someCollator)
}

// RemoveSudo is a paid mutator transaction binding the contract method 0x075efca9.
//
// Solidity: function removeSudo(uint256 code, address _itsProxy, address _someCollator) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) RemoveSudo(code *big.Int, _itsProxy common.Address, _someCollator common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.RemoveSudo(&_MoonriverDelegatorCoverOracle.TransactOpts, code, _itsProxy, _someCollator)
}

// ReportPara is a paid mutator transaction binding the contract method 0x9e4c5019.
//
// Solidity: function reportPara(address _collator, uint128 _eraId, uint128 _eraNonce, (uint256,uint128,uint128,uint128,uint128,bytes32,uint128,bool,(uint256,uint256,uint128,address,bool,(uint256,address)[])[]) _report) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) ReportPara(opts *bind.TransactOpts, _collator common.Address, _eraId *big.Int, _eraNonce *big.Int, _report TypesOracleData) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "reportPara", _collator, _eraId, _eraNonce, _report)
}

// ReportPara is a paid mutator transaction binding the contract method 0x9e4c5019.
//
// Solidity: function reportPara(address _collator, uint128 _eraId, uint128 _eraNonce, (uint256,uint128,uint128,uint128,uint128,bytes32,uint128,bool,(uint256,uint256,uint128,address,bool,(uint256,address)[])[]) _report) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ReportPara(_collator common.Address, _eraId *big.Int, _eraNonce *big.Int, _report TypesOracleData) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ReportPara(&_MoonriverDelegatorCoverOracle.TransactOpts, _collator, _eraId, _eraNonce, _report)
}

// ReportPara is a paid mutator transaction binding the contract method 0x9e4c5019.
//
// Solidity: function reportPara(address _collator, uint128 _eraId, uint128 _eraNonce, (uint256,uint128,uint128,uint128,uint128,bytes32,uint128,bool,(uint256,uint256,uint128,address,bool,(uint256,address)[])[]) _report) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) ReportPara(_collator common.Address, _eraId *big.Int, _eraNonce *big.Int, _report TypesOracleData) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ReportPara(&_MoonriverDelegatorCoverOracle.TransactOpts, _collator, _eraId, _eraNonce, _report)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Resume() (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Resume(&_MoonriverDelegatorCoverOracle.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) Resume() (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Resume(&_MoonriverDelegatorCoverOracle.TransactOpts)
}

// SetOracleBinaryConfig is a paid mutator transaction binding the contract method 0x5d349fa9.
//
// Solidity: function setOracleBinaryConfig(uint128 maxDelegatorsInReport, uint128 runPeriodInSeconds) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) SetOracleBinaryConfig(opts *bind.TransactOpts, maxDelegatorsInReport *big.Int, runPeriodInSeconds *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "setOracleBinaryConfig", maxDelegatorsInReport, runPeriodInSeconds)
}

// SetOracleBinaryConfig is a paid mutator transaction binding the contract method 0x5d349fa9.
//
// Solidity: function setOracleBinaryConfig(uint128 maxDelegatorsInReport, uint128 runPeriodInSeconds) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) SetOracleBinaryConfig(maxDelegatorsInReport *big.Int, runPeriodInSeconds *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetOracleBinaryConfig(&_MoonriverDelegatorCoverOracle.TransactOpts, maxDelegatorsInReport, runPeriodInSeconds)
}

// SetOracleBinaryConfig is a paid mutator transaction binding the contract method 0x5d349fa9.
//
// Solidity: function setOracleBinaryConfig(uint128 maxDelegatorsInReport, uint128 runPeriodInSeconds) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) SetOracleBinaryConfig(maxDelegatorsInReport *big.Int, runPeriodInSeconds *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetOracleBinaryConfig(&_MoonriverDelegatorCoverOracle.TransactOpts, maxDelegatorsInReport, runPeriodInSeconds)
}

// SetQuorum is a paid mutator transaction binding the contract method 0x6dcab78f.
//
// Solidity: function setQuorum(uint8 _quorum) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) SetQuorum(opts *bind.TransactOpts, _quorum uint8) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "setQuorum", _quorum)
}

// SetQuorum is a paid mutator transaction binding the contract method 0x6dcab78f.
//
// Solidity: function setQuorum(uint8 _quorum) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) SetQuorum(_quorum uint8) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetQuorum(&_MoonriverDelegatorCoverOracle.TransactOpts, _quorum)
}

// SetQuorum is a paid mutator transaction binding the contract method 0x6dcab78f.
//
// Solidity: function setQuorum(uint8 _quorum) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) SetQuorum(_quorum uint8) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetQuorum(&_MoonriverDelegatorCoverOracle.TransactOpts, _quorum)
}

// SetVetoOracleMember is a paid mutator transaction binding the contract method 0x24939907.
//
// Solidity: function setVetoOracleMember(address _vetoOracleMember) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) SetVetoOracleMember(opts *bind.TransactOpts, _vetoOracleMember common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "setVetoOracleMember", _vetoOracleMember)
}

// SetVetoOracleMember is a paid mutator transaction binding the contract method 0x24939907.
//
// Solidity: function setVetoOracleMember(address _vetoOracleMember) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) SetVetoOracleMember(_vetoOracleMember common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetVetoOracleMember(&_MoonriverDelegatorCoverOracle.TransactOpts, _vetoOracleMember)
}

// SetVetoOracleMember is a paid mutator transaction binding the contract method 0x24939907.
//
// Solidity: function setVetoOracleMember(address _vetoOracleMember) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) SetVetoOracleMember(_vetoOracleMember common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetVetoOracleMember(&_MoonriverDelegatorCoverOracle.TransactOpts, _vetoOracleMember)
}

// UnregisterOracleMember is a paid mutator transaction binding the contract method 0x6c5f47a0.
//
// Solidity: function unregisterOracleMember(address _oracleMember, address _collator) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) UnregisterOracleMember(opts *bind.TransactOpts, _oracleMember common.Address, _collator common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "unregisterOracleMember", _oracleMember, _collator)
}

// UnregisterOracleMember is a paid mutator transaction binding the contract method 0x6c5f47a0.
//
// Solidity: function unregisterOracleMember(address _oracleMember, address _collator) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) UnregisterOracleMember(_oracleMember common.Address, _collator common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.UnregisterOracleMember(&_MoonriverDelegatorCoverOracle.TransactOpts, _oracleMember, _collator)
}

// UnregisterOracleMember is a paid mutator transaction binding the contract method 0x6c5f47a0.
//
// Solidity: function unregisterOracleMember(address _oracleMember, address _collator) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) UnregisterOracleMember(_oracleMember common.Address, _collator common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.UnregisterOracleMember(&_MoonriverDelegatorCoverOracle.TransactOpts, _oracleMember, _collator)
}

// MoonriverDelegatorCoverOracleMemberAddedIterator is returned from FilterMemberAdded and is used to iterate over the raw logs and unpacked data for MemberAdded events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberAddedIterator struct {
	Event *MoonriverDelegatorCoverOracleMemberAdded // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleMemberAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleMemberAdded)
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
		it.Event = new(MoonriverDelegatorCoverOracleMemberAdded)
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
func (it *MoonriverDelegatorCoverOracleMemberAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleMemberAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleMemberAdded represents a MemberAdded event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberAdded struct {
	Member common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMemberAdded is a free log retrieval operation binding the contract event 0xb251eb052afc73ffd02ffe85ad79990a8b3fed60d76dbc2fa2fdd7123dffd914.
//
// Solidity: event MemberAdded(address member)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterMemberAdded(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleMemberAddedIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "MemberAdded")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleMemberAddedIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "MemberAdded", logs: logs, sub: sub}, nil
}

// WatchMemberAdded is a free log subscription operation binding the contract event 0xb251eb052afc73ffd02ffe85ad79990a8b3fed60d76dbc2fa2fdd7123dffd914.
//
// Solidity: event MemberAdded(address member)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchMemberAdded(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleMemberAdded) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "MemberAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleMemberAdded)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberAdded", log); err != nil {
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

// ParseMemberAdded is a log parse operation binding the contract event 0xb251eb052afc73ffd02ffe85ad79990a8b3fed60d76dbc2fa2fdd7123dffd914.
//
// Solidity: event MemberAdded(address member)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseMemberAdded(log types.Log) (*MoonriverDelegatorCoverOracleMemberAdded, error) {
	event := new(MoonriverDelegatorCoverOracleMemberAdded)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleMemberRemovedIterator is returned from FilterMemberRemoved and is used to iterate over the raw logs and unpacked data for MemberRemoved events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberRemovedIterator struct {
	Event *MoonriverDelegatorCoverOracleMemberRemoved // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleMemberRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleMemberRemoved)
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
		it.Event = new(MoonriverDelegatorCoverOracleMemberRemoved)
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
func (it *MoonriverDelegatorCoverOracleMemberRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleMemberRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleMemberRemoved represents a MemberRemoved event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberRemoved struct {
	Member common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMemberRemoved is a free log retrieval operation binding the contract event 0x6e76fb4c77256006d9c38ec7d82b45a8c8f3c27b1d6766fffc42dfb8de684492.
//
// Solidity: event MemberRemoved(address member)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterMemberRemoved(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleMemberRemovedIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "MemberRemoved")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleMemberRemovedIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "MemberRemoved", logs: logs, sub: sub}, nil
}

// WatchMemberRemoved is a free log subscription operation binding the contract event 0x6e76fb4c77256006d9c38ec7d82b45a8c8f3c27b1d6766fffc42dfb8de684492.
//
// Solidity: event MemberRemoved(address member)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchMemberRemoved(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleMemberRemoved) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "MemberRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleMemberRemoved)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberRemoved", log); err != nil {
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

// ParseMemberRemoved is a log parse operation binding the contract event 0x6e76fb4c77256006d9c38ec7d82b45a8c8f3c27b1d6766fffc42dfb8de684492.
//
// Solidity: event MemberRemoved(address member)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseMemberRemoved(log types.Log) (*MoonriverDelegatorCoverOracleMemberRemoved, error) {
	event := new(MoonriverDelegatorCoverOracleMemberRemoved)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOraclePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOraclePausedIterator struct {
	Event *MoonriverDelegatorCoverOraclePaused // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOraclePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOraclePaused)
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
		it.Event = new(MoonriverDelegatorCoverOraclePaused)
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
func (it *MoonriverDelegatorCoverOraclePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOraclePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOraclePaused represents a Paused event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOraclePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterPaused(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOraclePausedIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOraclePausedIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOraclePaused) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOraclePaused)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParsePaused(log types.Log) (*MoonriverDelegatorCoverOraclePaused, error) {
	event := new(MoonriverDelegatorCoverOraclePaused)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleQuorumChangedIterator is returned from FilterQuorumChanged and is used to iterate over the raw logs and unpacked data for QuorumChanged events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleQuorumChangedIterator struct {
	Event *MoonriverDelegatorCoverOracleQuorumChanged // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleQuorumChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleQuorumChanged)
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
		it.Event = new(MoonriverDelegatorCoverOracleQuorumChanged)
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
func (it *MoonriverDelegatorCoverOracleQuorumChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleQuorumChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleQuorumChanged represents a QuorumChanged event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleQuorumChanged struct {
	QUORUM uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterQuorumChanged is a free log retrieval operation binding the contract event 0xd6ecd940f5784ceceeeb048268900f819851986dbe41c49f6e754b988868df87.
//
// Solidity: event QuorumChanged(uint8 QUORUM)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterQuorumChanged(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleQuorumChangedIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "QuorumChanged")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleQuorumChangedIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "QuorumChanged", logs: logs, sub: sub}, nil
}

// WatchQuorumChanged is a free log subscription operation binding the contract event 0xd6ecd940f5784ceceeeb048268900f819851986dbe41c49f6e754b988868df87.
//
// Solidity: event QuorumChanged(uint8 QUORUM)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchQuorumChanged(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleQuorumChanged) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "QuorumChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleQuorumChanged)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "QuorumChanged", log); err != nil {
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

// ParseQuorumChanged is a log parse operation binding the contract event 0xd6ecd940f5784ceceeeb048268900f819851986dbe41c49f6e754b988868df87.
//
// Solidity: event QuorumChanged(uint8 QUORUM)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseQuorumChanged(log types.Log) (*MoonriverDelegatorCoverOracleQuorumChanged, error) {
	event := new(MoonriverDelegatorCoverOracleQuorumChanged)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "QuorumChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleSudoRemovedIterator is returned from FilterSudoRemoved and is used to iterate over the raw logs and unpacked data for SudoRemoved events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleSudoRemovedIterator struct {
	Event *MoonriverDelegatorCoverOracleSudoRemoved // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleSudoRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleSudoRemoved)
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
		it.Event = new(MoonriverDelegatorCoverOracleSudoRemoved)
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
func (it *MoonriverDelegatorCoverOracleSudoRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleSudoRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleSudoRemoved represents a SudoRemoved event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleSudoRemoved struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSudoRemoved is a free log retrieval operation binding the contract event 0x6bca71b4113b31893ecdb4c4b3dfc3bd62c1bab475715e9b46009104aa9a2582.
//
// Solidity: event SudoRemoved()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterSudoRemoved(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleSudoRemovedIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "SudoRemoved")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleSudoRemovedIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "SudoRemoved", logs: logs, sub: sub}, nil
}

// WatchSudoRemoved is a free log subscription operation binding the contract event 0x6bca71b4113b31893ecdb4c4b3dfc3bd62c1bab475715e9b46009104aa9a2582.
//
// Solidity: event SudoRemoved()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchSudoRemoved(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleSudoRemoved) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "SudoRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleSudoRemoved)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "SudoRemoved", log); err != nil {
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

// ParseSudoRemoved is a log parse operation binding the contract event 0x6bca71b4113b31893ecdb4c4b3dfc3bd62c1bab475715e9b46009104aa9a2582.
//
// Solidity: event SudoRemoved()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseSudoRemoved(log types.Log) (*MoonriverDelegatorCoverOracleSudoRemoved, error) {
	event := new(MoonriverDelegatorCoverOracleSudoRemoved)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "SudoRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleUnpausedIterator struct {
	Event *MoonriverDelegatorCoverOracleUnpaused // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleUnpaused)
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
		it.Event = new(MoonriverDelegatorCoverOracleUnpaused)
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
func (it *MoonriverDelegatorCoverOracleUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleUnpaused represents a Unpaused event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleUnpausedIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleUnpausedIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleUnpaused) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleUnpaused)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseUnpaused(log types.Log) (*MoonriverDelegatorCoverOracleUnpaused, error) {
	event := new(MoonriverDelegatorCoverOracleUnpaused)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
