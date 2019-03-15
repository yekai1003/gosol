// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// PdbankABI is the input ABI used to generate the binding from.
const PdbankABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PdbankBin is the compiled bytecode used for deploying new contracts.
const PdbankBin = `0x608060405234801561001057600080fd5b50600080546001600160a01b031916331781556002556101db806100356000396000f3fe60806040526004361061004a5760003560e01c80631a39d8ef1461004f57806327e235e3146100765780632e1a7d4d146100a95780638da5cb5b146100c8578063d0e30db0146100f9575b600080fd5b34801561005b57600080fd5b50610064610101565b60408051918252519081900360200190f35b34801561008257600080fd5b506100646004803603602081101561009957600080fd5b50356001600160a01b0316610107565b6100c6600480360360208110156100bf57600080fd5b5035610119565b005b3480156100d457600080fd5b506100dd61017d565b604080516001600160a01b039092168252519081900360200190f35b6100c661018c565b60025481565b60016020526000908152604090205481565b3360009081526001602052604090205481101561017a5733600081815260016020526040808220805485900390555183156108fc0291849190818181858888f1935050505015801561016f573d6000803e3d6000fd5b506002805482900390555b50565b6000546001600160a01b031681565b60028054349081019091553360009081526001602052604090208054909101905556fea165627a7a72305820731c2a8b00f69921a1d9e2bfddbcf0b6b9e78887870fc0305f2f66d13d09a6860029`

// DeployPdbank deploys a new Ethereum contract, binding an instance of Pdbank to it.
func DeployPdbank(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pdbank, error) {
	parsed, err := abi.JSON(strings.NewReader(PdbankABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PdbankBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pdbank{PdbankCaller: PdbankCaller{contract: contract}, PdbankTransactor: PdbankTransactor{contract: contract}, PdbankFilterer: PdbankFilterer{contract: contract}}, nil
}

// Pdbank is an auto generated Go binding around an Ethereum contract.
type Pdbank struct {
	PdbankCaller     // Read-only binding to the contract
	PdbankTransactor // Write-only binding to the contract
	PdbankFilterer   // Log filterer for contract events
}

// PdbankCaller is an auto generated read-only Go binding around an Ethereum contract.
type PdbankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PdbankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PdbankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PdbankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PdbankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PdbankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PdbankSession struct {
	Contract     *Pdbank           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PdbankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PdbankCallerSession struct {
	Contract *PdbankCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PdbankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PdbankTransactorSession struct {
	Contract     *PdbankTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PdbankRaw is an auto generated low-level Go binding around an Ethereum contract.
type PdbankRaw struct {
	Contract *Pdbank // Generic contract binding to access the raw methods on
}

// PdbankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PdbankCallerRaw struct {
	Contract *PdbankCaller // Generic read-only contract binding to access the raw methods on
}

// PdbankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PdbankTransactorRaw struct {
	Contract *PdbankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPdbank creates a new instance of Pdbank, bound to a specific deployed contract.
func NewPdbank(address common.Address, backend bind.ContractBackend) (*Pdbank, error) {
	contract, err := bindPdbank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pdbank{PdbankCaller: PdbankCaller{contract: contract}, PdbankTransactor: PdbankTransactor{contract: contract}, PdbankFilterer: PdbankFilterer{contract: contract}}, nil
}

// NewPdbankCaller creates a new read-only instance of Pdbank, bound to a specific deployed contract.
func NewPdbankCaller(address common.Address, caller bind.ContractCaller) (*PdbankCaller, error) {
	contract, err := bindPdbank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PdbankCaller{contract: contract}, nil
}

// NewPdbankTransactor creates a new write-only instance of Pdbank, bound to a specific deployed contract.
func NewPdbankTransactor(address common.Address, transactor bind.ContractTransactor) (*PdbankTransactor, error) {
	contract, err := bindPdbank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PdbankTransactor{contract: contract}, nil
}

// NewPdbankFilterer creates a new log filterer instance of Pdbank, bound to a specific deployed contract.
func NewPdbankFilterer(address common.Address, filterer bind.ContractFilterer) (*PdbankFilterer, error) {
	contract, err := bindPdbank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PdbankFilterer{contract: contract}, nil
}

// bindPdbank binds a generic wrapper to an already deployed contract.
func bindPdbank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PdbankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pdbank *PdbankRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pdbank.Contract.PdbankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pdbank *PdbankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pdbank.Contract.PdbankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pdbank *PdbankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pdbank.Contract.PdbankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pdbank *PdbankCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pdbank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pdbank *PdbankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pdbank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pdbank *PdbankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pdbank.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_Pdbank *PdbankCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pdbank.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_Pdbank *PdbankSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Pdbank.Contract.Balances(&_Pdbank.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_Pdbank *PdbankCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Pdbank.Contract.Balances(&_Pdbank.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pdbank *PdbankCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pdbank.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pdbank *PdbankSession) Owner() (common.Address, error) {
	return _Pdbank.Contract.Owner(&_Pdbank.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pdbank *PdbankCallerSession) Owner() (common.Address, error) {
	return _Pdbank.Contract.Owner(&_Pdbank.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() constant returns(uint256)
func (_Pdbank *PdbankCaller) TotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pdbank.contract.Call(opts, out, "totalAmount")
	return *ret0, err
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() constant returns(uint256)
func (_Pdbank *PdbankSession) TotalAmount() (*big.Int, error) {
	return _Pdbank.Contract.TotalAmount(&_Pdbank.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() constant returns(uint256)
func (_Pdbank *PdbankCallerSession) TotalAmount() (*big.Int, error) {
	return _Pdbank.Contract.TotalAmount(&_Pdbank.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Pdbank *PdbankTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pdbank.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Pdbank *PdbankSession) Deposit() (*types.Transaction, error) {
	return _Pdbank.Contract.Deposit(&_Pdbank.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Pdbank *PdbankTransactorSession) Deposit() (*types.Transaction, error) {
	return _Pdbank.Contract.Deposit(&_Pdbank.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(_amount uint256) returns()
func (_Pdbank *PdbankTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Pdbank.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(_amount uint256) returns()
func (_Pdbank *PdbankSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Pdbank.Contract.Withdraw(&_Pdbank.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(_amount uint256) returns()
func (_Pdbank *PdbankTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Pdbank.Contract.Withdraw(&_Pdbank.TransactOpts, _amount)
}
