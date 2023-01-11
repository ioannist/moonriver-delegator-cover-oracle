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
}

// MoonriverDelegatorCoverOracleMetaData contains all meta data concerning the MoonriverDelegatorCoverOracle contract.
var MoonriverDelegatorCoverOracleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"CancelDecreaseCoverEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DecreaseCoverEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DecreaseCoverScheduledEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegatorNotPaidEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"eraId\",\"type\":\"uint128\"}],\"name\":\"MemberHasZeroPointsEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"eraId\",\"type\":\"uint128\"}],\"name\":\"MemberNotActiveEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MemberNotPaidEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PayoutEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"eraId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"ReportPushedEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUTH_MANAGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPOSIT_STAKING\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERAS_BETWEEN_FORCED_UNDELEGATION\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_DEPOSIT_TOTAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_DEPOSIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_PAYOUT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE_MASTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKE_UNIT_COVER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelDecreaseCover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"coverOwedTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"candidateDelegationCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorDelegationCount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegatorNotPaid\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"more\",\"type\":\"uint256\"}],\"name\":\"delegator_bond_more\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collator\",\"type\":\"address\"}],\"name\":\"depositCover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eraId\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"erasCovered\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"executeDelegationRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"collator\",\"type\":\"address\"}],\"name\":\"executeScheduled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"getErasCovered\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"getMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collator\",\"type\":\"address\"}],\"name\":\"getPayoutAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"getScheduledDecrease\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auth_manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle_master\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_deposit_staking\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_min_deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_max_deposit_total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stake_unit_cover\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_min_payout\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_eras_between_forced_undelegation\",\"type\":\"uint128\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"memberAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"memberNotPaid\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_noZeroPtsCoverAfterEra\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_noActiveSetCoverAfterEra\",\"type\":\"bool\"}],\"name\":\"memberSetCoverTypes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_max_covered\",\"type\":\"uint256\"}],\"name\":\"memberSetMaxCoveredDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"members\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isMember\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDefaulted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxCoveredDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"lastPushedEra\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"noZeroPtsCoverAfterEra\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"noActiveSetCoverAfterEra\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"membersDepositTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"delegators\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"collators\",\"type\":\"address[]\"}],\"name\":\"payOutCover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_eraId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"totalSelected\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"orbitersCount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"blockNumber\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"awarded\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collatorAccount\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"points\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"bond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegationsTotal\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"ownerAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.DelegationsData[]\",\"name\":\"topActiveDelegations\",\"type\":\"tuple[]\"}],\"internalType\":\"structTypes.CollatorData[]\",\"name\":\"collators\",\"type\":\"tuple[]\"}],\"internalType\":\"structTypes.OracleData\",\"name\":\"_report\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"pushData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetNotPaid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"scheduleDecreaseCover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"scheduleDecreaseCoverManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"less\",\"type\":\"uint256\"}],\"name\":\"schedule_delegator_bond_less\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"schedule_delegator_revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"scheduledDecreasesMap\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"era\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_eras_between_forced_undelegation\",\"type\":\"uint128\"}],\"name\":\"setErasBetweenForcedUndelegations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_execute_delay\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"setExecuteDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_max_deposit_total\",\"type\":\"uint256\"}],\"name\":\"setMaxDepositTotal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_min_deposit\",\"type\":\"uint256\"}],\"name\":\"setMinDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_min_payout\",\"type\":\"uint256\"}],\"name\":\"setMinPayout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_refundOracleGasPrice\",\"type\":\"uint256\"}],\"name\":\"setRefundOracleGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stake_unit_cover\",\"type\":\"uint256\"}],\"name\":\"setStakeUnitCover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"staking\",\"outputs\":[{\"internalType\":\"contractParachainStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalPayouts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMember\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"whitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"withdrawRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// DEPOSITSTAKING is a free data retrieval call binding the contract method 0x92e5e977.
//
// Solidity: function DEPOSIT_STAKING() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) DEPOSITSTAKING(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "DEPOSIT_STAKING")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPOSITSTAKING is a free data retrieval call binding the contract method 0x92e5e977.
//
// Solidity: function DEPOSIT_STAKING() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) DEPOSITSTAKING() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.DEPOSITSTAKING(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// DEPOSITSTAKING is a free data retrieval call binding the contract method 0x92e5e977.
//
// Solidity: function DEPOSIT_STAKING() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) DEPOSITSTAKING() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.DEPOSITSTAKING(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// ERASBETWEENFORCEDUNDELEGATION is a free data retrieval call binding the contract method 0x2123c222.
//
// Solidity: function ERAS_BETWEEN_FORCED_UNDELEGATION() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) ERASBETWEENFORCEDUNDELEGATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "ERAS_BETWEEN_FORCED_UNDELEGATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ERASBETWEENFORCEDUNDELEGATION is a free data retrieval call binding the contract method 0x2123c222.
//
// Solidity: function ERAS_BETWEEN_FORCED_UNDELEGATION() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ERASBETWEENFORCEDUNDELEGATION() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ERASBETWEENFORCEDUNDELEGATION(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// ERASBETWEENFORCEDUNDELEGATION is a free data retrieval call binding the contract method 0x2123c222.
//
// Solidity: function ERAS_BETWEEN_FORCED_UNDELEGATION() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) ERASBETWEENFORCEDUNDELEGATION() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ERASBETWEENFORCEDUNDELEGATION(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// MAXDEPOSITTOTAL is a free data retrieval call binding the contract method 0x7891233e.
//
// Solidity: function MAX_DEPOSIT_TOTAL() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) MAXDEPOSITTOTAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "MAX_DEPOSIT_TOTAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXDEPOSITTOTAL is a free data retrieval call binding the contract method 0x7891233e.
//
// Solidity: function MAX_DEPOSIT_TOTAL() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) MAXDEPOSITTOTAL() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MAXDEPOSITTOTAL(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// MAXDEPOSITTOTAL is a free data retrieval call binding the contract method 0x7891233e.
//
// Solidity: function MAX_DEPOSIT_TOTAL() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) MAXDEPOSITTOTAL() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MAXDEPOSITTOTAL(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// MINDEPOSIT is a free data retrieval call binding the contract method 0xe1e158a5.
//
// Solidity: function MIN_DEPOSIT() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) MINDEPOSIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "MIN_DEPOSIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINDEPOSIT is a free data retrieval call binding the contract method 0xe1e158a5.
//
// Solidity: function MIN_DEPOSIT() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) MINDEPOSIT() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MINDEPOSIT(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// MINDEPOSIT is a free data retrieval call binding the contract method 0xe1e158a5.
//
// Solidity: function MIN_DEPOSIT() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) MINDEPOSIT() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MINDEPOSIT(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// MINPAYOUT is a free data retrieval call binding the contract method 0x3bedcafe.
//
// Solidity: function MIN_PAYOUT() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) MINPAYOUT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "MIN_PAYOUT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINPAYOUT is a free data retrieval call binding the contract method 0x3bedcafe.
//
// Solidity: function MIN_PAYOUT() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) MINPAYOUT() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MINPAYOUT(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// MINPAYOUT is a free data retrieval call binding the contract method 0x3bedcafe.
//
// Solidity: function MIN_PAYOUT() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) MINPAYOUT() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MINPAYOUT(&_MoonriverDelegatorCoverOracle.CallOpts)
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

