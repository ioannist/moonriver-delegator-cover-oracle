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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"CancelDecreaseCoverEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DecreaseCoverEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"eraId\",\"type\":\"uint128\"}],\"name\":\"DecreaseCoverScheduledEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelegatorNotPaidEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"DelegatorPayoutLessThanMinEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"MemberActiveStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"MemberDefaultEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"eraId\",\"type\":\"uint128\"}],\"name\":\"MemberHasZeroPointsEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"eraId\",\"type\":\"uint128\"}],\"name\":\"MemberInvoicedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"eraId\",\"type\":\"uint128\"}],\"name\":\"MemberNotInActiveSetEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MemberNotPaidEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"noZeroPtsCoverAfterEra\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"noActiveSetCoverAfterEra\",\"type\":\"bool\"}],\"name\":\"MemberSetCoverTypesEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MemberSetMaxCoveredDelegationEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"eraId\",\"type\":\"uint128\"}],\"name\":\"OraclePaidEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PayoutEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetErasBetweenForcedUndelegationsEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetMaxEraMemberPayoutEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetMaxErasCoveredEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetMemberFeeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetMinPayoutEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"notRequired\",\"type\":\"bool\"}],\"name\":\"SetNoManualWhitelistingRequiredEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetRefundOracleGasPriceEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SetStakeUnitCoverEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"UpgradedToV2Event\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"WhitelistEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUTH_MANAGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"DEPOSIT_STAKING\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"ERAS_BETWEEN_FORCED_UNDELEGATION\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"MAX_DEPOSIT_TOTAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"MAX_ERAS_COVERED\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"MAX_ERA_MEMBER_PAYOUT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"MIN_DEPOSIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"MIN_PAYOUT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"ORACLE_MASTER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"STAKE_UNIT_COVER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"delegatorNotPaid\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"eraId\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"erasCovered\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"memberAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"memberFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"memberNotPaid\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"members\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxCoveredDelegation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorsReportedInEra\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastDelegationsTotal\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"noZeroPtsCoverAfterEra\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"noActiveSetCoverAfterEra\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"defaultCount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"lastPushedEra\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"wentInactiveEra\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"wentActiveEra\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isMember\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"membersDepositTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"membersInvoicedLastEra\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"noManualWhitelistingRequired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"payoutAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"payoutsOwedTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"internalType\":\"contractIProxy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"refundOracleGasPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"runtime2100\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"scheduledDecreasesMap\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"era\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"staking\",\"outputs\":[{\"internalType\":\"contractParachainStaking\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalPayouts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelisted\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"stateMutability\":\"payable\",\"type\":\"receive\",\"payable\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auth_manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle_master\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_deposit_staking\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_min_deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_max_deposit_total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stake_unit_cover\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_min_payout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_max_era_member_payout\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_eras_between_forced_undelegation\",\"type\":\"uint128\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"}],\"name\":\"depositCover\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"scheduleDecreaseCover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"}],\"name\":\"cancelDecreaseCover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_max_covered\",\"type\":\"uint256\"}],\"name\":\"memberSetMaxCoveredDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_noZeroPtsCoverAfterEra\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_noActiveSetCoverAfterEra\",\"type\":\"bool\"}],\"name\":\"memberSetCoverTypes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proxyAccount\",\"type\":\"address\"}],\"name\":\"transferMemberAuth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_member\",\"type\":\"address\"}],\"name\":\"executeScheduled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"delegators\",\"type\":\"address[]\"}],\"name\":\"payOutCover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"executeDelegationRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"invoiceMembers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"withdrawRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proxyAccount\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_paused\",\"type\":\"bool\"}],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_min_deposit\",\"type\":\"uint256\"}],\"name\":\"setMinDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setRuntime2100\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_max_deposit_total\",\"type\":\"uint256\"}],\"name\":\"setMaxDepositTotal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_erasCovered\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"setErasCovered\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_defaultCount\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"}],\"name\":\"setDefaultCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stake_unit_cover\",\"type\":\"uint256\"}],\"name\":\"setStakeUnitCover\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_min_payout\",\"type\":\"uint256\"}],\"name\":\"setMinPayout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_refundOracleGasPrice\",\"type\":\"uint256\"}],\"name\":\"setRefundOracleGasPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_eras_between_forced_undelegation\",\"type\":\"uint128\"}],\"name\":\"setErasBetweenForcedUndelegations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_max_era_member_payout\",\"type\":\"uint256\"}],\"name\":\"setMaxEraMemberPayout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_max_eras_covered\",\"type\":\"uint128\"}],\"name\":\"setMaxErasCovered\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_noManualWhitelistingRequired\",\"type\":\"bool\"}],\"name\":\"setNoManualWhitelistingRequired\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_memberFee\",\"type\":\"uint256\"}],\"name\":\"setMemberFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"getMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"name\":\"getScheduledDecrease\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getMembersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_eraId\",\"type\":\"uint128\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"totalSelected\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"orbitersCount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"round\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"blockNumber\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"awarded\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"finalize\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"bond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegationsTotal\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"points\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"collatorAccount\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ownerAccount\",\"type\":\"address\"}],\"internalType\":\"structTypes.DelegationsData[]\",\"name\":\"topActiveDelegations\",\"type\":\"tuple[]\"}],\"internalType\":\"structTypes.CollatorData[]\",\"name\":\"collators\",\"type\":\"tuple[]\"}],\"internalType\":\"structTypes.OracleData\",\"name\":\"_report\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_oracleCollator\",\"type\":\"address\"}],\"name\":\"pushData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"candidateDelegationCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegatorDelegationCount\",\"type\":\"uint256\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"more\",\"type\":\"uint256\"}],\"name\":\"delegator_bond_more\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"less\",\"type\":\"uint256\"}],\"name\":\"schedule_delegator_bond_less\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"}],\"name\":\"schedule_delegator_revoke\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetNotPaid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// MAXERASCOVERED is a free data retrieval call binding the contract method 0x8b811aa5.
//
// Solidity: function MAX_ERAS_COVERED() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) MAXERASCOVERED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "MAX_ERAS_COVERED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXERASCOVERED is a free data retrieval call binding the contract method 0x8b811aa5.
//
// Solidity: function MAX_ERAS_COVERED() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) MAXERASCOVERED() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MAXERASCOVERED(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// MAXERASCOVERED is a free data retrieval call binding the contract method 0x8b811aa5.
//
// Solidity: function MAX_ERAS_COVERED() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) MAXERASCOVERED() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MAXERASCOVERED(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// MAXERAMEMBERPAYOUT is a free data retrieval call binding the contract method 0xe33052bd.
//
// Solidity: function MAX_ERA_MEMBER_PAYOUT() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) MAXERAMEMBERPAYOUT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "MAX_ERA_MEMBER_PAYOUT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXERAMEMBERPAYOUT is a free data retrieval call binding the contract method 0xe33052bd.
//
// Solidity: function MAX_ERA_MEMBER_PAYOUT() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) MAXERAMEMBERPAYOUT() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MAXERAMEMBERPAYOUT(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// MAXERAMEMBERPAYOUT is a free data retrieval call binding the contract method 0xe33052bd.
//
// Solidity: function MAX_ERA_MEMBER_PAYOUT() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) MAXERAMEMBERPAYOUT() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MAXERAMEMBERPAYOUT(&_MoonriverDelegatorCoverOracle.CallOpts)
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

// GetMember is a free data retrieval call binding the contract method 0x2ada2596.
//
// Solidity: function getMember(address member) view returns(bool, bool, uint256, uint256, uint256, uint128, uint128, uint128, uint128, uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) GetMember(opts *bind.CallOpts, member common.Address) (bool, bool, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "getMember", member)

	if err != nil {
		return *new(bool), *new(bool), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	out8 := *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	out9 := *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, out9, err

}

// GetMember is a free data retrieval call binding the contract method 0x2ada2596.
//
// Solidity: function getMember(address member) view returns(bool, bool, uint256, uint256, uint256, uint128, uint128, uint128, uint128, uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) GetMember(member common.Address) (bool, bool, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetMember(&_MoonriverDelegatorCoverOracle.CallOpts, member)
}

// GetMember is a free data retrieval call binding the contract method 0x2ada2596.
//
// Solidity: function getMember(address member) view returns(bool, bool, uint256, uint256, uint256, uint128, uint128, uint128, uint128, uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) GetMember(member common.Address) (bool, bool, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.GetMember(&_MoonriverDelegatorCoverOracle.CallOpts, member)
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

