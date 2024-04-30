// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abis

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
	_ = abi.ConvertType
)

// RocketNodeStakingMetaData contains all meta data concerning the RocketNodeStaking contract.
var RocketNodeStakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRocketStorageInterface\",\"name\":\"_rocketStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"RPLSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"RPLStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"RPLWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"StakeRPLForAllowed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeETHCollateralisationRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeETHMatched\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeETHMatchedLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeETHProvided\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeEffectiveRPLStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeMaximumRPLStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeMinimumRPLStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeRPLStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"}],\"name\":\"getNodeRPLStakedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalRPLStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_allowed\",\"type\":\"bool\"}],\"name\":\"setStakeRPLForAllowed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ethSlashAmount\",\"type\":\"uint256\"}],\"name\":\"slashRPL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stakeRPL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stakeRPLFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawRPL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RocketNodeStakingABI is the input ABI used to generate the binding from.
// Deprecated: Use RocketNodeStakingMetaData.ABI instead.
var RocketNodeStakingABI = RocketNodeStakingMetaData.ABI

// RocketNodeStaking is an auto generated Go binding around an Ethereum contract.
type RocketNodeStaking struct {
	RocketNodeStakingCaller     // Read-only binding to the contract
	RocketNodeStakingTransactor // Write-only binding to the contract
	RocketNodeStakingFilterer   // Log filterer for contract events
}

// RocketNodeStakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type RocketNodeStakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RocketNodeStakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RocketNodeStakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RocketNodeStakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RocketNodeStakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RocketNodeStakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RocketNodeStakingSession struct {
	Contract     *RocketNodeStaking // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RocketNodeStakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RocketNodeStakingCallerSession struct {
	Contract *RocketNodeStakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// RocketNodeStakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RocketNodeStakingTransactorSession struct {
	Contract     *RocketNodeStakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// RocketNodeStakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type RocketNodeStakingRaw struct {
	Contract *RocketNodeStaking // Generic contract binding to access the raw methods on
}

// RocketNodeStakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RocketNodeStakingCallerRaw struct {
	Contract *RocketNodeStakingCaller // Generic read-only contract binding to access the raw methods on
}

// RocketNodeStakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RocketNodeStakingTransactorRaw struct {
	Contract *RocketNodeStakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRocketNodeStaking creates a new instance of RocketNodeStaking, bound to a specific deployed contract.
func NewRocketNodeStaking(address common.Address, backend bind.ContractBackend) (*RocketNodeStaking, error) {
	contract, err := bindRocketNodeStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RocketNodeStaking{RocketNodeStakingCaller: RocketNodeStakingCaller{contract: contract}, RocketNodeStakingTransactor: RocketNodeStakingTransactor{contract: contract}, RocketNodeStakingFilterer: RocketNodeStakingFilterer{contract: contract}}, nil
}

// NewRocketNodeStakingCaller creates a new read-only instance of RocketNodeStaking, bound to a specific deployed contract.
func NewRocketNodeStakingCaller(address common.Address, caller bind.ContractCaller) (*RocketNodeStakingCaller, error) {
	contract, err := bindRocketNodeStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RocketNodeStakingCaller{contract: contract}, nil
}

// NewRocketNodeStakingTransactor creates a new write-only instance of RocketNodeStaking, bound to a specific deployed contract.
func NewRocketNodeStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*RocketNodeStakingTransactor, error) {
	contract, err := bindRocketNodeStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RocketNodeStakingTransactor{contract: contract}, nil
}

// NewRocketNodeStakingFilterer creates a new log filterer instance of RocketNodeStaking, bound to a specific deployed contract.
func NewRocketNodeStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*RocketNodeStakingFilterer, error) {
	contract, err := bindRocketNodeStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RocketNodeStakingFilterer{contract: contract}, nil
}