// STAKEUNITCOVER is a free data retrieval call binding the contract method 0x7a8c9c6b.
//
// Solidity: function STAKE_UNIT_COVER() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) STAKEUNITCOVER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "STAKE_UNIT_COVER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STAKEUNITCOVER is a free data retrieval call binding the contract method 0x7a8c9c6b.
//
// Solidity: function STAKE_UNIT_COVER() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) STAKEUNITCOVER() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.STAKEUNITCOVER(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// STAKEUNITCOVER is a free data retrieval call binding the contract method 0x7a8c9c6b.
//
// Solidity: function STAKE_UNIT_COVER() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) STAKEUNITCOVER() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.STAKEUNITCOVER(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// CoverOwedTotal is a free data retrieval call binding the contract method 0x2f433ab1.
//
// Solidity: function coverOwedTotal() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) CoverOwedTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "coverOwedTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CoverOwedTotal is a free data retrieval call binding the contract method 0x2f433ab1.
//
// Solidity: function coverOwedTotal() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) CoverOwedTotal() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.CoverOwedTotal(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// CoverOwedTotal is a free data retrieval call binding the contract method 0x2f433ab1.
//
// Solidity: function coverOwedTotal() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) CoverOwedTotal() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.CoverOwedTotal(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// DelegatorNotPaid is a free data retrieval call binding the contract method 0xfaf5b40e.
//
// Solidity: function delegatorNotPaid() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) DelegatorNotPaid(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "delegatorNotPaid")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelegatorNotPaid is a free data retrieval call binding the contract method 0xfaf5b40e.
//
// Solidity: function delegatorNotPaid() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) DelegatorNotPaid() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.DelegatorNotPaid(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// DelegatorNotPaid is a free data retrieval call binding the contract method 0xfaf5b40e.
//
// Solidity: function delegatorNotPaid() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) DelegatorNotPaid() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.DelegatorNotPaid(&_MoonriverDelegatorCoverOracle.CallOpts)
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

// ErasCovered is a free data retrieval call binding the contract method 0x3b10f611.
//
// Solidity: function erasCovered(address ) view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) ErasCovered(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "erasCovered", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ErasCovered is a free data retrieval call binding the contract method 0x3b10f611.
//
// Solidity: function erasCovered(address ) view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ErasCovered(arg0 common.Address) (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ErasCovered(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// ErasCovered is a free data retrieval call binding the contract method 0x3b10f611.
//
// Solidity: function erasCovered(address ) view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) ErasCovered(arg0 common.Address) (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ErasCovered(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// GetErasCovered is a free data retrieval call binding the contract method 0xf5551b9a.
//
// Solidity: function getErasCovered(address member) view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) GetErasCovered(opts *bind.CallOpts, member common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "getErasCovered", member)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetErasCovered is a free data retrieval call binding the contract method 0xf5551b9a.
//
// Solidity: function getErasCovered(address member) view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) GetErasCovered(member common.Address) (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetErasCovered(&_MoonriverDelegatorCoverOracle.CallOpts, member)
}

// GetErasCovered is a free data retrieval call binding the contract method 0xf5551b9a.
//
// Solidity: function getErasCovered(address member) view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) GetErasCovered(member common.Address) (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetErasCovered(&_MoonriverDelegatorCoverOracle.CallOpts, member)
}

// GetMember is a free data retrieval call binding the contract method 0x2ada2596.
//
// Solidity: function getMember(address member) view returns(bool, bool, uint256, uint256, uint256, uint128, uint128, uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) GetMember(opts *bind.CallOpts, member common.Address) (bool, bool, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "getMember", member)

	if err != nil {
		return *new(bool), *new(bool), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, err

}

// GetMember is a free data retrieval call binding the contract method 0x2ada2596.
//
// Solidity: function getMember(address member) view returns(bool, bool, uint256, uint256, uint256, uint128, uint128, uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) GetMember(member common.Address) (bool, bool, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetMember(&_MoonriverDelegatorCoverOracle.CallOpts, member)
}

// GetMember is a free data retrieval call binding the contract method 0x2ada2596.
//
// Solidity: function getMember(address member) view returns(bool, bool, uint256, uint256, uint256, uint128, uint128, uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) GetMember(member common.Address) (bool, bool, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetMember(&_MoonriverDelegatorCoverOracle.CallOpts, member)
}

// GetPayoutAmount is a free data retrieval call binding the contract method 0xc94e5024.
//
// Solidity: function getPayoutAmount(address delegator, address collator) view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) GetPayoutAmount(opts *bind.CallOpts, delegator common.Address, collator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "getPayoutAmount", delegator, collator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPayoutAmount is a free data retrieval call binding the contract method 0xc94e5024.
//
// Solidity: function getPayoutAmount(address delegator, address collator) view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) GetPayoutAmount(delegator common.Address, collator common.Address) (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetPayoutAmount(&_MoonriverDelegatorCoverOracle.CallOpts, delegator, collator)
}

// GetPayoutAmount is a free data retrieval call binding the contract method 0xc94e5024.
//
// Solidity: function getPayoutAmount(address delegator, address collator) view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) GetPayoutAmount(delegator common.Address, collator common.Address) (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetPayoutAmount(&_MoonriverDelegatorCoverOracle.CallOpts, delegator, collator)
}

// GetScheduledDecrease is a free data retrieval call binding the contract method 0xb15188f4.
//
// Solidity: function getScheduledDecrease(address member) view returns(uint128, uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) GetScheduledDecrease(opts *bind.CallOpts, member common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "getScheduledDecrease", member)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetScheduledDecrease is a free data retrieval call binding the contract method 0xb15188f4.
//
// Solidity: function getScheduledDecrease(address member) view returns(uint128, uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) GetScheduledDecrease(member common.Address) (*big.Int, *big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetScheduledDecrease(&_MoonriverDelegatorCoverOracle.CallOpts, member)
}

// GetScheduledDecrease is a free data retrieval call binding the contract method 0xb15188f4.
//
// Solidity: function getScheduledDecrease(address member) view returns(uint128, uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) GetScheduledDecrease(member common.Address) (*big.Int, *big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetScheduledDecrease(&_MoonriverDelegatorCoverOracle.CallOpts, member)
}

