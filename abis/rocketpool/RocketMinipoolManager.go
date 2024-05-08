// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abis

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

type Cache interface {
	Set(key string, value any)
	Get(key string) (string, error)
}

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
	_ = abi.ConvertType
)

// RocketMinipoolManagerMetaData contains all meta data concerning the RocketMinipoolManager contract.
var RocketMinipoolManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRocketStorageInterface\",\"name\":\"_rocketStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minipool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"BeginBondReduction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minipool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"CancelReductionVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minipool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"MinipoolCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minipool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"MinipoolDestroyed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minipool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ReductionCancelled\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"createMinipool\",\"outputs\":[{\"internalType\":\"contractRocketMinipoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_validatorPubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_bondAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_currentBalance\",\"type\":\"uint256\"}],\"name\":\"createVacantMinipool\",\"outputs\":[{\"internalType\":\"contractRocketMinipoolInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"decrementNodeStakingMinipoolCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroyMinipool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFinalisedMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getMinipoolAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"getMinipoolByPubkey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getMinipoolCountPerStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"initialisedCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prelaunchCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakingCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawableCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dissolvedCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minipoolAddress\",\"type\":\"address\"}],\"name\":\"getMinipoolDepositType\",\"outputs\":[{\"internalType\":\"enumMinipoolDeposit\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minipoolAddress\",\"type\":\"address\"}],\"name\":\"getMinipoolDestroyed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minipoolAddress\",\"type\":\"address\"}],\"name\":\"getMinipoolExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minipoolAddress\",\"type\":\"address\"}],\"name\":\"getMinipoolPubkey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minipoolAddress\",\"type\":\"address\"}],\"name\":\"getMinipoolRPLSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minipoolAddress\",\"type\":\"address\"}],\"name\":\"getMinipoolWithdrawalCredentials\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeActiveMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeFinalisedMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNodeMinipoolAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeStakingMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_depositSize\",\"type\":\"uint256\"}],\"name\":\"getNodeStakingMinipoolCountBySize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNodeValidatingMinipoolAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeValidatingMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getPrelaunchMinipools\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getVacantMinipoolAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVacantMinipoolCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"incrementNodeFinalisedMinipoolCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"incrementNodeStakingMinipoolCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeVacantMinipool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_pubkey\",\"type\":\"bytes\"}],\"name\":\"setMinipoolPubkey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_previousBond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newBond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_previousFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"updateNodeStakingMinipoolCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RocketMinipoolManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use RocketMinipoolManagerMetaData.ABI instead.
var RocketMinipoolManagerABI = RocketMinipoolManagerMetaData.ABI

// RocketMinipoolManager is an auto generated Go binding around an Ethereum contract.
type RocketMinipoolManager struct {
	RocketMinipoolManagerCaller     // Read-only binding to the contract
	RocketMinipoolManagerTransactor // Write-only binding to the contract
	RocketMinipoolManagerFilterer   // Log filterer for contract events
}

// RocketMinipoolManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RocketMinipoolManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RocketMinipoolManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RocketMinipoolManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RocketMinipoolManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RocketMinipoolManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RocketMinipoolManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RocketMinipoolManagerSession struct {
	Contract     *RocketMinipoolManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RocketMinipoolManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RocketMinipoolManagerCallerSession struct {
	Contract *RocketMinipoolManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// RocketMinipoolManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RocketMinipoolManagerTransactorSession struct {
	Contract     *RocketMinipoolManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// RocketMinipoolManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RocketMinipoolManagerRaw struct {
	Contract *RocketMinipoolManager // Generic contract binding to access the raw methods on
}

// RocketMinipoolManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RocketMinipoolManagerCallerRaw struct {
	Contract *RocketMinipoolManagerCaller // Generic read-only contract binding to access the raw methods on
}

// RocketMinipoolManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RocketMinipoolManagerTransactorRaw struct {
	Contract *RocketMinipoolManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRocketMinipoolManager creates a new instance of RocketMinipoolManager, bound to a specific deployed contract.
func NewRocketMinipoolManager(address common.Address, backend bind.ContractBackend) (*RocketMinipoolManager, error) {
	contract, err := bindRocketMinipoolManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RocketMinipoolManager{RocketMinipoolManagerCaller: RocketMinipoolManagerCaller{contract: contract}, RocketMinipoolManagerTransactor: RocketMinipoolManagerTransactor{contract: contract}, RocketMinipoolManagerFilterer: RocketMinipoolManagerFilterer{contract: contract}}, nil
}

// NewRocketMinipoolManagerCaller creates a new read-only instance of RocketMinipoolManager, bound to a specific deployed contract.
func NewRocketMinipoolManagerCaller(address common.Address, caller bind.ContractCaller) (*RocketMinipoolManagerCaller, error) {
	contract, err := bindRocketMinipoolManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RocketMinipoolManagerCaller{contract: contract}, nil
}

// NewRocketMinipoolManagerTransactor creates a new write-only instance of RocketMinipoolManager, bound to a specific deployed contract.
func NewRocketMinipoolManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*RocketMinipoolManagerTransactor, error) {
	contract, err := bindRocketMinipoolManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RocketMinipoolManagerTransactor{contract: contract}, nil
}

// NewRocketMinipoolManagerFilterer creates a new log filterer instance of RocketMinipoolManager, bound to a specific deployed contract.
func NewRocketMinipoolManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*RocketMinipoolManagerFilterer, error) {
	contract, err := bindRocketMinipoolManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RocketMinipoolManagerFilterer{contract: contract}, nil
}

// bindRocketMinipoolManager binds a generic wrapper to an already deployed contract.
func bindRocketMinipoolManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RocketMinipoolManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RocketMinipoolManager *RocketMinipoolManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RocketMinipoolManager.Contract.RocketMinipoolManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RocketMinipoolManager *RocketMinipoolManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.RocketMinipoolManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RocketMinipoolManager *RocketMinipoolManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.RocketMinipoolManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RocketMinipoolManager *RocketMinipoolManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RocketMinipoolManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RocketMinipoolManager *RocketMinipoolManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RocketMinipoolManager *RocketMinipoolManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.contract.Transact(opts, method, params...)
}

// Solidity: function getActiveMinipoolCount() view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetActiveMinipoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getActiveMinipoolCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetActiveMinipoolCountWithCache(opts *bind.CallOpts, cache Cache) (*big.Int, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetActiveMinipoolCount.%v", []interface{}{opts.BlockNumber})

	result, err := cache.Get(cacheKey)
	var cachedValue *big.Int

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getActiveMinipoolCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetActiveMinipoolCount is a free data retrieval call binding the contract method 0xce9b79ad.
//
// Solidity: function getActiveMinipoolCount() view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetActiveMinipoolCount() (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetActiveMinipoolCount(&_RocketMinipoolManager.CallOpts)
}

// GetActiveMinipoolCount is a free data retrieval call binding the contract method 0xce9b79ad.
//
// Solidity: function getActiveMinipoolCount() view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetActiveMinipoolCount() (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetActiveMinipoolCount(&_RocketMinipoolManager.CallOpts)
}

// Solidity: function getFinalisedMinipoolCount() view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetFinalisedMinipoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getFinalisedMinipoolCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetFinalisedMinipoolCountWithCache(opts *bind.CallOpts, cache Cache) (*big.Int, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetFinalisedMinipoolCount.%v", []interface{}{opts.BlockNumber})

	result, err := cache.Get(cacheKey)
	var cachedValue *big.Int

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getFinalisedMinipoolCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetFinalisedMinipoolCount is a free data retrieval call binding the contract method 0xd1ea6ce0.
//
// Solidity: function getFinalisedMinipoolCount() view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetFinalisedMinipoolCount() (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetFinalisedMinipoolCount(&_RocketMinipoolManager.CallOpts)
}

// GetFinalisedMinipoolCount is a free data retrieval call binding the contract method 0xd1ea6ce0.
//
// Solidity: function getFinalisedMinipoolCount() view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetFinalisedMinipoolCount() (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetFinalisedMinipoolCount(&_RocketMinipoolManager.CallOpts)
}

// Solidity: function getMinipoolAt(uint256 _index) view returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetMinipoolAt(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getMinipoolAt", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetMinipoolAtWithCache(opts *bind.CallOpts, _index *big.Int, cache Cache) (common.Address, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetMinipoolAt.%v", []interface{}{opts.BlockNumber, _index})

	result, err := cache.Get(cacheKey)
	var cachedValue common.Address

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getMinipoolAt", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetMinipoolAt is a free data retrieval call binding the contract method 0xeff7319f.
//
// Solidity: function getMinipoolAt(uint256 _index) view returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetMinipoolAt(_index *big.Int) (common.Address, error) {
	return _RocketMinipoolManager.Contract.GetMinipoolAt(&_RocketMinipoolManager.CallOpts, _index)
}