// bindRocketNodeStaking binds a generic wrapper to an already deployed contract.
func bindRocketNodeStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RocketNodeStakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RocketNodeStaking *RocketNodeStakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RocketNodeStaking.Contract.RocketNodeStakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RocketNodeStaking *RocketNodeStakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketNodeStaking.Contract.RocketNodeStakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RocketNodeStaking *RocketNodeStakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RocketNodeStaking.Contract.RocketNodeStakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RocketNodeStaking *RocketNodeStakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RocketNodeStaking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RocketNodeStaking *RocketNodeStakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketNodeStaking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RocketNodeStaking *RocketNodeStakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RocketNodeStaking.Contract.contract.Transact(opts, method, params...)
}

// GetNodeETHCollateralisationRatio is a free data retrieval call binding the contract method 0x97be2143.
//
// Solidity: function getNodeETHCollateralisationRatio(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingCaller) GetNodeETHCollateralisationRatio(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketNodeStaking.contract.Call(opts, &out, "getNodeETHCollateralisationRatio", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeETHCollateralisationRatio is a free data retrieval call binding the contract method 0x97be2143.
//
// Solidity: function getNodeETHCollateralisationRatio(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingSession) GetNodeETHCollateralisationRatio(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketNodeStaking.Contract.GetNodeETHCollateralisationRatio(&_RocketNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeETHCollateralisationRatio is a free data retrieval call binding the contract method 0x97be2143.
//
// Solidity: function getNodeETHCollateralisationRatio(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingCallerSession) GetNodeETHCollateralisationRatio(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketNodeStaking.Contract.GetNodeETHCollateralisationRatio(&_RocketNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeETHMatched is a free data retrieval call binding the contract method 0xa493e6a2.
//
// Solidity: function getNodeETHMatched(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingCaller) GetNodeETHMatched(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketNodeStaking.contract.Call(opts, &out, "getNodeETHMatched", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeETHMatched is a free data retrieval call binding the contract method 0xa493e6a2.
//
// Solidity: function getNodeETHMatched(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingSession) GetNodeETHMatched(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketNodeStaking.Contract.GetNodeETHMatched(&_RocketNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeETHMatched is a free data retrieval call binding the contract method 0xa493e6a2.
//
// Solidity: function getNodeETHMatched(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingCallerSession) GetNodeETHMatched(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketNodeStaking.Contract.GetNodeETHMatched(&_RocketNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeETHMatchedLimit is a free data retrieval call binding the contract method 0x48aeedf5.
//
// Solidity: function getNodeETHMatchedLimit(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingCaller) GetNodeETHMatchedLimit(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketNodeStaking.contract.Call(opts, &out, "getNodeETHMatchedLimit", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeETHMatchedLimit is a free data retrieval call binding the contract method 0x48aeedf5.
//
// Solidity: function getNodeETHMatchedLimit(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingSession) GetNodeETHMatchedLimit(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketNodeStaking.Contract.GetNodeETHMatchedLimit(&_RocketNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeETHMatchedLimit is a free data retrieval call binding the contract method 0x48aeedf5.
//
// Solidity: function getNodeETHMatchedLimit(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingCallerSession) GetNodeETHMatchedLimit(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketNodeStaking.Contract.GetNodeETHMatchedLimit(&_RocketNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeETHProvided is a free data retrieval call binding the contract method 0xe5e82b7c.
//
// Solidity: function getNodeETHProvided(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingCaller) GetNodeETHProvided(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketNodeStaking.contract.Call(opts, &out, "getNodeETHProvided", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeETHProvided is a free data retrieval call binding the contract method 0xe5e82b7c.
//
// Solidity: function getNodeETHProvided(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingSession) GetNodeETHProvided(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketNodeStaking.Contract.GetNodeETHProvided(&_RocketNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeETHProvided is a free data retrieval call binding the contract method 0xe5e82b7c.
//
// Solidity: function getNodeETHProvided(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingCallerSession) GetNodeETHProvided(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketNodeStaking.Contract.GetNodeETHProvided(&_RocketNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeEffectiveRPLStake is a free data retrieval call binding the contract method 0xf0d19b89.
//
// Solidity: function getNodeEffectiveRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingCaller) GetNodeEffectiveRPLStake(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketNodeStaking.contract.Call(opts, &out, "getNodeEffectiveRPLStake", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeEffectiveRPLStake is a free data retrieval call binding the contract method 0xf0d19b89.
//
// Solidity: function getNodeEffectiveRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingSession) GetNodeEffectiveRPLStake(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketNodeStaking.Contract.GetNodeEffectiveRPLStake(&_RocketNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeEffectiveRPLStake is a free data retrieval call binding the contract method 0xf0d19b89.
//
// Solidity: function getNodeEffectiveRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingCallerSession) GetNodeEffectiveRPLStake(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketNodeStaking.Contract.GetNodeEffectiveRPLStake(&_RocketNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeMaximumRPLStake is a free data retrieval call binding the contract method 0x4e58ff6e.
//
// Solidity: function getNodeMaximumRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingCaller) GetNodeMaximumRPLStake(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketNodeStaking.contract.Call(opts, &out, "getNodeMaximumRPLStake", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeMaximumRPLStake is a free data retrieval call binding the contract method 0x4e58ff6e.
//
// Solidity: function getNodeMaximumRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingSession) GetNodeMaximumRPLStake(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketNodeStaking.Contract.GetNodeMaximumRPLStake(&_RocketNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeMaximumRPLStake is a free data retrieval call binding the contract method 0x4e58ff6e.
//
// Solidity: function getNodeMaximumRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingCallerSession) GetNodeMaximumRPLStake(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketNodeStaking.Contract.GetNodeMaximumRPLStake(&_RocketNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeMinimumRPLStake is a free data retrieval call binding the contract method 0x03fa87b4.
//
// Solidity: function getNodeMinimumRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingCaller) GetNodeMinimumRPLStake(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketNodeStaking.contract.Call(opts, &out, "getNodeMinimumRPLStake", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeMinimumRPLStake is a free data retrieval call binding the contract method 0x03fa87b4.
//
// Solidity: function getNodeMinimumRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingSession) GetNodeMinimumRPLStake(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketNodeStaking.Contract.GetNodeMinimumRPLStake(&_RocketNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeMinimumRPLStake is a free data retrieval call binding the contract method 0x03fa87b4.
//
// Solidity: function getNodeMinimumRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingCallerSession) GetNodeMinimumRPLStake(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketNodeStaking.Contract.GetNodeMinimumRPLStake(&_RocketNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeRPLStake is a free data retrieval call binding the contract method 0x9961cee4.
//
// Solidity: function getNodeRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingCaller) GetNodeRPLStake(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketNodeStaking.contract.Call(opts, &out, "getNodeRPLStake", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeRPLStake is a free data retrieval call binding the contract method 0x9961cee4.
//
// Solidity: function getNodeRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingSession) GetNodeRPLStake(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketNodeStaking.Contract.GetNodeRPLStake(&_RocketNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeRPLStake is a free data retrieval call binding the contract method 0x9961cee4.
//
// Solidity: function getNodeRPLStake(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingCallerSession) GetNodeRPLStake(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketNodeStaking.Contract.GetNodeRPLStake(&_RocketNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeRPLStakedTime is a free data retrieval call binding the contract method 0xc0d05dd8.
//
// Solidity: function getNodeRPLStakedTime(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingCaller) GetNodeRPLStakedTime(opts *bind.CallOpts, _nodeAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketNodeStaking.contract.Call(opts, &out, "getNodeRPLStakedTime", _nodeAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeRPLStakedTime is a free data retrieval call binding the contract method 0xc0d05dd8.
//
// Solidity: function getNodeRPLStakedTime(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingSession) GetNodeRPLStakedTime(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketNodeStaking.Contract.GetNodeRPLStakedTime(&_RocketNodeStaking.CallOpts, _nodeAddress)
}