// MemberAddresses is a free data retrieval call binding the contract method 0xcc4e6299.
//
// Solidity: function memberAddresses(uint256 ) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) MemberAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "memberAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MemberAddresses is a free data retrieval call binding the contract method 0xcc4e6299.
//
// Solidity: function memberAddresses(uint256 ) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) MemberAddresses(arg0 *big.Int) (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MemberAddresses(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// MemberAddresses is a free data retrieval call binding the contract method 0xcc4e6299.
//
// Solidity: function memberAddresses(uint256 ) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) MemberAddresses(arg0 *big.Int) (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MemberAddresses(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// MemberNotPaid is a free data retrieval call binding the contract method 0x3fb6829f.
//
// Solidity: function memberNotPaid() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) MemberNotPaid(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "memberNotPaid")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MemberNotPaid is a free data retrieval call binding the contract method 0x3fb6829f.
//
// Solidity: function memberNotPaid() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) MemberNotPaid() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MemberNotPaid(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// MemberNotPaid is a free data retrieval call binding the contract method 0x3fb6829f.
//
// Solidity: function memberNotPaid() view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) MemberNotPaid() (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MemberNotPaid(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bool isMember, bool active, uint256 deposit, uint256 maxDefaulted, uint256 maxCoveredDelegation, uint128 lastPushedEra, uint128 noZeroPtsCoverAfterEra, uint128 noActiveSetCoverAfterEra)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) Members(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsMember                 bool
	Active                   bool
	Deposit                  *big.Int
	MaxDefaulted             *big.Int
	MaxCoveredDelegation     *big.Int
	LastPushedEra            *big.Int
	NoZeroPtsCoverAfterEra   *big.Int
	NoActiveSetCoverAfterEra *big.Int
}, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "members", arg0)

	outstruct := new(struct {
		IsMember                 bool
		Active                   bool
		Deposit                  *big.Int
		MaxDefaulted             *big.Int
		MaxCoveredDelegation     *big.Int
		LastPushedEra            *big.Int
		NoZeroPtsCoverAfterEra   *big.Int
		NoActiveSetCoverAfterEra *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsMember = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Active = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Deposit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxDefaulted = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxCoveredDelegation = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LastPushedEra = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.NoZeroPtsCoverAfterEra = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.NoActiveSetCoverAfterEra = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bool isMember, bool active, uint256 deposit, uint256 maxDefaulted, uint256 maxCoveredDelegation, uint128 lastPushedEra, uint128 noZeroPtsCoverAfterEra, uint128 noActiveSetCoverAfterEra)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Members(arg0 common.Address) (struct {
	IsMember                 bool
	Active                   bool
	Deposit                  *big.Int
	MaxDefaulted             *big.Int
	MaxCoveredDelegation     *big.Int
	LastPushedEra            *big.Int
	NoZeroPtsCoverAfterEra   *big.Int
	NoActiveSetCoverAfterEra *big.Int
}, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Members(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bool isMember, bool active, uint256 deposit, uint256 maxDefaulted, uint256 maxCoveredDelegation, uint128 lastPushedEra, uint128 noZeroPtsCoverAfterEra, uint128 noActiveSetCoverAfterEra)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) Members(arg0 common.Address) (struct {
	IsMember                 bool
	Active                   bool
	Deposit                  *big.Int
	MaxDefaulted             *big.Int
	MaxCoveredDelegation     *big.Int
	LastPushedEra            *big.Int
	NoZeroPtsCoverAfterEra   *big.Int
	NoActiveSetCoverAfterEra *big.Int
}, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Members(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// MembersDepositTotal is a free data retrieval call binding the contract method 0xc4d1f77d.
//
// Solidity: function membersDepositTotal() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) MembersDepositTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "membersDepositTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MembersDepositTotal is a free data retrieval call binding the contract method 0xc4d1f77d.
//
// Solidity: function membersDepositTotal() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) MembersDepositTotal() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MembersDepositTotal(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// MembersDepositTotal is a free data retrieval call binding the contract method 0xc4d1f77d.
//
// Solidity: function membersDepositTotal() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) MembersDepositTotal() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MembersDepositTotal(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// ScheduledDecreasesMap is a free data retrieval call binding the contract method 0xedaba84d.
//
// Solidity: function scheduledDecreasesMap(address ) view returns(uint128 era, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) ScheduledDecreasesMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	Era    *big.Int
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "scheduledDecreasesMap", arg0)

	outstruct := new(struct {
		Era    *big.Int
		Amount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Era = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ScheduledDecreasesMap is a free data retrieval call binding the contract method 0xedaba84d.
//
// Solidity: function scheduledDecreasesMap(address ) view returns(uint128 era, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ScheduledDecreasesMap(arg0 common.Address) (struct {
	Era    *big.Int
	Amount *big.Int
}, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ScheduledDecreasesMap(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// ScheduledDecreasesMap is a free data retrieval call binding the contract method 0xedaba84d.
//
// Solidity: function scheduledDecreasesMap(address ) view returns(uint128 era, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) ScheduledDecreasesMap(arg0 common.Address) (struct {
	Era    *big.Int
	Amount *big.Int
}, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ScheduledDecreasesMap(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
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

// TotalPayouts is a free data retrieval call binding the contract method 0x7fd7a012.
//
// Solidity: function totalPayouts(address ) view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) TotalPayouts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "totalPayouts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPayouts is a free data retrieval call binding the contract method 0x7fd7a012.
//
// Solidity: function totalPayouts(address ) view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) TotalPayouts(arg0 common.Address) (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.TotalPayouts(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// TotalPayouts is a free data retrieval call binding the contract method 0x7fd7a012.
//
// Solidity: function totalPayouts(address ) view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) TotalPayouts(arg0 common.Address) (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.TotalPayouts(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) Whitelisted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "whitelisted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Whitelisted(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) Whitelisted(arg0 common.Address) (bool, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Whitelisted(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// CancelDecreaseCover is a paid mutator transaction binding the contract method 0x9018a5b1.
//
// Solidity: function cancelDecreaseCover() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) CancelDecreaseCover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "cancelDecreaseCover")
}

// CancelDecreaseCover is a paid mutator transaction binding the contract method 0x9018a5b1.
//
// Solidity: function cancelDecreaseCover() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) CancelDecreaseCover() (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.CancelDecreaseCover(&_MoonriverDelegatorCoverOracle.TransactOpts)
}

// CancelDecreaseCover is a paid mutator transaction binding the contract method 0x9018a5b1.
//
// Solidity: function cancelDecreaseCover() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) CancelDecreaseCover() (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.CancelDecreaseCover(&_MoonriverDelegatorCoverOracle.TransactOpts)
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

// DelegatorBondMore is a paid mutator transaction binding the contract method 0xf8331108.
//
// Solidity: function delegator_bond_more(address candidate, uint256 more) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) DelegatorBondMore(opts *bind.TransactOpts, candidate common.Address, more *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "delegator_bond_more", candidate, more)
}

// DelegatorBondMore is a paid mutator transaction binding the contract method 0xf8331108.
//
// Solidity: function delegator_bond_more(address candidate, uint256 more) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) DelegatorBondMore(candidate common.Address, more *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.DelegatorBondMore(&_MoonriverDelegatorCoverOracle.TransactOpts, candidate, more)
}

// DelegatorBondMore is a paid mutator transaction binding the contract method 0xf8331108.
//
// Solidity: function delegator_bond_more(address candidate, uint256 more) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) DelegatorBondMore(candidate common.Address, more *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.DelegatorBondMore(&_MoonriverDelegatorCoverOracle.TransactOpts, candidate, more)
}

// DepositCover is a paid mutator transaction binding the contract method 0xfbe4b783.
//
// Solidity: function depositCover(address collator) payable returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) DepositCover(opts *bind.TransactOpts, collator common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "depositCover", collator)
}

// DepositCover is a paid mutator transaction binding the contract method 0xfbe4b783.
//
// Solidity: function depositCover(address collator) payable returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) DepositCover(collator common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.DepositCover(&_MoonriverDelegatorCoverOracle.TransactOpts, collator)
}

// DepositCover is a paid mutator transaction binding the contract method 0xfbe4b783.
//
// Solidity: function depositCover(address collator) payable returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) DepositCover(collator common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.DepositCover(&_MoonriverDelegatorCoverOracle.TransactOpts, collator)
}

// ExecuteDelegationRequest is a paid mutator transaction binding the contract method 0xe98c8abe.
//
// Solidity: function executeDelegationRequest(address delegator, address candidate) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) ExecuteDelegationRequest(opts *bind.TransactOpts, delegator common.Address, candidate common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "executeDelegationRequest", delegator, candidate)
}

// ExecuteDelegationRequest is a paid mutator transaction binding the contract method 0xe98c8abe.
//
// Solidity: function executeDelegationRequest(address delegator, address candidate) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ExecuteDelegationRequest(delegator common.Address, candidate common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ExecuteDelegationRequest(&_MoonriverDelegatorCoverOracle.TransactOpts, delegator, candidate)
}

// ExecuteDelegationRequest is a paid mutator transaction binding the contract method 0xe98c8abe.
//
// Solidity: function executeDelegationRequest(address delegator, address candidate) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) ExecuteDelegationRequest(delegator common.Address, candidate common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ExecuteDelegationRequest(&_MoonriverDelegatorCoverOracle.TransactOpts, delegator, candidate)
}

// ExecuteScheduled is a paid mutator transaction binding the contract method 0x43fc576d.
//
// Solidity: function executeScheduled(address collator) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) ExecuteScheduled(opts *bind.TransactOpts, collator common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "executeScheduled", collator)
}

// ExecuteScheduled is a paid mutator transaction binding the contract method 0x43fc576d.
//
// Solidity: function executeScheduled(address collator) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ExecuteScheduled(collator common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ExecuteScheduled(&_MoonriverDelegatorCoverOracle.TransactOpts, collator)
}

// ExecuteScheduled is a paid mutator transaction binding the contract method 0x43fc576d.
//
// Solidity: function executeScheduled(address collator) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) ExecuteScheduled(collator common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ExecuteScheduled(&_MoonriverDelegatorCoverOracle.TransactOpts, collator)
}

// Initialize is a paid mutator transaction binding the contract method 0x47486a28.
//
// Solidity: function initialize(address _auth_manager, address _oracle_master, address _deposit_staking, uint256 _min_deposit, uint256 _max_deposit_total, uint256 _stake_unit_cover, uint256 _min_payout, uint128 _eras_between_forced_undelegation) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) Initialize(opts *bind.TransactOpts, _auth_manager common.Address, _oracle_master common.Address, _deposit_staking common.Address, _min_deposit *big.Int, _max_deposit_total *big.Int, _stake_unit_cover *big.Int, _min_payout *big.Int, _eras_between_forced_undelegation *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "initialize", _auth_manager, _oracle_master, _deposit_staking, _min_deposit, _max_deposit_total, _stake_unit_cover, _min_payout, _eras_between_forced_undelegation)
}