// MemberFee is a free data retrieval call binding the contract method 0xd3098883.
//
// Solidity: function memberFee() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) MemberFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "memberFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MemberFee is a free data retrieval call binding the contract method 0xd3098883.
//
// Solidity: function memberFee() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) MemberFee() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MemberFee(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// MemberFee is a free data retrieval call binding the contract method 0xd3098883.
//
// Solidity: function memberFee() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) MemberFee() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MemberFee(&_MoonriverDelegatorCoverOracle.CallOpts)
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
// Solidity: function members(address ) view returns(uint256 deposit, uint256 maxCoveredDelegation, uint256 delegatorsReportedInEra, uint256 lastDelegationsTotal, uint128 noZeroPtsCoverAfterEra, uint128 noActiveSetCoverAfterEra, uint128 defaultCount, uint128 lastPushedEra, uint128 wentInactiveEra, uint128 wentActiveEra, bool isMember, bool active)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) Members(opts *bind.CallOpts, arg0 common.Address) (struct {
	Deposit                  *big.Int
	MaxCoveredDelegation     *big.Int
	DelegatorsReportedInEra  *big.Int
	LastDelegationsTotal     *big.Int
	NoZeroPtsCoverAfterEra   *big.Int
	NoActiveSetCoverAfterEra *big.Int
	DefaultCount             *big.Int
	LastPushedEra            *big.Int
	WentInactiveEra          *big.Int
	WentActiveEra            *big.Int
	IsMember                 bool
	Active                   bool
}, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "members", arg0)

	outstruct := new(struct {
		Deposit                  *big.Int
		MaxCoveredDelegation     *big.Int
		DelegatorsReportedInEra  *big.Int
		LastDelegationsTotal     *big.Int
		NoZeroPtsCoverAfterEra   *big.Int
		NoActiveSetCoverAfterEra *big.Int
		DefaultCount             *big.Int
		LastPushedEra            *big.Int
		WentInactiveEra          *big.Int
		WentActiveEra            *big.Int
		IsMember                 bool
		Active                   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Deposit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxCoveredDelegation = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DelegatorsReportedInEra = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LastDelegationsTotal = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.NoZeroPtsCoverAfterEra = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.NoActiveSetCoverAfterEra = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.DefaultCount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.LastPushedEra = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.WentInactiveEra = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.WentActiveEra = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.IsMember = *abi.ConvertType(out[10], new(bool)).(*bool)
	outstruct.Active = *abi.ConvertType(out[11], new(bool)).(*bool)

	return *outstruct, err

}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(uint256 deposit, uint256 maxCoveredDelegation, uint256 delegatorsReportedInEra, uint256 lastDelegationsTotal, uint128 noZeroPtsCoverAfterEra, uint128 noActiveSetCoverAfterEra, uint128 defaultCount, uint128 lastPushedEra, uint128 wentInactiveEra, uint128 wentActiveEra, bool isMember, bool active)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Members(arg0 common.Address) (struct {
	Deposit                  *big.Int
	MaxCoveredDelegation     *big.Int
	DelegatorsReportedInEra  *big.Int
	LastDelegationsTotal     *big.Int
	NoZeroPtsCoverAfterEra   *big.Int
	NoActiveSetCoverAfterEra *big.Int
	DefaultCount             *big.Int
	LastPushedEra            *big.Int
	WentInactiveEra          *big.Int
	WentActiveEra            *big.Int
	IsMember                 bool
	Active                   bool
}, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Members(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(uint256 deposit, uint256 maxCoveredDelegation, uint256 delegatorsReportedInEra, uint256 lastDelegationsTotal, uint128 noZeroPtsCoverAfterEra, uint128 noActiveSetCoverAfterEra, uint128 defaultCount, uint128 lastPushedEra, uint128 wentInactiveEra, uint128 wentActiveEra, bool isMember, bool active)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) Members(arg0 common.Address) (struct {
	Deposit                  *big.Int
	MaxCoveredDelegation     *big.Int
	DelegatorsReportedInEra  *big.Int
	LastDelegationsTotal     *big.Int
	NoZeroPtsCoverAfterEra   *big.Int
	NoActiveSetCoverAfterEra *big.Int
	DefaultCount             *big.Int
	LastPushedEra            *big.Int
	WentInactiveEra          *big.Int
	WentActiveEra            *big.Int
	IsMember                 bool
	Active                   bool
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

// MembersInvoicedLastEra is a free data retrieval call binding the contract method 0xb3230f9d.
//
// Solidity: function membersInvoicedLastEra() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) MembersInvoicedLastEra(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "membersInvoicedLastEra")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MembersInvoicedLastEra is a free data retrieval call binding the contract method 0xb3230f9d.
//
// Solidity: function membersInvoicedLastEra() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) MembersInvoicedLastEra() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MembersInvoicedLastEra(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// MembersInvoicedLastEra is a free data retrieval call binding the contract method 0xb3230f9d.
//
// Solidity: function membersInvoicedLastEra() view returns(uint128)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) MembersInvoicedLastEra() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MembersInvoicedLastEra(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// NoManualWhitelistingRequired is a free data retrieval call binding the contract method 0xcc9b2516.
//
// Solidity: function noManualWhitelistingRequired() view returns(bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) NoManualWhitelistingRequired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "noManualWhitelistingRequired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NoManualWhitelistingRequired is a free data retrieval call binding the contract method 0xcc9b2516.
//
// Solidity: function noManualWhitelistingRequired() view returns(bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) NoManualWhitelistingRequired() (bool, error) {
	return _MoonriverDelegatorCoverOracle.Contract.NoManualWhitelistingRequired(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// NoManualWhitelistingRequired is a free data retrieval call binding the contract method 0xcc9b2516.
//
// Solidity: function noManualWhitelistingRequired() view returns(bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) NoManualWhitelistingRequired() (bool, error) {
	return _MoonriverDelegatorCoverOracle.Contract.NoManualWhitelistingRequired(&_MoonriverDelegatorCoverOracle.CallOpts)
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

// PayoutAmounts is a free data retrieval call binding the contract method 0xd1b3fda0.
//
// Solidity: function payoutAmounts(address ) view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) PayoutAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "payoutAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PayoutAmounts is a free data retrieval call binding the contract method 0xd1b3fda0.
//
// Solidity: function payoutAmounts(address ) view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) PayoutAmounts(arg0 common.Address) (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.PayoutAmounts(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// PayoutAmounts is a free data retrieval call binding the contract method 0xd1b3fda0.
//
// Solidity: function payoutAmounts(address ) view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) PayoutAmounts(arg0 common.Address) (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.PayoutAmounts(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// PayoutsOwedTotal is a free data retrieval call binding the contract method 0x434104a6.
//
// Solidity: function payoutsOwedTotal() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) PayoutsOwedTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "payoutsOwedTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PayoutsOwedTotal is a free data retrieval call binding the contract method 0x434104a6.
//
// Solidity: function payoutsOwedTotal() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) PayoutsOwedTotal() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.PayoutsOwedTotal(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// PayoutsOwedTotal is a free data retrieval call binding the contract method 0x434104a6.
//
// Solidity: function payoutsOwedTotal() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) PayoutsOwedTotal() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.PayoutsOwedTotal(&_MoonriverDelegatorCoverOracle.CallOpts)
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

// RefundOracleGasPrice is a free data retrieval call binding the contract method 0x6d58c80d.
//
// Solidity: function refundOracleGasPrice() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) RefundOracleGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "refundOracleGasPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RefundOracleGasPrice is a free data retrieval call binding the contract method 0x6d58c80d.
//
// Solidity: function refundOracleGasPrice() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) RefundOracleGasPrice() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.RefundOracleGasPrice(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// RefundOracleGasPrice is a free data retrieval call binding the contract method 0x6d58c80d.
//
// Solidity: function refundOracleGasPrice() view returns(uint256)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) RefundOracleGasPrice() (*big.Int, error) {
	return _MoonriverDelegatorCoverOracle.Contract.RefundOracleGasPrice(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// Runtime2100 is a free data retrieval call binding the contract method 0x4dfa5ac6.
//
// Solidity: function runtime2100() view returns(bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) Runtime2100(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "runtime2100")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Runtime2100 is a free data retrieval call binding the contract method 0x4dfa5ac6.
//
// Solidity: function runtime2100() view returns(bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Runtime2100() (bool, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Runtime2100(&_MoonriverDelegatorCoverOracle.CallOpts)
}

// Runtime2100 is a free data retrieval call binding the contract method 0x4dfa5ac6.
//
// Solidity: function runtime2100() view returns(bool)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) Runtime2100() (bool, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Runtime2100(&_MoonriverDelegatorCoverOracle.CallOpts)
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
// Solidity: function whitelisted(address ) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCaller) Whitelisted(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _MoonriverDelegatorCoverOracle.contract.Call(opts, &out, "whitelisted", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Whitelisted(arg0 common.Address) (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Whitelisted(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// Whitelisted is a free data retrieval call binding the contract method 0xd936547e.
//
// Solidity: function whitelisted(address ) view returns(address)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleCallerSession) Whitelisted(arg0 common.Address) (common.Address, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Whitelisted(&_MoonriverDelegatorCoverOracle.CallOpts, arg0)
}

// CancelDecreaseCover is a paid mutator transaction binding the contract method 0xc8aec3bf.
//
// Solidity: function cancelDecreaseCover(address _member) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) CancelDecreaseCover(opts *bind.TransactOpts, _member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "cancelDecreaseCover", _member)
}

// CancelDecreaseCover is a paid mutator transaction binding the contract method 0xc8aec3bf.
//
// Solidity: function cancelDecreaseCover(address _member) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) CancelDecreaseCover(_member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.CancelDecreaseCover(&_MoonriverDelegatorCoverOracle.TransactOpts, _member)
}

// CancelDecreaseCover is a paid mutator transaction binding the contract method 0xc8aec3bf.
//
// Solidity: function cancelDecreaseCover(address _member) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) CancelDecreaseCover(_member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.CancelDecreaseCover(&_MoonriverDelegatorCoverOracle.TransactOpts, _member)
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
// Solidity: function depositCover(address _member) payable returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) DepositCover(opts *bind.TransactOpts, _member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "depositCover", _member)
}

// DepositCover is a paid mutator transaction binding the contract method 0xfbe4b783.
//
// Solidity: function depositCover(address _member) payable returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) DepositCover(_member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.DepositCover(&_MoonriverDelegatorCoverOracle.TransactOpts, _member)
}

// DepositCover is a paid mutator transaction binding the contract method 0xfbe4b783.
//
// Solidity: function depositCover(address _member) payable returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) DepositCover(_member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.DepositCover(&_MoonriverDelegatorCoverOracle.TransactOpts, _member)
}

// ExecuteDelegationRequest is a paid mutator transaction binding the contract method 0x13a5f393.
//
// Solidity: function executeDelegationRequest(address candidate) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) ExecuteDelegationRequest(opts *bind.TransactOpts, candidate common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "executeDelegationRequest", candidate)
}

// ExecuteDelegationRequest is a paid mutator transaction binding the contract method 0x13a5f393.
//
// Solidity: function executeDelegationRequest(address candidate) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ExecuteDelegationRequest(candidate common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ExecuteDelegationRequest(&_MoonriverDelegatorCoverOracle.TransactOpts, candidate)
}

// ExecuteDelegationRequest is a paid mutator transaction binding the contract method 0x13a5f393.
//
// Solidity: function executeDelegationRequest(address candidate) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) ExecuteDelegationRequest(candidate common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ExecuteDelegationRequest(&_MoonriverDelegatorCoverOracle.TransactOpts, candidate)
}

// ExecuteScheduled is a paid mutator transaction binding the contract method 0x43fc576d.
//
// Solidity: function executeScheduled(address _member) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) ExecuteScheduled(opts *bind.TransactOpts, _member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "executeScheduled", _member)
}

// ExecuteScheduled is a paid mutator transaction binding the contract method 0x43fc576d.
//
// Solidity: function executeScheduled(address _member) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ExecuteScheduled(_member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ExecuteScheduled(&_MoonriverDelegatorCoverOracle.TransactOpts, _member)
}

// ExecuteScheduled is a paid mutator transaction binding the contract method 0x43fc576d.
//
// Solidity: function executeScheduled(address _member) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) ExecuteScheduled(_member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ExecuteScheduled(&_MoonriverDelegatorCoverOracle.TransactOpts, _member)
}

