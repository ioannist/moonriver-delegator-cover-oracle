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

// MoonriverDelegatorCoverOracleMetaData contains all meta data concerning the MoonriverDelegatorCoverOracle contract.
var MoonriverDelegatorCoverOracleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegateEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegatorBondLessEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegatorBondMoreEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"eraId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collator\",\"type\":\"address\"}],\"name\":\"ForceRevokeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collator\",\"type\":\"address\"}],\"name\":\"RevokeEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUTH_MANAGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INACTIVITY_COVER\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isDelegated\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForcedUndelegationEra\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPercentStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakedTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staking\",\"outputs\":[{\"internalType\":\"contractParachainStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auth_manager\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_inactivity_cover\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"candidateDelegationCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorDelegationCount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"more\",\"type\":\"uint256\"}],\"name\":\"delegatorBondMore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"less\",\"type\":\"uint256\"}],\"name\":\"scheduleDelegatorBondLess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"scheduleDelegatorRevoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxPercentStaked\",\"type\":\"uint256\"}],\"name\":\"setMaxPercentStaked\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceScheduleRevoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"getIsDelegated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"getDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getCollatorsDelegated\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// Delegations is a free data retrieval call binding the contract method 0xbffe3486.
//
// Solidity: function delegations(address ) view returns(bool isDelegated, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) Delegations(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsDelegated bool
	Amount      *big.Int
}, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "delegations", arg0)

	outstruct := new(struct {
		IsDelegated bool
		Amount      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsDelegated = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Delegations is a free data retrieval call binding the contract method 0xbffe3486.
//
// Solidity: function delegations(address ) view returns(bool isDelegated, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Delegations(arg0 common.Address) (struct {
	IsDelegated bool
	Amount      *big.Int
}, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Delegations(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// Delegations is a free data retrieval call binding the contract method 0xbffe3486.
//
// Solidity: function delegations(address ) view returns(bool isDelegated, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) Delegations(arg0 common.Address) (struct {
	IsDelegated bool
	Amount      *big.Int
}, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Delegations(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// GetCollatorsDelegated is a free data retrieval call binding the contract method 0x425fc098.
//
// Solidity: function getCollatorsDelegated(uint256 index) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) GetCollatorsDelegated(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "getCollatorsDelegated", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCollatorsDelegated is a free data retrieval call binding the contract method 0x425fc098.
//
// Solidity: function getCollatorsDelegated(uint256 index) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) GetCollatorsDelegated(index *big.Int) (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetCollatorsDelegated(&_MoonriverDelegatorCoverOracle.CallOpts, index)
}

// GetCollatorsDelegated is a free data retrieval call binding the contract method 0x425fc098.
//
// Solidity: function getCollatorsDelegated(uint256 index) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) GetCollatorsDelegated(index *big.Int) (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetCollatorsDelegated(&_MoonriverDelegatorCoverOracle.CallOpts, index)
}

// GetDelegation is a free data retrieval call binding the contract method 0x2b293768.
//
// Solidity: function getDelegation(address candidate) view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) GetDelegation(opts *bind.CallOpts, candidate common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "getDelegation", candidate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDelegation is a free data retrieval call binding the contract method 0x2b293768.
//
// Solidity: function getDelegation(address candidate) view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) GetDelegation(candidate common.Address) (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetDelegation(&_MoonriverDelegatorCoverOracle.CallOpts, candidate)
}

// GetDelegation is a free data retrieval call binding the contract method 0x2b293768.
//
// Solidity: function getDelegation(address candidate) view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) GetDelegation(candidate common.Address) (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetDelegation(&_MoonriverDelegatorCoverOracle.CallOpts, candidate)
}

// GetIsDelegated is a free data retrieval call binding the contract method 0x3fb279fa.
//
// Solidity: function getIsDelegated(address candidate) view returns(bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) GetIsDelegated(opts *bind.CallOpts, candidate common.Address) (bool, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "getIsDelegated", candidate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsDelegated is a free data retrieval call binding the contract method 0x3fb279fa.
//
// Solidity: function getIsDelegated(address candidate) view returns(bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) GetIsDelegated(candidate common.Address) (bool, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetIsDelegated(&_MoonriverDelegatorCoverOracle.CallOpts, candidate)
}

// GetIsDelegated is a free data retrieval call binding the contract method 0x3fb279fa.
//
// Solidity: function getIsDelegated(address candidate) view returns(bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) GetIsDelegated(candidate common.Address) (bool, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetIsDelegated(&_MoonriverDelegatorCoverOracle.CallOpts, candidate)
}