// GetMinipoolAt is a free data retrieval call binding the contract method 0xeff7319f.
//
// Solidity: function getMinipoolAt(uint256 _index) view returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetMinipoolAt(_index *big.Int) (common.Address, error) {
	return _RocketMinipoolManager.Contract.GetMinipoolAt(&_RocketMinipoolManager.CallOpts, _index)
}

// Solidity: function getMinipoolByPubkey(bytes _pubkey) view returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetMinipoolByPubkey(opts *bind.CallOpts, _pubkey []byte) (common.Address, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getMinipoolByPubkey", _pubkey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetMinipoolByPubkeyWithCache(opts *bind.CallOpts, _pubkey []byte, cache Cache) (common.Address, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetMinipoolByPubkey.%v", []interface{}{opts.BlockNumber, _pubkey})

	result, err := cache.Get(cacheKey)
	var cachedValue common.Address

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getMinipoolByPubkey", _pubkey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetMinipoolByPubkey is a free data retrieval call binding the contract method 0xcf6a4763.
//
// Solidity: function getMinipoolByPubkey(bytes _pubkey) view returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetMinipoolByPubkey(_pubkey []byte) (common.Address, error) {
	return _RocketMinipoolManager.Contract.GetMinipoolByPubkey(&_RocketMinipoolManager.CallOpts, _pubkey)
}

// GetMinipoolByPubkey is a free data retrieval call binding the contract method 0xcf6a4763.
//
// Solidity: function getMinipoolByPubkey(bytes _pubkey) view returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetMinipoolByPubkey(_pubkey []byte) (common.Address, error) {
	return _RocketMinipoolManager.Contract.GetMinipoolByPubkey(&_RocketMinipoolManager.CallOpts, _pubkey)
}

// Solidity: function getMinipoolCount() view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetMinipoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getMinipoolCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetMinipoolCountWithCache(opts *bind.CallOpts, cache Cache) (*big.Int, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetMinipoolCount.%v", []interface{}{opts.BlockNumber})

	result, err := cache.Get(cacheKey)
	var cachedValue *big.Int

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getMinipoolCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetMinipoolCount is a free data retrieval call binding the contract method 0xae4d0bed.
//
// Solidity: function getMinipoolCount() view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetMinipoolCount() (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetMinipoolCount(&_RocketMinipoolManager.CallOpts)
}

// GetMinipoolCount is a free data retrieval call binding the contract method 0xae4d0bed.
//
// Solidity: function getMinipoolCount() view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetMinipoolCount() (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetMinipoolCount(&_RocketMinipoolManager.CallOpts)
}

// GetMinipoolCountPerStatusOutput is the output type for the GetMinipoolCountPerStatus method.
type GetMinipoolCountPerStatusOutput struct {
	InitialisedCount  *big.Int
	PrelaunchCount    *big.Int
	StakingCount      *big.Int
	WithdrawableCount *big.Int
	DissolvedCount    *big.Int
}

// GetMinipoolCountPerStatus is a free data retrieval call binding the contract method 0x3b5ecefa.
//
// Solidity: function getMinipoolCountPerStatus(uint256 _offset, uint256 _limit) view returns(uint256 initialisedCount, uint256 prelaunchCount, uint256 stakingCount, uint256 withdrawableCount, uint256 dissolvedCount)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetMinipoolCountPerStatus(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) (GetMinipoolCountPerStatusOutput, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getMinipoolCountPerStatus", _offset, _limit)

	outstruct := new(GetMinipoolCountPerStatusOutput)
	if err != nil {
		return *outstruct, err
	}

	outstruct.InitialisedCount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PrelaunchCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StakingCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.WithdrawableCount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.DissolvedCount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetMinipoolCountPerStatusWithCache(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int, cache Cache) (GetMinipoolCountPerStatusOutput, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetMinipoolCountPerStatus.%v", []interface{}{opts.BlockNumber, _offset, _limit})

	result, err := cache.Get(cacheKey)
	var cachedValue GetMinipoolCountPerStatusOutput

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getMinipoolCountPerStatus", _offset, _limit)

	outstruct := new(GetMinipoolCountPerStatusOutput)
	if err != nil {
		return *outstruct, err
	}

	outstruct.InitialisedCount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PrelaunchCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StakingCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.WithdrawableCount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.DissolvedCount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	cache.Set(cacheKey, outstruct)
	return *outstruct, err

}

// GetMinipoolCountPerStatus is a free data retrieval call binding the contract method 0x3b5ecefa.
//
// Solidity: function getMinipoolCountPerStatus(uint256 _offset, uint256 _limit) view returns(uint256 initialisedCount, uint256 prelaunchCount, uint256 stakingCount, uint256 withdrawableCount, uint256 dissolvedCount)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetMinipoolCountPerStatus(_offset *big.Int, _limit *big.Int) (struct {
	InitialisedCount  *big.Int
	PrelaunchCount    *big.Int
	StakingCount      *big.Int
	WithdrawableCount *big.Int
	DissolvedCount    *big.Int
}, error) {
	return _RocketMinipoolManager.Contract.GetMinipoolCountPerStatus(&_RocketMinipoolManager.CallOpts, _offset, _limit)
}

// GetMinipoolCountPerStatus is a free data retrieval call binding the contract method 0x3b5ecefa.
//
// Solidity: function getMinipoolCountPerStatus(uint256 _offset, uint256 _limit) view returns(uint256 initialisedCount, uint256 prelaunchCount, uint256 stakingCount, uint256 withdrawableCount, uint256 dissolvedCount)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetMinipoolCountPerStatus(_offset *big.Int, _limit *big.Int) (struct {
	InitialisedCount  *big.Int
	PrelaunchCount    *big.Int
	StakingCount      *big.Int
	WithdrawableCount *big.Int
	DissolvedCount    *big.Int
}, error) {
	return _RocketMinipoolManager.Contract.GetMinipoolCountPerStatus(&_RocketMinipoolManager.CallOpts, _offset, _limit)
}

// Solidity: function getMinipoolDepositType(address _minipoolAddress) view returns(uint8)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetMinipoolDepositType(opts *bind.CallOpts, _minipoolAddress common.Address) (uint8, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getMinipoolDepositType", _minipoolAddress)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetMinipoolDepositTypeWithCache(opts *bind.CallOpts, _minipoolAddress common.Address, cache Cache) (uint8, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetMinipoolDepositType.%v", []interface{}{opts.BlockNumber, _minipoolAddress})

	result, err := cache.Get(cacheKey)
	var cachedValue uint8

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getMinipoolDepositType", _minipoolAddress)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetMinipoolDepositType is a free data retrieval call binding the contract method 0x5ea1a6e2.
//
// Solidity: function getMinipoolDepositType(address _minipoolAddress) view returns(uint8)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetMinipoolDepositType(_minipoolAddress common.Address) (uint8, error) {
	return _RocketMinipoolManager.Contract.GetMinipoolDepositType(&_RocketMinipoolManager.CallOpts, _minipoolAddress)
}

// GetMinipoolDepositType is a free data retrieval call binding the contract method 0x5ea1a6e2.
//
// Solidity: function getMinipoolDepositType(address _minipoolAddress) view returns(uint8)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetMinipoolDepositType(_minipoolAddress common.Address) (uint8, error) {
	return _RocketMinipoolManager.Contract.GetMinipoolDepositType(&_RocketMinipoolManager.CallOpts, _minipoolAddress)
}

// Solidity: function getMinipoolDestroyed(address _minipoolAddress) view returns(bool)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetMinipoolDestroyed(opts *bind.CallOpts, _minipoolAddress common.Address) (bool, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getMinipoolDestroyed", _minipoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetMinipoolDestroyedWithCache(opts *bind.CallOpts, _minipoolAddress common.Address, cache Cache) (bool, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetMinipoolDestroyed.%v", []interface{}{opts.BlockNumber, _minipoolAddress})

	result, err := cache.Get(cacheKey)
	var cachedValue bool

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getMinipoolDestroyed", _minipoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetMinipoolDestroyed is a free data retrieval call binding the contract method 0xa757987a.
//
// Solidity: function getMinipoolDestroyed(address _minipoolAddress) view returns(bool)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetMinipoolDestroyed(_minipoolAddress common.Address) (bool, error) {
	return _RocketMinipoolManager.Contract.GetMinipoolDestroyed(&_RocketMinipoolManager.CallOpts, _minipoolAddress)
}

// GetMinipoolDestroyed is a free data retrieval call binding the contract method 0xa757987a.
//
// Solidity: function getMinipoolDestroyed(address _minipoolAddress) view returns(bool)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetMinipoolDestroyed(_minipoolAddress common.Address) (bool, error) {
	return _RocketMinipoolManager.Contract.GetMinipoolDestroyed(&_RocketMinipoolManager.CallOpts, _minipoolAddress)
}