// Initialize is a paid mutator transaction binding the contract method 0x47486a28.
//
// Solidity: function initialize(address _auth_manager, address _oracle_master, address _deposit_staking, uint256 _min_deposit, uint256 _max_deposit_total, uint256 _stake_unit_cover, uint256 _min_payout, uint128 _eras_between_forced_undelegation) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Initialize(_auth_manager common.Address, _oracle_master common.Address, _deposit_staking common.Address, _min_deposit *big.Int, _max_deposit_total *big.Int, _stake_unit_cover *big.Int, _min_payout *big.Int, _eras_between_forced_undelegation *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Initialize(&_MoonriverDelegatorCoverOracle.TransactOpts, _auth_manager, _oracle_master, _deposit_staking, _min_deposit, _max_deposit_total, _stake_unit_cover, _min_payout, _eras_between_forced_undelegation)
}

// Initialize is a paid mutator transaction binding the contract method 0x47486a28.
//
// Solidity: function initialize(address _auth_manager, address _oracle_master, address _deposit_staking, uint256 _min_deposit, uint256 _max_deposit_total, uint256 _stake_unit_cover, uint256 _min_payout, uint128 _eras_between_forced_undelegation) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) Initialize(_auth_manager common.Address, _oracle_master common.Address, _deposit_staking common.Address, _min_deposit *big.Int, _max_deposit_total *big.Int, _stake_unit_cover *big.Int, _min_payout *big.Int, _eras_between_forced_undelegation *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Initialize(&_MoonriverDelegatorCoverOracle.TransactOpts, _auth_manager, _oracle_master, _deposit_staking, _min_deposit, _max_deposit_total, _stake_unit_cover, _min_payout, _eras_between_forced_undelegation)
}

// MemberSetCoverTypes is a paid mutator transaction binding the contract method 0x8f45b2ff.
//
// Solidity: function memberSetCoverTypes(bool _noZeroPtsCoverAfterEra, bool _noActiveSetCoverAfterEra) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) MemberSetCoverTypes(opts *bind.TransactOpts, _noZeroPtsCoverAfterEra bool, _noActiveSetCoverAfterEra bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "memberSetCoverTypes", _noZeroPtsCoverAfterEra, _noActiveSetCoverAfterEra)
}

// MemberSetCoverTypes is a paid mutator transaction binding the contract method 0x8f45b2ff.
//
// Solidity: function memberSetCoverTypes(bool _noZeroPtsCoverAfterEra, bool _noActiveSetCoverAfterEra) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) MemberSetCoverTypes(_noZeroPtsCoverAfterEra bool, _noActiveSetCoverAfterEra bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MemberSetCoverTypes(&_MoonriverDelegatorCoverOracle.TransactOpts, _noZeroPtsCoverAfterEra, _noActiveSetCoverAfterEra)
}

// MemberSetCoverTypes is a paid mutator transaction binding the contract method 0x8f45b2ff.
//
// Solidity: function memberSetCoverTypes(bool _noZeroPtsCoverAfterEra, bool _noActiveSetCoverAfterEra) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) MemberSetCoverTypes(_noZeroPtsCoverAfterEra bool, _noActiveSetCoverAfterEra bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MemberSetCoverTypes(&_MoonriverDelegatorCoverOracle.TransactOpts, _noZeroPtsCoverAfterEra, _noActiveSetCoverAfterEra)
}

// MemberSetMaxCoveredDelegation is a paid mutator transaction binding the contract method 0xd4b39845.
//
// Solidity: function memberSetMaxCoveredDelegation(uint256 _max_covered) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) MemberSetMaxCoveredDelegation(opts *bind.TransactOpts, _max_covered *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "memberSetMaxCoveredDelegation", _max_covered)
}

// MemberSetMaxCoveredDelegation is a paid mutator transaction binding the contract method 0xd4b39845.
//
// Solidity: function memberSetMaxCoveredDelegation(uint256 _max_covered) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) MemberSetMaxCoveredDelegation(_max_covered *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MemberSetMaxCoveredDelegation(&_MoonriverDelegatorCoverOracle.TransactOpts, _max_covered)
}

// MemberSetMaxCoveredDelegation is a paid mutator transaction binding the contract method 0xd4b39845.
//
// Solidity: function memberSetMaxCoveredDelegation(uint256 _max_covered) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) MemberSetMaxCoveredDelegation(_max_covered *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MemberSetMaxCoveredDelegation(&_MoonriverDelegatorCoverOracle.TransactOpts, _max_covered)
}

// PayOutCover is a paid mutator transaction binding the contract method 0x37963fb1.
//
// Solidity: function payOutCover(address[] delegators, address[] collators) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) PayOutCover(opts *bind.TransactOpts, delegators []common.Address, collators []common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "payOutCover", delegators, collators)
}

// PayOutCover is a paid mutator transaction binding the contract method 0x37963fb1.
//
// Solidity: function payOutCover(address[] delegators, address[] collators) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) PayOutCover(delegators []common.Address, collators []common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.PayOutCover(&_MoonriverDelegatorCoverOracle.TransactOpts, delegators, collators)
}

// PayOutCover is a paid mutator transaction binding the contract method 0x37963fb1.
//
// Solidity: function payOutCover(address[] delegators, address[] collators) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) PayOutCover(delegators []common.Address, collators []common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.PayOutCover(&_MoonriverDelegatorCoverOracle.TransactOpts, delegators, collators)
}

// PushData is a paid mutator transaction binding the contract method 0x567ea961.
//
// Solidity: function pushData(uint128 _eraId, (uint256,uint128,uint128,uint128,bytes32,uint128,uint128,(address,uint128,bool,uint256,uint256,(address,uint256)[])[]) _report, address _oracle) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) PushData(opts *bind.TransactOpts, _eraId *big.Int, _report TypesOracleData, _oracle common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "pushData", _eraId, _report, _oracle)
}

// PushData is a paid mutator transaction binding the contract method 0x567ea961.
//
// Solidity: function pushData(uint128 _eraId, (uint256,uint128,uint128,uint128,bytes32,uint128,uint128,(address,uint128,bool,uint256,uint256,(address,uint256)[])[]) _report, address _oracle) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) PushData(_eraId *big.Int, _report TypesOracleData, _oracle common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.PushData(&_MoonriverDelegatorCoverOracle.TransactOpts, _eraId, _report, _oracle)
}

// PushData is a paid mutator transaction binding the contract method 0x567ea961.
//
// Solidity: function pushData(uint128 _eraId, (uint256,uint128,uint128,uint128,bytes32,uint128,uint128,(address,uint128,bool,uint256,uint256,(address,uint256)[])[]) _report, address _oracle) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) PushData(_eraId *big.Int, _report TypesOracleData, _oracle common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.PushData(&_MoonriverDelegatorCoverOracle.TransactOpts, _eraId, _report, _oracle)
}

// ResetNotPaid is a paid mutator transaction binding the contract method 0x1aa1f01f.
//
// Solidity: function resetNotPaid() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) ResetNotPaid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "resetNotPaid")
}

// ResetNotPaid is a paid mutator transaction binding the contract method 0x1aa1f01f.
//
// Solidity: function resetNotPaid() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ResetNotPaid() (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ResetNotPaid(&_MoonriverDelegatorCoverOracle.TransactOpts)
}

// ResetNotPaid is a paid mutator transaction binding the contract method 0x1aa1f01f.
//
// Solidity: function resetNotPaid() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) ResetNotPaid() (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ResetNotPaid(&_MoonriverDelegatorCoverOracle.TransactOpts)
}

// ScheduleDecreaseCover is a paid mutator transaction binding the contract method 0xec3fc0f4.
//
// Solidity: function scheduleDecreaseCover(uint256 amount) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) ScheduleDecreaseCover(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "scheduleDecreaseCover", amount)
}