// LastForcedUndelegationEra is a free data retrieval call binding the contract method 0x5db107e4.
//
// Solidity: function lastForcedUndelegationEra() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) LastForcedUndelegationEra(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "lastForcedUndelegationEra")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastForcedUndelegationEra is a free data retrieval call binding the contract method 0x5db107e4.
//
// Solidity: function lastForcedUndelegationEra() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) LastForcedUndelegationEra() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.LastForcedUndelegationEra(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// LastForcedUndelegationEra is a free data retrieval call binding the contract method 0x5db107e4.
//
// Solidity: function lastForcedUndelegationEra() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) LastForcedUndelegationEra() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.LastForcedUndelegationEra(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// MaxPercentStaked is a free data retrieval call binding the contract method 0xee2aecb5.
//
// Solidity: function maxPercentStaked() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) MaxPercentStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "maxPercentStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPercentStaked is a free data retrieval call binding the contract method 0xee2aecb5.
//
// Solidity: function maxPercentStaked() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) MaxPercentStaked() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MaxPercentStaked(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// MaxPercentStaked is a free data retrieval call binding the contract method 0xee2aecb5.
//
// Solidity: function maxPercentStaked() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) MaxPercentStaked() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MaxPercentStaked(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// StakedTotal is a free data retrieval call binding the contract method 0xd66692a7.
//
// Solidity: function stakedTotal() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) StakedTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "stakedTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedTotal is a free data retrieval call binding the contract method 0xd66692a7.
//
// Solidity: function stakedTotal() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) StakedTotal() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.StakedTotal(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// StakedTotal is a free data retrieval call binding the contract method 0xd66692a7.
//
// Solidity: function stakedTotal() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) StakedTotal() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.StakedTotal(&_MoonriverDelegatorCoverOracle.CallOpts)
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

// Delegate is a paid mutator transaction binding the contract method 0x829f5ee3.
//
// Solidity: function delegate(address candidate, uint256 amount, uint256 candidateDelegationCount, uint256 delegatorDelegationCount) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) Delegate(opts *bind.TransactOpts, candidate common.Address, amount *big.Int, candidateDelegationCount *big.Int, delegatorDelegationCount *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "delegate", candidate, amount, candidateDelegationCount, delegatorDelegationCount)
}

// Delegate is a paid mutator transaction binding the contract method 0x829f5ee3.
//
// Solidity: function delegate(address candidate, uint256 amount, uint256 candidateDelegationCount, uint256 delegatorDelegationCount) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Delegate(candidate common.Address, amount *big.Int, candidateDelegationCount *big.Int, delegatorDelegationCount *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Delegate(&_MoonriverDelegatorCoverOracle.TransactOpts, candidate, amount, candidateDelegationCount, delegatorDelegationCount)
}

// Delegate is a paid mutator transaction binding the contract method 0x829f5ee3.
//
// Solidity: function delegate(address candidate, uint256 amount, uint256 candidateDelegationCount, uint256 delegatorDelegationCount) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) Delegate(candidate common.Address, amount *big.Int, candidateDelegationCount *big.Int, delegatorDelegationCount *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Delegate(&_MoonriverDelegatorCoverOracle.TransactOpts, candidate, amount, candidateDelegationCount, delegatorDelegationCount)
}

// DelegatorBondMore is a paid mutator transaction binding the contract method 0x0465135b.
//
// Solidity: function delegatorBondMore(address candidate, uint256 more) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) DelegatorBondMore(opts *bind.TransactOpts, candidate common.Address, more *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "delegatorBondMore", candidate, more)
}

// DelegatorBondMore is a paid mutator transaction binding the contract method 0x0465135b.
//
// Solidity: function delegatorBondMore(address candidate, uint256 more) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) DelegatorBondMore(candidate common.Address, more *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.DelegatorBondMore(&_MoonriverDelegatorCoverOracle.TransactOpts, candidate, more)
}

// DelegatorBondMore is a paid mutator transaction binding the contract method 0x0465135b.
//
// Solidity: function delegatorBondMore(address candidate, uint256 more) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) DelegatorBondMore(candidate common.Address, more *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.DelegatorBondMore(&_MoonriverDelegatorCoverOracle.TransactOpts, candidate, more)
}

// ForceScheduleRevoke is a paid mutator transaction binding the contract method 0xd4cd2f33.
//
// Solidity: function forceScheduleRevoke() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) ForceScheduleRevoke(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "forceScheduleRevoke")
}