// Solidity: function getMinipoolExists(address _minipoolAddress) view returns(bool)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetMinipoolExists(opts *bind.CallOpts, _minipoolAddress common.Address) (bool, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getMinipoolExists", _minipoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetMinipoolExistsWithCache(opts *bind.CallOpts, _minipoolAddress common.Address, cache Cache) (bool, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetMinipoolExists.%v", []interface{}{opts.BlockNumber, _minipoolAddress})

	result, err := cache.Get(cacheKey)
	var cachedValue bool

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getMinipoolExists", _minipoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetMinipoolExists is a free data retrieval call binding the contract method 0x606bb62e.
//
// Solidity: function getMinipoolExists(address _minipoolAddress) view returns(bool)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetMinipoolExists(_minipoolAddress common.Address) (bool, error) {
	return _RocketMinipoolManager.Contract.GetMinipoolExists(&_RocketMinipoolManager.CallOpts, _minipoolAddress)
}

// GetMinipoolExists is a free data retrieval call binding the contract method 0x606bb62e.
//
// Solidity: function getMinipoolExists(address _minipoolAddress) view returns(bool)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetMinipoolExists(_minipoolAddress common.Address) (bool, error) {
	return _RocketMinipoolManager.Contract.GetMinipoolExists(&_RocketMinipoolManager.CallOpts, _minipoolAddress)
}

// Solidity: function getMinipoolPubkey(address _minipoolAddress) view returns(bytes)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetMinipoolPubkey(opts *bind.CallOpts, _minipoolAddress common.Address) ([]byte, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getMinipoolPubkey", _minipoolAddress)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetMinipoolPubkeyWithCache(opts *bind.CallOpts, _minipoolAddress common.Address, cache Cache) ([]byte, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetMinipoolPubkey.%v", []interface{}{opts.BlockNumber, _minipoolAddress})

	result, err := cache.Get(cacheKey)
	var cachedValue []byte

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getMinipoolPubkey", _minipoolAddress)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetMinipoolPubkey is a free data retrieval call binding the contract method 0x3eb535e9.
//
// Solidity: function getMinipoolPubkey(address _minipoolAddress) view returns(bytes)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetMinipoolPubkey(_minipoolAddress common.Address) ([]byte, error) {
	return _RocketMinipoolManager.Contract.GetMinipoolPubkey(&_RocketMinipoolManager.CallOpts, _minipoolAddress)
}

// GetMinipoolPubkey is a free data retrieval call binding the contract method 0x3eb535e9.
//
// Solidity: function getMinipoolPubkey(address _minipoolAddress) view returns(bytes)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetMinipoolPubkey(_minipoolAddress common.Address) ([]byte, error) {
	return _RocketMinipoolManager.Contract.GetMinipoolPubkey(&_RocketMinipoolManager.CallOpts, _minipoolAddress)
}

// Solidity: function getMinipoolRPLSlashed(address _minipoolAddress) view returns(bool)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetMinipoolRPLSlashed(opts *bind.CallOpts, _minipoolAddress common.Address) (bool, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getMinipoolRPLSlashed", _minipoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetMinipoolRPLSlashedWithCache(opts *bind.CallOpts, _minipoolAddress common.Address, cache Cache) (bool, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetMinipoolRPLSlashed.%v", []interface{}{opts.BlockNumber, _minipoolAddress})

	result, err := cache.Get(cacheKey)
	var cachedValue bool

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getMinipoolRPLSlashed", _minipoolAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetMinipoolRPLSlashed is a free data retrieval call binding the contract method 0x0c21b8a7.
//
// Solidity: function getMinipoolRPLSlashed(address _minipoolAddress) view returns(bool)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetMinipoolRPLSlashed(_minipoolAddress common.Address) (bool, error) {
	return _RocketMinipoolManager.Contract.GetMinipoolRPLSlashed(&_RocketMinipoolManager.CallOpts, _minipoolAddress)
}

// GetMinipoolRPLSlashed is a free data retrieval call binding the contract method 0x0c21b8a7.
//
// Solidity: function getMinipoolRPLSlashed(address _minipoolAddress) view returns(bool)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetMinipoolRPLSlashed(_minipoolAddress common.Address) (bool, error) {
	return _RocketMinipoolManager.Contract.GetMinipoolRPLSlashed(&_RocketMinipoolManager.CallOpts, _minipoolAddress)
}

// Solidity: function getMinipoolWithdrawalCredentials(address _minipoolAddress) pure returns(bytes)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetMinipoolWithdrawalCredentials(opts *bind.CallOpts, _minipoolAddress common.Address) ([]byte, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getMinipoolWithdrawalCredentials", _minipoolAddress)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetMinipoolWithdrawalCredentialsWithCache(opts *bind.CallOpts, _minipoolAddress common.Address, cache Cache) ([]byte, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetMinipoolWithdrawalCredentials.%v", []interface{}{opts.BlockNumber, _minipoolAddress})

	result, err := cache.Get(cacheKey)
	var cachedValue []byte

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getMinipoolWithdrawalCredentials", _minipoolAddress)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetMinipoolWithdrawalCredentials is a free data retrieval call binding the contract method 0x2cb76c37.
//
// Solidity: function getMinipoolWithdrawalCredentials(address _minipoolAddress) pure returns(bytes)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetMinipoolWithdrawalCredentials(_minipoolAddress common.Address) ([]byte, error) {
	return _RocketMinipoolManager.Contract.GetMinipoolWithdrawalCredentials(&_RocketMinipoolManager.CallOpts, _minipoolAddress)
}

// GetMinipoolWithdrawalCredentials is a free data retrieval call binding the contract method 0x2cb76c37.
//
// Solidity: function getMinipoolWithdrawalCredentials(address _minipoolAddress) pure returns(bytes)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetMinipoolWithdrawalCredentials(_minipoolAddress common.Address) ([]byte, error) {
	return _RocketMinipoolManager.Contract.GetMinipoolWithdrawalCredentials(&_RocketMinipoolManager.CallOpts, _minipoolAddress)
}

// Solidity: function getNodeActiveMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetNodeActiveMinipoolCount(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getNodeActiveMinipoolCount", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetNodeActiveMinipoolCountWithCache(opts *bind.CallOpts, _nodeAddress common.Address, cache Cache) (*big.Int, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetNodeActiveMinipoolCount.%v", []interface{}{opts.BlockNumber, _nodeAddress})

	result, err := cache.Get(cacheKey)
	var cachedValue *big.Int

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getNodeActiveMinipoolCount", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetNodeActiveMinipoolCount is a free data retrieval call binding the contract method 0x1844ec01.
//
// Solidity: function getNodeActiveMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetNodeActiveMinipoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetNodeActiveMinipoolCount(&_RocketMinipoolManager.CallOpts, _nodeAddress)
}

// GetNodeActiveMinipoolCount is a free data retrieval call binding the contract method 0x1844ec01.
//
// Solidity: function getNodeActiveMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetNodeActiveMinipoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetNodeActiveMinipoolCount(&_RocketMinipoolManager.CallOpts, _nodeAddress)
}

// Solidity: function getNodeFinalisedMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetNodeFinalisedMinipoolCount(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getNodeFinalisedMinipoolCount", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetNodeFinalisedMinipoolCountWithCache(opts *bind.CallOpts, _nodeAddress common.Address, cache Cache) (*big.Int, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetNodeFinalisedMinipoolCount.%v", []interface{}{opts.BlockNumber, _nodeAddress})

	result, err := cache.Get(cacheKey)
	var cachedValue *big.Int

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getNodeFinalisedMinipoolCount", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetNodeFinalisedMinipoolCount is a free data retrieval call binding the contract method 0xb88a89f7.
//
// Solidity: function getNodeFinalisedMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetNodeFinalisedMinipoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetNodeFinalisedMinipoolCount(&_RocketMinipoolManager.CallOpts, _nodeAddress)
}

// GetNodeFinalisedMinipoolCount is a free data retrieval call binding the contract method 0xb88a89f7.
//
// Solidity: function getNodeFinalisedMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetNodeFinalisedMinipoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetNodeFinalisedMinipoolCount(&_RocketMinipoolManager.CallOpts, _nodeAddress)
}