// ScheduleDecreaseCover is a paid mutator transaction binding the contract method 0xec3fc0f4.
//
// Solidity: function scheduleDecreaseCover(uint256 amount) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ScheduleDecreaseCover(amount *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ScheduleDecreaseCover(&_MoonriverDelegatorCoverOracle.TransactOpts, amount)
}

// ScheduleDecreaseCover is a paid mutator transaction binding the contract method 0xec3fc0f4.
//
// Solidity: function scheduleDecreaseCover(uint256 amount) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) ScheduleDecreaseCover(amount *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ScheduleDecreaseCover(&_MoonriverDelegatorCoverOracle.TransactOpts, amount)
}

// ScheduleDecreaseCoverManager is a paid mutator transaction binding the contract method 0x2c92bc31.
//
// Solidity: function scheduleDecreaseCoverManager(uint256 amount, address member) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) ScheduleDecreaseCoverManager(opts *bind.TransactOpts, amount *big.Int, member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "scheduleDecreaseCoverManager", amount, member)
}

// ScheduleDecreaseCoverManager is a paid mutator transaction binding the contract method 0x2c92bc31.
//
// Solidity: function scheduleDecreaseCoverManager(uint256 amount, address member) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ScheduleDecreaseCoverManager(amount *big.Int, member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ScheduleDecreaseCoverManager(&_MoonriverDelegatorCoverOracle.TransactOpts, amount, member)
}

// ScheduleDecreaseCoverManager is a paid mutator transaction binding the contract method 0x2c92bc31.
//
// Solidity: function scheduleDecreaseCoverManager(uint256 amount, address member) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) ScheduleDecreaseCoverManager(amount *big.Int, member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ScheduleDecreaseCoverManager(&_MoonriverDelegatorCoverOracle.TransactOpts, amount, member)
}

// ScheduleDelegatorBondLess is a paid mutator transaction binding the contract method 0x00043acf.
//
// Solidity: function schedule_delegator_bond_less(address candidate, uint256 less) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) ScheduleDelegatorBondLess(opts *bind.TransactOpts, candidate common.Address, less *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "schedule_delegator_bond_less", candidate, less)
}

// ScheduleDelegatorBondLess is a paid mutator transaction binding the contract method 0x00043acf.
//
// Solidity: function schedule_delegator_bond_less(address candidate, uint256 less) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ScheduleDelegatorBondLess(candidate common.Address, less *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ScheduleDelegatorBondLess(&_MoonriverDelegatorCoverOracle.TransactOpts, candidate, less)
}

// ScheduleDelegatorBondLess is a paid mutator transaction binding the contract method 0x00043acf.
//
// Solidity: function schedule_delegator_bond_less(address candidate, uint256 less) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) ScheduleDelegatorBondLess(candidate common.Address, less *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ScheduleDelegatorBondLess(&_MoonriverDelegatorCoverOracle.TransactOpts, candidate, less)
}

// ScheduleDelegatorRevoke is a paid mutator transaction binding the contract method 0xc64c50b1.
//
// Solidity: function schedule_delegator_revoke(address candidate) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) ScheduleDelegatorRevoke(opts *bind.TransactOpts, candidate common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "schedule_delegator_revoke", candidate)
}

// ScheduleDelegatorRevoke is a paid mutator transaction binding the contract method 0xc64c50b1.
//
// Solidity: function schedule_delegator_revoke(address candidate) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ScheduleDelegatorRevoke(candidate common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ScheduleDelegatorRevoke(&_MoonriverDelegatorCoverOracle.TransactOpts, candidate)
}

// ScheduleDelegatorRevoke is a paid mutator transaction binding the contract method 0xc64c50b1.
//
// Solidity: function schedule_delegator_revoke(address candidate) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) ScheduleDelegatorRevoke(candidate common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ScheduleDelegatorRevoke(&_MoonriverDelegatorCoverOracle.TransactOpts, candidate)
}

// SetErasBetweenForcedUndelegations is a paid mutator transaction binding the contract method 0x0a74522b.
//
// Solidity: function setErasBetweenForcedUndelegations(uint128 _eras_between_forced_undelegation) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) SetErasBetweenForcedUndelegations(opts *bind.TransactOpts, _eras_between_forced_undelegation *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "setErasBetweenForcedUndelegations", _eras_between_forced_undelegation)
}

// SetErasBetweenForcedUndelegations is a paid mutator transaction binding the contract method 0x0a74522b.
//
// Solidity: function setErasBetweenForcedUndelegations(uint128 _eras_between_forced_undelegation) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) SetErasBetweenForcedUndelegations(_eras_between_forced_undelegation *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetErasBetweenForcedUndelegations(&_MoonriverDelegatorCoverOracle.TransactOpts, _eras_between_forced_undelegation)
}

// SetErasBetweenForcedUndelegations is a paid mutator transaction binding the contract method 0x0a74522b.
//
// Solidity: function setErasBetweenForcedUndelegations(uint128 _eras_between_forced_undelegation) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) SetErasBetweenForcedUndelegations(_eras_between_forced_undelegation *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetErasBetweenForcedUndelegations(&_MoonriverDelegatorCoverOracle.TransactOpts, _eras_between_forced_undelegation)
}

// SetExecuteDelay is a paid mutator transaction binding the contract method 0x4ff7cde7.
//
// Solidity: function setExecuteDelay(uint128 _execute_delay, address member) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) SetExecuteDelay(opts *bind.TransactOpts, _execute_delay *big.Int, member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "setExecuteDelay", _execute_delay, member)
}

// SetExecuteDelay is a paid mutator transaction binding the contract method 0x4ff7cde7.
//
// Solidity: function setExecuteDelay(uint128 _execute_delay, address member) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) SetExecuteDelay(_execute_delay *big.Int, member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetExecuteDelay(&_MoonriverDelegatorCoverOracle.TransactOpts, _execute_delay, member)
}

// SetExecuteDelay is a paid mutator transaction binding the contract method 0x4ff7cde7.
//
// Solidity: function setExecuteDelay(uint128 _execute_delay, address member) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) SetExecuteDelay(_execute_delay *big.Int, member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetExecuteDelay(&_MoonriverDelegatorCoverOracle.TransactOpts, _execute_delay, member)
}

// SetMaxDepositTotal is a paid mutator transaction binding the contract method 0x70062c7f.
//
// Solidity: function setMaxDepositTotal(uint256 _max_deposit_total) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) SetMaxDepositTotal(opts *bind.TransactOpts, _max_deposit_total *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "setMaxDepositTotal", _max_deposit_total)
}

// SetMaxDepositTotal is a paid mutator transaction binding the contract method 0x70062c7f.
//
// Solidity: function setMaxDepositTotal(uint256 _max_deposit_total) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) SetMaxDepositTotal(_max_deposit_total *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetMaxDepositTotal(&_MoonriverDelegatorCoverOracle.TransactOpts, _max_deposit_total)
}

// SetMaxDepositTotal is a paid mutator transaction binding the contract method 0x70062c7f.
//
// Solidity: function setMaxDepositTotal(uint256 _max_deposit_total) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) SetMaxDepositTotal(_max_deposit_total *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetMaxDepositTotal(&_MoonriverDelegatorCoverOracle.TransactOpts, _max_deposit_total)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(uint256 _min_deposit) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) SetMinDeposit(opts *bind.TransactOpts, _min_deposit *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "setMinDeposit", _min_deposit)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(uint256 _min_deposit) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) SetMinDeposit(_min_deposit *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetMinDeposit(&_MoonriverDelegatorCoverOracle.TransactOpts, _min_deposit)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(uint256 _min_deposit) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) SetMinDeposit(_min_deposit *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetMinDeposit(&_MoonriverDelegatorCoverOracle.TransactOpts, _min_deposit)
}

// SetMinPayout is a paid mutator transaction binding the contract method 0x081b9da4.
//
// Solidity: function setMinPayout(uint256 _min_payout) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) SetMinPayout(opts *bind.TransactOpts, _min_payout *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "setMinPayout", _min_payout)
}