// ForceScheduleRevoke is a paid mutator transaction binding the contract method 0xd4cd2f33.
//
// Solidity: function forceScheduleRevoke() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ForceScheduleRevoke() (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ForceScheduleRevoke(&_MoonriverDelegatorCoverOracle.TransactOpts)
}

// ForceScheduleRevoke is a paid mutator transaction binding the contract method 0xd4cd2f33.
//
// Solidity: function forceScheduleRevoke() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) ForceScheduleRevoke() (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ForceScheduleRevoke(&_MoonriverDelegatorCoverOracle.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _auth_manager, address _inactivity_cover) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) Initialize(opts *bind.TransactOpts, _auth_manager common.Address, _inactivity_cover common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "initialize", _auth_manager, _inactivity_cover)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _auth_manager, address _inactivity_cover) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Initialize(_auth_manager common.Address, _inactivity_cover common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Initialize(&_MoonriverDelegatorCoverOracle.TransactOpts, _auth_manager, _inactivity_cover)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _auth_manager, address _inactivity_cover) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) Initialize(_auth_manager common.Address, _inactivity_cover common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Initialize(&_MoonriverDelegatorCoverOracle.TransactOpts, _auth_manager, _inactivity_cover)
}

// ScheduleDelegatorBondLess is a paid mutator transaction binding the contract method 0xc172fd2b.
//
// Solidity: function scheduleDelegatorBondLess(address candidate, uint256 less) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) ScheduleDelegatorBondLess(opts *bind.TransactOpts, candidate common.Address, less *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "scheduleDelegatorBondLess", candidate, less)
}

// ScheduleDelegatorBondLess is a paid mutator transaction binding the contract method 0xc172fd2b.
//
// Solidity: function scheduleDelegatorBondLess(address candidate, uint256 less) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ScheduleDelegatorBondLess(candidate common.Address, less *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ScheduleDelegatorBondLess(&_MoonriverDelegatorCoverOracle.TransactOpts, candidate, less)
}

// ScheduleDelegatorBondLess is a paid mutator transaction binding the contract method 0xc172fd2b.
//
// Solidity: function scheduleDelegatorBondLess(address candidate, uint256 less) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) ScheduleDelegatorBondLess(candidate common.Address, less *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ScheduleDelegatorBondLess(&_MoonriverDelegatorCoverOracle.TransactOpts, candidate, less)
}

// ScheduleDelegatorRevoke is a paid mutator transaction binding the contract method 0xcba6829d.
//
// Solidity: function scheduleDelegatorRevoke(address candidate) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) ScheduleDelegatorRevoke(opts *bind.TransactOpts, candidate common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "scheduleDelegatorRevoke", candidate)
}

// ScheduleDelegatorRevoke is a paid mutator transaction binding the contract method 0xcba6829d.
//
// Solidity: function scheduleDelegatorRevoke(address candidate) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ScheduleDelegatorRevoke(candidate common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ScheduleDelegatorRevoke(&_MoonriverDelegatorCoverOracle.TransactOpts, candidate)
}

// ScheduleDelegatorRevoke is a paid mutator transaction binding the contract method 0xcba6829d.
//
// Solidity: function scheduleDelegatorRevoke(address candidate) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) ScheduleDelegatorRevoke(candidate common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ScheduleDelegatorRevoke(&_MoonriverDelegatorCoverOracle.TransactOpts, candidate)
}

// SetMaxPercentStaked is a paid mutator transaction binding the contract method 0x819a4bea.
//
// Solidity: function setMaxPercentStaked(uint256 _maxPercentStaked) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) SetMaxPercentStaked(opts *bind.TransactOpts, _maxPercentStaked *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "setMaxPercentStaked", _maxPercentStaked)
}

// SetMaxPercentStaked is a paid mutator transaction binding the contract method 0x819a4bea.
//
// Solidity: function setMaxPercentStaked(uint256 _maxPercentStaked) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) SetMaxPercentStaked(_maxPercentStaked *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetMaxPercentStaked(&_MoonriverDelegatorCoverOracle.TransactOpts, _maxPercentStaked)
}

// SetMaxPercentStaked is a paid mutator transaction binding the contract method 0x819a4bea.
//
// Solidity: function setMaxPercentStaked(uint256 _maxPercentStaked) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) SetMaxPercentStaked(_maxPercentStaked *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetMaxPercentStaked(&_MoonriverDelegatorCoverOracle.TransactOpts, _maxPercentStaked)
}

