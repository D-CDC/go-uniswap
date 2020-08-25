// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenm

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TokenmABI is the input ABI used to generate the binding from.
const TokenmABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"_users\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TokenmBin is the compiled bytecode used for deploying new contracts.
var TokenmBin = "0x60806040523480156100115760006000fd5b505b6012600a0a6402540be40002600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055505b610070565b6114d08061007f6000396000f3fe60806040523480156100115760006000fd5b50600436106100a35760003560e01c806370a082311161006757806370a082311461025757806388d695b2146102b057806395d89b4114610423578063a9059cbb146104a7578063dd62ed3e1461050e576100a3565b806306fdde03146100a9578063095ea7b31461012d57806318160ddd1461019457806323b872dd146101b2578063313ce56714610239576100a3565b60006000fd5b6100b1610587565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100f25780820151818401525b6020810190506100d6565b50505050905090810190601f16801561011f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61017a600480360360408110156101445760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506105c3565b604051808215151515815260200191505060405180910390f35b61019c6106c5565b6040518082815260200191505060405180910390f35b61021f600480360360608110156101c95760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506106d4565b604051808215151515815260200191505060405180910390f35b610241610bad565b6040518082815260200191505060405180910390f35b61029a6004803603602081101561026e5760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610bb2565b6040518082815260200191505060405180910390f35b610409600480360360408110156102c75760006000fd5b81019080803590602001906401000000008111156102e55760006000fd5b8201836020820111156102f85760006000fd5b8035906020019184602083028401116401000000008311171561031b5760006000fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050909091929090919290803590602001906401000000008111156103805760006000fd5b8201836020820111156103935760006000fd5b803590602001918460208302840111640100000000831117156103b65760006000fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050909091929090919290505050610c06565b604051808215151515815260200191505060405180910390f35b61042b611088565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561046c5780820151818401525b602081019050610450565b50505050905090810190601f1680156104995780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6104f4600480360360408110156104be5760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506110c4565b604051808215151515815260200191505060405180910390f35b610571600480360360408110156105255760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506113c2565b6040518082815260200191505060405180910390f35b6040518060400160405280601281526020017f4d6172636f506f6c6f2050726f746f636f6c000000000000000000000000000081526020015081565b600081600160005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a3600190506106bf565b92915050565b6012600a0a6402540be4000281565b6000600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561077e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6e6f7420656e6f7567682062616c616e6365202100000000000000000000000081526020015060200191505060405180910390fd5b81600060005060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505410158015610858575081600160005060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505410155b15156108cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601a8152602001807f6e6f7420656e6f75676820616c6c6f7765642062616c616e636500000000000081526020015060200191505060405180910390fd5b61096782600160005060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505461145790919063ffffffff16565b600160005060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050819090905550610a4a82600060005060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505461145790919063ffffffff16565b600060005060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050819090905550610aed82600060005060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505461147590919063ffffffff16565b600060005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a360019050610ba6565b9392505050565b601281565b6000600060005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050549050610c01565b919050565b600081518351141515610c84576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f6e6f742073616d65206c656e677468000000000000000000000000000000000081526020015060200191505060405180910390fd5b6000600090505b83518160ff16101561107857600073ffffffffffffffffffffffffffffffffffffffff16848260ff16815181101515610cc057fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1614151515610d57576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f61646472657373206973207a65726f000000000000000000000000000000000081526020015060200191505060405180910390fd5b828160ff16815181101515610d6857fe5b6020026020010151600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505410151515610e2f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6e6f7420656e6f7567682062616c616e6365202100000000000000000000000081526020015060200191505060405180910390fd5b610e9f838260ff16815181101515610e4357fe5b6020026020010151600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505461145790919063ffffffff16565b600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050819090905550610f72838260ff16815181101515610efe57fe5b602002602001015160006000506000878560ff16815181101515610f1e57fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505461147590919063ffffffff16565b60006000506000868460ff16815181101515610f8a57fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050819090905550838160ff16815181101515610fe657fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef858460ff1681518110151561104d57fe5b60200260200101516040518082815260200191505060405180910390a35b8080600101915050610c8b565b5060019050611082565b92915050565b6040518060400160405280600381526020017f4d4150000000000000000000000000000000000000000000000000000000000081526020015081565b6000600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561116e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260148152602001807f6e6f7420656e6f7567682062616c616e6365202100000000000000000000000081526020015060200191505060405180910390fd5b81600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505410151515611208576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260008152602001805060200191505060405180910390fd5b61126082600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505461145790919063ffffffff16565b600060005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005081909090555061130382600060005060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505461147590919063ffffffff16565b600060005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190506113bc565b92915050565b6000600160005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050549050611451565b92915050565b600082821115151561146557fe5b818303905061146f565b92915050565b60006000828401905083811015151561148a57fe5b8091505061149456505b9291505056fea264697066735822122081b97a7de624abcf34d0954078b534a66d7901e9ea1c883e87ed5496def596a764736f6c63430006000033"