// SetMinPayout is a paid mutator transaction binding the contract method 0x081b9da4.
//
// Solidity: function setMinPayout(uint256 _min_payout) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) SetMinPayout(_min_payout *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetMinPayout(&_MoonriverDelegatorCoverOracle.TransactOpts, _min_payout)
}

// SetMinPayout is a paid mutator transaction binding the contract method 0x081b9da4.
//
// Solidity: function setMinPayout(uint256 _min_payout) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) SetMinPayout(_min_payout *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetMinPayout(&_MoonriverDelegatorCoverOracle.TransactOpts, _min_payout)
}

// SetRefundOracleGasPrice is a paid mutator transaction binding the contract method 0x8adb0ee9.
//
// Solidity: function setRefundOracleGasPrice(uint256 _refundOracleGasPrice) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) SetRefundOracleGasPrice(opts *bind.TransactOpts, _refundOracleGasPrice *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "setRefundOracleGasPrice", _refundOracleGasPrice)
}

// SetRefundOracleGasPrice is a paid mutator transaction binding the contract method 0x8adb0ee9.
//
// Solidity: function setRefundOracleGasPrice(uint256 _refundOracleGasPrice) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) SetRefundOracleGasPrice(_refundOracleGasPrice *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetRefundOracleGasPrice(&_MoonriverDelegatorCoverOracle.TransactOpts, _refundOracleGasPrice)
}

// SetRefundOracleGasPrice is a paid mutator transaction binding the contract method 0x8adb0ee9.
//
// Solidity: function setRefundOracleGasPrice(uint256 _refundOracleGasPrice) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) SetRefundOracleGasPrice(_refundOracleGasPrice *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetRefundOracleGasPrice(&_MoonriverDelegatorCoverOracle.TransactOpts, _refundOracleGasPrice)
}

// SetStakeUnitCover is a paid mutator transaction binding the contract method 0x64ac6d4b.
//
// Solidity: function setStakeUnitCover(uint256 _stake_unit_cover) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) SetStakeUnitCover(opts *bind.TransactOpts, _stake_unit_cover *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "setStakeUnitCover", _stake_unit_cover)
}

// SetStakeUnitCover is a paid mutator transaction binding the contract method 0x64ac6d4b.
//
// Solidity: function setStakeUnitCover(uint256 _stake_unit_cover) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) SetStakeUnitCover(_stake_unit_cover *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetStakeUnitCover(&_MoonriverDelegatorCoverOracle.TransactOpts, _stake_unit_cover)
}

// SetStakeUnitCover is a paid mutator transaction binding the contract method 0x64ac6d4b.
//
// Solidity: function setStakeUnitCover(uint256 _stake_unit_cover) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) SetStakeUnitCover(_stake_unit_cover *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetStakeUnitCover(&_MoonriverDelegatorCoverOracle.TransactOpts, _stake_unit_cover)
}

// Whitelist is a paid mutator transaction binding the contract method 0xf59c3708.
//
// Solidity: function whitelist(address newMember, bool status) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) Whitelist(opts *bind.TransactOpts, newMember common.Address, status bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "whitelist", newMember, status)
}

// Whitelist is a paid mutator transaction binding the contract method 0xf59c3708.
//
// Solidity: function whitelist(address newMember, bool status) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Whitelist(newMember common.Address, status bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Whitelist(&_MoonriverDelegatorCoverOracle.TransactOpts, newMember, status)
}

// Whitelist is a paid mutator transaction binding the contract method 0xf59c3708.
//
// Solidity: function whitelist(address newMember, bool status) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) Whitelist(newMember common.Address, status bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Whitelist(&_MoonriverDelegatorCoverOracle.TransactOpts, newMember, status)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xf3301f6b.
//
// Solidity: function withdrawRewards(uint256 amount, address receiver) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) WithdrawRewards(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "withdrawRewards", amount, receiver)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xf3301f6b.
//
// Solidity: function withdrawRewards(uint256 amount, address receiver) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) WithdrawRewards(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.WithdrawRewards(&_MoonriverDelegatorCoverOracle.TransactOpts, amount, receiver)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xf3301f6b.
//
// Solidity: function withdrawRewards(uint256 amount, address receiver) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) WithdrawRewards(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.WithdrawRewards(&_MoonriverDelegatorCoverOracle.TransactOpts, amount, receiver)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Receive() (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Receive(&_MoonriverDelegatorCoverOracle.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) Receive() (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Receive(&_MoonriverDelegatorCoverOracle.TransactOpts)
}