// Initialize is a paid mutator transaction binding the contract method 0x5fbbc222.
//
// Solidity: function initialize(address _auth_manager, address _oracle_master, address _deposit_staking, uint256 _min_deposit, uint256 _max_deposit_total, uint256 _stake_unit_cover, uint256 _min_payout, uint256 _max_era_member_payout, uint128 _eras_between_forced_undelegation) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) Initialize(opts *bind.TransactOpts, _auth_manager common.Address, _oracle_master common.Address, _deposit_staking common.Address, _min_deposit *big.Int, _max_deposit_total *big.Int, _stake_unit_cover *big.Int, _min_payout *big.Int, _max_era_member_payout *big.Int, _eras_between_forced_undelegation *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "initialize", _auth_manager, _oracle_master, _deposit_staking, _min_deposit, _max_deposit_total, _stake_unit_cover, _min_payout, _max_era_member_payout, _eras_between_forced_undelegation)
}

// Initialize is a paid mutator transaction binding the contract method 0x5fbbc222.
//
// Solidity: function initialize(address _auth_manager, address _oracle_master, address _deposit_staking, uint256 _min_deposit, uint256 _max_deposit_total, uint256 _stake_unit_cover, uint256 _min_payout, uint256 _max_era_member_payout, uint128 _eras_between_forced_undelegation) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Initialize(_auth_manager common.Address, _oracle_master common.Address, _deposit_staking common.Address, _min_deposit *big.Int, _max_deposit_total *big.Int, _stake_unit_cover *big.Int, _min_payout *big.Int, _max_era_member_payout *big.Int, _eras_between_forced_undelegation *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Initialize(&_MoonriverDelegatorCoverOracle.TransactOpts, _auth_manager, _oracle_master, _deposit_staking, _min_deposit, _max_deposit_total, _stake_unit_cover, _min_payout, _max_era_member_payout, _eras_between_forced_undelegation)
}

// Initialize is a paid mutator transaction binding the contract method 0x5fbbc222.
//
// Solidity: function initialize(address _auth_manager, address _oracle_master, address _deposit_staking, uint256 _min_deposit, uint256 _max_deposit_total, uint256 _stake_unit_cover, uint256 _min_payout, uint256 _max_era_member_payout, uint128 _eras_between_forced_undelegation) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) Initialize(_auth_manager common.Address, _oracle_master common.Address, _deposit_staking common.Address, _min_deposit *big.Int, _max_deposit_total *big.Int, _stake_unit_cover *big.Int, _min_payout *big.Int, _max_era_member_payout *big.Int, _eras_between_forced_undelegation *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Initialize(&_MoonriverDelegatorCoverOracle.TransactOpts, _auth_manager, _oracle_master, _deposit_staking, _min_deposit, _max_deposit_total, _stake_unit_cover, _min_payout, _max_era_member_payout, _eras_between_forced_undelegation)
}

// InvoiceMembers is a paid mutator transaction binding the contract method 0xc1a62315.
//
// Solidity: function invoiceMembers() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) InvoiceMembers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "invoiceMembers")
}

// InvoiceMembers is a paid mutator transaction binding the contract method 0xc1a62315.
//
// Solidity: function invoiceMembers() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) InvoiceMembers() (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.InvoiceMembers(&_MoonriverDelegatorCoverOracle.TransactOpts)
}

// InvoiceMembers is a paid mutator transaction binding the contract method 0xc1a62315.
//
// Solidity: function invoiceMembers() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) InvoiceMembers() (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.InvoiceMembers(&_MoonriverDelegatorCoverOracle.TransactOpts)
}

// MemberSetCoverTypes is a paid mutator transaction binding the contract method 0x94db6a6c.
//
// Solidity: function memberSetCoverTypes(address _member, bool _noZeroPtsCoverAfterEra, bool _noActiveSetCoverAfterEra) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) MemberSetCoverTypes(opts *bind.TransactOpts, _member common.Address, _noZeroPtsCoverAfterEra bool, _noActiveSetCoverAfterEra bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "memberSetCoverTypes", _member, _noZeroPtsCoverAfterEra, _noActiveSetCoverAfterEra)
}

// MemberSetCoverTypes is a paid mutator transaction binding the contract method 0x94db6a6c.
//
// Solidity: function memberSetCoverTypes(address _member, bool _noZeroPtsCoverAfterEra, bool _noActiveSetCoverAfterEra) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) MemberSetCoverTypes(_member common.Address, _noZeroPtsCoverAfterEra bool, _noActiveSetCoverAfterEra bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MemberSetCoverTypes(&_MoonriverDelegatorCoverOracle.TransactOpts, _member, _noZeroPtsCoverAfterEra, _noActiveSetCoverAfterEra)
}

// MemberSetCoverTypes is a paid mutator transaction binding the contract method 0x94db6a6c.
//
// Solidity: function memberSetCoverTypes(address _member, bool _noZeroPtsCoverAfterEra, bool _noActiveSetCoverAfterEra) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) MemberSetCoverTypes(_member common.Address, _noZeroPtsCoverAfterEra bool, _noActiveSetCoverAfterEra bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MemberSetCoverTypes(&_MoonriverDelegatorCoverOracle.TransactOpts, _member, _noZeroPtsCoverAfterEra, _noActiveSetCoverAfterEra)
}

// MemberSetMaxCoveredDelegation is a paid mutator transaction binding the contract method 0x58b84395.
//
// Solidity: function memberSetMaxCoveredDelegation(address _member, uint256 _max_covered) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) MemberSetMaxCoveredDelegation(opts *bind.TransactOpts, _member common.Address, _max_covered *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "memberSetMaxCoveredDelegation", _member, _max_covered)
}

// MemberSetMaxCoveredDelegation is a paid mutator transaction binding the contract method 0x58b84395.
//
// Solidity: function memberSetMaxCoveredDelegation(address _member, uint256 _max_covered) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) MemberSetMaxCoveredDelegation(_member common.Address, _max_covered *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MemberSetMaxCoveredDelegation(&_MoonriverDelegatorCoverOracle.TransactOpts, _member, _max_covered)
}

// MemberSetMaxCoveredDelegation is a paid mutator transaction binding the contract method 0x58b84395.
//
// Solidity: function memberSetMaxCoveredDelegation(address _member, uint256 _max_covered) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) MemberSetMaxCoveredDelegation(_member common.Address, _max_covered *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.MemberSetMaxCoveredDelegation(&_MoonriverDelegatorCoverOracle.TransactOpts, _member, _max_covered)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _paused) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) Pause(opts *bind.TransactOpts, _paused bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "pause", _paused)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _paused) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Pause(_paused bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Pause(&_MoonriverDelegatorCoverOracle.TransactOpts, _paused)
}

// Pause is a paid mutator transaction binding the contract method 0x02329a29.
//
// Solidity: function pause(bool _paused) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) Pause(_paused bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Pause(&_MoonriverDelegatorCoverOracle.TransactOpts, _paused)
}

// PayOutCover is a paid mutator transaction binding the contract method 0xf6997666.
//
// Solidity: function payOutCover(address[] delegators) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) PayOutCover(opts *bind.TransactOpts, delegators []common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "payOutCover", delegators)
}

// PayOutCover is a paid mutator transaction binding the contract method 0xf6997666.
//
// Solidity: function payOutCover(address[] delegators) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) PayOutCover(delegators []common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.PayOutCover(&_MoonriverDelegatorCoverOracle.TransactOpts, delegators)
}

// PayOutCover is a paid mutator transaction binding the contract method 0xf6997666.
//
// Solidity: function payOutCover(address[] delegators) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) PayOutCover(delegators []common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.PayOutCover(&_MoonriverDelegatorCoverOracle.TransactOpts, delegators)
}

// PushData is a paid mutator transaction binding the contract method 0xee06578f.
//
// Solidity: function pushData(uint128 _eraId, (uint256,uint128,uint128,uint128,uint128,bytes32,uint128,bool,(uint256,uint256,uint128,address,bool,(uint256,address)[])[]) _report, address _oracleCollator) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) PushData(opts *bind.TransactOpts, _eraId *big.Int, _report TypesOracleData, _oracleCollator common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "pushData", _eraId, _report, _oracleCollator)
}

// PushData is a paid mutator transaction binding the contract method 0xee06578f.
//
// Solidity: function pushData(uint128 _eraId, (uint256,uint128,uint128,uint128,uint128,bytes32,uint128,bool,(uint256,uint256,uint128,address,bool,(uint256,address)[])[]) _report, address _oracleCollator) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) PushData(_eraId *big.Int, _report TypesOracleData, _oracleCollator common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.PushData(&_MoonriverDelegatorCoverOracle.TransactOpts, _eraId, _report, _oracleCollator)
}

// PushData is a paid mutator transaction binding the contract method 0xee06578f.
//
// Solidity: function pushData(uint128 _eraId, (uint256,uint128,uint128,uint128,uint128,bytes32,uint128,bool,(uint256,uint256,uint128,address,bool,(uint256,address)[])[]) _report, address _oracleCollator) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) PushData(_eraId *big.Int, _report TypesOracleData, _oracleCollator common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.PushData(&_MoonriverDelegatorCoverOracle.TransactOpts, _eraId, _report, _oracleCollator)
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

// ScheduleDecreaseCover is a paid mutator transaction binding the contract method 0xa131791d.
//
// Solidity: function scheduleDecreaseCover(address _member, uint256 _amount) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) ScheduleDecreaseCover(opts *bind.TransactOpts, _member common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "scheduleDecreaseCover", _member, _amount)
}

// ScheduleDecreaseCover is a paid mutator transaction binding the contract method 0xa131791d.
//
// Solidity: function scheduleDecreaseCover(address _member, uint256 _amount) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) ScheduleDecreaseCover(_member common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ScheduleDecreaseCover(&_MoonriverDelegatorCoverOracle.TransactOpts, _member, _amount)
}

// ScheduleDecreaseCover is a paid mutator transaction binding the contract method 0xa131791d.
//
// Solidity: function scheduleDecreaseCover(address _member, uint256 _amount) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) ScheduleDecreaseCover(_member common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.ScheduleDecreaseCover(&_MoonriverDelegatorCoverOracle.TransactOpts, _member, _amount)
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