// GetNodeRPLStakedTime is a free data retrieval call binding the contract method 0xc0d05dd8.
//
// Solidity: function getNodeRPLStakedTime(address _nodeAddress) view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingCallerSession) GetNodeRPLStakedTime(_nodeAddress common.Address) (*big.Int, error) {
	return _RocketNodeStaking.Contract.GetNodeRPLStakedTime(&_RocketNodeStaking.CallOpts, _nodeAddress)
}

// GetTotalRPLStake is a free data retrieval call binding the contract method 0x9a206c8e.
//
// Solidity: function getTotalRPLStake() view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingCaller) GetTotalRPLStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RocketNodeStaking.contract.Call(opts, &out, "getTotalRPLStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalRPLStake is a free data retrieval call binding the contract method 0x9a206c8e.
//
// Solidity: function getTotalRPLStake() view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingSession) GetTotalRPLStake() (*big.Int, error) {
	return _RocketNodeStaking.Contract.GetTotalRPLStake(&_RocketNodeStaking.CallOpts)
}

// GetTotalRPLStake is a free data retrieval call binding the contract method 0x9a206c8e.
//
// Solidity: function getTotalRPLStake() view returns(uint256)
func (_RocketNodeStaking *RocketNodeStakingCallerSession) GetTotalRPLStake() (*big.Int, error) {
	return _RocketNodeStaking.Contract.GetTotalRPLStake(&_RocketNodeStaking.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_RocketNodeStaking *RocketNodeStakingCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RocketNodeStaking.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_RocketNodeStaking *RocketNodeStakingSession) Version() (uint8, error) {
	return _RocketNodeStaking.Contract.Version(&_RocketNodeStaking.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_RocketNodeStaking *RocketNodeStakingCallerSession) Version() (uint8, error) {
	return _RocketNodeStaking.Contract.Version(&_RocketNodeStaking.CallOpts)
}

// SetStakeRPLForAllowed is a paid mutator transaction binding the contract method 0x088903a4.
//
// Solidity: function setStakeRPLForAllowed(address _caller, bool _allowed) returns()
func (_RocketNodeStaking *RocketNodeStakingTransactor) SetStakeRPLForAllowed(opts *bind.TransactOpts, _caller common.Address, _allowed bool) (*types.Transaction, error) {
	return _RocketNodeStaking.contract.Transact(opts, "setStakeRPLForAllowed", _caller, _allowed)
}

// SetStakeRPLForAllowed is a paid mutator transaction binding the contract method 0x088903a4.
//
// Solidity: function setStakeRPLForAllowed(address _caller, bool _allowed) returns()
func (_RocketNodeStaking *RocketNodeStakingSession) SetStakeRPLForAllowed(_caller common.Address, _allowed bool) (*types.Transaction, error) {
	return _RocketNodeStaking.Contract.SetStakeRPLForAllowed(&_RocketNodeStaking.TransactOpts, _caller, _allowed)
}

// SetStakeRPLForAllowed is a paid mutator transaction binding the contract method 0x088903a4.
//
// Solidity: function setStakeRPLForAllowed(address _caller, bool _allowed) returns()
func (_RocketNodeStaking *RocketNodeStakingTransactorSession) SetStakeRPLForAllowed(_caller common.Address, _allowed bool) (*types.Transaction, error) {
	return _RocketNodeStaking.Contract.SetStakeRPLForAllowed(&_RocketNodeStaking.TransactOpts, _caller, _allowed)
}

// SlashRPL is a paid mutator transaction binding the contract method 0x245395a6.
//
// Solidity: function slashRPL(address _nodeAddress, uint256 _ethSlashAmount) returns()
func (_RocketNodeStaking *RocketNodeStakingTransactor) SlashRPL(opts *bind.TransactOpts, _nodeAddress common.Address, _ethSlashAmount *big.Int) (*types.Transaction, error) {
	return _RocketNodeStaking.contract.Transact(opts, "slashRPL", _nodeAddress, _ethSlashAmount)
}

// SlashRPL is a paid mutator transaction binding the contract method 0x245395a6.
//
// Solidity: function slashRPL(address _nodeAddress, uint256 _ethSlashAmount) returns()
func (_RocketNodeStaking *RocketNodeStakingSession) SlashRPL(_nodeAddress common.Address, _ethSlashAmount *big.Int) (*types.Transaction, error) {
	return _RocketNodeStaking.Contract.SlashRPL(&_RocketNodeStaking.TransactOpts, _nodeAddress, _ethSlashAmount)
}

// SlashRPL is a paid mutator transaction binding the contract method 0x245395a6.
//
// Solidity: function slashRPL(address _nodeAddress, uint256 _ethSlashAmount) returns()
func (_RocketNodeStaking *RocketNodeStakingTransactorSession) SlashRPL(_nodeAddress common.Address, _ethSlashAmount *big.Int) (*types.Transaction, error) {
	return _RocketNodeStaking.Contract.SlashRPL(&_RocketNodeStaking.TransactOpts, _nodeAddress, _ethSlashAmount)
}

// StakeRPL is a paid mutator transaction binding the contract method 0x3e200d4b.
//
// Solidity: function stakeRPL(uint256 _amount) returns()
func (_RocketNodeStaking *RocketNodeStakingTransactor) StakeRPL(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RocketNodeStaking.contract.Transact(opts, "stakeRPL", _amount)
}

// StakeRPL is a paid mutator transaction binding the contract method 0x3e200d4b.
//
// Solidity: function stakeRPL(uint256 _amount) returns()
func (_RocketNodeStaking *RocketNodeStakingSession) StakeRPL(_amount *big.Int) (*types.Transaction, error) {
	return _RocketNodeStaking.Contract.StakeRPL(&_RocketNodeStaking.TransactOpts, _amount)
}

// StakeRPL is a paid mutator transaction binding the contract method 0x3e200d4b.
//
// Solidity: function stakeRPL(uint256 _amount) returns()
func (_RocketNodeStaking *RocketNodeStakingTransactorSession) StakeRPL(_amount *big.Int) (*types.Transaction, error) {
	return _RocketNodeStaking.Contract.StakeRPL(&_RocketNodeStaking.TransactOpts, _amount)
}

// StakeRPLFor is a paid mutator transaction binding the contract method 0xcb1c8321.
//
// Solidity: function stakeRPLFor(address _nodeAddress, uint256 _amount) returns()
func (_RocketNodeStaking *RocketNodeStakingTransactor) StakeRPLFor(opts *bind.TransactOpts, _nodeAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RocketNodeStaking.contract.Transact(opts, "stakeRPLFor", _nodeAddress, _amount)
}

// StakeRPLFor is a paid mutator transaction binding the contract method 0xcb1c8321.
//
// Solidity: function stakeRPLFor(address _nodeAddress, uint256 _amount) returns()
func (_RocketNodeStaking *RocketNodeStakingSession) StakeRPLFor(_nodeAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RocketNodeStaking.Contract.StakeRPLFor(&_RocketNodeStaking.TransactOpts, _nodeAddress, _amount)
}

// StakeRPLFor is a paid mutator transaction binding the contract method 0xcb1c8321.
//
// Solidity: function stakeRPLFor(address _nodeAddress, uint256 _amount) returns()
func (_RocketNodeStaking *RocketNodeStakingTransactorSession) StakeRPLFor(_nodeAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RocketNodeStaking.Contract.StakeRPLFor(&_RocketNodeStaking.TransactOpts, _nodeAddress, _amount)
}

// WithdrawRPL is a paid mutator transaction binding the contract method 0x6b088d5c.
//
// Solidity: function withdrawRPL(uint256 _amount) returns()
func (_RocketNodeStaking *RocketNodeStakingTransactor) WithdrawRPL(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RocketNodeStaking.contract.Transact(opts, "withdrawRPL", _amount)
}

// WithdrawRPL is a paid mutator transaction binding the contract method 0x6b088d5c.
//
// Solidity: function withdrawRPL(uint256 _amount) returns()
func (_RocketNodeStaking *RocketNodeStakingSession) WithdrawRPL(_amount *big.Int) (*types.Transaction, error) {
	return _RocketNodeStaking.Contract.WithdrawRPL(&_RocketNodeStaking.TransactOpts, _amount)
}

// WithdrawRPL is a paid mutator transaction binding the contract method 0x6b088d5c.
//
// Solidity: function withdrawRPL(uint256 _amount) returns()
func (_RocketNodeStaking *RocketNodeStakingTransactorSession) WithdrawRPL(_amount *big.Int) (*types.Transaction, error) {
	return _RocketNodeStaking.Contract.WithdrawRPL(&_RocketNodeStaking.TransactOpts, _amount)
}

// RocketNodeStakingRPLSlashedIterator is returned from FilterRPLSlashed and is used to iterate over the raw logs and unpacked data for RPLSlashed events raised by the RocketNodeStaking contract.
type RocketNodeStakingRPLSlashedIterator struct {
	Event *RocketNodeStakingRPLSlashed // Event containing the contract specifics and raw log

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
func (it *RocketNodeStakingRPLSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketNodeStakingRPLSlashed)
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
		it.Event = new(RocketNodeStakingRPLSlashed)
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
func (it *RocketNodeStakingRPLSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketNodeStakingRPLSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketNodeStakingRPLSlashed represents a RPLSlashed event raised by the RocketNodeStaking contract.
type RocketNodeStakingRPLSlashed struct {
	Node     common.Address
	Amount   *big.Int
	EthValue *big.Int
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRPLSlashed is a free log retrieval operation binding the contract event 0x38a2777b6a84fdb3fc375fe8ade69fdad1afdcdd93c79e7ae2319b806a626c4d.
//
// Solidity: event RPLSlashed(address indexed node, uint256 amount, uint256 ethValue, uint256 time)
func (_RocketNodeStaking *RocketNodeStakingFilterer) FilterRPLSlashed(opts *bind.FilterOpts, node []common.Address) (*RocketNodeStakingRPLSlashedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _RocketNodeStaking.contract.FilterLogs(opts, "RPLSlashed", nodeRule)
	if err != nil {
		return nil, err
	}
	return &RocketNodeStakingRPLSlashedIterator{contract: _RocketNodeStaking.contract, event: "RPLSlashed", logs: logs, sub: sub}, nil
}

// WatchRPLSlashed is a free log subscription operation binding the contract event 0x38a2777b6a84fdb3fc375fe8ade69fdad1afdcdd93c79e7ae2319b806a626c4d.
//
// Solidity: event RPLSlashed(address indexed node, uint256 amount, uint256 ethValue, uint256 time)
func (_RocketNodeStaking *RocketNodeStakingFilterer) WatchRPLSlashed(opts *bind.WatchOpts, sink chan<- *RocketNodeStakingRPLSlashed, node []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}

	logs, sub, err := _RocketNodeStaking.contract.WatchLogs(opts, "RPLSlashed", nodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketNodeStakingRPLSlashed)
				if err := _RocketNodeStaking.contract.UnpackLog(event, "RPLSlashed", log); err != nil {
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

// ParseRPLSlashed is a log parse operation binding the contract event 0x38a2777b6a84fdb3fc375fe8ade69fdad1afdcdd93c79e7ae2319b806a626c4d.
//
// Solidity: event RPLSlashed(address indexed node, uint256 amount, uint256 ethValue, uint256 time)
func (_RocketNodeStaking *RocketNodeStakingFilterer) ParseRPLSlashed(log types.Log) (*RocketNodeStakingRPLSlashed, error) {
	event := new(RocketNodeStakingRPLSlashed)
	if err := _RocketNodeStaking.contract.UnpackLog(event, "RPLSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketNodeStakingRPLStakedIterator is returned from FilterRPLStaked and is used to iterate over the raw logs and unpacked data for RPLStaked events raised by the RocketNodeStaking contract.
type RocketNodeStakingRPLStakedIterator struct {
	Event *RocketNodeStakingRPLStaked // Event containing the contract specifics and raw log

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
func (it *RocketNodeStakingRPLStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketNodeStakingRPLStaked)
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
		it.Event = new(RocketNodeStakingRPLStaked)
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
func (it *RocketNodeStakingRPLStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketNodeStakingRPLStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketNodeStakingRPLStaked represents a RPLStaked event raised by the RocketNodeStaking contract.
type RocketNodeStakingRPLStaked struct {
	From   common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRPLStaked is a free log retrieval operation binding the contract event 0x4e3bcb61bb8e63cb9ed2c46d47eeb6ae847c629e909fbb32b9d17874affb4a89.
//
// Solidity: event RPLStaked(address indexed from, uint256 amount, uint256 time)
func (_RocketNodeStaking *RocketNodeStakingFilterer) FilterRPLStaked(opts *bind.FilterOpts, from []common.Address) (*RocketNodeStakingRPLStakedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RocketNodeStaking.contract.FilterLogs(opts, "RPLStaked", fromRule)
	if err != nil {
		return nil, err
	}
	return &RocketNodeStakingRPLStakedIterator{contract: _RocketNodeStaking.contract, event: "RPLStaked", logs: logs, sub: sub}, nil
}

// WatchRPLStaked is a free log subscription operation binding the contract event 0x4e3bcb61bb8e63cb9ed2c46d47eeb6ae847c629e909fbb32b9d17874affb4a89.
//
// Solidity: event RPLStaked(address indexed from, uint256 amount, uint256 time)
func (_RocketNodeStaking *RocketNodeStakingFilterer) WatchRPLStaked(opts *bind.WatchOpts, sink chan<- *RocketNodeStakingRPLStaked, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RocketNodeStaking.contract.WatchLogs(opts, "RPLStaked", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketNodeStakingRPLStaked)
				if err := _RocketNodeStaking.contract.UnpackLog(event, "RPLStaked", log); err != nil {
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

// ParseRPLStaked is a log parse operation binding the contract event 0x4e3bcb61bb8e63cb9ed2c46d47eeb6ae847c629e909fbb32b9d17874affb4a89.
//
// Solidity: event RPLStaked(address indexed from, uint256 amount, uint256 time)
func (_RocketNodeStaking *RocketNodeStakingFilterer) ParseRPLStaked(log types.Log) (*RocketNodeStakingRPLStaked, error) {
	event := new(RocketNodeStakingRPLStaked)
	if err := _RocketNodeStaking.contract.UnpackLog(event, "RPLStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketNodeStakingRPLWithdrawnIterator is returned from FilterRPLWithdrawn and is used to iterate over the raw logs and unpacked data for RPLWithdrawn events raised by the RocketNodeStaking contract.
type RocketNodeStakingRPLWithdrawnIterator struct {
	Event *RocketNodeStakingRPLWithdrawn // Event containing the contract specifics and raw log

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
func (it *RocketNodeStakingRPLWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketNodeStakingRPLWithdrawn)
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
		it.Event = new(RocketNodeStakingRPLWithdrawn)
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
func (it *RocketNodeStakingRPLWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketNodeStakingRPLWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketNodeStakingRPLWithdrawn represents a RPLWithdrawn event raised by the RocketNodeStaking contract.
type RocketNodeStakingRPLWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRPLWithdrawn is a free log retrieval operation binding the contract event 0x9947063f70b076145616018b82ed1dd5585e15b7ae0a0b17a8b06bec4c4c31e2.
//
// Solidity: event RPLWithdrawn(address indexed to, uint256 amount, uint256 time)
func (_RocketNodeStaking *RocketNodeStakingFilterer) FilterRPLWithdrawn(opts *bind.FilterOpts, to []common.Address) (*RocketNodeStakingRPLWithdrawnIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RocketNodeStaking.contract.FilterLogs(opts, "RPLWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return &RocketNodeStakingRPLWithdrawnIterator{contract: _RocketNodeStaking.contract, event: "RPLWithdrawn", logs: logs, sub: sub}, nil
}

// WatchRPLWithdrawn is a free log subscription operation binding the contract event 0x9947063f70b076145616018b82ed1dd5585e15b7ae0a0b17a8b06bec4c4c31e2.
//
// Solidity: event RPLWithdrawn(address indexed to, uint256 amount, uint256 time)
func (_RocketNodeStaking *RocketNodeStakingFilterer) WatchRPLWithdrawn(opts *bind.WatchOpts, sink chan<- *RocketNodeStakingRPLWithdrawn, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RocketNodeStaking.contract.WatchLogs(opts, "RPLWithdrawn", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketNodeStakingRPLWithdrawn)
				if err := _RocketNodeStaking.contract.UnpackLog(event, "RPLWithdrawn", log); err != nil {
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

// ParseRPLWithdrawn is a log parse operation binding the contract event 0x9947063f70b076145616018b82ed1dd5585e15b7ae0a0b17a8b06bec4c4c31e2.
//
// Solidity: event RPLWithdrawn(address indexed to, uint256 amount, uint256 time)
func (_RocketNodeStaking *RocketNodeStakingFilterer) ParseRPLWithdrawn(log types.Log) (*RocketNodeStakingRPLWithdrawn, error) {
	event := new(RocketNodeStakingRPLWithdrawn)
	if err := _RocketNodeStaking.contract.UnpackLog(event, "RPLWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketNodeStakingStakeRPLForAllowedIterator is returned from FilterStakeRPLForAllowed and is used to iterate over the raw logs and unpacked data for StakeRPLForAllowed events raised by the RocketNodeStaking contract.
type RocketNodeStakingStakeRPLForAllowedIterator struct {
	Event *RocketNodeStakingStakeRPLForAllowed // Event containing the contract specifics and raw log

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
func (it *RocketNodeStakingStakeRPLForAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketNodeStakingStakeRPLForAllowed)
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
		it.Event = new(RocketNodeStakingStakeRPLForAllowed)
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
func (it *RocketNodeStakingStakeRPLForAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketNodeStakingStakeRPLForAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketNodeStakingStakeRPLForAllowed represents a StakeRPLForAllowed event raised by the RocketNodeStaking contract.
type RocketNodeStakingStakeRPLForAllowed struct {
	Node    common.Address
	Caller  common.Address
	Allowed bool
	Time    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeRPLForAllowed is a free log retrieval operation binding the contract event 0xb8502fe170368d1312ca3c9feac7aba9cd92406753d7eca9f11df9757081aec5.
//
// Solidity: event StakeRPLForAllowed(address indexed node, address indexed caller, bool allowed, uint256 time)
func (_RocketNodeStaking *RocketNodeStakingFilterer) FilterStakeRPLForAllowed(opts *bind.FilterOpts, node []common.Address, caller []common.Address) (*RocketNodeStakingStakeRPLForAllowedIterator, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _RocketNodeStaking.contract.FilterLogs(opts, "StakeRPLForAllowed", nodeRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &RocketNodeStakingStakeRPLForAllowedIterator{contract: _RocketNodeStaking.contract, event: "StakeRPLForAllowed", logs: logs, sub: sub}, nil
}

// WatchStakeRPLForAllowed is a free log subscription operation binding the contract event 0xb8502fe170368d1312ca3c9feac7aba9cd92406753d7eca9f11df9757081aec5.
//
// Solidity: event StakeRPLForAllowed(address indexed node, address indexed caller, bool allowed, uint256 time)
func (_RocketNodeStaking *RocketNodeStakingFilterer) WatchStakeRPLForAllowed(opts *bind.WatchOpts, sink chan<- *RocketNodeStakingStakeRPLForAllowed, node []common.Address, caller []common.Address) (event.Subscription, error) {

	var nodeRule []interface{}
	for _, nodeItem := range node {
		nodeRule = append(nodeRule, nodeItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _RocketNodeStaking.contract.WatchLogs(opts, "StakeRPLForAllowed", nodeRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketNodeStakingStakeRPLForAllowed)
				if err := _RocketNodeStaking.contract.UnpackLog(event, "StakeRPLForAllowed", log); err != nil {
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

// ParseStakeRPLForAllowed is a log parse operation binding the contract event 0xb8502fe170368d1312ca3c9feac7aba9cd92406753d7eca9f11df9757081aec5.
//
// Solidity: event StakeRPLForAllowed(address indexed node, address indexed caller, bool allowed, uint256 time)
func (_RocketNodeStaking *RocketNodeStakingFilterer) ParseStakeRPLForAllowed(log types.Log) (*RocketNodeStakingStakeRPLForAllowed, error) {
	event := new(RocketNodeStakingStakeRPLForAllowed)
	if err := _RocketNodeStaking.contract.UnpackLog(event, "StakeRPLForAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