// MoonriverDelegatorCoverOracleCancelDecreaseCoverEventIterator is returned from FilterCancelDecreaseCoverEvent and is used to iterate over the raw logs and unpacked data for CancelDecreaseCoverEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleCancelDecreaseCoverEventIterator struct {
	Event *MoonriverDelegatorCoverOracleCancelDecreaseCoverEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleCancelDecreaseCoverEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleCancelDecreaseCoverEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleCancelDecreaseCoverEvent)
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
func (it *MoonriverDelegatorCoverOracleCancelDecreaseCoverEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleCancelDecreaseCoverEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleCancelDecreaseCoverEvent represents a CancelDecreaseCoverEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleCancelDecreaseCoverEvent struct {
	Member common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCancelDecreaseCoverEvent is a free log retrieval operation binding the contract event 0xeae932c34d2c04988e3c36ec47086143f196df3deeb1f0423e92886b42640881.
//
// Solidity: event CancelDecreaseCoverEvent(address member)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterCancelDecreaseCoverEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleCancelDecreaseCoverEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "CancelDecreaseCoverEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleCancelDecreaseCoverEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "CancelDecreaseCoverEvent", logs: logs, sub: sub}, nil
}

// WatchCancelDecreaseCoverEvent is a free log subscription operation binding the contract event 0xeae932c34d2c04988e3c36ec47086143f196df3deeb1f0423e92886b42640881.
//
// Solidity: event CancelDecreaseCoverEvent(address member)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchCancelDecreaseCoverEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleCancelDecreaseCoverEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "CancelDecreaseCoverEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleCancelDecreaseCoverEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "CancelDecreaseCoverEvent", log); err != nil {
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

// ParseCancelDecreaseCoverEvent is a log parse operation binding the contract event 0xeae932c34d2c04988e3c36ec47086143f196df3deeb1f0423e92886b42640881.
//
// Solidity: event CancelDecreaseCoverEvent(address member)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseCancelDecreaseCoverEvent(log types.Log) (*MoonriverDelegatorCoverOracleCancelDecreaseCoverEvent, error) {
	event := new(MoonriverDelegatorCoverOracleCancelDecreaseCoverEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "CancelDecreaseCoverEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleDecreaseCoverEventIterator is returned from FilterDecreaseCoverEvent and is used to iterate over the raw logs and unpacked data for DecreaseCoverEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleDecreaseCoverEventIterator struct {
	Event *MoonriverDelegatorCoverOracleDecreaseCoverEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleDecreaseCoverEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleDecreaseCoverEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleDecreaseCoverEvent)
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
func (it *MoonriverDelegatorCoverOracleDecreaseCoverEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleDecreaseCoverEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleDecreaseCoverEvent represents a DecreaseCoverEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleDecreaseCoverEvent struct {
	Member common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreaseCoverEvent is a free log retrieval operation binding the contract event 0x3b96cbb2914cf59559b99ef15348a6ae28b7219c4c8c4779f2229436f921dbf7.
//
// Solidity: event DecreaseCoverEvent(address member, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterDecreaseCoverEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleDecreaseCoverEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "DecreaseCoverEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleDecreaseCoverEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "DecreaseCoverEvent", logs: logs, sub: sub}, nil
}

// WatchDecreaseCoverEvent is a free log subscription operation binding the contract event 0x3b96cbb2914cf59559b99ef15348a6ae28b7219c4c8c4779f2229436f921dbf7.
//
// Solidity: event DecreaseCoverEvent(address member, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchDecreaseCoverEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleDecreaseCoverEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "DecreaseCoverEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleDecreaseCoverEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "DecreaseCoverEvent", log); err != nil {
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

// ParseDecreaseCoverEvent is a log parse operation binding the contract event 0x3b96cbb2914cf59559b99ef15348a6ae28b7219c4c8c4779f2229436f921dbf7.
//
// Solidity: event DecreaseCoverEvent(address member, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseDecreaseCoverEvent(log types.Log) (*MoonriverDelegatorCoverOracleDecreaseCoverEvent, error) {
	event := new(MoonriverDelegatorCoverOracleDecreaseCoverEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "DecreaseCoverEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleDecreaseCoverScheduledEventIterator is returned from FilterDecreaseCoverScheduledEvent and is used to iterate over the raw logs and unpacked data for DecreaseCoverScheduledEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleDecreaseCoverScheduledEventIterator struct {
	Event *MoonriverDelegatorCoverOracleDecreaseCoverScheduledEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleDecreaseCoverScheduledEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleDecreaseCoverScheduledEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleDecreaseCoverScheduledEvent)
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
func (it *MoonriverDelegatorCoverOracleDecreaseCoverScheduledEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleDecreaseCoverScheduledEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleDecreaseCoverScheduledEvent represents a DecreaseCoverScheduledEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleDecreaseCoverScheduledEvent struct {
	Member common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreaseCoverScheduledEvent is a free log retrieval operation binding the contract event 0xfa335089c1b5adb15492b5e0c08cde8c37f0f8283dc633dde417e87db3201d9d.
//
// Solidity: event DecreaseCoverScheduledEvent(address member, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterDecreaseCoverScheduledEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleDecreaseCoverScheduledEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "DecreaseCoverScheduledEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleDecreaseCoverScheduledEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "DecreaseCoverScheduledEvent", logs: logs, sub: sub}, nil
}

// WatchDecreaseCoverScheduledEvent is a free log subscription operation binding the contract event 0xfa335089c1b5adb15492b5e0c08cde8c37f0f8283dc633dde417e87db3201d9d.
//
// Solidity: event DecreaseCoverScheduledEvent(address member, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchDecreaseCoverScheduledEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleDecreaseCoverScheduledEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "DecreaseCoverScheduledEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleDecreaseCoverScheduledEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "DecreaseCoverScheduledEvent", log); err != nil {
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

// ParseDecreaseCoverScheduledEvent is a log parse operation binding the contract event 0xfa335089c1b5adb15492b5e0c08cde8c37f0f8283dc633dde417e87db3201d9d.
//
// Solidity: event DecreaseCoverScheduledEvent(address member, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseDecreaseCoverScheduledEvent(log types.Log) (*MoonriverDelegatorCoverOracleDecreaseCoverScheduledEvent, error) {
	event := new(MoonriverDelegatorCoverOracleDecreaseCoverScheduledEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "DecreaseCoverScheduledEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleDelegatorNotPaidEventIterator is returned from FilterDelegatorNotPaidEvent and is used to iterate over the raw logs and unpacked data for DelegatorNotPaidEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleDelegatorNotPaidEventIterator struct {
	Event *MoonriverDelegatorCoverOracleDelegatorNotPaidEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleDelegatorNotPaidEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleDelegatorNotPaidEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleDelegatorNotPaidEvent)
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
func (it *MoonriverDelegatorCoverOracleDelegatorNotPaidEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleDelegatorNotPaidEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleDelegatorNotPaidEvent represents a DelegatorNotPaidEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleDelegatorNotPaidEvent struct {
	Delegator common.Address
	Collator  common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegatorNotPaidEvent is a free log retrieval operation binding the contract event 0x671a7c8283e391d50fc75124737f4a55976686284f13bceb231061ff2ba12ba0.
//
// Solidity: event DelegatorNotPaidEvent(address delegator, address collator, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterDelegatorNotPaidEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleDelegatorNotPaidEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "DelegatorNotPaidEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleDelegatorNotPaidEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "DelegatorNotPaidEvent", logs: logs, sub: sub}, nil
}

// WatchDelegatorNotPaidEvent is a free log subscription operation binding the contract event 0x671a7c8283e391d50fc75124737f4a55976686284f13bceb231061ff2ba12ba0.
//
// Solidity: event DelegatorNotPaidEvent(address delegator, address collator, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchDelegatorNotPaidEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleDelegatorNotPaidEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "DelegatorNotPaidEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleDelegatorNotPaidEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "DelegatorNotPaidEvent", log); err != nil {
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

// ParseDelegatorNotPaidEvent is a log parse operation binding the contract event 0x671a7c8283e391d50fc75124737f4a55976686284f13bceb231061ff2ba12ba0.
//
// Solidity: event DelegatorNotPaidEvent(address delegator, address collator, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseDelegatorNotPaidEvent(log types.Log) (*MoonriverDelegatorCoverOracleDelegatorNotPaidEvent, error) {
	event := new(MoonriverDelegatorCoverOracleDelegatorNotPaidEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "DelegatorNotPaidEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleDepositEventIterator is returned from FilterDepositEvent and is used to iterate over the raw logs and unpacked data for DepositEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleDepositEventIterator struct {
	Event *MoonriverDelegatorCoverOracleDepositEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleDepositEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleDepositEvent)
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
func (it *MoonriverDelegatorCoverOracleDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleDepositEvent represents a DepositEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleDepositEvent struct {
	Member common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositEvent is a free log retrieval operation binding the contract event 0x2d8a08b6430a894aea608bcaa6013d5d3e263bc49110605e4d4ba76930ae5c29.
//
// Solidity: event DepositEvent(address member, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterDepositEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleDepositEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "DepositEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleDepositEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "DepositEvent", logs: logs, sub: sub}, nil
}

// WatchDepositEvent is a free log subscription operation binding the contract event 0x2d8a08b6430a894aea608bcaa6013d5d3e263bc49110605e4d4ba76930ae5c29.
//
// Solidity: event DepositEvent(address member, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchDepositEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleDepositEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "DepositEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleDepositEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "DepositEvent", log); err != nil {
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

// ParseDepositEvent is a log parse operation binding the contract event 0x2d8a08b6430a894aea608bcaa6013d5d3e263bc49110605e4d4ba76930ae5c29.
//
// Solidity: event DepositEvent(address member, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseDepositEvent(log types.Log) (*MoonriverDelegatorCoverOracleDepositEvent, error) {
	event := new(MoonriverDelegatorCoverOracleDepositEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "DepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleMemberHasZeroPointsEventIterator is returned from FilterMemberHasZeroPointsEvent and is used to iterate over the raw logs and unpacked data for MemberHasZeroPointsEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberHasZeroPointsEventIterator struct {
	Event *MoonriverDelegatorCoverOracleMemberHasZeroPointsEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleMemberHasZeroPointsEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleMemberHasZeroPointsEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleMemberHasZeroPointsEvent)
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
func (it *MoonriverDelegatorCoverOracleMemberHasZeroPointsEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleMemberHasZeroPointsEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleMemberHasZeroPointsEvent represents a MemberHasZeroPointsEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberHasZeroPointsEvent struct {
	Member common.Address
	EraId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMemberHasZeroPointsEvent is a free log retrieval operation binding the contract event 0xd0b69d5408a7f1f0c3e0b6c8edf3004e816a324eb2736a17e979c9f290d24f28.
//
// Solidity: event MemberHasZeroPointsEvent(address member, uint128 eraId)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterMemberHasZeroPointsEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleMemberHasZeroPointsEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "MemberHasZeroPointsEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleMemberHasZeroPointsEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "MemberHasZeroPointsEvent", logs: logs, sub: sub}, nil
}

// WatchMemberHasZeroPointsEvent is a free log subscription operation binding the contract event 0xd0b69d5408a7f1f0c3e0b6c8edf3004e816a324eb2736a17e979c9f290d24f28.
//
// Solidity: event MemberHasZeroPointsEvent(address member, uint128 eraId)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchMemberHasZeroPointsEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleMemberHasZeroPointsEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "MemberHasZeroPointsEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleMemberHasZeroPointsEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberHasZeroPointsEvent", log); err != nil {
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

// ParseMemberHasZeroPointsEvent is a log parse operation binding the contract event 0xd0b69d5408a7f1f0c3e0b6c8edf3004e816a324eb2736a17e979c9f290d24f28.
//
// Solidity: event MemberHasZeroPointsEvent(address member, uint128 eraId)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseMemberHasZeroPointsEvent(log types.Log) (*MoonriverDelegatorCoverOracleMemberHasZeroPointsEvent, error) {
	event := new(MoonriverDelegatorCoverOracleMemberHasZeroPointsEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberHasZeroPointsEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleMemberNotActiveEventIterator is returned from FilterMemberNotActiveEvent and is used to iterate over the raw logs and unpacked data for MemberNotActiveEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberNotActiveEventIterator struct {
	Event *MoonriverDelegatorCoverOracleMemberNotActiveEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleMemberNotActiveEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleMemberNotActiveEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleMemberNotActiveEvent)
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
func (it *MoonriverDelegatorCoverOracleMemberNotActiveEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleMemberNotActiveEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleMemberNotActiveEvent represents a MemberNotActiveEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberNotActiveEvent struct {
	Member common.Address
	EraId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMemberNotActiveEvent is a free log retrieval operation binding the contract event 0xc7516b46dc9ef54bff5d26cb78aa87e62b9fe4f2d0eff8dab9978d873fa81d4a.
//
// Solidity: event MemberNotActiveEvent(address member, uint128 eraId)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterMemberNotActiveEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleMemberNotActiveEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "MemberNotActiveEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleMemberNotActiveEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "MemberNotActiveEvent", logs: logs, sub: sub}, nil
}

// WatchMemberNotActiveEvent is a free log subscription operation binding the contract event 0xc7516b46dc9ef54bff5d26cb78aa87e62b9fe4f2d0eff8dab9978d873fa81d4a.
//
// Solidity: event MemberNotActiveEvent(address member, uint128 eraId)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchMemberNotActiveEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleMemberNotActiveEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "MemberNotActiveEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleMemberNotActiveEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberNotActiveEvent", log); err != nil {
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

// ParseMemberNotActiveEvent is a log parse operation binding the contract event 0xc7516b46dc9ef54bff5d26cb78aa87e62b9fe4f2d0eff8dab9978d873fa81d4a.
//
// Solidity: event MemberNotActiveEvent(address member, uint128 eraId)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseMemberNotActiveEvent(log types.Log) (*MoonriverDelegatorCoverOracleMemberNotActiveEvent, error) {
	event := new(MoonriverDelegatorCoverOracleMemberNotActiveEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberNotActiveEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleMemberNotPaidEventIterator is returned from FilterMemberNotPaidEvent and is used to iterate over the raw logs and unpacked data for MemberNotPaidEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberNotPaidEventIterator struct {
	Event *MoonriverDelegatorCoverOracleMemberNotPaidEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleMemberNotPaidEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleMemberNotPaidEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleMemberNotPaidEvent)
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
func (it *MoonriverDelegatorCoverOracleMemberNotPaidEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleMemberNotPaidEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleMemberNotPaidEvent represents a MemberNotPaidEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberNotPaidEvent struct {
	Member common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMemberNotPaidEvent is a free log retrieval operation binding the contract event 0xacd9d5583b898d830400af5be0b38cb4d963dc701594491b34e441f6e15dedb5.
//
// Solidity: event MemberNotPaidEvent(address member, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterMemberNotPaidEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleMemberNotPaidEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "MemberNotPaidEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleMemberNotPaidEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "MemberNotPaidEvent", logs: logs, sub: sub}, nil
}

// WatchMemberNotPaidEvent is a free log subscription operation binding the contract event 0xacd9d5583b898d830400af5be0b38cb4d963dc701594491b34e441f6e15dedb5.
//
// Solidity: event MemberNotPaidEvent(address member, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchMemberNotPaidEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleMemberNotPaidEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "MemberNotPaidEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleMemberNotPaidEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberNotPaidEvent", log); err != nil {
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

// ParseMemberNotPaidEvent is a log parse operation binding the contract event 0xacd9d5583b898d830400af5be0b38cb4d963dc701594491b34e441f6e15dedb5.
//
// Solidity: event MemberNotPaidEvent(address member, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseMemberNotPaidEvent(log types.Log) (*MoonriverDelegatorCoverOracleMemberNotPaidEvent, error) {
	event := new(MoonriverDelegatorCoverOracleMemberNotPaidEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberNotPaidEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOraclePayoutEventIterator is returned from FilterPayoutEvent and is used to iterate over the raw logs and unpacked data for PayoutEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOraclePayoutEventIterator struct {
	Event *MoonriverDelegatorCoverOraclePayoutEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOraclePayoutEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOraclePayoutEvent)
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
		it.Event = new(MoonriverDelegatorCoverOraclePayoutEvent)
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
func (it *MoonriverDelegatorCoverOraclePayoutEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOraclePayoutEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOraclePayoutEvent represents a PayoutEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOraclePayoutEvent struct {
	Delegator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPayoutEvent is a free log retrieval operation binding the contract event 0x80e009436f30a4f32133e25b51b5349f2842ea9e93ea4fb6cdaaa4b4e518c687.
//
// Solidity: event PayoutEvent(address delegator, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterPayoutEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOraclePayoutEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "PayoutEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOraclePayoutEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "PayoutEvent", logs: logs, sub: sub}, nil
}

// WatchPayoutEvent is a free log subscription operation binding the contract event 0x80e009436f30a4f32133e25b51b5349f2842ea9e93ea4fb6cdaaa4b4e518c687.
//
// Solidity: event PayoutEvent(address delegator, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchPayoutEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOraclePayoutEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "PayoutEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOraclePayoutEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "PayoutEvent", log); err != nil {
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

// ParsePayoutEvent is a log parse operation binding the contract event 0x80e009436f30a4f32133e25b51b5349f2842ea9e93ea4fb6cdaaa4b4e518c687.
//
// Solidity: event PayoutEvent(address delegator, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParsePayoutEvent(log types.Log) (*MoonriverDelegatorCoverOraclePayoutEvent, error) {
	event := new(MoonriverDelegatorCoverOraclePayoutEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "PayoutEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleReportPushedEventIterator is returned from FilterReportPushedEvent and is used to iterate over the raw logs and unpacked data for ReportPushedEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleReportPushedEventIterator struct {
	Event *MoonriverDelegatorCoverOracleReportPushedEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleReportPushedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleReportPushedEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleReportPushedEvent)
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
func (it *MoonriverDelegatorCoverOracleReportPushedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleReportPushedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleReportPushedEvent represents a ReportPushedEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleReportPushedEvent struct {
	EraId  *big.Int
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReportPushedEvent is a free log retrieval operation binding the contract event 0xa9d39cef10d54be71f13a25fb1651fde21e9f2f4b9890dcc06936dceaf098ec5.
//
// Solidity: event ReportPushedEvent(uint128 eraId, address oracle)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterReportPushedEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleReportPushedEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "ReportPushedEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleReportPushedEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "ReportPushedEvent", logs: logs, sub: sub}, nil
}

// WatchReportPushedEvent is a free log subscription operation binding the contract event 0xa9d39cef10d54be71f13a25fb1651fde21e9f2f4b9890dcc06936dceaf098ec5.
//
// Solidity: event ReportPushedEvent(uint128 eraId, address oracle)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchReportPushedEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleReportPushedEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "ReportPushedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleReportPushedEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "ReportPushedEvent", log); err != nil {
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

// ParseReportPushedEvent is a log parse operation binding the contract event 0xa9d39cef10d54be71f13a25fb1651fde21e9f2f4b9890dcc06936dceaf098ec5.
//
// Solidity: event ReportPushedEvent(uint128 eraId, address oracle)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseReportPushedEvent(log types.Log) (*MoonriverDelegatorCoverOracleReportPushedEvent, error) {
	event := new(MoonriverDelegatorCoverOracleReportPushedEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "ReportPushedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