// SetDefaultCount is a paid mutator transaction binding the contract method 0x53f832bc.
//
// Solidity: function setDefaultCount(uint128 _defaultCount, address _member) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) SetDefaultCount(opts *bind.TransactOpts, _defaultCount *big.Int, _member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "setDefaultCount", _defaultCount, _member)
}

// SetDefaultCount is a paid mutator transaction binding the contract method 0x53f832bc.
//
// Solidity: function setDefaultCount(uint128 _defaultCount, address _member) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) SetDefaultCount(_defaultCount *big.Int, _member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetDefaultCount(&_MoonriverDelegatorCoverOracle.TransactOpts, _defaultCount, _member)
}

// SetDefaultCount is a paid mutator transaction binding the contract method 0x53f832bc.
//
// Solidity: function setDefaultCount(uint128 _defaultCount, address _member) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) SetDefaultCount(_defaultCount *big.Int, _member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetDefaultCount(&_MoonriverDelegatorCoverOracle.TransactOpts, _defaultCount, _member)
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

// SetErasCovered is a paid mutator transaction binding the contract method 0xf6cc210d.
//
// Solidity: function setErasCovered(uint128 _erasCovered, address member) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) SetErasCovered(opts *bind.TransactOpts, _erasCovered *big.Int, member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "setErasCovered", _erasCovered, member)
}

// SetErasCovered is a paid mutator transaction binding the contract method 0xf6cc210d.
//
// Solidity: function setErasCovered(uint128 _erasCovered, address member) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) SetErasCovered(_erasCovered *big.Int, member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetErasCovered(&_MoonriverDelegatorCoverOracle.TransactOpts, _erasCovered, member)
}

// SetErasCovered is a paid mutator transaction binding the contract method 0xf6cc210d.
//
// Solidity: function setErasCovered(uint128 _erasCovered, address member) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) SetErasCovered(_erasCovered *big.Int, member common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetErasCovered(&_MoonriverDelegatorCoverOracle.TransactOpts, _erasCovered, member)
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

// SetMaxEraMemberPayout is a paid mutator transaction binding the contract method 0x46baab6d.
//
// Solidity: function setMaxEraMemberPayout(uint256 _max_era_member_payout) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) SetMaxEraMemberPayout(opts *bind.TransactOpts, _max_era_member_payout *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "setMaxEraMemberPayout", _max_era_member_payout)
}

// SetMaxEraMemberPayout is a paid mutator transaction binding the contract method 0x46baab6d.
//
// Solidity: function setMaxEraMemberPayout(uint256 _max_era_member_payout) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) SetMaxEraMemberPayout(_max_era_member_payout *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetMaxEraMemberPayout(&_MoonriverDelegatorCoverOracle.TransactOpts, _max_era_member_payout)
}

// SetMaxEraMemberPayout is a paid mutator transaction binding the contract method 0x46baab6d.
//
// Solidity: function setMaxEraMemberPayout(uint256 _max_era_member_payout) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) SetMaxEraMemberPayout(_max_era_member_payout *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetMaxEraMemberPayout(&_MoonriverDelegatorCoverOracle.TransactOpts, _max_era_member_payout)
}

// SetMaxErasCovered is a paid mutator transaction binding the contract method 0x39603d5a.
//
// Solidity: function setMaxErasCovered(uint128 _max_eras_covered) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) SetMaxErasCovered(opts *bind.TransactOpts, _max_eras_covered *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "setMaxErasCovered", _max_eras_covered)
}

// SetMaxErasCovered is a paid mutator transaction binding the contract method 0x39603d5a.
//
// Solidity: function setMaxErasCovered(uint128 _max_eras_covered) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) SetMaxErasCovered(_max_eras_covered *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetMaxErasCovered(&_MoonriverDelegatorCoverOracle.TransactOpts, _max_eras_covered)
}

// SetMaxErasCovered is a paid mutator transaction binding the contract method 0x39603d5a.
//
// Solidity: function setMaxErasCovered(uint128 _max_eras_covered) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) SetMaxErasCovered(_max_eras_covered *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetMaxErasCovered(&_MoonriverDelegatorCoverOracle.TransactOpts, _max_eras_covered)
}

// SetMemberFee is a paid mutator transaction binding the contract method 0x5f6fe4fd.
//
// Solidity: function setMemberFee(uint256 _memberFee) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) SetMemberFee(opts *bind.TransactOpts, _memberFee *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "setMemberFee", _memberFee)
}

// SetMemberFee is a paid mutator transaction binding the contract method 0x5f6fe4fd.
//
// Solidity: function setMemberFee(uint256 _memberFee) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) SetMemberFee(_memberFee *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetMemberFee(&_MoonriverDelegatorCoverOracle.TransactOpts, _memberFee)
}

// SetMemberFee is a paid mutator transaction binding the contract method 0x5f6fe4fd.
//
// Solidity: function setMemberFee(uint256 _memberFee) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) SetMemberFee(_memberFee *big.Int) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetMemberFee(&_MoonriverDelegatorCoverOracle.TransactOpts, _memberFee)
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

// SetNoManualWhitelistingRequired is a paid mutator transaction binding the contract method 0xcbfb3015.
//
// Solidity: function setNoManualWhitelistingRequired(bool _noManualWhitelistingRequired) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) SetNoManualWhitelistingRequired(opts *bind.TransactOpts, _noManualWhitelistingRequired bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "setNoManualWhitelistingRequired", _noManualWhitelistingRequired)
}

// SetNoManualWhitelistingRequired is a paid mutator transaction binding the contract method 0xcbfb3015.
//
// Solidity: function setNoManualWhitelistingRequired(bool _noManualWhitelistingRequired) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) SetNoManualWhitelistingRequired(_noManualWhitelistingRequired bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetNoManualWhitelistingRequired(&_MoonriverDelegatorCoverOracle.TransactOpts, _noManualWhitelistingRequired)
}

// SetNoManualWhitelistingRequired is a paid mutator transaction binding the contract method 0xcbfb3015.
//
// Solidity: function setNoManualWhitelistingRequired(bool _noManualWhitelistingRequired) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) SetNoManualWhitelistingRequired(_noManualWhitelistingRequired bool) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetNoManualWhitelistingRequired(&_MoonriverDelegatorCoverOracle.TransactOpts, _noManualWhitelistingRequired)
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

// SetRuntime2100 is a paid mutator transaction binding the contract method 0xe3faa04e.
//
// Solidity: function setRuntime2100() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) SetRuntime2100(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "setRuntime2100")
}

// SetRuntime2100 is a paid mutator transaction binding the contract method 0xe3faa04e.
//
// Solidity: function setRuntime2100() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) SetRuntime2100() (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetRuntime2100(&_MoonriverDelegatorCoverOracle.TransactOpts)
}

// SetRuntime2100 is a paid mutator transaction binding the contract method 0xe3faa04e.
//
// Solidity: function setRuntime2100() returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) SetRuntime2100() (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.SetRuntime2100(&_MoonriverDelegatorCoverOracle.TransactOpts)
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

// TransferMemberAuth is a paid mutator transaction binding the contract method 0xbc8519bf.
//
// Solidity: function transferMemberAuth(address _member, address proxyAccount) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) TransferMemberAuth(opts *bind.TransactOpts, _member common.Address, proxyAccount common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "transferMemberAuth", _member, proxyAccount)
}

// TransferMemberAuth is a paid mutator transaction binding the contract method 0xbc8519bf.
//
// Solidity: function transferMemberAuth(address _member, address proxyAccount) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) TransferMemberAuth(_member common.Address, proxyAccount common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.TransferMemberAuth(&_MoonriverDelegatorCoverOracle.TransactOpts, _member, proxyAccount)
}

// TransferMemberAuth is a paid mutator transaction binding the contract method 0xbc8519bf.
//
// Solidity: function transferMemberAuth(address _member, address proxyAccount) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) TransferMemberAuth(_member common.Address, proxyAccount common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.TransferMemberAuth(&_MoonriverDelegatorCoverOracle.TransactOpts, _member, proxyAccount)
}

// Whitelist is a paid mutator transaction binding the contract method 0xb092145e.
//
// Solidity: function whitelist(address _member, address proxyAccount) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactor) Whitelist(opts *bind.TransactOpts, _member common.Address, proxyAccount common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.contract.Transact(opts, "whitelist", _member, proxyAccount)
}

// Whitelist is a paid mutator transaction binding the contract method 0xb092145e.
//
// Solidity: function whitelist(address _member, address proxyAccount) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleSession) Whitelist(_member common.Address, proxyAccount common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Whitelist(&_MoonriverDelegatorCoverOracle.TransactOpts, _member, proxyAccount)
}

// Whitelist is a paid mutator transaction binding the contract method 0xb092145e.
//
// Solidity: function whitelist(address _member, address proxyAccount) returns()
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleTransactorSession) Whitelist(_member common.Address, proxyAccount common.Address) (*types.Transaction, error) {
	return _MoonriverDelegatorCoverOracle.Contract.Whitelist(&_MoonriverDelegatorCoverOracle.TransactOpts, _member, proxyAccount)
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
	EraId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreaseCoverScheduledEvent is a free log retrieval operation binding the contract event 0xcc71ed28070c44cc6c961d9bc3529576b26945a791dfc8911ad08def5c212d99.
//
// Solidity: event DecreaseCoverScheduledEvent(address member, uint256 amount, uint128 eraId)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterDecreaseCoverScheduledEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleDecreaseCoverScheduledEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "DecreaseCoverScheduledEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleDecreaseCoverScheduledEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "DecreaseCoverScheduledEvent", logs: logs, sub: sub}, nil
}

// WatchDecreaseCoverScheduledEvent is a free log subscription operation binding the contract event 0xcc71ed28070c44cc6c961d9bc3529576b26945a791dfc8911ad08def5c212d99.
//
// Solidity: event DecreaseCoverScheduledEvent(address member, uint256 amount, uint128 eraId)
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

// ParseDecreaseCoverScheduledEvent is a log parse operation binding the contract event 0xcc71ed28070c44cc6c961d9bc3529576b26945a791dfc8911ad08def5c212d99.
//
// Solidity: event DecreaseCoverScheduledEvent(address member, uint256 amount, uint128 eraId)
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
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegatorNotPaidEvent is a free log retrieval operation binding the contract event 0x672619ed690cc62ad4832c1bdca2c8f5e9927c7d419c2383ce6cea71edeac8bb.
//
// Solidity: event DelegatorNotPaidEvent(address delegator, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterDelegatorNotPaidEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleDelegatorNotPaidEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "DelegatorNotPaidEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleDelegatorNotPaidEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "DelegatorNotPaidEvent", logs: logs, sub: sub}, nil
}