// MoonriverDelegatorCoverOracleDelegateEventIterator is returned from FilterDelegateEvent and is used to iterate over the raw logs and unpacked data for DelegateEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleDelegateEventIterator struct {
	Event *MoonriverDelegatorCoverOracleDelegateEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleDelegateEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleDelegateEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleDelegateEvent)
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
func (it *MoonriverDelegatorCoverOracleDelegateEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleDelegateEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleDelegateEvent represents a DelegateEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleDelegateEvent struct {
	Collator common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelegateEvent is a free log retrieval operation binding the contract event 0x9779ce23d31f8efd9af6e8c72e3038d97df168b146bdd24008d90793c7b06611.
//
// Solidity: event DelegateEvent(address collator, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterDelegateEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleDelegateEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "DelegateEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleDelegateEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "DelegateEvent", logs: logs, sub: sub}, nil
}

// WatchDelegateEvent is a free log subscription operation binding the contract event 0x9779ce23d31f8efd9af6e8c72e3038d97df168b146bdd24008d90793c7b06611.
//
// Solidity: event DelegateEvent(address collator, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchDelegateEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleDelegateEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "DelegateEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleDelegateEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "DelegateEvent", log); err != nil {
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

// ParseDelegateEvent is a log parse operation binding the contract event 0x9779ce23d31f8efd9af6e8c72e3038d97df168b146bdd24008d90793c7b06611.
//
// Solidity: event DelegateEvent(address collator, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseDelegateEvent(log types.Log) (*MoonriverDelegatorCoverOracleDelegateEvent, error) {
	event := new(MoonriverDelegatorCoverOracleDelegateEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "DelegateEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleDelegatorBondLessEventIterator is returned from FilterDelegatorBondLessEvent and is used to iterate over the raw logs and unpacked data for DelegatorBondLessEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleDelegatorBondLessEventIterator struct {
	Event *MoonriverDelegatorCoverOracleDelegatorBondLessEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleDelegatorBondLessEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleDelegatorBondLessEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleDelegatorBondLessEvent)
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
func (it *MoonriverDelegatorCoverOracleDelegatorBondLessEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleDelegatorBondLessEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleDelegatorBondLessEvent represents a DelegatorBondLessEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleDelegatorBondLessEvent struct {
	Collator common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelegatorBondLessEvent is a free log retrieval operation binding the contract event 0x6404908a037b1eb70f1aec40b4369a4098d1abf3fa102d0560c7e47da4dd610d.
//
// Solidity: event DelegatorBondLessEvent(address collator, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterDelegatorBondLessEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleDelegatorBondLessEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "DelegatorBondLessEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleDelegatorBondLessEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "DelegatorBondLessEvent", logs: logs, sub: sub}, nil
}

// WatchDelegatorBondLessEvent is a free log subscription operation binding the contract event 0x6404908a037b1eb70f1aec40b4369a4098d1abf3fa102d0560c7e47da4dd610d.
//
// Solidity: event DelegatorBondLessEvent(address collator, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchDelegatorBondLessEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleDelegatorBondLessEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "DelegatorBondLessEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleDelegatorBondLessEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "DelegatorBondLessEvent", log); err != nil {
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

// ParseDelegatorBondLessEvent is a log parse operation binding the contract event 0x6404908a037b1eb70f1aec40b4369a4098d1abf3fa102d0560c7e47da4dd610d.
//
// Solidity: event DelegatorBondLessEvent(address collator, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseDelegatorBondLessEvent(log types.Log) (*MoonriverDelegatorCoverOracleDelegatorBondLessEvent, error) {
	event := new(MoonriverDelegatorCoverOracleDelegatorBondLessEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "DelegatorBondLessEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleDelegatorBondMoreEventIterator is returned from FilterDelegatorBondMoreEvent and is used to iterate over the raw logs and unpacked data for DelegatorBondMoreEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleDelegatorBondMoreEventIterator struct {
	Event *MoonriverDelegatorCoverOracleDelegatorBondMoreEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleDelegatorBondMoreEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleDelegatorBondMoreEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleDelegatorBondMoreEvent)
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
func (it *MoonriverDelegatorCoverOracleDelegatorBondMoreEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleDelegatorBondMoreEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleDelegatorBondMoreEvent represents a DelegatorBondMoreEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleDelegatorBondMoreEvent struct {
	Collator common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelegatorBondMoreEvent is a free log retrieval operation binding the contract event 0x7c62f80460f0f5dec6a624e12b8825780248cd2ff0c0746ec24713e9ead343e6.
//
// Solidity: event DelegatorBondMoreEvent(address collator, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterDelegatorBondMoreEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleDelegatorBondMoreEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "DelegatorBondMoreEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleDelegatorBondMoreEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "DelegatorBondMoreEvent", logs: logs, sub: sub}, nil
}

// WatchDelegatorBondMoreEvent is a free log subscription operation binding the contract event 0x7c62f80460f0f5dec6a624e12b8825780248cd2ff0c0746ec24713e9ead343e6.
//
// Solidity: event DelegatorBondMoreEvent(address collator, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchDelegatorBondMoreEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleDelegatorBondMoreEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "DelegatorBondMoreEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleDelegatorBondMoreEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "DelegatorBondMoreEvent", log); err != nil {
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

// ParseDelegatorBondMoreEvent is a log parse operation binding the contract event 0x7c62f80460f0f5dec6a624e12b8825780248cd2ff0c0746ec24713e9ead343e6.
//
// Solidity: event DelegatorBondMoreEvent(address collator, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseDelegatorBondMoreEvent(log types.Log) (*MoonriverDelegatorCoverOracleDelegatorBondMoreEvent, error) {
	event := new(MoonriverDelegatorCoverOracleDelegatorBondMoreEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "DelegatorBondMoreEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleForceRevokeEventIterator is returned from FilterForceRevokeEvent and is used to iterate over the raw logs and unpacked data for ForceRevokeEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleForceRevokeEventIterator struct {
	Event *MoonriverDelegatorCoverOracleForceRevokeEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleForceRevokeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleForceRevokeEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleForceRevokeEvent)
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
func (it *MoonriverDelegatorCoverOracleForceRevokeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleForceRevokeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleForceRevokeEvent represents a ForceRevokeEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleForceRevokeEvent struct {
	EraId    *big.Int
	Collator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterForceRevokeEvent is a free log retrieval operation binding the contract event 0xb1858ea5d1bfa2ca34ac4a7759add1b928ee9eb105dc2ccb1eccd9dcab26fef3.
//
// Solidity: event ForceRevokeEvent(uint128 eraId, address collator)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterForceRevokeEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleForceRevokeEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "ForceRevokeEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleForceRevokeEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "ForceRevokeEvent", logs: logs, sub: sub}, nil
}

// WatchForceRevokeEvent is a free log subscription operation binding the contract event 0xb1858ea5d1bfa2ca34ac4a7759add1b928ee9eb105dc2ccb1eccd9dcab26fef3.
//
// Solidity: event ForceRevokeEvent(uint128 eraId, address collator)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchForceRevokeEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleForceRevokeEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "ForceRevokeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleForceRevokeEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "ForceRevokeEvent", log); err != nil {
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

// ParseForceRevokeEvent is a log parse operation binding the contract event 0xb1858ea5d1bfa2ca34ac4a7759add1b928ee9eb105dc2ccb1eccd9dcab26fef3.
//
// Solidity: event ForceRevokeEvent(uint128 eraId, address collator)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseForceRevokeEvent(log types.Log) (*MoonriverDelegatorCoverOracleForceRevokeEvent, error) {
	event := new(MoonriverDelegatorCoverOracleForceRevokeEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "ForceRevokeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleRevokeEventIterator is returned from FilterRevokeEvent and is used to iterate over the raw logs and unpacked data for RevokeEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleRevokeEventIterator struct {
	Event *MoonriverDelegatorCoverOracleRevokeEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleRevokeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleRevokeEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleRevokeEvent)
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
func (it *MoonriverDelegatorCoverOracleRevokeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleRevokeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleRevokeEvent represents a RevokeEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleRevokeEvent struct {
	Collator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRevokeEvent is a free log retrieval operation binding the contract event 0x6a11aacc472a6128c5d0726bc7872c23210ffcdab041e7f0e3630c72cb86f501.
//
// Solidity: event RevokeEvent(address collator)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterRevokeEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleRevokeEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "RevokeEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleRevokeEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "RevokeEvent", logs: logs, sub: sub}, nil
}

// WatchRevokeEvent is a free log subscription operation binding the contract event 0x6a11aacc472a6128c5d0726bc7872c23210ffcdab041e7f0e3630c72cb86f501.
//
// Solidity: event RevokeEvent(address collator)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchRevokeEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleRevokeEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "RevokeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleRevokeEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "RevokeEvent", log); err != nil {
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

// ParseRevokeEvent is a log parse operation binding the contract event 0x6a11aacc472a6128c5d0726bc7872c23210ffcdab041e7f0e3630c72cb86f501.
//
// Solidity: event RevokeEvent(address collator)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseRevokeEvent(log types.Log) (*MoonriverDelegatorCoverOracleRevokeEvent, error) {
	event := new(MoonriverDelegatorCoverOracleRevokeEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "RevokeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