// Solidity: function getNodeMinipoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetNodeMinipoolAt(opts *bind.CallOpts, _nodeAddress common.Address, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getNodeMinipoolAt", _nodeAddress, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetNodeMinipoolAtWithCache(opts *bind.CallOpts, _nodeAddress common.Address, _index *big.Int, cache Cache) (common.Address, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetNodeMinipoolAt.%v", []interface{}{opts.BlockNumber, _nodeAddress, _index})

	result, err := cache.Get(cacheKey)
	var cachedValue common.Address

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getNodeMinipoolAt", _nodeAddress, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetNodeMinipoolAt is a free data retrieval call binding the contract method 0x8b300029.
//
// Solidity: function getNodeMinipoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetNodeMinipoolAt(_nodeAddress common.Address, _index *big.Int) (common.Address, error) {
	return _RocketMinipoolManager.Contract.GetNodeMinipoolAt(&_RocketMinipoolManager.CallOpts, _nodeAddress, _index)
}

// GetNodeMinipoolAt is a free data retrieval call binding the contract method 0x8b300029.
//
// Solidity: function getNodeMinipoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetNodeMinipoolAt(_nodeAddress common.Address, _index *big.Int) (common.Address, error) {
	return _RocketMinipoolManager.Contract.GetNodeMinipoolAt(&_RocketMinipoolManager.CallOpts, _nodeAddress, _index)
}

// Solidity: function getNodeMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetNodeMinipoolCount(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getNodeMinipoolCount", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetNodeMinipoolCountWithCache(opts *bind.CallOpts, _nodeAddress common.Address, cache Cache) (*big.Int, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetNodeMinipoolCount.%v", []interface{}{opts.BlockNumber, _nodeAddress})

	result, err := cache.Get(cacheKey)
	var cachedValue *big.Int

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getNodeMinipoolCount", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetNodeMinipoolCount is a free data retrieval call binding the contract method 0x1ce9ec33.
//
// Solidity: function getNodeMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetNodeMinipoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetNodeMinipoolCount(&_RocketMinipoolManager.CallOpts, _nodeAddress)
}

// GetNodeMinipoolCount is a free data retrieval call binding the contract method 0x1ce9ec33.
//
// Solidity: function getNodeMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetNodeMinipoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetNodeMinipoolCount(&_RocketMinipoolManager.CallOpts, _nodeAddress)
}

// Solidity: function getNodeStakingMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetNodeStakingMinipoolCount(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getNodeStakingMinipoolCount", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetNodeStakingMinipoolCountWithCache(opts *bind.CallOpts, _nodeAddress common.Address, cache Cache) (*big.Int, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetNodeStakingMinipoolCount.%v", []interface{}{opts.BlockNumber, _nodeAddress})

	result, err := cache.Get(cacheKey)
	var cachedValue *big.Int

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getNodeStakingMinipoolCount", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetNodeStakingMinipoolCount is a free data retrieval call binding the contract method 0x57b4ef6b.
//
// Solidity: function getNodeStakingMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetNodeStakingMinipoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetNodeStakingMinipoolCount(&_RocketMinipoolManager.CallOpts, _nodeAddress)
}

// GetNodeStakingMinipoolCount is a free data retrieval call binding the contract method 0x57b4ef6b.
//
// Solidity: function getNodeStakingMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetNodeStakingMinipoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetNodeStakingMinipoolCount(&_RocketMinipoolManager.CallOpts, _nodeAddress)
}

// Solidity: function getNodeStakingMinipoolCountBySize(address _nodeAddress, uint256 _depositSize) view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetNodeStakingMinipoolCountBySize(opts *bind.CallOpts, _nodeAddress common.Address, _depositSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getNodeStakingMinipoolCountBySize", _nodeAddress, _depositSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetNodeStakingMinipoolCountBySizeWithCache(opts *bind.CallOpts, _nodeAddress common.Address, _depositSize *big.Int, cache Cache) (*big.Int, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetNodeStakingMinipoolCountBySize.%v", []interface{}{opts.BlockNumber, _nodeAddress, _depositSize})

	result, err := cache.Get(cacheKey)
	var cachedValue *big.Int

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getNodeStakingMinipoolCountBySize", _nodeAddress, _depositSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetNodeStakingMinipoolCountBySize is a free data retrieval call binding the contract method 0x240eb330.
//
// Solidity: function getNodeStakingMinipoolCountBySize(address _nodeAddress, uint256 _depositSize) view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetNodeStakingMinipoolCountBySize(_nodeAddress common.Address, _depositSize *big.Int) (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetNodeStakingMinipoolCountBySize(&_RocketMinipoolManager.CallOpts, _nodeAddress, _depositSize)
}

// GetNodeStakingMinipoolCountBySize is a free data retrieval call binding the contract method 0x240eb330.
//
// Solidity: function getNodeStakingMinipoolCountBySize(address _nodeAddress, uint256 _depositSize) view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetNodeStakingMinipoolCountBySize(_nodeAddress common.Address, _depositSize *big.Int) (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetNodeStakingMinipoolCountBySize(&_RocketMinipoolManager.CallOpts, _nodeAddress, _depositSize)
}

// Solidity: function getNodeValidatingMinipoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetNodeValidatingMinipoolAt(opts *bind.CallOpts, _nodeAddress common.Address, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getNodeValidatingMinipoolAt", _nodeAddress, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetNodeValidatingMinipoolAtWithCache(opts *bind.CallOpts, _nodeAddress common.Address, _index *big.Int, cache Cache) (common.Address, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetNodeValidatingMinipoolAt.%v", []interface{}{opts.BlockNumber, _nodeAddress, _index})

	result, err := cache.Get(cacheKey)
	var cachedValue common.Address

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getNodeValidatingMinipoolAt", _nodeAddress, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetNodeValidatingMinipoolAt is a free data retrieval call binding the contract method 0x9da0700f.
//
// Solidity: function getNodeValidatingMinipoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetNodeValidatingMinipoolAt(_nodeAddress common.Address, _index *big.Int) (common.Address, error) {
	return _RocketMinipoolManager.Contract.GetNodeValidatingMinipoolAt(&_RocketMinipoolManager.CallOpts, _nodeAddress, _index)
}

// GetNodeValidatingMinipoolAt is a free data retrieval call binding the contract method 0x9da0700f.
//
// Solidity: function getNodeValidatingMinipoolAt(address _nodeAddress, uint256 _index) view returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetNodeValidatingMinipoolAt(_nodeAddress common.Address, _index *big.Int) (common.Address, error) {
	return _RocketMinipoolManager.Contract.GetNodeValidatingMinipoolAt(&_RocketMinipoolManager.CallOpts, _nodeAddress, _index)
}

// Solidity: function getNodeValidatingMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetNodeValidatingMinipoolCount(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getNodeValidatingMinipoolCount", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetNodeValidatingMinipoolCountWithCache(opts *bind.CallOpts, _nodeAddress common.Address, cache Cache) (*big.Int, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetNodeValidatingMinipoolCount.%v", []interface{}{opts.BlockNumber, _nodeAddress})

	result, err := cache.Get(cacheKey)
	var cachedValue *big.Int

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getNodeValidatingMinipoolCount", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetNodeValidatingMinipoolCount is a free data retrieval call binding the contract method 0xf90267c4.
//
// Solidity: function getNodeValidatingMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetNodeValidatingMinipoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetNodeValidatingMinipoolCount(&_RocketMinipoolManager.CallOpts, _nodeAddress)
}

// GetNodeValidatingMinipoolCount is a free data retrieval call binding the contract method 0xf90267c4.
//
// Solidity: function getNodeValidatingMinipoolCount(address _nodeAddress) view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetNodeValidatingMinipoolCount(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetNodeValidatingMinipoolCount(&_RocketMinipoolManager.CallOpts, _nodeAddress)
}

// Solidity: function getPrelaunchMinipools(uint256 _offset, uint256 _limit) view returns(address[])
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetPrelaunchMinipools(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getPrelaunchMinipools", _offset, _limit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetPrelaunchMinipoolsWithCache(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int, cache Cache) ([]common.Address, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetPrelaunchMinipools.%v", []interface{}{opts.BlockNumber, _offset, _limit})

	result, err := cache.Get(cacheKey)
	var cachedValue []common.Address

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getPrelaunchMinipools", _offset, _limit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetPrelaunchMinipools is a free data retrieval call binding the contract method 0x5dfef965.
//
// Solidity: function getPrelaunchMinipools(uint256 _offset, uint256 _limit) view returns(address[])
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetPrelaunchMinipools(_offset *big.Int, _limit *big.Int) ([]common.Address, error) {
	return _RocketMinipoolManager.Contract.GetPrelaunchMinipools(&_RocketMinipoolManager.CallOpts, _offset, _limit)
}

// GetPrelaunchMinipools is a free data retrieval call binding the contract method 0x5dfef965.
//
// Solidity: function getPrelaunchMinipools(uint256 _offset, uint256 _limit) view returns(address[])
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetPrelaunchMinipools(_offset *big.Int, _limit *big.Int) ([]common.Address, error) {
	return _RocketMinipoolManager.Contract.GetPrelaunchMinipools(&_RocketMinipoolManager.CallOpts, _offset, _limit)
}