// WatchDelegatorNotPaidEvent is a free log subscription operation binding the contract event 0x672619ed690cc62ad4832c1bdca2c8f5e9927c7d419c2383ce6cea71edeac8bb.
//
// Solidity: event DelegatorNotPaidEvent(address delegator, uint256 amount)
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

// ParseDelegatorNotPaidEvent is a log parse operation binding the contract event 0x672619ed690cc62ad4832c1bdca2c8f5e9927c7d419c2383ce6cea71edeac8bb.
//
// Solidity: event DelegatorNotPaidEvent(address delegator, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseDelegatorNotPaidEvent(log types.Log) (*MoonriverDelegatorCoverOracleDelegatorNotPaidEvent, error) {
	event := new(MoonriverDelegatorCoverOracleDelegatorNotPaidEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "DelegatorNotPaidEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleDelegatorPayoutLessThanMinEventIterator is returned from FilterDelegatorPayoutLessThanMinEvent and is used to iterate over the raw logs and unpacked data for DelegatorPayoutLessThanMinEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleDelegatorPayoutLessThanMinEventIterator struct {
	Event *MoonriverDelegatorCoverOracleDelegatorPayoutLessThanMinEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleDelegatorPayoutLessThanMinEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleDelegatorPayoutLessThanMinEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleDelegatorPayoutLessThanMinEvent)
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
func (it *MoonriverDelegatorCoverOracleDelegatorPayoutLessThanMinEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleDelegatorPayoutLessThanMinEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleDelegatorPayoutLessThanMinEvent represents a DelegatorPayoutLessThanMinEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleDelegatorPayoutLessThanMinEvent struct {
	Delegator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegatorPayoutLessThanMinEvent is a free log retrieval operation binding the contract event 0xeb44d2cde8a0e5f492334395012c62f89aaa4098db03cd671073b2bfff4b0df2.
//
// Solidity: event DelegatorPayoutLessThanMinEvent(address delegator)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterDelegatorPayoutLessThanMinEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleDelegatorPayoutLessThanMinEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "DelegatorPayoutLessThanMinEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleDelegatorPayoutLessThanMinEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "DelegatorPayoutLessThanMinEvent", logs: logs, sub: sub}, nil
}

// WatchDelegatorPayoutLessThanMinEvent is a free log subscription operation binding the contract event 0xeb44d2cde8a0e5f492334395012c62f89aaa4098db03cd671073b2bfff4b0df2.
//
// Solidity: event DelegatorPayoutLessThanMinEvent(address delegator)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchDelegatorPayoutLessThanMinEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleDelegatorPayoutLessThanMinEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "DelegatorPayoutLessThanMinEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleDelegatorPayoutLessThanMinEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "DelegatorPayoutLessThanMinEvent", log); err != nil {
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

// ParseDelegatorPayoutLessThanMinEvent is a log parse operation binding the contract event 0xeb44d2cde8a0e5f492334395012c62f89aaa4098db03cd671073b2bfff4b0df2.
//
// Solidity: event DelegatorPayoutLessThanMinEvent(address delegator)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseDelegatorPayoutLessThanMinEvent(log types.Log) (*MoonriverDelegatorCoverOracleDelegatorPayoutLessThanMinEvent, error) {
	event := new(MoonriverDelegatorCoverOracleDelegatorPayoutLessThanMinEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "DelegatorPayoutLessThanMinEvent", log); err != nil {
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

// MoonriverDelegatorCoverOracleMemberActiveStatusChangedIterator is returned from FilterMemberActiveStatusChanged and is used to iterate over the raw logs and unpacked data for MemberActiveStatusChanged events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberActiveStatusChangedIterator struct {
	Event *MoonriverDelegatorCoverOracleMemberActiveStatusChanged // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleMemberActiveStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleMemberActiveStatusChanged)
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
		it.Event = new(MoonriverDelegatorCoverOracleMemberActiveStatusChanged)
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
func (it *MoonriverDelegatorCoverOracleMemberActiveStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleMemberActiveStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleMemberActiveStatusChanged represents a MemberActiveStatusChanged event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberActiveStatusChanged struct {
	Member common.Address
	Active bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMemberActiveStatusChanged is a free log retrieval operation binding the contract event 0xdc7a9bc5567ca66968b8bb475cba88e830154e12e6541280a5519c87eab15f47.
//
// Solidity: event MemberActiveStatusChanged(address member, bool active)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterMemberActiveStatusChanged(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleMemberActiveStatusChangedIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "MemberActiveStatusChanged")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleMemberActiveStatusChangedIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "MemberActiveStatusChanged", logs: logs, sub: sub}, nil
}

// WatchMemberActiveStatusChanged is a free log subscription operation binding the contract event 0xdc7a9bc5567ca66968b8bb475cba88e830154e12e6541280a5519c87eab15f47.
//
// Solidity: event MemberActiveStatusChanged(address member, bool active)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchMemberActiveStatusChanged(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleMemberActiveStatusChanged) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "MemberActiveStatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleMemberActiveStatusChanged)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberActiveStatusChanged", log); err != nil {
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

// ParseMemberActiveStatusChanged is a log parse operation binding the contract event 0xdc7a9bc5567ca66968b8bb475cba88e830154e12e6541280a5519c87eab15f47.
//
// Solidity: event MemberActiveStatusChanged(address member, bool active)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseMemberActiveStatusChanged(log types.Log) (*MoonriverDelegatorCoverOracleMemberActiveStatusChanged, error) {
	event := new(MoonriverDelegatorCoverOracleMemberActiveStatusChanged)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberActiveStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleMemberDefaultEventIterator is returned from FilterMemberDefaultEvent and is used to iterate over the raw logs and unpacked data for MemberDefaultEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberDefaultEventIterator struct {
	Event *MoonriverDelegatorCoverOracleMemberDefaultEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleMemberDefaultEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleMemberDefaultEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleMemberDefaultEvent)
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
func (it *MoonriverDelegatorCoverOracleMemberDefaultEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleMemberDefaultEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleMemberDefaultEvent represents a MemberDefaultEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberDefaultEvent struct {
	Member common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMemberDefaultEvent is a free log retrieval operation binding the contract event 0x83e5adb5e507b374e44f493d7c3894d4c36b7df3d426e4370d7edf300bf7eead.
//
// Solidity: event MemberDefaultEvent(address member)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterMemberDefaultEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleMemberDefaultEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "MemberDefaultEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleMemberDefaultEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "MemberDefaultEvent", logs: logs, sub: sub}, nil
}

// WatchMemberDefaultEvent is a free log subscription operation binding the contract event 0x83e5adb5e507b374e44f493d7c3894d4c36b7df3d426e4370d7edf300bf7eead.
//
// Solidity: event MemberDefaultEvent(address member)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchMemberDefaultEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleMemberDefaultEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "MemberDefaultEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleMemberDefaultEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberDefaultEvent", log); err != nil {
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

// ParseMemberDefaultEvent is a log parse operation binding the contract event 0x83e5adb5e507b374e44f493d7c3894d4c36b7df3d426e4370d7edf300bf7eead.
//
// Solidity: event MemberDefaultEvent(address member)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseMemberDefaultEvent(log types.Log) (*MoonriverDelegatorCoverOracleMemberDefaultEvent, error) {
	event := new(MoonriverDelegatorCoverOracleMemberDefaultEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberDefaultEvent", log); err != nil {
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

// MoonriverDelegatorCoverOracleMemberInvoicedEventIterator is returned from FilterMemberInvoicedEvent and is used to iterate over the raw logs and unpacked data for MemberInvoicedEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberInvoicedEventIterator struct {
	Event *MoonriverDelegatorCoverOracleMemberInvoicedEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleMemberInvoicedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleMemberInvoicedEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleMemberInvoicedEvent)
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
func (it *MoonriverDelegatorCoverOracleMemberInvoicedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleMemberInvoicedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleMemberInvoicedEvent represents a MemberInvoicedEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberInvoicedEvent struct {
	Member common.Address
	Amount *big.Int
	EraId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMemberInvoicedEvent is a free log retrieval operation binding the contract event 0x76e080a95b1e02d44cc11c2f38cf031d5d6a5c764a760db58a9cfdbb11e05c41.
//
// Solidity: event MemberInvoicedEvent(address member, uint256 amount, uint128 eraId)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterMemberInvoicedEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleMemberInvoicedEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "MemberInvoicedEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleMemberInvoicedEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "MemberInvoicedEvent", logs: logs, sub: sub}, nil
}

// WatchMemberInvoicedEvent is a free log subscription operation binding the contract event 0x76e080a95b1e02d44cc11c2f38cf031d5d6a5c764a760db58a9cfdbb11e05c41.
//
// Solidity: event MemberInvoicedEvent(address member, uint256 amount, uint128 eraId)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchMemberInvoicedEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleMemberInvoicedEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "MemberInvoicedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleMemberInvoicedEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberInvoicedEvent", log); err != nil {
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

// ParseMemberInvoicedEvent is a log parse operation binding the contract event 0x76e080a95b1e02d44cc11c2f38cf031d5d6a5c764a760db58a9cfdbb11e05c41.
//
// Solidity: event MemberInvoicedEvent(address member, uint256 amount, uint128 eraId)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseMemberInvoicedEvent(log types.Log) (*MoonriverDelegatorCoverOracleMemberInvoicedEvent, error) {
	event := new(MoonriverDelegatorCoverOracleMemberInvoicedEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberInvoicedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleMemberNotInActiveSetEventIterator is returned from FilterMemberNotInActiveSetEvent and is used to iterate over the raw logs and unpacked data for MemberNotInActiveSetEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberNotInActiveSetEventIterator struct {
	Event *MoonriverDelegatorCoverOracleMemberNotInActiveSetEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleMemberNotInActiveSetEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleMemberNotInActiveSetEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleMemberNotInActiveSetEvent)
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
func (it *MoonriverDelegatorCoverOracleMemberNotInActiveSetEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleMemberNotInActiveSetEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleMemberNotInActiveSetEvent represents a MemberNotInActiveSetEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberNotInActiveSetEvent struct {
	Member common.Address
	EraId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMemberNotInActiveSetEvent is a free log retrieval operation binding the contract event 0x58080af508644502de96ea2dba06534ceb1f0aa8299ae575fdc0c16d34003d27.
//
// Solidity: event MemberNotInActiveSetEvent(address member, uint128 eraId)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterMemberNotInActiveSetEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleMemberNotInActiveSetEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "MemberNotInActiveSetEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleMemberNotInActiveSetEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "MemberNotInActiveSetEvent", logs: logs, sub: sub}, nil
}

// WatchMemberNotInActiveSetEvent is a free log subscription operation binding the contract event 0x58080af508644502de96ea2dba06534ceb1f0aa8299ae575fdc0c16d34003d27.
//
// Solidity: event MemberNotInActiveSetEvent(address member, uint128 eraId)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchMemberNotInActiveSetEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleMemberNotInActiveSetEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "MemberNotInActiveSetEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleMemberNotInActiveSetEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberNotInActiveSetEvent", log); err != nil {
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

// ParseMemberNotInActiveSetEvent is a log parse operation binding the contract event 0x58080af508644502de96ea2dba06534ceb1f0aa8299ae575fdc0c16d34003d27.
//
// Solidity: event MemberNotInActiveSetEvent(address member, uint128 eraId)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseMemberNotInActiveSetEvent(log types.Log) (*MoonriverDelegatorCoverOracleMemberNotInActiveSetEvent, error) {
	event := new(MoonriverDelegatorCoverOracleMemberNotInActiveSetEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberNotInActiveSetEvent", log); err != nil {
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

// MoonriverDelegatorCoverOracleMemberSetCoverTypesEventIterator is returned from FilterMemberSetCoverTypesEvent and is used to iterate over the raw logs and unpacked data for MemberSetCoverTypesEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberSetCoverTypesEventIterator struct {
	Event *MoonriverDelegatorCoverOracleMemberSetCoverTypesEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleMemberSetCoverTypesEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleMemberSetCoverTypesEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleMemberSetCoverTypesEvent)
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
func (it *MoonriverDelegatorCoverOracleMemberSetCoverTypesEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleMemberSetCoverTypesEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleMemberSetCoverTypesEvent represents a MemberSetCoverTypesEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberSetCoverTypesEvent struct {
	Member                   common.Address
	NoZeroPtsCoverAfterEra   bool
	NoActiveSetCoverAfterEra bool
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterMemberSetCoverTypesEvent is a free log retrieval operation binding the contract event 0x64ff8b21509b656bfbb1face101a20453b99ce6df8679f1f16e075e0fb5e5d44.
//
// Solidity: event MemberSetCoverTypesEvent(address member, bool noZeroPtsCoverAfterEra, bool noActiveSetCoverAfterEra)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterMemberSetCoverTypesEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleMemberSetCoverTypesEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "MemberSetCoverTypesEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleMemberSetCoverTypesEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "MemberSetCoverTypesEvent", logs: logs, sub: sub}, nil
}

// WatchMemberSetCoverTypesEvent is a free log subscription operation binding the contract event 0x64ff8b21509b656bfbb1face101a20453b99ce6df8679f1f16e075e0fb5e5d44.
//
// Solidity: event MemberSetCoverTypesEvent(address member, bool noZeroPtsCoverAfterEra, bool noActiveSetCoverAfterEra)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchMemberSetCoverTypesEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleMemberSetCoverTypesEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "MemberSetCoverTypesEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleMemberSetCoverTypesEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberSetCoverTypesEvent", log); err != nil {
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

// ParseMemberSetCoverTypesEvent is a log parse operation binding the contract event 0x64ff8b21509b656bfbb1face101a20453b99ce6df8679f1f16e075e0fb5e5d44.
//
// Solidity: event MemberSetCoverTypesEvent(address member, bool noZeroPtsCoverAfterEra, bool noActiveSetCoverAfterEra)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseMemberSetCoverTypesEvent(log types.Log) (*MoonriverDelegatorCoverOracleMemberSetCoverTypesEvent, error) {
	event := new(MoonriverDelegatorCoverOracleMemberSetCoverTypesEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberSetCoverTypesEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleMemberSetMaxCoveredDelegationEventIterator is returned from FilterMemberSetMaxCoveredDelegationEvent and is used to iterate over the raw logs and unpacked data for MemberSetMaxCoveredDelegationEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberSetMaxCoveredDelegationEventIterator struct {
	Event *MoonriverDelegatorCoverOracleMemberSetMaxCoveredDelegationEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleMemberSetMaxCoveredDelegationEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleMemberSetMaxCoveredDelegationEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleMemberSetMaxCoveredDelegationEvent)
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
func (it *MoonriverDelegatorCoverOracleMemberSetMaxCoveredDelegationEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleMemberSetMaxCoveredDelegationEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleMemberSetMaxCoveredDelegationEvent represents a MemberSetMaxCoveredDelegationEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleMemberSetMaxCoveredDelegationEvent struct {
	Member common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMemberSetMaxCoveredDelegationEvent is a free log retrieval operation binding the contract event 0xc03853b905a3fab5b1be1eca777cdfec528edd4a33e4e44613d67ae29995044f.
//
// Solidity: event MemberSetMaxCoveredDelegationEvent(address member, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterMemberSetMaxCoveredDelegationEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleMemberSetMaxCoveredDelegationEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "MemberSetMaxCoveredDelegationEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleMemberSetMaxCoveredDelegationEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "MemberSetMaxCoveredDelegationEvent", logs: logs, sub: sub}, nil
}

// WatchMemberSetMaxCoveredDelegationEvent is a free log subscription operation binding the contract event 0xc03853b905a3fab5b1be1eca777cdfec528edd4a33e4e44613d67ae29995044f.
//
// Solidity: event MemberSetMaxCoveredDelegationEvent(address member, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchMemberSetMaxCoveredDelegationEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleMemberSetMaxCoveredDelegationEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "MemberSetMaxCoveredDelegationEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleMemberSetMaxCoveredDelegationEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberSetMaxCoveredDelegationEvent", log); err != nil {
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

// ParseMemberSetMaxCoveredDelegationEvent is a log parse operation binding the contract event 0xc03853b905a3fab5b1be1eca777cdfec528edd4a33e4e44613d67ae29995044f.
//
// Solidity: event MemberSetMaxCoveredDelegationEvent(address member, uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseMemberSetMaxCoveredDelegationEvent(log types.Log) (*MoonriverDelegatorCoverOracleMemberSetMaxCoveredDelegationEvent, error) {
	event := new(MoonriverDelegatorCoverOracleMemberSetMaxCoveredDelegationEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "MemberSetMaxCoveredDelegationEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleOraclePaidEventIterator is returned from FilterOraclePaidEvent and is used to iterate over the raw logs and unpacked data for OraclePaidEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleOraclePaidEventIterator struct {
	Event *MoonriverDelegatorCoverOracleOraclePaidEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleOraclePaidEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleOraclePaidEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleOraclePaidEvent)
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
func (it *MoonriverDelegatorCoverOracleOraclePaidEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleOraclePaidEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleOraclePaidEvent represents a OraclePaidEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleOraclePaidEvent struct {
	Member common.Address
	Amount *big.Int
	EraId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOraclePaidEvent is a free log retrieval operation binding the contract event 0xf1e225ae1c6ca105ef73d6b12f4af0e6895c2c14ccac1c2796f8bf6c773478a6.
//
// Solidity: event OraclePaidEvent(address member, uint256 amount, uint128 eraId)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterOraclePaidEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleOraclePaidEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "OraclePaidEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleOraclePaidEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "OraclePaidEvent", logs: logs, sub: sub}, nil
}

// WatchOraclePaidEvent is a free log subscription operation binding the contract event 0xf1e225ae1c6ca105ef73d6b12f4af0e6895c2c14ccac1c2796f8bf6c773478a6.
//
// Solidity: event OraclePaidEvent(address member, uint256 amount, uint128 eraId)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchOraclePaidEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleOraclePaidEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "OraclePaidEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleOraclePaidEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "OraclePaidEvent", log); err != nil {
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

// ParseOraclePaidEvent is a log parse operation binding the contract event 0xf1e225ae1c6ca105ef73d6b12f4af0e6895c2c14ccac1c2796f8bf6c773478a6.
//
// Solidity: event OraclePaidEvent(address member, uint256 amount, uint128 eraId)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseOraclePaidEvent(log types.Log) (*MoonriverDelegatorCoverOracleOraclePaidEvent, error) {
	event := new(MoonriverDelegatorCoverOracleOraclePaidEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "OraclePaidEvent", log); err != nil {
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

// MoonriverDelegatorCoverOracleSetErasBetweenForcedUndelegationsEventIterator is returned from FilterSetErasBetweenForcedUndelegationsEvent and is used to iterate over the raw logs and unpacked data for SetErasBetweenForcedUndelegationsEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleSetErasBetweenForcedUndelegationsEventIterator struct {
	Event *MoonriverDelegatorCoverOracleSetErasBetweenForcedUndelegationsEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleSetErasBetweenForcedUndelegationsEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleSetErasBetweenForcedUndelegationsEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleSetErasBetweenForcedUndelegationsEvent)
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
func (it *MoonriverDelegatorCoverOracleSetErasBetweenForcedUndelegationsEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleSetErasBetweenForcedUndelegationsEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleSetErasBetweenForcedUndelegationsEvent represents a SetErasBetweenForcedUndelegationsEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleSetErasBetweenForcedUndelegationsEvent struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetErasBetweenForcedUndelegationsEvent is a free log retrieval operation binding the contract event 0x174cd5ac781094d736baee7a962c3c4822a1d65aac0027c76e395f7bd2e44742.
//
// Solidity: event SetErasBetweenForcedUndelegationsEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterSetErasBetweenForcedUndelegationsEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleSetErasBetweenForcedUndelegationsEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "SetErasBetweenForcedUndelegationsEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleSetErasBetweenForcedUndelegationsEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "SetErasBetweenForcedUndelegationsEvent", logs: logs, sub: sub}, nil
}

// WatchSetErasBetweenForcedUndelegationsEvent is a free log subscription operation binding the contract event 0x174cd5ac781094d736baee7a962c3c4822a1d65aac0027c76e395f7bd2e44742.
//
// Solidity: event SetErasBetweenForcedUndelegationsEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchSetErasBetweenForcedUndelegationsEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleSetErasBetweenForcedUndelegationsEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "SetErasBetweenForcedUndelegationsEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleSetErasBetweenForcedUndelegationsEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "SetErasBetweenForcedUndelegationsEvent", log); err != nil {
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

// ParseSetErasBetweenForcedUndelegationsEvent is a log parse operation binding the contract event 0x174cd5ac781094d736baee7a962c3c4822a1d65aac0027c76e395f7bd2e44742.
//
// Solidity: event SetErasBetweenForcedUndelegationsEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseSetErasBetweenForcedUndelegationsEvent(log types.Log) (*MoonriverDelegatorCoverOracleSetErasBetweenForcedUndelegationsEvent, error) {
	event := new(MoonriverDelegatorCoverOracleSetErasBetweenForcedUndelegationsEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "SetErasBetweenForcedUndelegationsEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleSetMaxEraMemberPayoutEventIterator is returned from FilterSetMaxEraMemberPayoutEvent and is used to iterate over the raw logs and unpacked data for SetMaxEraMemberPayoutEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleSetMaxEraMemberPayoutEventIterator struct {
	Event *MoonriverDelegatorCoverOracleSetMaxEraMemberPayoutEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleSetMaxEraMemberPayoutEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleSetMaxEraMemberPayoutEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleSetMaxEraMemberPayoutEvent)
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
func (it *MoonriverDelegatorCoverOracleSetMaxEraMemberPayoutEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleSetMaxEraMemberPayoutEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleSetMaxEraMemberPayoutEvent represents a SetMaxEraMemberPayoutEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleSetMaxEraMemberPayoutEvent struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMaxEraMemberPayoutEvent is a free log retrieval operation binding the contract event 0x58e9601abd3b183c1e27f1e9e7594fd13c2581af0505517712b7c6814f4dcd94.
//
// Solidity: event SetMaxEraMemberPayoutEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterSetMaxEraMemberPayoutEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleSetMaxEraMemberPayoutEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "SetMaxEraMemberPayoutEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleSetMaxEraMemberPayoutEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "SetMaxEraMemberPayoutEvent", logs: logs, sub: sub}, nil
}

// WatchSetMaxEraMemberPayoutEvent is a free log subscription operation binding the contract event 0x58e9601abd3b183c1e27f1e9e7594fd13c2581af0505517712b7c6814f4dcd94.
//
// Solidity: event SetMaxEraMemberPayoutEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchSetMaxEraMemberPayoutEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleSetMaxEraMemberPayoutEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "SetMaxEraMemberPayoutEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleSetMaxEraMemberPayoutEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "SetMaxEraMemberPayoutEvent", log); err != nil {
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

// ParseSetMaxEraMemberPayoutEvent is a log parse operation binding the contract event 0x58e9601abd3b183c1e27f1e9e7594fd13c2581af0505517712b7c6814f4dcd94.
//
// Solidity: event SetMaxEraMemberPayoutEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseSetMaxEraMemberPayoutEvent(log types.Log) (*MoonriverDelegatorCoverOracleSetMaxEraMemberPayoutEvent, error) {
	event := new(MoonriverDelegatorCoverOracleSetMaxEraMemberPayoutEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "SetMaxEraMemberPayoutEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleSetMaxErasCoveredEventIterator is returned from FilterSetMaxErasCoveredEvent and is used to iterate over the raw logs and unpacked data for SetMaxErasCoveredEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleSetMaxErasCoveredEventIterator struct {
	Event *MoonriverDelegatorCoverOracleSetMaxErasCoveredEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleSetMaxErasCoveredEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleSetMaxErasCoveredEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleSetMaxErasCoveredEvent)
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
func (it *MoonriverDelegatorCoverOracleSetMaxErasCoveredEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleSetMaxErasCoveredEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleSetMaxErasCoveredEvent represents a SetMaxErasCoveredEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleSetMaxErasCoveredEvent struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMaxErasCoveredEvent is a free log retrieval operation binding the contract event 0x9f6e19ec1579c3e4d73721acdee391987e26a55ffb2f83fa8321aff53a54d986.
//
// Solidity: event SetMaxErasCoveredEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterSetMaxErasCoveredEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleSetMaxErasCoveredEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "SetMaxErasCoveredEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleSetMaxErasCoveredEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "SetMaxErasCoveredEvent", logs: logs, sub: sub}, nil
}

// WatchSetMaxErasCoveredEvent is a free log subscription operation binding the contract event 0x9f6e19ec1579c3e4d73721acdee391987e26a55ffb2f83fa8321aff53a54d986.
//
// Solidity: event SetMaxErasCoveredEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchSetMaxErasCoveredEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleSetMaxErasCoveredEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "SetMaxErasCoveredEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleSetMaxErasCoveredEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "SetMaxErasCoveredEvent", log); err != nil {
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

// ParseSetMaxErasCoveredEvent is a log parse operation binding the contract event 0x9f6e19ec1579c3e4d73721acdee391987e26a55ffb2f83fa8321aff53a54d986.
//
// Solidity: event SetMaxErasCoveredEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseSetMaxErasCoveredEvent(log types.Log) (*MoonriverDelegatorCoverOracleSetMaxErasCoveredEvent, error) {
	event := new(MoonriverDelegatorCoverOracleSetMaxErasCoveredEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "SetMaxErasCoveredEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleSetMemberFeeEventIterator is returned from FilterSetMemberFeeEvent and is used to iterate over the raw logs and unpacked data for SetMemberFeeEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleSetMemberFeeEventIterator struct {
	Event *MoonriverDelegatorCoverOracleSetMemberFeeEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleSetMemberFeeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleSetMemberFeeEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleSetMemberFeeEvent)
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
func (it *MoonriverDelegatorCoverOracleSetMemberFeeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleSetMemberFeeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleSetMemberFeeEvent represents a SetMemberFeeEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleSetMemberFeeEvent struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMemberFeeEvent is a free log retrieval operation binding the contract event 0x853f7e26cee23aeebfb0247a2bb47f82591b1c9abdf199717127f6324dacf75c.
//
// Solidity: event SetMemberFeeEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterSetMemberFeeEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleSetMemberFeeEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "SetMemberFeeEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleSetMemberFeeEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "SetMemberFeeEvent", logs: logs, sub: sub}, nil
}

// WatchSetMemberFeeEvent is a free log subscription operation binding the contract event 0x853f7e26cee23aeebfb0247a2bb47f82591b1c9abdf199717127f6324dacf75c.
//
// Solidity: event SetMemberFeeEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchSetMemberFeeEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleSetMemberFeeEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "SetMemberFeeEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleSetMemberFeeEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "SetMemberFeeEvent", log); err != nil {
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

// ParseSetMemberFeeEvent is a log parse operation binding the contract event 0x853f7e26cee23aeebfb0247a2bb47f82591b1c9abdf199717127f6324dacf75c.
//
// Solidity: event SetMemberFeeEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseSetMemberFeeEvent(log types.Log) (*MoonriverDelegatorCoverOracleSetMemberFeeEvent, error) {
	event := new(MoonriverDelegatorCoverOracleSetMemberFeeEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "SetMemberFeeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleSetMinPayoutEventIterator is returned from FilterSetMinPayoutEvent and is used to iterate over the raw logs and unpacked data for SetMinPayoutEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleSetMinPayoutEventIterator struct {
	Event *MoonriverDelegatorCoverOracleSetMinPayoutEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleSetMinPayoutEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleSetMinPayoutEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleSetMinPayoutEvent)
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
func (it *MoonriverDelegatorCoverOracleSetMinPayoutEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleSetMinPayoutEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleSetMinPayoutEvent represents a SetMinPayoutEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleSetMinPayoutEvent struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMinPayoutEvent is a free log retrieval operation binding the contract event 0x9da8207d0c438296097f0bbd97936b3736a136533b318c8ebd3b24b38d85170d.
//
// Solidity: event SetMinPayoutEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterSetMinPayoutEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleSetMinPayoutEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "SetMinPayoutEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleSetMinPayoutEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "SetMinPayoutEvent", logs: logs, sub: sub}, nil
}

// WatchSetMinPayoutEvent is a free log subscription operation binding the contract event 0x9da8207d0c438296097f0bbd97936b3736a136533b318c8ebd3b24b38d85170d.
//
// Solidity: event SetMinPayoutEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchSetMinPayoutEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleSetMinPayoutEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "SetMinPayoutEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleSetMinPayoutEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "SetMinPayoutEvent", log); err != nil {
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

// ParseSetMinPayoutEvent is a log parse operation binding the contract event 0x9da8207d0c438296097f0bbd97936b3736a136533b318c8ebd3b24b38d85170d.
//
// Solidity: event SetMinPayoutEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseSetMinPayoutEvent(log types.Log) (*MoonriverDelegatorCoverOracleSetMinPayoutEvent, error) {
	event := new(MoonriverDelegatorCoverOracleSetMinPayoutEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "SetMinPayoutEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleSetNoManualWhitelistingRequiredEventIterator is returned from FilterSetNoManualWhitelistingRequiredEvent and is used to iterate over the raw logs and unpacked data for SetNoManualWhitelistingRequiredEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleSetNoManualWhitelistingRequiredEventIterator struct {
	Event *MoonriverDelegatorCoverOracleSetNoManualWhitelistingRequiredEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleSetNoManualWhitelistingRequiredEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleSetNoManualWhitelistingRequiredEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleSetNoManualWhitelistingRequiredEvent)
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
func (it *MoonriverDelegatorCoverOracleSetNoManualWhitelistingRequiredEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleSetNoManualWhitelistingRequiredEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleSetNoManualWhitelistingRequiredEvent represents a SetNoManualWhitelistingRequiredEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleSetNoManualWhitelistingRequiredEvent struct {
	NotRequired bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetNoManualWhitelistingRequiredEvent is a free log retrieval operation binding the contract event 0x4e54d84ee1beafdff538cbbe35bf50d461fa812c3c57cff010490c3b56cfe7f0.
//
// Solidity: event SetNoManualWhitelistingRequiredEvent(bool notRequired)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterSetNoManualWhitelistingRequiredEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleSetNoManualWhitelistingRequiredEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "SetNoManualWhitelistingRequiredEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleSetNoManualWhitelistingRequiredEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "SetNoManualWhitelistingRequiredEvent", logs: logs, sub: sub}, nil
}

// WatchSetNoManualWhitelistingRequiredEvent is a free log subscription operation binding the contract event 0x4e54d84ee1beafdff538cbbe35bf50d461fa812c3c57cff010490c3b56cfe7f0.
//
// Solidity: event SetNoManualWhitelistingRequiredEvent(bool notRequired)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchSetNoManualWhitelistingRequiredEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleSetNoManualWhitelistingRequiredEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "SetNoManualWhitelistingRequiredEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleSetNoManualWhitelistingRequiredEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "SetNoManualWhitelistingRequiredEvent", log); err != nil {
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

// ParseSetNoManualWhitelistingRequiredEvent is a log parse operation binding the contract event 0x4e54d84ee1beafdff538cbbe35bf50d461fa812c3c57cff010490c3b56cfe7f0.
//
// Solidity: event SetNoManualWhitelistingRequiredEvent(bool notRequired)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseSetNoManualWhitelistingRequiredEvent(log types.Log) (*MoonriverDelegatorCoverOracleSetNoManualWhitelistingRequiredEvent, error) {
	event := new(MoonriverDelegatorCoverOracleSetNoManualWhitelistingRequiredEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "SetNoManualWhitelistingRequiredEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleSetRefundOracleGasPriceEventIterator is returned from FilterSetRefundOracleGasPriceEvent and is used to iterate over the raw logs and unpacked data for SetRefundOracleGasPriceEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleSetRefundOracleGasPriceEventIterator struct {
	Event *MoonriverDelegatorCoverOracleSetRefundOracleGasPriceEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleSetRefundOracleGasPriceEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleSetRefundOracleGasPriceEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleSetRefundOracleGasPriceEvent)
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
func (it *MoonriverDelegatorCoverOracleSetRefundOracleGasPriceEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleSetRefundOracleGasPriceEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleSetRefundOracleGasPriceEvent represents a SetRefundOracleGasPriceEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleSetRefundOracleGasPriceEvent struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetRefundOracleGasPriceEvent is a free log retrieval operation binding the contract event 0xe202eb3e92c6098e8c56494d605f469e5bde19c0fcce47f7f42a5d43a2454467.
//
// Solidity: event SetRefundOracleGasPriceEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterSetRefundOracleGasPriceEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleSetRefundOracleGasPriceEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "SetRefundOracleGasPriceEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleSetRefundOracleGasPriceEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "SetRefundOracleGasPriceEvent", logs: logs, sub: sub}, nil
}

// WatchSetRefundOracleGasPriceEvent is a free log subscription operation binding the contract event 0xe202eb3e92c6098e8c56494d605f469e5bde19c0fcce47f7f42a5d43a2454467.
//
// Solidity: event SetRefundOracleGasPriceEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchSetRefundOracleGasPriceEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleSetRefundOracleGasPriceEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "SetRefundOracleGasPriceEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleSetRefundOracleGasPriceEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "SetRefundOracleGasPriceEvent", log); err != nil {
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

// ParseSetRefundOracleGasPriceEvent is a log parse operation binding the contract event 0xe202eb3e92c6098e8c56494d605f469e5bde19c0fcce47f7f42a5d43a2454467.
//
// Solidity: event SetRefundOracleGasPriceEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseSetRefundOracleGasPriceEvent(log types.Log) (*MoonriverDelegatorCoverOracleSetRefundOracleGasPriceEvent, error) {
	event := new(MoonriverDelegatorCoverOracleSetRefundOracleGasPriceEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "SetRefundOracleGasPriceEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleSetStakeUnitCoverEventIterator is returned from FilterSetStakeUnitCoverEvent and is used to iterate over the raw logs and unpacked data for SetStakeUnitCoverEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleSetStakeUnitCoverEventIterator struct {
	Event *MoonriverDelegatorCoverOracleSetStakeUnitCoverEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleSetStakeUnitCoverEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleSetStakeUnitCoverEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleSetStakeUnitCoverEvent)
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
func (it *MoonriverDelegatorCoverOracleSetStakeUnitCoverEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleSetStakeUnitCoverEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleSetStakeUnitCoverEvent represents a SetStakeUnitCoverEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleSetStakeUnitCoverEvent struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetStakeUnitCoverEvent is a free log retrieval operation binding the contract event 0x2303e1be65f6d0f77548e74695ab90965b8b94893257b02add910386852ee90d.
//
// Solidity: event SetStakeUnitCoverEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterSetStakeUnitCoverEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleSetStakeUnitCoverEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "SetStakeUnitCoverEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleSetStakeUnitCoverEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "SetStakeUnitCoverEvent", logs: logs, sub: sub}, nil
}

// WatchSetStakeUnitCoverEvent is a free log subscription operation binding the contract event 0x2303e1be65f6d0f77548e74695ab90965b8b94893257b02add910386852ee90d.
//
// Solidity: event SetStakeUnitCoverEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchSetStakeUnitCoverEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleSetStakeUnitCoverEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "SetStakeUnitCoverEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleSetStakeUnitCoverEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "SetStakeUnitCoverEvent", log); err != nil {
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

// ParseSetStakeUnitCoverEvent is a log parse operation binding the contract event 0x2303e1be65f6d0f77548e74695ab90965b8b94893257b02add910386852ee90d.
//
// Solidity: event SetStakeUnitCoverEvent(uint256 amount)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseSetStakeUnitCoverEvent(log types.Log) (*MoonriverDelegatorCoverOracleSetStakeUnitCoverEvent, error) {
	event := new(MoonriverDelegatorCoverOracleSetStakeUnitCoverEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "SetStakeUnitCoverEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleUpgradedToV2EventIterator is returned from FilterUpgradedToV2Event and is used to iterate over the raw logs and unpacked data for UpgradedToV2Event events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleUpgradedToV2EventIterator struct {
	Event *MoonriverDelegatorCoverOracleUpgradedToV2Event // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleUpgradedToV2EventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleUpgradedToV2Event)
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
		it.Event = new(MoonriverDelegatorCoverOracleUpgradedToV2Event)
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
func (it *MoonriverDelegatorCoverOracleUpgradedToV2EventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleUpgradedToV2EventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleUpgradedToV2Event represents a UpgradedToV2Event event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleUpgradedToV2Event struct {
	Member common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpgradedToV2Event is a free log retrieval operation binding the contract event 0x71513e15e4b32c19d588a6ab40e3a88b81642c8a87b7a07894dc69380699baaf.
//
// Solidity: event UpgradedToV2Event(address member)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterUpgradedToV2Event(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleUpgradedToV2EventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "UpgradedToV2Event")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleUpgradedToV2EventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "UpgradedToV2Event", logs: logs, sub: sub}, nil
}

// WatchUpgradedToV2Event is a free log subscription operation binding the contract event 0x71513e15e4b32c19d588a6ab40e3a88b81642c8a87b7a07894dc69380699baaf.
//
// Solidity: event UpgradedToV2Event(address member)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchUpgradedToV2Event(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleUpgradedToV2Event) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "UpgradedToV2Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleUpgradedToV2Event)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "UpgradedToV2Event", log); err != nil {
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

// ParseUpgradedToV2Event is a log parse operation binding the contract event 0x71513e15e4b32c19d588a6ab40e3a88b81642c8a87b7a07894dc69380699baaf.
//
// Solidity: event UpgradedToV2Event(address member)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseUpgradedToV2Event(log types.Log) (*MoonriverDelegatorCoverOracleUpgradedToV2Event, error) {
	event := new(MoonriverDelegatorCoverOracleUpgradedToV2Event)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "UpgradedToV2Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonriverDelegatorCoverOracleWhitelistEventIterator is returned from FilterWhitelistEvent and is used to iterate over the raw logs and unpacked data for WhitelistEvent events raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleWhitelistEventIterator struct {
	Event *MoonriverDelegatorCoverOracleWhitelistEvent // Event containing the contract specifics and raw log

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
func (it *MoonriverDelegatorCoverOracleWhitelistEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonriverDelegatorCoverOracleWhitelistEvent)
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
		it.Event = new(MoonriverDelegatorCoverOracleWhitelistEvent)
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
func (it *MoonriverDelegatorCoverOracleWhitelistEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonriverDelegatorCoverOracleWhitelistEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonriverDelegatorCoverOracleWhitelistEvent represents a WhitelistEvent event raised by the MoonriverDelegatorCoverOracle contract.
type MoonriverDelegatorCoverOracleWhitelistEvent struct {
	Member common.Address
	Proxy  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWhitelistEvent is a free log retrieval operation binding the contract event 0x6a8e2d7cbb198acec5bd0846eded92ab56885ce18ee8a0a1864daf24bfd54267.
//
// Solidity: event WhitelistEvent(address member, address proxy)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) FilterWhitelistEvent(opts *bind.FilterOpts) (*MoonriverDelegatorCoverOracleWhitelistEventIterator, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.FilterLogs(opts, "WhitelistEvent")
	if err != nil {
		return nil, err
	}
	return &MoonriverDelegatorCoverOracleWhitelistEventIterator{contract: _MoonriverDelegatorCoverOracle.contract, event: "WhitelistEvent", logs: logs, sub: sub}, nil
}

// WatchWhitelistEvent is a free log subscription operation binding the contract event 0x6a8e2d7cbb198acec5bd0846eded92ab56885ce18ee8a0a1864daf24bfd54267.
//
// Solidity: event WhitelistEvent(address member, address proxy)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) WatchWhitelistEvent(opts *bind.WatchOpts, sink chan<- *MoonriverDelegatorCoverOracleWhitelistEvent) (event.Subscription, error) {

	logs, sub, err := _MoonriverDelegatorCoverOracle.contract.WatchLogs(opts, "WhitelistEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonriverDelegatorCoverOracleWhitelistEvent)
				if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "WhitelistEvent", log); err != nil {
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

// ParseWhitelistEvent is a log parse operation binding the contract event 0x6a8e2d7cbb198acec5bd0846eded92ab56885ce18ee8a0a1864daf24bfd54267.
//
// Solidity: event WhitelistEvent(address member, address proxy)
func (_MoonriverDelegatorCoverOracle *MoonriverDelegatorCoverOracleFilterer) ParseWhitelistEvent(log types.Log) (*MoonriverDelegatorCoverOracleWhitelistEvent, error) {
	event := new(MoonriverDelegatorCoverOracleWhitelistEvent)
	if err := _MoonriverDelegatorCoverOracle.contract.UnpackLog(event, "WhitelistEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