// DeployTokenm deploys a new Ethereum contract, binding an instance of Tokenm to it.
func DeployTokenm(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tokenm, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenmABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenmBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tokenm{TokenmCaller: TokenmCaller{contract: contract}, TokenmTransactor: TokenmTransactor{contract: contract}, TokenmFilterer: TokenmFilterer{contract: contract}}, nil
}

// Tokenm is an auto generated Go binding around an Ethereum contract.
type Tokenm struct {
	TokenmCaller     // Read-only binding to the contract
	TokenmTransactor // Write-only binding to the contract
	TokenmFilterer   // Log filterer for contract events
}

// TokenmCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenmSession struct {
	Contract     *Tokenm           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenmCallerSession struct {
	Contract *TokenmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenmTransactorSession struct {
	Contract     *TokenmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenmRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenmRaw struct {
	Contract *Tokenm // Generic contract binding to access the raw methods on
}

// TokenmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenmCallerRaw struct {
	Contract *TokenmCaller // Generic read-only contract binding to access the raw methods on
}

// TokenmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenmTransactorRaw struct {
	Contract *TokenmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenm creates a new instance of Tokenm, bound to a specific deployed contract.
func NewTokenm(address common.Address, backend bind.ContractBackend) (*Tokenm, error) {
	contract, err := bindTokenm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokenm{TokenmCaller: TokenmCaller{contract: contract}, TokenmTransactor: TokenmTransactor{contract: contract}, TokenmFilterer: TokenmFilterer{contract: contract}}, nil
}

// NewTokenmCaller creates a new read-only instance of Tokenm, bound to a specific deployed contract.
func NewTokenmCaller(address common.Address, caller bind.ContractCaller) (*TokenmCaller, error) {
	contract, err := bindTokenm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenmCaller{contract: contract}, nil
}

// NewTokenmTransactor creates a new write-only instance of Tokenm, bound to a specific deployed contract.
func NewTokenmTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenmTransactor, error) {
	contract, err := bindTokenm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenmTransactor{contract: contract}, nil
}

// NewTokenmFilterer creates a new log filterer instance of Tokenm, bound to a specific deployed contract.
func NewTokenmFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenmFilterer, error) {
	contract, err := bindTokenm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenmFilterer{contract: contract}, nil
}

