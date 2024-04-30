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

// RocketVaultMetaData contains all meta data concerning the RocketVault contract.
var RocketVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRocketStorageInterface\",\"name\":\"_rocketStorageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"by\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"EtherDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"by\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"EtherWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"by\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TokenBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"by\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TokenDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"by\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"to\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TokenTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"by\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TokenWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_networkContractName\",\"type\":\"string\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_networkContractName\",\"type\":\"string\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"balanceOfToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractERC20Burnable\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burnToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEther\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_networkContractName\",\"type\":\"string\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"depositToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_networkContractName\",\"type\":\"string\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_withdrawalAddress\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RocketVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use RocketVaultMetaData.ABI instead.
var RocketVaultABI = RocketVaultMetaData.ABI

// RocketVault is an auto generated Go binding around an Ethereum contract.
type RocketVault struct {
	RocketVaultCaller     // Read-only binding to the contract
	RocketVaultTransactor // Write-only binding to the contract
	RocketVaultFilterer   // Log filterer for contract events
}

// RocketVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type RocketVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RocketVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RocketVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RocketVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RocketVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RocketVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RocketVaultSession struct {
	Contract     *RocketVault      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RocketVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RocketVaultCallerSession struct {
	Contract *RocketVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RocketVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RocketVaultTransactorSession struct {
	Contract     *RocketVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RocketVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type RocketVaultRaw struct {
	Contract *RocketVault // Generic contract binding to access the raw methods on
}

// RocketVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RocketVaultCallerRaw struct {
	Contract *RocketVaultCaller // Generic read-only contract binding to access the raw methods on
}

// RocketVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RocketVaultTransactorRaw struct {
	Contract *RocketVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRocketVault creates a new instance of RocketVault, bound to a specific deployed contract.
func NewRocketVault(address common.Address, backend bind.ContractBackend) (*RocketVault, error) {
	contract, err := bindRocketVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RocketVault{RocketVaultCaller: RocketVaultCaller{contract: contract}, RocketVaultTransactor: RocketVaultTransactor{contract: contract}, RocketVaultFilterer: RocketVaultFilterer{contract: contract}}, nil
}

// NewRocketVaultCaller creates a new read-only instance of RocketVault, bound to a specific deployed contract.
func NewRocketVaultCaller(address common.Address, caller bind.ContractCaller) (*RocketVaultCaller, error) {
	contract, err := bindRocketVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RocketVaultCaller{contract: contract}, nil
}

// NewRocketVaultTransactor creates a new write-only instance of RocketVault, bound to a specific deployed contract.
func NewRocketVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*RocketVaultTransactor, error) {
	contract, err := bindRocketVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RocketVaultTransactor{contract: contract}, nil
}

// NewRocketVaultFilterer creates a new log filterer instance of RocketVault, bound to a specific deployed contract.
func NewRocketVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*RocketVaultFilterer, error) {
	contract, err := bindRocketVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RocketVaultFilterer{contract: contract}, nil
}

// bindRocketVault binds a generic wrapper to an already deployed contract.
func bindRocketVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RocketVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RocketVault *RocketVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RocketVault.Contract.RocketVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RocketVault *RocketVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketVault.Contract.RocketVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RocketVault *RocketVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RocketVault.Contract.RocketVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RocketVault *RocketVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RocketVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RocketVault *RocketVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RocketVault *RocketVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RocketVault.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string _networkContractName) view returns(uint256)
func (_RocketVault *RocketVaultCaller) BalanceOf(opts *bind.CallOpts, _networkContractName string) (*big.Int, error) {
	var out []interface{}
	err := _RocketVault.contract.Call(opts, &out, "balanceOf", _networkContractName)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string _networkContractName) view returns(uint256)
func (_RocketVault *RocketVaultSession) BalanceOf(_networkContractName string) (*big.Int, error) {
	return _RocketVault.Contract.BalanceOf(&_RocketVault.CallOpts, _networkContractName)
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string _networkContractName) view returns(uint256)
func (_RocketVault *RocketVaultCallerSession) BalanceOf(_networkContractName string) (*big.Int, error) {
	return _RocketVault.Contract.BalanceOf(&_RocketVault.CallOpts, _networkContractName)
}