// Solidity: function getStakingMinipoolCount() view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetStakingMinipoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getStakingMinipoolCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetStakingMinipoolCountWithCache(opts *bind.CallOpts, cache Cache) (*big.Int, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetStakingMinipoolCount.%v", []interface{}{opts.BlockNumber})

	result, err := cache.Get(cacheKey)
	var cachedValue *big.Int

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getStakingMinipoolCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetStakingMinipoolCount is a free data retrieval call binding the contract method 0x67bca235.
//
// Solidity: function getStakingMinipoolCount() view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetStakingMinipoolCount() (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetStakingMinipoolCount(&_RocketMinipoolManager.CallOpts)
}

// GetStakingMinipoolCount is a free data retrieval call binding the contract method 0x67bca235.
//
// Solidity: function getStakingMinipoolCount() view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetStakingMinipoolCount() (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetStakingMinipoolCount(&_RocketMinipoolManager.CallOpts)
}

// Solidity: function getVacantMinipoolAt(uint256 _index) view returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetVacantMinipoolAt(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getVacantMinipoolAt", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetVacantMinipoolAtWithCache(opts *bind.CallOpts, _index *big.Int, cache Cache) (common.Address, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetVacantMinipoolAt.%v", []interface{}{opts.BlockNumber, _index})

	result, err := cache.Get(cacheKey)
	var cachedValue common.Address

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getVacantMinipoolAt", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetVacantMinipoolAt is a free data retrieval call binding the contract method 0xd1401991.
//
// Solidity: function getVacantMinipoolAt(uint256 _index) view returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetVacantMinipoolAt(_index *big.Int) (common.Address, error) {
	return _RocketMinipoolManager.Contract.GetVacantMinipoolAt(&_RocketMinipoolManager.CallOpts, _index)
}

// GetVacantMinipoolAt is a free data retrieval call binding the contract method 0xd1401991.
//
// Solidity: function getVacantMinipoolAt(uint256 _index) view returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetVacantMinipoolAt(_index *big.Int) (common.Address, error) {
	return _RocketMinipoolManager.Contract.GetVacantMinipoolAt(&_RocketMinipoolManager.CallOpts, _index)
}

// Solidity: function getVacantMinipoolCount() view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetVacantMinipoolCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "getVacantMinipoolCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) GetVacantMinipoolCountWithCache(opts *bind.CallOpts, cache Cache) (*big.Int, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("GetVacantMinipoolCount.%v", []interface{}{opts.BlockNumber})

	result, err := cache.Get(cacheKey)
	var cachedValue *big.Int

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "getVacantMinipoolCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	cache.Set(cacheKey, out0)
	return out0, err

}

// GetVacantMinipoolCount is a free data retrieval call binding the contract method 0x1286377e.
//
// Solidity: function getVacantMinipoolCount() view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) GetVacantMinipoolCount() (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetVacantMinipoolCount(&_RocketMinipoolManager.CallOpts)
}

// GetVacantMinipoolCount is a free data retrieval call binding the contract method 0x1286377e.
//
// Solidity: function getVacantMinipoolCount() view returns(uint256)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) GetVacantMinipoolCount() (*big.Int, error) {
	return _RocketMinipoolManager.Contract.GetVacantMinipoolCount(&_RocketMinipoolManager.CallOpts)
}

// Solidity: function version() view returns(uint8)
func (_RocketMinipoolManager *RocketMinipoolManagerCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RocketMinipoolManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

func (_RocketMinipoolManager *RocketMinipoolManagerCaller) VersionWithCache(opts *bind.CallOpts, cache Cache) (uint8, error) {
	// Cache key is created with the method name and the parameters
	cacheKey := fmt.Sprintf("Version.%v", []interface{}{opts.BlockNumber})

	result, err := cache.Get(cacheKey)
	var cachedValue uint8

	if err == nil {
		err = json.Unmarshal([]byte(result), &cachedValue)
		if err == nil {
			return cachedValue, nil
		}
	}

	var out []interface{}
	err = _RocketMinipoolManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	cache.Set(cacheKey, out0)
	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) Version() (uint8, error) {
	return _RocketMinipoolManager.Contract.Version(&_RocketMinipoolManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_RocketMinipoolManager *RocketMinipoolManagerCallerSession) Version() (uint8, error) {
	return _RocketMinipoolManager.Contract.Version(&_RocketMinipoolManager.CallOpts)
}

// CreateMinipool is a paid mutator transaction binding the contract method 0xc64372bb.
//
// Solidity: function createMinipool(address _nodeAddress, uint256 _salt) returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerTransactor) CreateMinipool(opts *bind.TransactOpts, _nodeAddress common.Address, _salt *big.Int) (*types.Transaction, error) {
	return _RocketMinipoolManager.contract.Transact(opts, "createMinipool", _nodeAddress, _salt)
}

// CreateMinipool is a paid mutator transaction binding the contract method 0xc64372bb.
//
// Solidity: function createMinipool(address _nodeAddress, uint256 _salt) returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) CreateMinipool(_nodeAddress common.Address, _salt *big.Int) (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.CreateMinipool(&_RocketMinipoolManager.TransactOpts, _nodeAddress, _salt)
}

// CreateMinipool is a paid mutator transaction binding the contract method 0xc64372bb.
//
// Solidity: function createMinipool(address _nodeAddress, uint256 _salt) returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerTransactorSession) CreateMinipool(_nodeAddress common.Address, _salt *big.Int) (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.CreateMinipool(&_RocketMinipoolManager.TransactOpts, _nodeAddress, _salt)
}

// CreateVacantMinipool is a paid mutator transaction binding the contract method 0xa179778b.
//
// Solidity: function createVacantMinipool(address _nodeAddress, uint256 _salt, bytes _validatorPubkey, uint256 _bondAmount, uint256 _currentBalance) returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerTransactor) CreateVacantMinipool(opts *bind.TransactOpts, _nodeAddress common.Address, _salt *big.Int, _validatorPubkey []byte, _bondAmount *big.Int, _currentBalance *big.Int) (*types.Transaction, error) {
	return _RocketMinipoolManager.contract.Transact(opts, "createVacantMinipool", _nodeAddress, _salt, _validatorPubkey, _bondAmount, _currentBalance)
}

// CreateVacantMinipool is a paid mutator transaction binding the contract method 0xa179778b.
//
// Solidity: function createVacantMinipool(address _nodeAddress, uint256 _salt, bytes _validatorPubkey, uint256 _bondAmount, uint256 _currentBalance) returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerSession) CreateVacantMinipool(_nodeAddress common.Address, _salt *big.Int, _validatorPubkey []byte, _bondAmount *big.Int, _currentBalance *big.Int) (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.CreateVacantMinipool(&_RocketMinipoolManager.TransactOpts, _nodeAddress, _salt, _validatorPubkey, _bondAmount, _currentBalance)
}

// CreateVacantMinipool is a paid mutator transaction binding the contract method 0xa179778b.
//
// Solidity: function createVacantMinipool(address _nodeAddress, uint256 _salt, bytes _validatorPubkey, uint256 _bondAmount, uint256 _currentBalance) returns(address)
func (_RocketMinipoolManager *RocketMinipoolManagerTransactorSession) CreateVacantMinipool(_nodeAddress common.Address, _salt *big.Int, _validatorPubkey []byte, _bondAmount *big.Int, _currentBalance *big.Int) (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.CreateVacantMinipool(&_RocketMinipoolManager.TransactOpts, _nodeAddress, _salt, _validatorPubkey, _bondAmount, _currentBalance)
}

// DecrementNodeStakingMinipoolCount is a paid mutator transaction binding the contract method 0x75b59c7f.
//
// Solidity: function decrementNodeStakingMinipoolCount(address _nodeAddress) returns()
func (_RocketMinipoolManager *RocketMinipoolManagerTransactor) DecrementNodeStakingMinipoolCount(opts *bind.TransactOpts, _nodeAddress common.Address) (*types.Transaction, error) {
	return _RocketMinipoolManager.contract.Transact(opts, "decrementNodeStakingMinipoolCount", _nodeAddress)
}

// DecrementNodeStakingMinipoolCount is a paid mutator transaction binding the contract method 0x75b59c7f.
//
// Solidity: function decrementNodeStakingMinipoolCount(address _nodeAddress) returns()
func (_RocketMinipoolManager *RocketMinipoolManagerSession) DecrementNodeStakingMinipoolCount(_nodeAddress common.Address) (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.DecrementNodeStakingMinipoolCount(&_RocketMinipoolManager.TransactOpts, _nodeAddress)
}