// bindTokenm binds a generic wrapper to an already deployed contract.
func bindTokenm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenmABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenm *TokenmRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tokenm.Contract.TokenmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenm *TokenmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenm.Contract.TokenmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenm *TokenmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenm.Contract.TokenmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenm *TokenmCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tokenm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenm *TokenmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenm *TokenmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenm.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Tokenm *TokenmCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Tokenm.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Tokenm *TokenmSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Tokenm.Contract.Allowance(&_Tokenm.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Tokenm *TokenmCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Tokenm.Contract.Allowance(&_Tokenm.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Tokenm *TokenmCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Tokenm.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Tokenm *TokenmSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Tokenm.Contract.BalanceOf(&_Tokenm.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_Tokenm *TokenmCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Tokenm.Contract.BalanceOf(&_Tokenm.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Tokenm *TokenmCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Tokenm.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Tokenm *TokenmSession) Decimals() (*big.Int, error) {
	return _Tokenm.Contract.Decimals(&_Tokenm.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Tokenm *TokenmCallerSession) Decimals() (*big.Int, error) {
	return _Tokenm.Contract.Decimals(&_Tokenm.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokenm *TokenmCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Tokenm.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokenm *TokenmSession) Name() (string, error) {
	return _Tokenm.Contract.Name(&_Tokenm.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokenm *TokenmCallerSession) Name() (string, error) {
	return _Tokenm.Contract.Name(&_Tokenm.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokenm *TokenmCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Tokenm.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokenm *TokenmSession) Symbol() (string, error) {
	return _Tokenm.Contract.Symbol(&_Tokenm.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokenm *TokenmCallerSession) Symbol() (string, error) {
	return _Tokenm.Contract.Symbol(&_Tokenm.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokenm *TokenmCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Tokenm.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokenm *TokenmSession) TotalSupply() (*big.Int, error) {
	return _Tokenm.Contract.TotalSupply(&_Tokenm.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokenm *TokenmCallerSession) TotalSupply() (*big.Int, error) {
	return _Tokenm.Contract.TotalSupply(&_Tokenm.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Tokenm *TokenmTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tokenm.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Tokenm *TokenmSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tokenm.Contract.Approve(&_Tokenm.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Tokenm *TokenmTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tokenm.Contract.Approve(&_Tokenm.TransactOpts, _spender, _value)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x88d695b2.
//
// Solidity: function batchTransfer(address[] _users, uint256[] _amounts) returns(bool)
func (_Tokenm *TokenmTransactor) BatchTransfer(opts *bind.TransactOpts, _users []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Tokenm.contract.Transact(opts, "batchTransfer", _users, _amounts)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x88d695b2.
//
// Solidity: function batchTransfer(address[] _users, uint256[] _amounts) returns(bool)
func (_Tokenm *TokenmSession) BatchTransfer(_users []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Tokenm.Contract.BatchTransfer(&_Tokenm.TransactOpts, _users, _amounts)
}

// BatchTransfer is a paid mutator transaction binding the contract method 0x88d695b2.
//
// Solidity: function batchTransfer(address[] _users, uint256[] _amounts) returns(bool)
func (_Tokenm *TokenmTransactorSession) BatchTransfer(_users []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Tokenm.Contract.BatchTransfer(&_Tokenm.TransactOpts, _users, _amounts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Tokenm *TokenmTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tokenm.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Tokenm *TokenmSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tokenm.Contract.Transfer(&_Tokenm.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Tokenm *TokenmTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tokenm.Contract.Transfer(&_Tokenm.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Tokenm *TokenmTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tokenm.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Tokenm *TokenmSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tokenm.Contract.TransferFrom(&_Tokenm.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool success)
func (_Tokenm *TokenmTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tokenm.Contract.TransferFrom(&_Tokenm.TransactOpts, _from, _to, _value)
}

// TokenmApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Tokenm contract.
type TokenmApprovalIterator struct {
	Event *TokenmApproval // Event containing the contract specifics and raw log

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
func (it *TokenmApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenmApproval)
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
		it.Event = new(TokenmApproval)
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
func (it *TokenmApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenmApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenmApproval represents a Approval event raised by the Tokenm contract.
type TokenmApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Tokenm *TokenmFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*TokenmApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Tokenm.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenmApprovalIterator{contract: _Tokenm.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Tokenm *TokenmFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenmApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Tokenm.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenmApproval)
				if err := _Tokenm.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Tokenm *TokenmFilterer) ParseApproval(log types.Log) (*TokenmApproval, error) {
	event := new(TokenmApproval)
	if err := _Tokenm.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenmTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Tokenm contract.
type TokenmTransferIterator struct {
	Event *TokenmTransfer // Event containing the contract specifics and raw log

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
func (it *TokenmTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenmTransfer)
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
		it.Event = new(TokenmTransfer)
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
func (it *TokenmTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenmTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenmTransfer represents a Transfer event raised by the Tokenm contract.
type TokenmTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Tokenm *TokenmFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*TokenmTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Tokenm.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &TokenmTransferIterator{contract: _Tokenm.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Tokenm *TokenmFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenmTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Tokenm.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenmTransfer)
				if err := _Tokenm.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Tokenm *TokenmFilterer) ParseTransfer(log types.Log) (*TokenmTransfer, error) {
	event := new(TokenmTransfer)
	if err := _Tokenm.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