// BalanceOfToken is a free data retrieval call binding the contract method 0xeccefff6.
//
// Solidity: function balanceOfToken(string _networkContractName, address _tokenAddress) view returns(uint256)
func (_RocketVault *RocketVaultCaller) BalanceOfToken(opts *bind.CallOpts, _networkContractName string, _tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RocketVault.contract.Call(opts, &out, "balanceOfToken", _networkContractName, _tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfToken is a free data retrieval call binding the contract method 0xeccefff6.
//
// Solidity: function balanceOfToken(string _networkContractName, address _tokenAddress) view returns(uint256)
func (_RocketVault *RocketVaultSession) BalanceOfToken(_networkContractName string, _tokenAddress common.Address) (*big.Int, error) {
	return _RocketVault.Contract.BalanceOfToken(&_RocketVault.CallOpts, _networkContractName, _tokenAddress)
}

// BalanceOfToken is a free data retrieval call binding the contract method 0xeccefff6.
//
// Solidity: function balanceOfToken(string _networkContractName, address _tokenAddress) view returns(uint256)
func (_RocketVault *RocketVaultCallerSession) BalanceOfToken(_networkContractName string, _tokenAddress common.Address) (*big.Int, error) {
	return _RocketVault.Contract.BalanceOfToken(&_RocketVault.CallOpts, _networkContractName, _tokenAddress)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_RocketVault *RocketVaultCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RocketVault.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_RocketVault *RocketVaultSession) Version() (uint8, error) {
	return _RocketVault.Contract.Version(&_RocketVault.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_RocketVault *RocketVaultCallerSession) Version() (uint8, error) {
	return _RocketVault.Contract.Version(&_RocketVault.CallOpts)
}

// BurnToken is a paid mutator transaction binding the contract method 0xd1df306c.
//
// Solidity: function burnToken(address _tokenAddress, uint256 _amount) returns()
func (_RocketVault *RocketVaultTransactor) BurnToken(opts *bind.TransactOpts, _tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RocketVault.contract.Transact(opts, "burnToken", _tokenAddress, _amount)
}

// BurnToken is a paid mutator transaction binding the contract method 0xd1df306c.
//
// Solidity: function burnToken(address _tokenAddress, uint256 _amount) returns()
func (_RocketVault *RocketVaultSession) BurnToken(_tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RocketVault.Contract.BurnToken(&_RocketVault.TransactOpts, _tokenAddress, _amount)
}

// BurnToken is a paid mutator transaction binding the contract method 0xd1df306c.
//
// Solidity: function burnToken(address _tokenAddress, uint256 _amount) returns()
func (_RocketVault *RocketVaultTransactorSession) BurnToken(_tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RocketVault.Contract.BurnToken(&_RocketVault.TransactOpts, _tokenAddress, _amount)
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() payable returns()
func (_RocketVault *RocketVaultTransactor) DepositEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RocketVault.contract.Transact(opts, "depositEther")
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() payable returns()
func (_RocketVault *RocketVaultSession) DepositEther() (*types.Transaction, error) {
	return _RocketVault.Contract.DepositEther(&_RocketVault.TransactOpts)
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() payable returns()
func (_RocketVault *RocketVaultTransactorSession) DepositEther() (*types.Transaction, error) {
	return _RocketVault.Contract.DepositEther(&_RocketVault.TransactOpts)
}

// DepositToken is a paid mutator transaction binding the contract method 0xf4442958.
//
// Solidity: function depositToken(string _networkContractName, address _tokenContract, uint256 _amount) returns()
func (_RocketVault *RocketVaultTransactor) DepositToken(opts *bind.TransactOpts, _networkContractName string, _tokenContract common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RocketVault.contract.Transact(opts, "depositToken", _networkContractName, _tokenContract, _amount)
}

// DepositToken is a paid mutator transaction binding the contract method 0xf4442958.
//
// Solidity: function depositToken(string _networkContractName, address _tokenContract, uint256 _amount) returns()
func (_RocketVault *RocketVaultSession) DepositToken(_networkContractName string, _tokenContract common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RocketVault.Contract.DepositToken(&_RocketVault.TransactOpts, _networkContractName, _tokenContract, _amount)
}

// DepositToken is a paid mutator transaction binding the contract method 0xf4442958.
//
// Solidity: function depositToken(string _networkContractName, address _tokenContract, uint256 _amount) returns()
func (_RocketVault *RocketVaultTransactorSession) DepositToken(_networkContractName string, _tokenContract common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RocketVault.Contract.DepositToken(&_RocketVault.TransactOpts, _networkContractName, _tokenContract, _amount)
}

// TransferToken is a paid mutator transaction binding the contract method 0xee91035e.
//
// Solidity: function transferToken(string _networkContractName, address _tokenAddress, uint256 _amount) returns()
func (_RocketVault *RocketVaultTransactor) TransferToken(opts *bind.TransactOpts, _networkContractName string, _tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RocketVault.contract.Transact(opts, "transferToken", _networkContractName, _tokenAddress, _amount)
}

// TransferToken is a paid mutator transaction binding the contract method 0xee91035e.
//
// Solidity: function transferToken(string _networkContractName, address _tokenAddress, uint256 _amount) returns()
func (_RocketVault *RocketVaultSession) TransferToken(_networkContractName string, _tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RocketVault.Contract.TransferToken(&_RocketVault.TransactOpts, _networkContractName, _tokenAddress, _amount)
}

// TransferToken is a paid mutator transaction binding the contract method 0xee91035e.
//
// Solidity: function transferToken(string _networkContractName, address _tokenAddress, uint256 _amount) returns()
func (_RocketVault *RocketVaultTransactorSession) TransferToken(_networkContractName string, _tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RocketVault.Contract.TransferToken(&_RocketVault.TransactOpts, _networkContractName, _tokenAddress, _amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 _amount) returns()
func (_RocketVault *RocketVaultTransactor) WithdrawEther(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RocketVault.contract.Transact(opts, "withdrawEther", _amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 _amount) returns()
func (_RocketVault *RocketVaultSession) WithdrawEther(_amount *big.Int) (*types.Transaction, error) {
	return _RocketVault.Contract.WithdrawEther(&_RocketVault.TransactOpts, _amount)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x3bed33ce.
//
// Solidity: function withdrawEther(uint256 _amount) returns()
func (_RocketVault *RocketVaultTransactorSession) WithdrawEther(_amount *big.Int) (*types.Transaction, error) {
	return _RocketVault.Contract.WithdrawEther(&_RocketVault.TransactOpts, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _withdrawalAddress, address _tokenAddress, uint256 _amount) returns()
func (_RocketVault *RocketVaultTransactor) WithdrawToken(opts *bind.TransactOpts, _withdrawalAddress common.Address, _tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RocketVault.contract.Transact(opts, "withdrawToken", _withdrawalAddress, _tokenAddress, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _withdrawalAddress, address _tokenAddress, uint256 _amount) returns()
func (_RocketVault *RocketVaultSession) WithdrawToken(_withdrawalAddress common.Address, _tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RocketVault.Contract.WithdrawToken(&_RocketVault.TransactOpts, _withdrawalAddress, _tokenAddress, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _withdrawalAddress, address _tokenAddress, uint256 _amount) returns()
func (_RocketVault *RocketVaultTransactorSession) WithdrawToken(_withdrawalAddress common.Address, _tokenAddress common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RocketVault.Contract.WithdrawToken(&_RocketVault.TransactOpts, _withdrawalAddress, _tokenAddress, _amount)
}

// RocketVaultEtherDepositedIterator is returned from FilterEtherDeposited and is used to iterate over the raw logs and unpacked data for EtherDeposited events raised by the RocketVault contract.
type RocketVaultEtherDepositedIterator struct {
	Event *RocketVaultEtherDeposited // Event containing the contract specifics and raw log

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
func (it *RocketVaultEtherDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketVaultEtherDeposited)
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
		it.Event = new(RocketVaultEtherDeposited)
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
func (it *RocketVaultEtherDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketVaultEtherDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketVaultEtherDeposited represents a EtherDeposited event raised by the RocketVault contract.
type RocketVaultEtherDeposited struct {
	By     common.Hash
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherDeposited is a free log retrieval operation binding the contract event 0xbf25d6cb74c97a403cfab1c4c0abc9ffe3edd964de9924de0565f5ffe3d6ca79.
//
// Solidity: event EtherDeposited(string indexed by, uint256 amount, uint256 time)
func (_RocketVault *RocketVaultFilterer) FilterEtherDeposited(opts *bind.FilterOpts, by []string) (*RocketVaultEtherDepositedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _RocketVault.contract.FilterLogs(opts, "EtherDeposited", byRule)
	if err != nil {
		return nil, err
	}
	return &RocketVaultEtherDepositedIterator{contract: _RocketVault.contract, event: "EtherDeposited", logs: logs, sub: sub}, nil
}

// WatchEtherDeposited is a free log subscription operation binding the contract event 0xbf25d6cb74c97a403cfab1c4c0abc9ffe3edd964de9924de0565f5ffe3d6ca79.
//
// Solidity: event EtherDeposited(string indexed by, uint256 amount, uint256 time)
func (_RocketVault *RocketVaultFilterer) WatchEtherDeposited(opts *bind.WatchOpts, sink chan<- *RocketVaultEtherDeposited, by []string) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _RocketVault.contract.WatchLogs(opts, "EtherDeposited", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketVaultEtherDeposited)
				if err := _RocketVault.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
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

// ParseEtherDeposited is a log parse operation binding the contract event 0xbf25d6cb74c97a403cfab1c4c0abc9ffe3edd964de9924de0565f5ffe3d6ca79.
//
// Solidity: event EtherDeposited(string indexed by, uint256 amount, uint256 time)
func (_RocketVault *RocketVaultFilterer) ParseEtherDeposited(log types.Log) (*RocketVaultEtherDeposited, error) {
	event := new(RocketVaultEtherDeposited)
	if err := _RocketVault.contract.UnpackLog(event, "EtherDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketVaultEtherWithdrawnIterator is returned from FilterEtherWithdrawn and is used to iterate over the raw logs and unpacked data for EtherWithdrawn events raised by the RocketVault contract.
type RocketVaultEtherWithdrawnIterator struct {
	Event *RocketVaultEtherWithdrawn // Event containing the contract specifics and raw log

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
func (it *RocketVaultEtherWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketVaultEtherWithdrawn)
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
		it.Event = new(RocketVaultEtherWithdrawn)
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
func (it *RocketVaultEtherWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketVaultEtherWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketVaultEtherWithdrawn represents a EtherWithdrawn event raised by the RocketVault contract.
type RocketVaultEtherWithdrawn struct {
	By     common.Hash
	Amount *big.Int
	Time   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEtherWithdrawn is a free log retrieval operation binding the contract event 0xfc06bd2dc22238bad571c0432cfc04aee3d074be8cd974c9a9151e99df57e72b.
//
// Solidity: event EtherWithdrawn(string indexed by, uint256 amount, uint256 time)
func (_RocketVault *RocketVaultFilterer) FilterEtherWithdrawn(opts *bind.FilterOpts, by []string) (*RocketVaultEtherWithdrawnIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _RocketVault.contract.FilterLogs(opts, "EtherWithdrawn", byRule)
	if err != nil {
		return nil, err
	}
	return &RocketVaultEtherWithdrawnIterator{contract: _RocketVault.contract, event: "EtherWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEtherWithdrawn is a free log subscription operation binding the contract event 0xfc06bd2dc22238bad571c0432cfc04aee3d074be8cd974c9a9151e99df57e72b.
//
// Solidity: event EtherWithdrawn(string indexed by, uint256 amount, uint256 time)
func (_RocketVault *RocketVaultFilterer) WatchEtherWithdrawn(opts *bind.WatchOpts, sink chan<- *RocketVaultEtherWithdrawn, by []string) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _RocketVault.contract.WatchLogs(opts, "EtherWithdrawn", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketVaultEtherWithdrawn)
				if err := _RocketVault.contract.UnpackLog(event, "EtherWithdrawn", log); err != nil {
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

// ParseEtherWithdrawn is a log parse operation binding the contract event 0xfc06bd2dc22238bad571c0432cfc04aee3d074be8cd974c9a9151e99df57e72b.
//
// Solidity: event EtherWithdrawn(string indexed by, uint256 amount, uint256 time)
func (_RocketVault *RocketVaultFilterer) ParseEtherWithdrawn(log types.Log) (*RocketVaultEtherWithdrawn, error) {
	event := new(RocketVaultEtherWithdrawn)
	if err := _RocketVault.contract.UnpackLog(event, "EtherWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketVaultTokenBurnedIterator is returned from FilterTokenBurned and is used to iterate over the raw logs and unpacked data for TokenBurned events raised by the RocketVault contract.
type RocketVaultTokenBurnedIterator struct {
	Event *RocketVaultTokenBurned // Event containing the contract specifics and raw log

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
func (it *RocketVaultTokenBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketVaultTokenBurned)
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
		it.Event = new(RocketVaultTokenBurned)
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
func (it *RocketVaultTokenBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketVaultTokenBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketVaultTokenBurned represents a TokenBurned event raised by the RocketVault contract.
type RocketVaultTokenBurned struct {
	By           [32]byte
	TokenAddress common.Address
	Amount       *big.Int
	Time         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenBurned is a free log retrieval operation binding the contract event 0x85ed341c0b07bd4923f06588bb1b1c46a4637853245129a55ebaf4a0e30bbd4b.
//
// Solidity: event TokenBurned(bytes32 indexed by, address indexed tokenAddress, uint256 amount, uint256 time)
func (_RocketVault *RocketVaultFilterer) FilterTokenBurned(opts *bind.FilterOpts, by [][32]byte, tokenAddress []common.Address) (*RocketVaultTokenBurnedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _RocketVault.contract.FilterLogs(opts, "TokenBurned", byRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &RocketVaultTokenBurnedIterator{contract: _RocketVault.contract, event: "TokenBurned", logs: logs, sub: sub}, nil
}

// WatchTokenBurned is a free log subscription operation binding the contract event 0x85ed341c0b07bd4923f06588bb1b1c46a4637853245129a55ebaf4a0e30bbd4b.
//
// Solidity: event TokenBurned(bytes32 indexed by, address indexed tokenAddress, uint256 amount, uint256 time)
func (_RocketVault *RocketVaultFilterer) WatchTokenBurned(opts *bind.WatchOpts, sink chan<- *RocketVaultTokenBurned, by [][32]byte, tokenAddress []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _RocketVault.contract.WatchLogs(opts, "TokenBurned", byRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketVaultTokenBurned)
				if err := _RocketVault.contract.UnpackLog(event, "TokenBurned", log); err != nil {
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

// ParseTokenBurned is a log parse operation binding the contract event 0x85ed341c0b07bd4923f06588bb1b1c46a4637853245129a55ebaf4a0e30bbd4b.
//
// Solidity: event TokenBurned(bytes32 indexed by, address indexed tokenAddress, uint256 amount, uint256 time)
func (_RocketVault *RocketVaultFilterer) ParseTokenBurned(log types.Log) (*RocketVaultTokenBurned, error) {
	event := new(RocketVaultTokenBurned)
	if err := _RocketVault.contract.UnpackLog(event, "TokenBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketVaultTokenDepositedIterator is returned from FilterTokenDeposited and is used to iterate over the raw logs and unpacked data for TokenDeposited events raised by the RocketVault contract.
type RocketVaultTokenDepositedIterator struct {
	Event *RocketVaultTokenDeposited // Event containing the contract specifics and raw log

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
func (it *RocketVaultTokenDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketVaultTokenDeposited)
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
		it.Event = new(RocketVaultTokenDeposited)
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
func (it *RocketVaultTokenDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketVaultTokenDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketVaultTokenDeposited represents a TokenDeposited event raised by the RocketVault contract.
type RocketVaultTokenDeposited struct {
	By           [32]byte
	TokenAddress common.Address
	Amount       *big.Int
	Time         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenDeposited is a free log retrieval operation binding the contract event 0x14c4e7cf1c77c463baf4198cf43a79854a178a5641b6e59a109d7ed539f410aa.
//
// Solidity: event TokenDeposited(bytes32 indexed by, address indexed tokenAddress, uint256 amount, uint256 time)
func (_RocketVault *RocketVaultFilterer) FilterTokenDeposited(opts *bind.FilterOpts, by [][32]byte, tokenAddress []common.Address) (*RocketVaultTokenDepositedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _RocketVault.contract.FilterLogs(opts, "TokenDeposited", byRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &RocketVaultTokenDepositedIterator{contract: _RocketVault.contract, event: "TokenDeposited", logs: logs, sub: sub}, nil
}

// WatchTokenDeposited is a free log subscription operation binding the contract event 0x14c4e7cf1c77c463baf4198cf43a79854a178a5641b6e59a109d7ed539f410aa.
//
// Solidity: event TokenDeposited(bytes32 indexed by, address indexed tokenAddress, uint256 amount, uint256 time)
func (_RocketVault *RocketVaultFilterer) WatchTokenDeposited(opts *bind.WatchOpts, sink chan<- *RocketVaultTokenDeposited, by [][32]byte, tokenAddress []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _RocketVault.contract.WatchLogs(opts, "TokenDeposited", byRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketVaultTokenDeposited)
				if err := _RocketVault.contract.UnpackLog(event, "TokenDeposited", log); err != nil {
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

// ParseTokenDeposited is a log parse operation binding the contract event 0x14c4e7cf1c77c463baf4198cf43a79854a178a5641b6e59a109d7ed539f410aa.
//
// Solidity: event TokenDeposited(bytes32 indexed by, address indexed tokenAddress, uint256 amount, uint256 time)
func (_RocketVault *RocketVaultFilterer) ParseTokenDeposited(log types.Log) (*RocketVaultTokenDeposited, error) {
	event := new(RocketVaultTokenDeposited)
	if err := _RocketVault.contract.UnpackLog(event, "TokenDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketVaultTokenTransferIterator is returned from FilterTokenTransfer and is used to iterate over the raw logs and unpacked data for TokenTransfer events raised by the RocketVault contract.
type RocketVaultTokenTransferIterator struct {
	Event *RocketVaultTokenTransfer // Event containing the contract specifics and raw log

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
func (it *RocketVaultTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketVaultTokenTransfer)
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
		it.Event = new(RocketVaultTokenTransfer)
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
func (it *RocketVaultTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketVaultTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketVaultTokenTransfer represents a TokenTransfer event raised by the RocketVault contract.
type RocketVaultTokenTransfer struct {
	By           [32]byte
	To           [32]byte
	TokenAddress common.Address
	Amount       *big.Int
	Time         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenTransfer is a free log retrieval operation binding the contract event 0x11e32afbaec5786fc0b9e69b002583dedb32cae9499deeade686c878871b8ec0.
//
// Solidity: event TokenTransfer(bytes32 indexed by, bytes32 indexed to, address indexed tokenAddress, uint256 amount, uint256 time)
func (_RocketVault *RocketVaultFilterer) FilterTokenTransfer(opts *bind.FilterOpts, by [][32]byte, to [][32]byte, tokenAddress []common.Address) (*RocketVaultTokenTransferIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _RocketVault.contract.FilterLogs(opts, "TokenTransfer", byRule, toRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &RocketVaultTokenTransferIterator{contract: _RocketVault.contract, event: "TokenTransfer", logs: logs, sub: sub}, nil
}

// WatchTokenTransfer is a free log subscription operation binding the contract event 0x11e32afbaec5786fc0b9e69b002583dedb32cae9499deeade686c878871b8ec0.
//
// Solidity: event TokenTransfer(bytes32 indexed by, bytes32 indexed to, address indexed tokenAddress, uint256 amount, uint256 time)
func (_RocketVault *RocketVaultFilterer) WatchTokenTransfer(opts *bind.WatchOpts, sink chan<- *RocketVaultTokenTransfer, by [][32]byte, to [][32]byte, tokenAddress []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _RocketVault.contract.WatchLogs(opts, "TokenTransfer", byRule, toRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketVaultTokenTransfer)
				if err := _RocketVault.contract.UnpackLog(event, "TokenTransfer", log); err != nil {
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

// ParseTokenTransfer is a log parse operation binding the contract event 0x11e32afbaec5786fc0b9e69b002583dedb32cae9499deeade686c878871b8ec0.
//
// Solidity: event TokenTransfer(bytes32 indexed by, bytes32 indexed to, address indexed tokenAddress, uint256 amount, uint256 time)
func (_RocketVault *RocketVaultFilterer) ParseTokenTransfer(log types.Log) (*RocketVaultTokenTransfer, error) {
	event := new(RocketVaultTokenTransfer)
	if err := _RocketVault.contract.UnpackLog(event, "TokenTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RocketVaultTokenWithdrawnIterator is returned from FilterTokenWithdrawn and is used to iterate over the raw logs and unpacked data for TokenWithdrawn events raised by the RocketVault contract.
type RocketVaultTokenWithdrawnIterator struct {
	Event *RocketVaultTokenWithdrawn // Event containing the contract specifics and raw log

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
func (it *RocketVaultTokenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RocketVaultTokenWithdrawn)
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
		it.Event = new(RocketVaultTokenWithdrawn)
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
func (it *RocketVaultTokenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RocketVaultTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RocketVaultTokenWithdrawn represents a TokenWithdrawn event raised by the RocketVault contract.
type RocketVaultTokenWithdrawn struct {
	By           [32]byte
	TokenAddress common.Address
	Amount       *big.Int
	Time         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawn is a free log retrieval operation binding the contract event 0xacd47105acf174819374e9d73c4c60bdaee9d4c7939b88a4cebb4dfe0500fbd8.
//
// Solidity: event TokenWithdrawn(bytes32 indexed by, address indexed tokenAddress, uint256 amount, uint256 time)
func (_RocketVault *RocketVaultFilterer) FilterTokenWithdrawn(opts *bind.FilterOpts, by [][32]byte, tokenAddress []common.Address) (*RocketVaultTokenWithdrawnIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _RocketVault.contract.FilterLogs(opts, "TokenWithdrawn", byRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &RocketVaultTokenWithdrawnIterator{contract: _RocketVault.contract, event: "TokenWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawn is a free log subscription operation binding the contract event 0xacd47105acf174819374e9d73c4c60bdaee9d4c7939b88a4cebb4dfe0500fbd8.
//
// Solidity: event TokenWithdrawn(bytes32 indexed by, address indexed tokenAddress, uint256 amount, uint256 time)
func (_RocketVault *RocketVaultFilterer) WatchTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *RocketVaultTokenWithdrawn, by [][32]byte, tokenAddress []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _RocketVault.contract.WatchLogs(opts, "TokenWithdrawn", byRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RocketVaultTokenWithdrawn)
				if err := _RocketVault.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
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

// ParseTokenWithdrawn is a log parse operation binding the contract event 0xacd47105acf174819374e9d73c4c60bdaee9d4c7939b88a4cebb4dfe0500fbd8.
//
// Solidity: event TokenWithdrawn(bytes32 indexed by, address indexed tokenAddress, uint256 amount, uint256 time)
func (_RocketVault *RocketVaultFilterer) ParseTokenWithdrawn(log types.Log) (*RocketVaultTokenWithdrawn, error) {
	event := new(RocketVaultTokenWithdrawn)
	if err := _RocketVault.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