// DecrementNodeStakingMinipoolCount is a paid mutator transaction binding the contract method 0x75b59c7f.
//
// Solidity: function decrementNodeStakingMinipoolCount(address _nodeAddress) returns()
func (_RocketMinipoolManager *RocketMinipoolManagerTransactorSession) DecrementNodeStakingMinipoolCount(_nodeAddress common.Address) (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.DecrementNodeStakingMinipoolCount(&_RocketMinipoolManager.TransactOpts, _nodeAddress)
}

// DestroyMinipool is a paid mutator transaction binding the contract method 0x7bb40aaf.
//
// Solidity: function destroyMinipool() returns()
func (_RocketMinipoolManager *RocketMinipoolManagerTransactor) DestroyMinipool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketMinipoolManager.contract.Transact(opts, "destroyMinipool")
}

// DestroyMinipool is a paid mutator transaction binding the contract method 0x7bb40aaf.
//
// Solidity: function destroyMinipool() returns()
func (_RocketMinipoolManager *RocketMinipoolManagerSession) DestroyMinipool() (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.DestroyMinipool(&_RocketMinipoolManager.TransactOpts)
}

// DestroyMinipool is a paid mutator transaction binding the contract method 0x7bb40aaf.
//
// Solidity: function destroyMinipool() returns()
func (_RocketMinipoolManager *RocketMinipoolManagerTransactorSession) DestroyMinipool() (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.DestroyMinipool(&_RocketMinipoolManager.TransactOpts)
}

// IncrementNodeFinalisedMinipoolCount is a paid mutator transaction binding the contract method 0xb04e8868.
//
// Solidity: function incrementNodeFinalisedMinipoolCount(address _nodeAddress) returns()
func (_RocketMinipoolManager *RocketMinipoolManagerTransactor) IncrementNodeFinalisedMinipoolCount(opts *bind.TransactOpts, _nodeAddress common.Address) (*types.Transaction, error) {
	return _RocketMinipoolManager.contract.Transact(opts, "incrementNodeFinalisedMinipoolCount", _nodeAddress)
}

// IncrementNodeFinalisedMinipoolCount is a paid mutator transaction binding the contract method 0xb04e8868.
//
// Solidity: function incrementNodeFinalisedMinipoolCount(address _nodeAddress) returns()
func (_RocketMinipoolManager *RocketMinipoolManagerSession) IncrementNodeFinalisedMinipoolCount(_nodeAddress common.Address) (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.IncrementNodeFinalisedMinipoolCount(&_RocketMinipoolManager.TransactOpts, _nodeAddress)
}

// IncrementNodeFinalisedMinipoolCount is a paid mutator transaction binding the contract method 0xb04e8868.
//
// Solidity: function incrementNodeFinalisedMinipoolCount(address _nodeAddress) returns()
func (_RocketMinipoolManager *RocketMinipoolManagerTransactorSession) IncrementNodeFinalisedMinipoolCount(_nodeAddress common.Address) (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.IncrementNodeFinalisedMinipoolCount(&_RocketMinipoolManager.TransactOpts, _nodeAddress)
}

// IncrementNodeStakingMinipoolCount is a paid mutator transaction binding the contract method 0x9907288c.
//
// Solidity: function incrementNodeStakingMinipoolCount(address _nodeAddress) returns()
func (_RocketMinipoolManager *RocketMinipoolManagerTransactor) IncrementNodeStakingMinipoolCount(opts *bind.TransactOpts, _nodeAddress common.Address) (*types.Transaction, error) {
	return _RocketMinipoolManager.contract.Transact(opts, "incrementNodeStakingMinipoolCount", _nodeAddress)
}

// IncrementNodeStakingMinipoolCount is a paid mutator transaction binding the contract method 0x9907288c.
//
// Solidity: function incrementNodeStakingMinipoolCount(address _nodeAddress) returns()
func (_RocketMinipoolManager *RocketMinipoolManagerSession) IncrementNodeStakingMinipoolCount(_nodeAddress common.Address) (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.IncrementNodeStakingMinipoolCount(&_RocketMinipoolManager.TransactOpts, _nodeAddress)
}

// IncrementNodeStakingMinipoolCount is a paid mutator transaction binding the contract method 0x9907288c.
//
// Solidity: function incrementNodeStakingMinipoolCount(address _nodeAddress) returns()
func (_RocketMinipoolManager *RocketMinipoolManagerTransactorSession) IncrementNodeStakingMinipoolCount(_nodeAddress common.Address) (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.IncrementNodeStakingMinipoolCount(&_RocketMinipoolManager.TransactOpts, _nodeAddress)
}

// RemoveVacantMinipool is a paid mutator transaction binding the contract method 0x44e51a03.
//
// Solidity: function removeVacantMinipool() returns()
func (_RocketMinipoolManager *RocketMinipoolManagerTransactor) RemoveVacantMinipool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketMinipoolManager.contract.Transact(opts, "removeVacantMinipool")
}

// RemoveVacantMinipool is a paid mutator transaction binding the contract method 0x44e51a03.
//
// Solidity: function removeVacantMinipool() returns()
func (_RocketMinipoolManager *RocketMinipoolManagerSession) RemoveVacantMinipool() (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.RemoveVacantMinipool(&_RocketMinipoolManager.TransactOpts)
}

// RemoveVacantMinipool is a paid mutator transaction binding the contract method 0x44e51a03.
//
// Solidity: function removeVacantMinipool() returns()
func (_RocketMinipoolManager *RocketMinipoolManagerTransactorSession) RemoveVacantMinipool() (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.RemoveVacantMinipool(&_RocketMinipoolManager.TransactOpts)
}

// SetMinipoolPubkey is a paid mutator transaction binding the contract method 0x2c7f64d4.
//
// Solidity: function setMinipoolPubkey(bytes _pubkey) returns()
func (_RocketMinipoolManager *RocketMinipoolManagerTransactor) SetMinipoolPubkey(opts *bind.TransactOpts, _pubkey []byte) (*types.Transaction, error) {
	return _RocketMinipoolManager.contract.Transact(opts, "setMinipoolPubkey", _pubkey)
}

// SetMinipoolPubkey is a paid mutator transaction binding the contract method 0x2c7f64d4.
//
// Solidity: function setMinipoolPubkey(bytes _pubkey) returns()
func (_RocketMinipoolManager *RocketMinipoolManagerSession) SetMinipoolPubkey(_pubkey []byte) (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.SetMinipoolPubkey(&_RocketMinipoolManager.TransactOpts, _pubkey)
}

// SetMinipoolPubkey is a paid mutator transaction binding the contract method 0x2c7f64d4.
//
// Solidity: function setMinipoolPubkey(bytes _pubkey) returns()
func (_RocketMinipoolManager *RocketMinipoolManagerTransactorSession) SetMinipoolPubkey(_pubkey []byte) (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.SetMinipoolPubkey(&_RocketMinipoolManager.TransactOpts, _pubkey)
}

// UpdateNodeStakingMinipoolCount is a paid mutator transaction binding the contract method 0x0fcc8178.
//
// Solidity: function updateNodeStakingMinipoolCount(uint256 _previousBond, uint256 _newBond, uint256 _previousFee, uint256 _newFee) returns()
func (_RocketMinipoolManager *RocketMinipoolManagerTransactor) UpdateNodeStakingMinipoolCount(opts *bind.TransactOpts, _previousBond *big.Int, _newBond *big.Int, _previousFee *big.Int, _newFee *big.Int) (*types.Transaction, error) {
	return _RocketMinipoolManager.contract.Transact(opts, "updateNodeStakingMinipoolCount", _previousBond, _newBond, _previousFee, _newFee)
}

// UpdateNodeStakingMinipoolCount is a paid mutator transaction binding the contract method 0x0fcc8178.
//
// Solidity: function updateNodeStakingMinipoolCount(uint256 _previousBond, uint256 _newBond, uint256 _previousFee, uint256 _newFee) returns()
func (_RocketMinipoolManager *RocketMinipoolManagerSession) UpdateNodeStakingMinipoolCount(_previousBond *big.Int, _newBond *big.Int, _previousFee *big.Int, _newFee *big.Int) (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.UpdateNodeStakingMinipoolCount(&_RocketMinipoolManager.TransactOpts, _previousBond, _newBond, _previousFee, _newFee)
}

// UpdateNodeStakingMinipoolCount is a paid mutator transaction binding the contract method 0x0fcc8178.
//
// Solidity: function updateNodeStakingMinipoolCount(uint256 _previousBond, uint256 _newBond, uint256 _previousFee, uint256 _newFee) returns()
func (_RocketMinipoolManager *RocketMinipoolManagerTransactorSession) UpdateNodeStakingMinipoolCount(_previousBond *big.Int, _newBond *big.Int, _previousFee *big.Int, _newFee *big.Int) (*types.Transaction, error) {
	return _RocketMinipoolManager.Contract.UpdateNodeStakingMinipoolCount(&_RocketMinipoolManager.TransactOpts, _previousBond, _newBond, _previousFee, _newFee)
}

