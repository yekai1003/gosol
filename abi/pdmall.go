// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pdmall

import (
	"math/big"
	"strings"
)

// PdmallABI is the input ABI used to generate the binding from.
const PdmallABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// PdmallBin is the compiled bytecode used for deploying new contracts.
const PdmallBin = `0x608060405234801561001057600080fd5b50600080546001600160a01b0319163317815560025561015e806100356000396000f3fe60806040526004361061003f5760003560e01c80631a39d8ef146100415780633ccfd60b146100685780638da5cb5b14610070578063d0e30db0146100a1575b005b34801561004d57600080fd5b506100566100a9565b60408051918252519081900360200190f35b61003f6100af565b34801561007c57600080fd5b50610085610100565b604080516001600160a01b039092168252519081900360200190f35b61003f61010f565b60025481565b600254156100fe5760028054600091829055815460405191926001600160a01b039091169183156108fc0291849190818181858888f193505050501580156100fb573d6000803e3d6000fd5b50505b565b6000546001600160a01b031681565b60028054349081019091553360009081526001602052604090208054909101905556fea165627a7a723058206b371b2130a4d099df45420ca363280eb2447d7dd2c3467439f59c93dfa10ca70029`

// DeployPdmall deploys a new Ethereum contract, binding an instance of Pdmall to it.
func DeployPdmall(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pdmall, error) {
	parsed, err := abi.JSON(strings.NewReader(PdmallABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PdmallBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pdmall{PdmallCaller: PdmallCaller{contract: contract}, PdmallTransactor: PdmallTransactor{contract: contract}, PdmallFilterer: PdmallFilterer{contract: contract}}, nil
}

// Pdmall is an auto generated Go binding around an Ethereum contract.
type Pdmall struct {
	PdmallCaller     // Read-only binding to the contract
	PdmallTransactor // Write-only binding to the contract
	PdmallFilterer   // Log filterer for contract events
}

// PdmallCaller is an auto generated read-only Go binding around an Ethereum contract.
type PdmallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PdmallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PdmallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PdmallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PdmallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PdmallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PdmallSession struct {
	Contract     *Pdmall           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PdmallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PdmallCallerSession struct {
	Contract *PdmallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PdmallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PdmallTransactorSession struct {
	Contract     *PdmallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PdmallRaw is an auto generated low-level Go binding around an Ethereum contract.
type PdmallRaw struct {
	Contract *Pdmall // Generic contract binding to access the raw methods on
}

// PdmallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PdmallCallerRaw struct {
	Contract *PdmallCaller // Generic read-only contract binding to access the raw methods on
}

// PdmallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PdmallTransactorRaw struct {
	Contract *PdmallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPdmall creates a new instance of Pdmall, bound to a specific deployed contract.
func NewPdmall(address common.Address, backend bind.ContractBackend) (*Pdmall, error) {
	contract, err := bindPdmall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pdmall{PdmallCaller: PdmallCaller{contract: contract}, PdmallTransactor: PdmallTransactor{contract: contract}, PdmallFilterer: PdmallFilterer{contract: contract}}, nil
}

// NewPdmallCaller creates a new read-only instance of Pdmall, bound to a specific deployed contract.
func NewPdmallCaller(address common.Address, caller bind.ContractCaller) (*PdmallCaller, error) {
	contract, err := bindPdmall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PdmallCaller{contract: contract}, nil
}

// NewPdmallTransactor creates a new write-only instance of Pdmall, bound to a specific deployed contract.
func NewPdmallTransactor(address common.Address, transactor bind.ContractTransactor) (*PdmallTransactor, error) {
	contract, err := bindPdmall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PdmallTransactor{contract: contract}, nil
}

// NewPdmallFilterer creates a new log filterer instance of Pdmall, bound to a specific deployed contract.
func NewPdmallFilterer(address common.Address, filterer bind.ContractFilterer) (*PdmallFilterer, error) {
	contract, err := bindPdmall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PdmallFilterer{contract: contract}, nil
}

// bindPdmall binds a generic wrapper to an already deployed contract.
func bindPdmall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PdmallABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pdmall *PdmallRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pdmall.Contract.PdmallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pdmall *PdmallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pdmall.Contract.PdmallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pdmall *PdmallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pdmall.Contract.PdmallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pdmall *PdmallCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pdmall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pdmall *PdmallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pdmall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pdmall *PdmallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pdmall.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pdmall *PdmallCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pdmall.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pdmall *PdmallSession) Owner() (common.Address, error) {
	return _Pdmall.Contract.Owner(&_Pdmall.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pdmall *PdmallCallerSession) Owner() (common.Address, error) {
	return _Pdmall.Contract.Owner(&_Pdmall.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() constant returns(uint256)
func (_Pdmall *PdmallCaller) TotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pdmall.contract.Call(opts, out, "totalAmount")
	return *ret0, err
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() constant returns(uint256)
func (_Pdmall *PdmallSession) TotalAmount() (*big.Int, error) {
	return _Pdmall.Contract.TotalAmount(&_Pdmall.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() constant returns(uint256)
func (_Pdmall *PdmallCallerSession) TotalAmount() (*big.Int, error) {
	return _Pdmall.Contract.TotalAmount(&_Pdmall.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Pdmall *PdmallTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pdmall.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Pdmall *PdmallSession) Deposit() (*types.Transaction, error) {
	return _Pdmall.Contract.Deposit(&_Pdmall.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Pdmall *PdmallTransactorSession) Deposit() (*types.Transaction, error) {
	return _Pdmall.Contract.Deposit(&_Pdmall.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Pdmall *PdmallTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pdmall.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Pdmall *PdmallSession) Withdraw() (*types.Transaction, error) {
	return _Pdmall.Contract.Withdraw(&_Pdmall.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Pdmall *PdmallTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Pdmall.Contract.Withdraw(&_Pdmall.TransactOpts)
}