// RocketMinipoolManagerBeginBondReductionIterator is returned from FilterBeginBondReduction and is used to iterate over the raw logs and unpacked data for BeginBondReduction events raised by the RocketMinipoolManager contract.
type RocketMinipoolManagerBeginBondReductionIterator struct {
	Event *RocketMinipoolManagerBeginBondReduction // Event containing the contract specifics and raw log

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
func (it *RocketMinipoolManagerBeginBondReductionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketMinipoolManagerBeginBondReduction)
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
		it.Event = new(RocketMinipoolManagerBeginBondReduction)
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
func (it *RocketMinipoolManagerBeginBondReductionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketMinipoolManagerBeginBondReductionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketMinipoolManagerBeginBondReduction represents a BeginBondReduction event raised by the RocketMinipoolManager contract.
type RocketMinipoolManagerBeginBondReduction struct {
	Minipool common.Address
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBeginBondReduction is a free log retrieval operation binding the contract event 0xebae778aeca850cfad423220d7978598785ef5905b868c482e03011b53808678.
//
// Solidity: event BeginBondReduction(address indexed minipool, uint256 time)
func (_RocketMinipoolManager *RocketMinipoolManagerFilterer) FilterBeginBondReduction(opts *bind.FilterOpts, minipool []common.Address) (*RocketMinipoolManagerBeginBondReductionIterator, error) {

	var minipoolRule []interface{}
	for _, minipoolItem := range minipool {
		minipoolRule = append(minipoolRule, minipoolItem)
	}

	logs, sub, err := _RocketMinipoolManager.contract.FilterLogs(opts, "BeginBondReduction", minipoolRule)
	if err != nil {
		return nil, err
	}
	return &RocketMinipoolManagerBeginBondReductionIterator{contract: _RocketMinipoolManager.contract, event: "BeginBondReduction", logs: logs, sub: sub}, nil
}

// WatchBeginBondReduction is a free log subscription operation binding the contract event 0xebae778aeca850cfad423220d7978598785ef5905b868c482e03011b53808678.
//
// Solidity: event BeginBondReduction(address indexed minipool, uint256 time)
func (_RocketMinipoolManager *RocketMinipoolManagerFilterer) WatchBeginBondReduction(opts *bind.WatchOpts, sink chan<- *RocketMinipoolManagerBeginBondReduction, minipool []common.Address) (event.Subscription, error) {

	var minipoolRule []interface{}
	for _, minipoolItem := range minipool {
		minipoolRule = append(minipoolRule, minipoolItem)
	}

	logs, sub, err := _RocketMinipoolManager.contract.WatchLogs(opts, "BeginBondReduction", minipoolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketMinipoolManagerBeginBondReduction)
				if err := _RocketMinipoolManager.contract.UnpackLog(event, "BeginBondReduction", log); err != nil {
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

// ParseBeginBondReduction is a log parse operation binding the contract event 0xebae778aeca850cfad423220d7978598785ef5905b868c482e03011b53808678.
//
// Solidity: event BeginBondReduction(address indexed minipool, uint256 time)
func (_RocketMinipoolManager *RocketMinipoolManagerFilterer) ParseBeginBondReduction(log types.Log) (*RocketMinipoolManagerBeginBondReduction, error) {
	event := new(RocketMinipoolManagerBeginBondReduction)
	if err := _RocketMinipoolManager.contract.UnpackLog(event, "BeginBondReduction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketMinipoolManagerCancelReductionVotedIterator is returned from FilterCancelReductionVoted and is used to iterate over the raw logs and unpacked data for CancelReductionVoted events raised by the RocketMinipoolManager contract.
type RocketMinipoolManagerCancelReductionVotedIterator struct {
	Event *RocketMinipoolManagerCancelReductionVoted // Event containing the contract specifics and raw log

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
func (it *RocketMinipoolManagerCancelReductionVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketMinipoolManagerCancelReductionVoted)
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
		it.Event = new(RocketMinipoolManagerCancelReductionVoted)
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
func (it *RocketMinipoolManagerCancelReductionVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketMinipoolManagerCancelReductionVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketMinipoolManagerCancelReductionVoted represents a CancelReductionVoted event raised by the RocketMinipoolManager contract.
type RocketMinipoolManagerCancelReductionVoted struct {
	Minipool common.Address
	Member   common.Address
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCancelReductionVoted is a free log retrieval operation binding the contract event 0x7cfaf8cd5e8153c2679a1100841445ab9926ef98867f41bd6e6afa9e3f09068e.
//
// Solidity: event CancelReductionVoted(address indexed minipool, address indexed member, uint256 time)
func (_RocketMinipoolManager *RocketMinipoolManagerFilterer) FilterCancelReductionVoted(opts *bind.FilterOpts, minipool []common.Address, member []common.Address) (*RocketMinipoolManagerCancelReductionVotedIterator, error) {

	var minipoolRule []interface{}
	for _, minipoolItem := range minipool {
		minipoolRule = append(minipoolRule, minipoolItem)
	}
	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _RocketMinipoolManager.contract.FilterLogs(opts, "CancelReductionVoted", minipoolRule, memberRule)
	if err != nil {
		return nil, err
	}
	return &RocketMinipoolManagerCancelReductionVotedIterator{contract: _RocketMinipoolManager.contract, event: "CancelReductionVoted", logs: logs, sub: sub}, nil
}

// WatchCancelReductionVoted is a free log subscription operation binding the contract event 0x7cfaf8cd5e8153c2679a1100841445ab9926ef98867f41bd6e6afa9e3f09068e.
//
// Solidity: event CancelReductionVoted(address indexed minipool, address indexed member, uint256 time)
func (_RocketMinipoolManager *RocketMinipoolManagerFilterer) WatchCancelReductionVoted(opts *bind.WatchOpts, sink chan<- *RocketMinipoolManagerCancelReductionVoted, minipool []common.Address, member []common.Address) (event.Subscription, error) {

	var minipoolRule []interface{}
	for _, minipoolItem := range minipool {
		minipoolRule = append(minipoolRule, minipoolItem)
	}
	var memberRule []interface{}
	for _, memberItem := range member {
		memberRule = append(memberRule, memberItem)
	}

	logs, sub, err := _RocketMinipoolManager.contract.WatchLogs(opts, "CancelReductionVoted", minipoolRule, memberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketMinipoolManagerCancelReductionVoted)
				if err := _RocketMinipoolManager.contract.UnpackLog(event, "CancelReductionVoted", log); err != nil {
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

// ParseCancelReductionVoted is a log parse operation binding the contract event 0x7cfaf8cd5e8153c2679a1100841445ab9926ef98867f41bd6e6afa9e3f09068e.
//
// Solidity: event CancelReductionVoted(address indexed minipool, address indexed member, uint256 time)
func (_RocketMinipoolManager *RocketMinipoolManagerFilterer) ParseCancelReductionVoted(log types.Log) (*RocketMinipoolManagerCancelReductionVoted, error) {
	event := new(RocketMinipoolManagerCancelReductionVoted)
	if err := _RocketMinipoolManager.contract.UnpackLog(event, "CancelReductionVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketMinipoolManagerMinipoolCreatedIterator is returned from FilterMinipoolCreated and is used to iterate over the raw logs and unpacked data for MinipoolCreated events raised by the RocketMinipoolManager contract.
type RocketMinipoolManagerMinipoolCreatedIterator struct {
	Event *RocketMinipoolManagerMinipoolCreated // Event containing the contract specifics and raw log

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
func (it *RocketMinipoolManagerMinipoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketMinipoolManagerMinipoolCreated)
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
		it.Event = new(RocketMinipoolManagerMinipoolCreated)
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
func (it *RocketMinipoolManagerMinipoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketMinipoolManagerMinipoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketMinipoolManagerMinipoolCreated represents a MinipoolCreated event raised by the RocketMinipoolManager contract.
type RocketMinipoolManagerMinipoolCreated struct {
	Minipool common.Address
	Node     common.Address
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinipoolCreated is a free log retrieval operation binding the contract event 0x08b4b91bafaf992145c5dd7e098dfcdb32f879714c154c651c2758a44c7aeae4.
//
// Solidity: event MinipoolCreated(address indexed minipool, address indexed node, uint256 time)
func (_RocketMinipoolManager *RocketMinipoolManagerFilterer) FilterMinipoolCreated(opts *bind.FilterOpts, minipool []common.Address, node []common.Address) (*RocketMinipoolManagerMinipoolCreatedIterator, error) {

	var minipoolRule []interface{}
	for _, minipoolItem := range minipool {
		minipoolRule = append(minipoolRule, minipoolItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _RocketMinipoolManager.contract.FilterLogs(opts, "MinipoolCreated", minipoolRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &RocketMinipoolManagerMinipoolCreatedIterator{contract: _RocketMinipoolManager.contract, event: "MinipoolCreated", logs: logs, sub: sub}, nil
}

// WatchMinipoolCreated is a free log subscription operation binding the contract event 0x08b4b91bafaf992145c5dd7e098dfcdb32f879714c154c651c2758a44c7aeae4.
//
// Solidity: event MinipoolCreated(address indexed minipool, address indexed node, uint256 time)
func (_RocketMinipoolManager *RocketMinipoolManagerFilterer) WatchMinipoolCreated(opts *bind.WatchOpts, sink chan<- *RocketMinipoolManagerMinipoolCreated, minipool []common.Address, node []common.Address) (event.Subscription, error) {

	var minipoolRule []interface{}
	for _, minipoolItem := range minipool {
		minipoolRule = append(minipoolRule, minipoolItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _RocketMinipoolManager.contract.WatchLogs(opts, "MinipoolCreated", minipoolRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketMinipoolManagerMinipoolCreated)
				if err := _RocketMinipoolManager.contract.UnpackLog(event, "MinipoolCreated", log); err != nil {
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

// ParseMinipoolCreated is a log parse operation binding the contract event 0x08b4b91bafaf992145c5dd7e098dfcdb32f879714c154c651c2758a44c7aeae4.
//
// Solidity: event MinipoolCreated(address indexed minipool, address indexed node, uint256 time)
func (_RocketMinipoolManager *RocketMinipoolManagerFilterer) ParseMinipoolCreated(log types.Log) (*RocketMinipoolManagerMinipoolCreated, error) {
	event := new(RocketMinipoolManagerMinipoolCreated)
	if err := _RocketMinipoolManager.contract.UnpackLog(event, "MinipoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketMinipoolManagerMinipoolDestroyedIterator is returned from FilterMinipoolDestroyed and is used to iterate over the raw logs and unpacked data for MinipoolDestroyed events raised by the RocketMinipoolManager contract.
type RocketMinipoolManagerMinipoolDestroyedIterator struct {
	Event *RocketMinipoolManagerMinipoolDestroyed // Event containing the contract specifics and raw log

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
func (it *RocketMinipoolManagerMinipoolDestroyedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketMinipoolManagerMinipoolDestroyed)
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
		it.Event = new(RocketMinipoolManagerMinipoolDestroyed)
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
func (it *RocketMinipoolManagerMinipoolDestroyedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketMinipoolManagerMinipoolDestroyedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketMinipoolManagerMinipoolDestroyed represents a MinipoolDestroyed event raised by the RocketMinipoolManager contract.
type RocketMinipoolManagerMinipoolDestroyed struct {
	Minipool common.Address
	Node     common.Address
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMinipoolDestroyed is a free log retrieval operation binding the contract event 0x3097cb0f536cd88115b814915d7030d2fe958943357cd2b1a9e1dba8a673ec69.
//
// Solidity: event MinipoolDestroyed(address indexed minipool, address indexed node, uint256 time)
func (_RocketMinipoolManager *RocketMinipoolManagerFilterer) FilterMinipoolDestroyed(opts *bind.FilterOpts, minipool []common.Address, node []common.Address) (*RocketMinipoolManagerMinipoolDestroyedIterator, error) {

	var minipoolRule []interface{}
	for _, minipoolItem := range minipool {
		minipoolRule = append(minipoolRule, minipoolItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _RocketMinipoolManager.contract.FilterLogs(opts, "MinipoolDestroyed", minipoolRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return &RocketMinipoolManagerMinipoolDestroyedIterator{contract: _RocketMinipoolManager.contract, event: "MinipoolDestroyed", logs: logs, sub: sub}, nil
}

// WatchMinipoolDestroyed is a free log subscription operation binding the contract event 0x3097cb0f536cd88115b814915d7030d2fe958943357cd2b1a9e1dba8a673ec69.
//
// Solidity: event MinipoolDestroyed(address indexed minipool, address indexed node, uint256 time)
func (_RocketMinipoolManager *RocketMinipoolManagerFilterer) WatchMinipoolDestroyed(opts *bind.WatchOpts, sink chan<- *RocketMinipoolManagerMinipoolDestroyed, minipool []common.Address, node []common.Address) (event.Subscription, error) {

	var minipoolRule []interface{}
	for _, minipoolItem := range minipool {
		minipoolRule = append(minipoolRule, minipoolItem)
	}
	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _RocketMinipoolManager.contract.WatchLogs(opts, "MinipoolDestroyed", minipoolRule, nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketMinipoolManagerMinipoolDestroyed)
				if err := _RocketMinipoolManager.contract.UnpackLog(event, "MinipoolDestroyed", log); err != nil {
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

// ParseMinipoolDestroyed is a log parse operation binding the contract event 0x3097cb0f536cd88115b814915d7030d2fe958943357cd2b1a9e1dba8a673ec69.
//
// Solidity: event MinipoolDestroyed(address indexed minipool, address indexed node, uint256 time)
func (_RocketMinipoolManager *RocketMinipoolManagerFilterer) ParseMinipoolDestroyed(log types.Log) (*RocketMinipoolManagerMinipoolDestroyed, error) {
	event := new(RocketMinipoolManagerMinipoolDestroyed)
	if err := _RocketMinipoolManager.contract.UnpackLog(event, "MinipoolDestroyed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketMinipoolManagerReductionCancelledIterator is returned from FilterReductionCancelled and is used to iterate over the raw logs and unpacked data for ReductionCancelled events raised by the RocketMinipoolManager contract.
type RocketMinipoolManagerReductionCancelledIterator struct {
	Event *RocketMinipoolManagerReductionCancelled // Event containing the contract specifics and raw log

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
func (it *RocketMinipoolManagerReductionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketMinipoolManagerReductionCancelled)
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
		it.Event = new(RocketMinipoolManagerReductionCancelled)
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
func (it *RocketMinipoolManagerReductionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketMinipoolManagerReductionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketMinipoolManagerReductionCancelled represents a ReductionCancelled event raised by the RocketMinipoolManager contract.
type RocketMinipoolManagerReductionCancelled struct {
	Minipool common.Address
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReductionCancelled is a free log retrieval operation binding the contract event 0xf5248f3aef129e8fd5ee9f0bd6dc051e7a9f3ad64e171a0acfb396298683bc24.
//
// Solidity: event ReductionCancelled(address indexed minipool, uint256 time)
func (_RocketMinipoolManager *RocketMinipoolManagerFilterer) FilterReductionCancelled(opts *bind.FilterOpts, minipool []common.Address) (*RocketMinipoolManagerReductionCancelledIterator, error) {

	var minipoolRule []interface{}
	for _, minipoolItem := range minipool {
		minipoolRule = append(minipoolRule, minipoolItem)
	}

	logs, sub, err := _RocketMinipoolManager.contract.FilterLogs(opts, "ReductionCancelled", minipoolRule)
	if err != nil {
		return nil, err
	}
	return &RocketMinipoolManagerReductionCancelledIterator{contract: _RocketMinipoolManager.contract, event: "ReductionCancelled", logs: logs, sub: sub}, nil
}

// WatchReductionCancelled is a free log subscription operation binding the contract event 0xf5248f3aef129e8fd5ee9f0bd6dc051e7a9f3ad64e171a0acfb396298683bc24.
//
// Solidity: event ReductionCancelled(address indexed minipool, uint256 time)
func (_RocketMinipoolManager *RocketMinipoolManagerFilterer) WatchReductionCancelled(opts *bind.WatchOpts, sink chan<- *RocketMinipoolManagerReductionCancelled, minipool []common.Address) (event.Subscription, error) {

	var minipoolRule []interface{}
	for _, minipoolItem := range minipool {
		minipoolRule = append(minipoolRule, minipoolItem)
	}

	logs, sub, err := _RocketMinipoolManager.contract.WatchLogs(opts, "ReductionCancelled", minipoolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketMinipoolManagerReductionCancelled)
				if err := _RocketMinipoolManager.contract.UnpackLog(event, "ReductionCancelled", log); err != nil {
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

// ParseReductionCancelled is a log parse operation binding the contract event 0xf5248f3aef129e8fd5ee9f0bd6dc051e7a9f3ad64e171a0acfb396298683bc24.
//
// Solidity: event ReductionCancelled(address indexed minipool, uint256 time)
func (_RocketMinipoolManager *RocketMinipoolManagerFilterer) ParseReductionCancelled(log types.Log) (*RocketMinipoolManagerReductionCancelled, error) {
	event := new(RocketMinipoolManagerReductionCancelled)
	if err := _RocketMinipoolManager.contract.UnpackLog(event, "ReductionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
