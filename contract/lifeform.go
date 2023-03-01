// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ICartoonMintRuleMintRule is an auto generated low-level Go binding around an user-defined struct.
type ICartoonMintRuleMintRule struct {
	MintRule         common.Address
	UdIndex          *big.Int
	StakeErc20       common.Address
	StakeErc20Amount *big.Int
	CostErc20        common.Address
	CostErc20Amount  *big.Int
	LimitTimes       *big.Int
	MintType         *big.Int
	SignCode         [32]byte
	WlSignature      []byte
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"SIGNER\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakeErc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"costErc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"costAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"eUpdateSigner\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TYPE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_SIGNER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_costAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_costErc20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_mintTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_onceSignCode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_stakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_stakeErc20\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rule\",\"type\":\"address\"}],\"name\":\"addMintRule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMintRules\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getTheMintTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"mintRule\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"udIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakeErc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeErc20Amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"costErc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"costErc20Amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitTimes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintType\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"signCode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"wlSignature\",\"type\":\"bytes\"}],\"internalType\":\"structICartoonMintRule.MintRule\",\"name\":\"mintData\",\"type\":\"tuple\"}],\"name\":\"hashCondition\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"mintRule\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"udIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakeErc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeErc20Amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"costErc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"costErc20Amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitTimes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintType\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"signCode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"wlSignature\",\"type\":\"bytes\"}],\"internalType\":\"structICartoonMintRule.MintRule\",\"name\":\"mintData\",\"type\":\"tuple\"}],\"name\":\"hashDigest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"signCode\",\"type\":\"bytes32\"}],\"name\":\"hashWhiteList\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signCode\",\"type\":\"bytes32\"}],\"name\":\"isExistSignCode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"mintRule\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"udIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakeErc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeErc20Amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"costErc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"costErc20Amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitTimes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintType\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"signCode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"wlSignature\",\"type\":\"bytes\"}],\"internalType\":\"structICartoonMintRule.MintRule\",\"name\":\"mintData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"dataSignature\",\"type\":\"bytes\"}],\"name\":\"mintAvatar721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rule\",\"type\":\"address\"}],\"name\":\"removeMintRule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"setOnceSignCode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"costErc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"costAmount\",\"type\":\"uint256\"}],\"name\":\"updateCostErc20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"SIGNER\",\"type\":\"address\"}],\"name\":\"updateSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakeErc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"updateStakeErc20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"mintRule\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"udIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakeErc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeErc20Amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"costErc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"costErc20Amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitTimes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintType\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"signCode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"wlSignature\",\"type\":\"bytes\"}],\"internalType\":\"structICartoonMintRule.MintRule\",\"name\":\"mintData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataSignature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"mintRule\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"udIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakeErc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeErc20Amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"costErc20\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"costErc20Amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitTimes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintType\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"signCode\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"wlSignature\",\"type\":\"bytes\"}],\"internalType\":\"structICartoonMintRule.MintRule\",\"name\":\"mintData\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"verifyCondition\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verifySignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Contract *ContractCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Contract *ContractSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Contract.Contract.DOMAINSEPARATOR(&_Contract.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Contract *ContractCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Contract.Contract.DOMAINSEPARATOR(&_Contract.CallOpts)
}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc49f91d3.
//
// Solidity: function EIP712DOMAIN_TYPEHASH() view returns(bytes32)
func (_Contract *ContractCaller) EIP712DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "EIP712DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc49f91d3.
//
// Solidity: function EIP712DOMAIN_TYPEHASH() view returns(bytes32)
func (_Contract *ContractSession) EIP712DOMAINTYPEHASH() ([32]byte, error) {
	return _Contract.Contract.EIP712DOMAINTYPEHASH(&_Contract.CallOpts)
}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc49f91d3.
//
// Solidity: function EIP712DOMAIN_TYPEHASH() view returns(bytes32)
func (_Contract *ContractCallerSession) EIP712DOMAINTYPEHASH() ([32]byte, error) {
	return _Contract.Contract.EIP712DOMAINTYPEHASH(&_Contract.CallOpts)
}

// TYPEHASH is a free data retrieval call binding the contract method 0x64d4c819.
//
// Solidity: function TYPE_HASH() view returns(bytes32)
func (_Contract *ContractCaller) TYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "TYPE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TYPEHASH is a free data retrieval call binding the contract method 0x64d4c819.
//
// Solidity: function TYPE_HASH() view returns(bytes32)
func (_Contract *ContractSession) TYPEHASH() ([32]byte, error) {
	return _Contract.Contract.TYPEHASH(&_Contract.CallOpts)
}

// TYPEHASH is a free data retrieval call binding the contract method 0x64d4c819.
//
// Solidity: function TYPE_HASH() view returns(bytes32)
func (_Contract *ContractCallerSession) TYPEHASH() ([32]byte, error) {
	return _Contract.Contract.TYPEHASH(&_Contract.CallOpts)
}

// SIGNER is a free data retrieval call binding the contract method 0x6bc054f1.
//
// Solidity: function _SIGNER() view returns(address)
func (_Contract *ContractCaller) SIGNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_SIGNER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SIGNER is a free data retrieval call binding the contract method 0x6bc054f1.
//
// Solidity: function _SIGNER() view returns(address)
func (_Contract *ContractSession) SIGNER() (common.Address, error) {
	return _Contract.Contract.SIGNER(&_Contract.CallOpts)
}

// SIGNER is a free data retrieval call binding the contract method 0x6bc054f1.
//
// Solidity: function _SIGNER() view returns(address)
func (_Contract *ContractCallerSession) SIGNER() (common.Address, error) {
	return _Contract.Contract.SIGNER(&_Contract.CallOpts)
}

// CostAmount is a free data retrieval call binding the contract method 0xce580ac9.
//
// Solidity: function _costAmount() view returns(uint256)
func (_Contract *ContractCaller) CostAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_costAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CostAmount is a free data retrieval call binding the contract method 0xce580ac9.
//
// Solidity: function _costAmount() view returns(uint256)
func (_Contract *ContractSession) CostAmount() (*big.Int, error) {
	return _Contract.Contract.CostAmount(&_Contract.CallOpts)
}

// CostAmount is a free data retrieval call binding the contract method 0xce580ac9.
//
// Solidity: function _costAmount() view returns(uint256)
func (_Contract *ContractCallerSession) CostAmount() (*big.Int, error) {
	return _Contract.Contract.CostAmount(&_Contract.CallOpts)
}

// CostErc20 is a free data retrieval call binding the contract method 0xf26480ec.
//
// Solidity: function _costErc20() view returns(address)
func (_Contract *ContractCaller) CostErc20(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_costErc20")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CostErc20 is a free data retrieval call binding the contract method 0xf26480ec.
//
// Solidity: function _costErc20() view returns(address)
func (_Contract *ContractSession) CostErc20() (common.Address, error) {
	return _Contract.Contract.CostErc20(&_Contract.CallOpts)
}

// CostErc20 is a free data retrieval call binding the contract method 0xf26480ec.
//
// Solidity: function _costErc20() view returns(address)
func (_Contract *ContractCallerSession) CostErc20() (common.Address, error) {
	return _Contract.Contract.CostErc20(&_Contract.CallOpts)
}

// MintTimes is a free data retrieval call binding the contract method 0x91ce3f50.
//
// Solidity: function _mintTimes(uint256 , address ) view returns(uint256)
func (_Contract *ContractCaller) MintTimes(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_mintTimes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintTimes is a free data retrieval call binding the contract method 0x91ce3f50.
//
// Solidity: function _mintTimes(uint256 , address ) view returns(uint256)
func (_Contract *ContractSession) MintTimes(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Contract.Contract.MintTimes(&_Contract.CallOpts, arg0, arg1)
}

// MintTimes is a free data retrieval call binding the contract method 0x91ce3f50.
//
// Solidity: function _mintTimes(uint256 , address ) view returns(uint256)
func (_Contract *ContractCallerSession) MintTimes(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Contract.Contract.MintTimes(&_Contract.CallOpts, arg0, arg1)
}

// OnceSignCode is a free data retrieval call binding the contract method 0x01e8cd5d.
//
// Solidity: function _onceSignCode() view returns(bool)
func (_Contract *ContractCaller) OnceSignCode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_onceSignCode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OnceSignCode is a free data retrieval call binding the contract method 0x01e8cd5d.
//
// Solidity: function _onceSignCode() view returns(bool)
func (_Contract *ContractSession) OnceSignCode() (bool, error) {
	return _Contract.Contract.OnceSignCode(&_Contract.CallOpts)
}

// OnceSignCode is a free data retrieval call binding the contract method 0x01e8cd5d.
//
// Solidity: function _onceSignCode() view returns(bool)
func (_Contract *ContractCallerSession) OnceSignCode() (bool, error) {
	return _Contract.Contract.OnceSignCode(&_Contract.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x061cc4a1.
//
// Solidity: function _stakeAmount() view returns(uint256)
func (_Contract *ContractCaller) StakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_stakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeAmount is a free data retrieval call binding the contract method 0x061cc4a1.
//
// Solidity: function _stakeAmount() view returns(uint256)
func (_Contract *ContractSession) StakeAmount() (*big.Int, error) {
	return _Contract.Contract.StakeAmount(&_Contract.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x061cc4a1.
//
// Solidity: function _stakeAmount() view returns(uint256)
func (_Contract *ContractCallerSession) StakeAmount() (*big.Int, error) {
	return _Contract.Contract.StakeAmount(&_Contract.CallOpts)
}

// StakeErc20 is a free data retrieval call binding the contract method 0xf0ec7fd7.
//
// Solidity: function _stakeErc20() view returns(address)
func (_Contract *ContractCaller) StakeErc20(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_stakeErc20")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeErc20 is a free data retrieval call binding the contract method 0xf0ec7fd7.
//
// Solidity: function _stakeErc20() view returns(address)
func (_Contract *ContractSession) StakeErc20() (common.Address, error) {
	return _Contract.Contract.StakeErc20(&_Contract.CallOpts)
}

// StakeErc20 is a free data retrieval call binding the contract method 0xf0ec7fd7.
//
// Solidity: function _stakeErc20() view returns(address)
func (_Contract *ContractCallerSession) StakeErc20() (common.Address, error) {
	return _Contract.Contract.StakeErc20(&_Contract.CallOpts)
}

// GetMintRules is a free data retrieval call binding the contract method 0xe5cacf1d.
//
// Solidity: function getMintRules() view returns(address[])
func (_Contract *ContractCaller) GetMintRules(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getMintRules")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetMintRules is a free data retrieval call binding the contract method 0xe5cacf1d.
//
// Solidity: function getMintRules() view returns(address[])
func (_Contract *ContractSession) GetMintRules() ([]common.Address, error) {
	return _Contract.Contract.GetMintRules(&_Contract.CallOpts)
}

// GetMintRules is a free data retrieval call binding the contract method 0xe5cacf1d.
//
// Solidity: function getMintRules() view returns(address[])
func (_Contract *ContractCallerSession) GetMintRules() ([]common.Address, error) {
	return _Contract.Contract.GetMintRules(&_Contract.CallOpts)
}

// GetTheMintTimes is a free data retrieval call binding the contract method 0x37c44f94.
//
// Solidity: function getTheMintTimes(uint256 mintType, address user) view returns(uint256)
func (_Contract *ContractCaller) GetTheMintTimes(opts *bind.CallOpts, mintType *big.Int, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTheMintTimes", mintType, user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTheMintTimes is a free data retrieval call binding the contract method 0x37c44f94.
//
// Solidity: function getTheMintTimes(uint256 mintType, address user) view returns(uint256)
func (_Contract *ContractSession) GetTheMintTimes(mintType *big.Int, user common.Address) (*big.Int, error) {
	return _Contract.Contract.GetTheMintTimes(&_Contract.CallOpts, mintType, user)
}

// GetTheMintTimes is a free data retrieval call binding the contract method 0x37c44f94.
//
// Solidity: function getTheMintTimes(uint256 mintType, address user) view returns(uint256)
func (_Contract *ContractCallerSession) GetTheMintTimes(mintType *big.Int, user common.Address) (*big.Int, error) {
	return _Contract.Contract.GetTheMintTimes(&_Contract.CallOpts, mintType, user)
}

// HashCondition is a free data retrieval call binding the contract method 0x78eaed6d.
//
// Solidity: function hashCondition((address,uint256,address,uint256,address,uint256,uint256,uint256,bytes32,bytes) mintData) pure returns(bytes32)
func (_Contract *ContractCaller) HashCondition(opts *bind.CallOpts, mintData ICartoonMintRuleMintRule) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "hashCondition", mintData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashCondition is a free data retrieval call binding the contract method 0x78eaed6d.
//
// Solidity: function hashCondition((address,uint256,address,uint256,address,uint256,uint256,uint256,bytes32,bytes) mintData) pure returns(bytes32)
func (_Contract *ContractSession) HashCondition(mintData ICartoonMintRuleMintRule) ([32]byte, error) {
	return _Contract.Contract.HashCondition(&_Contract.CallOpts, mintData)
}

// HashCondition is a free data retrieval call binding the contract method 0x78eaed6d.
//
// Solidity: function hashCondition((address,uint256,address,uint256,address,uint256,uint256,uint256,bytes32,bytes) mintData) pure returns(bytes32)
func (_Contract *ContractCallerSession) HashCondition(mintData ICartoonMintRuleMintRule) ([32]byte, error) {
	return _Contract.Contract.HashCondition(&_Contract.CallOpts, mintData)
}

// HashDigest is a free data retrieval call binding the contract method 0x3cde64b8.
//
// Solidity: function hashDigest((address,uint256,address,uint256,address,uint256,uint256,uint256,bytes32,bytes) mintData) view returns(bytes32)
func (_Contract *ContractCaller) HashDigest(opts *bind.CallOpts, mintData ICartoonMintRuleMintRule) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "hashDigest", mintData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashDigest is a free data retrieval call binding the contract method 0x3cde64b8.
//
// Solidity: function hashDigest((address,uint256,address,uint256,address,uint256,uint256,uint256,bytes32,bytes) mintData) view returns(bytes32)
func (_Contract *ContractSession) HashDigest(mintData ICartoonMintRuleMintRule) ([32]byte, error) {
	return _Contract.Contract.HashDigest(&_Contract.CallOpts, mintData)
}

// HashDigest is a free data retrieval call binding the contract method 0x3cde64b8.
//
// Solidity: function hashDigest((address,uint256,address,uint256,address,uint256,uint256,uint256,bytes32,bytes) mintData) view returns(bytes32)
func (_Contract *ContractCallerSession) HashDigest(mintData ICartoonMintRuleMintRule) ([32]byte, error) {
	return _Contract.Contract.HashDigest(&_Contract.CallOpts, mintData)
}

// HashWhiteList is a free data retrieval call binding the contract method 0xc569e060.
//
// Solidity: function hashWhiteList(address user, bytes32 signCode) pure returns(bytes32)
func (_Contract *ContractCaller) HashWhiteList(opts *bind.CallOpts, user common.Address, signCode [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "hashWhiteList", user, signCode)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashWhiteList is a free data retrieval call binding the contract method 0xc569e060.
//
// Solidity: function hashWhiteList(address user, bytes32 signCode) pure returns(bytes32)
func (_Contract *ContractSession) HashWhiteList(user common.Address, signCode [32]byte) ([32]byte, error) {
	return _Contract.Contract.HashWhiteList(&_Contract.CallOpts, user, signCode)
}

// HashWhiteList is a free data retrieval call binding the contract method 0xc569e060.
//
// Solidity: function hashWhiteList(address user, bytes32 signCode) pure returns(bytes32)
func (_Contract *ContractCallerSession) HashWhiteList(user common.Address, signCode [32]byte) ([32]byte, error) {
	return _Contract.Contract.HashWhiteList(&_Contract.CallOpts, user, signCode)
}

// IsExistSignCode is a free data retrieval call binding the contract method 0xf2dc79a2.
//
// Solidity: function isExistSignCode(bytes32 signCode) view returns(bool)
func (_Contract *ContractCaller) IsExistSignCode(opts *bind.CallOpts, signCode [32]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isExistSignCode", signCode)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExistSignCode is a free data retrieval call binding the contract method 0xf2dc79a2.
//
// Solidity: function isExistSignCode(bytes32 signCode) view returns(bool)
func (_Contract *ContractSession) IsExistSignCode(signCode [32]byte) (bool, error) {
	return _Contract.Contract.IsExistSignCode(&_Contract.CallOpts, signCode)
}

// IsExistSignCode is a free data retrieval call binding the contract method 0xf2dc79a2.
//
// Solidity: function isExistSignCode(bytes32 signCode) view returns(bool)
func (_Contract *ContractCallerSession) IsExistSignCode(signCode [32]byte) (bool, error) {
	return _Contract.Contract.IsExistSignCode(&_Contract.CallOpts, signCode)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0x3a2acf44.
//
// Solidity: function verify((address,uint256,address,uint256,address,uint256,uint256,uint256,bytes32,bytes) mintData, address user, bytes dataSignature) view returns(bool)
func (_Contract *ContractCaller) Verify(opts *bind.CallOpts, mintData ICartoonMintRuleMintRule, user common.Address, dataSignature []byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "verify", mintData, user, dataSignature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x3a2acf44.
//
// Solidity: function verify((address,uint256,address,uint256,address,uint256,uint256,uint256,bytes32,bytes) mintData, address user, bytes dataSignature) view returns(bool)
func (_Contract *ContractSession) Verify(mintData ICartoonMintRuleMintRule, user common.Address, dataSignature []byte) (bool, error) {
	return _Contract.Contract.Verify(&_Contract.CallOpts, mintData, user, dataSignature)
}

// Verify is a free data retrieval call binding the contract method 0x3a2acf44.
//
// Solidity: function verify((address,uint256,address,uint256,address,uint256,uint256,uint256,bytes32,bytes) mintData, address user, bytes dataSignature) view returns(bool)
func (_Contract *ContractCallerSession) Verify(mintData ICartoonMintRuleMintRule, user common.Address, dataSignature []byte) (bool, error) {
	return _Contract.Contract.Verify(&_Contract.CallOpts, mintData, user, dataSignature)
}

// VerifyCondition is a free data retrieval call binding the contract method 0x35f49a05.
//
// Solidity: function verifyCondition((address,uint256,address,uint256,address,uint256,uint256,uint256,bytes32,bytes) mintData, uint8 v, bytes32 r, bytes32 s) view returns(bool)
func (_Contract *ContractCaller) VerifyCondition(opts *bind.CallOpts, mintData ICartoonMintRuleMintRule, v uint8, r [32]byte, s [32]byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "verifyCondition", mintData, v, r, s)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyCondition is a free data retrieval call binding the contract method 0x35f49a05.
//
// Solidity: function verifyCondition((address,uint256,address,uint256,address,uint256,uint256,uint256,bytes32,bytes) mintData, uint8 v, bytes32 r, bytes32 s) view returns(bool)
func (_Contract *ContractSession) VerifyCondition(mintData ICartoonMintRuleMintRule, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _Contract.Contract.VerifyCondition(&_Contract.CallOpts, mintData, v, r, s)
}

// VerifyCondition is a free data retrieval call binding the contract method 0x35f49a05.
//
// Solidity: function verifyCondition((address,uint256,address,uint256,address,uint256,uint256,uint256,bytes32,bytes) mintData, uint8 v, bytes32 r, bytes32 s) view returns(bool)
func (_Contract *ContractCallerSession) VerifyCondition(mintData ICartoonMintRuleMintRule, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _Contract.Contract.VerifyCondition(&_Contract.CallOpts, mintData, v, r, s)
}

// VerifySignature is a free data retrieval call binding the contract method 0xdaca6f78.
//
// Solidity: function verifySignature(bytes32 hash, bytes signature) view returns(bool)
func (_Contract *ContractCaller) VerifySignature(opts *bind.CallOpts, hash [32]byte, signature []byte) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "verifySignature", hash, signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySignature is a free data retrieval call binding the contract method 0xdaca6f78.
//
// Solidity: function verifySignature(bytes32 hash, bytes signature) view returns(bool)
func (_Contract *ContractSession) VerifySignature(hash [32]byte, signature []byte) (bool, error) {
	return _Contract.Contract.VerifySignature(&_Contract.CallOpts, hash, signature)
}

// VerifySignature is a free data retrieval call binding the contract method 0xdaca6f78.
//
// Solidity: function verifySignature(bytes32 hash, bytes signature) view returns(bool)
func (_Contract *ContractCallerSession) VerifySignature(hash [32]byte, signature []byte) (bool, error) {
	return _Contract.Contract.VerifySignature(&_Contract.CallOpts, hash, signature)
}

// AddMintRule is a paid mutator transaction binding the contract method 0x510c5c6f.
//
// Solidity: function addMintRule(address rule) returns()
func (_Contract *ContractTransactor) AddMintRule(opts *bind.TransactOpts, rule common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addMintRule", rule)
}

// AddMintRule is a paid mutator transaction binding the contract method 0x510c5c6f.
//
// Solidity: function addMintRule(address rule) returns()
func (_Contract *ContractSession) AddMintRule(rule common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddMintRule(&_Contract.TransactOpts, rule)
}

// AddMintRule is a paid mutator transaction binding the contract method 0x510c5c6f.
//
// Solidity: function addMintRule(address rule) returns()
func (_Contract *ContractTransactorSession) AddMintRule(rule common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddMintRule(&_Contract.TransactOpts, rule)
}

// MintAvatar721 is a paid mutator transaction binding the contract method 0x5e194c37.
//
// Solidity: function mintAvatar721((address,uint256,address,uint256,address,uint256,uint256,uint256,bytes32,bytes) mintData, bytes dataSignature) returns()
func (_Contract *ContractTransactor) MintAvatar721(opts *bind.TransactOpts, mintData ICartoonMintRuleMintRule, dataSignature []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "mintAvatar721", mintData, dataSignature)
}

// MintAvatar721 is a paid mutator transaction binding the contract method 0x5e194c37.
//
// Solidity: function mintAvatar721((address,uint256,address,uint256,address,uint256,uint256,uint256,bytes32,bytes) mintData, bytes dataSignature) returns()
func (_Contract *ContractSession) MintAvatar721(mintData ICartoonMintRuleMintRule, dataSignature []byte) (*types.Transaction, error) {
	return _Contract.Contract.MintAvatar721(&_Contract.TransactOpts, mintData, dataSignature)
}

// MintAvatar721 is a paid mutator transaction binding the contract method 0x5e194c37.
//
// Solidity: function mintAvatar721((address,uint256,address,uint256,address,uint256,uint256,uint256,bytes32,bytes) mintData, bytes dataSignature) returns()
func (_Contract *ContractTransactorSession) MintAvatar721(mintData ICartoonMintRuleMintRule, dataSignature []byte) (*types.Transaction, error) {
	return _Contract.Contract.MintAvatar721(&_Contract.TransactOpts, mintData, dataSignature)
}

// RemoveMintRule is a paid mutator transaction binding the contract method 0xe66bbb38.
//
// Solidity: function removeMintRule(address rule) returns()
func (_Contract *ContractTransactor) RemoveMintRule(opts *bind.TransactOpts, rule common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "removeMintRule", rule)
}

// RemoveMintRule is a paid mutator transaction binding the contract method 0xe66bbb38.
//
// Solidity: function removeMintRule(address rule) returns()
func (_Contract *ContractSession) RemoveMintRule(rule common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveMintRule(&_Contract.TransactOpts, rule)
}

// RemoveMintRule is a paid mutator transaction binding the contract method 0xe66bbb38.
//
// Solidity: function removeMintRule(address rule) returns()
func (_Contract *ContractTransactorSession) RemoveMintRule(rule common.Address) (*types.Transaction, error) {
	return _Contract.Contract.RemoveMintRule(&_Contract.TransactOpts, rule)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// SetOnceSignCode is a paid mutator transaction binding the contract method 0xfb345076.
//
// Solidity: function setOnceSignCode(bool enable) returns()
func (_Contract *ContractTransactor) SetOnceSignCode(opts *bind.TransactOpts, enable bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setOnceSignCode", enable)
}

// SetOnceSignCode is a paid mutator transaction binding the contract method 0xfb345076.
//
// Solidity: function setOnceSignCode(bool enable) returns()
func (_Contract *ContractSession) SetOnceSignCode(enable bool) (*types.Transaction, error) {
	return _Contract.Contract.SetOnceSignCode(&_Contract.TransactOpts, enable)
}

// SetOnceSignCode is a paid mutator transaction binding the contract method 0xfb345076.
//
// Solidity: function setOnceSignCode(bool enable) returns()
func (_Contract *ContractTransactorSession) SetOnceSignCode(enable bool) (*types.Transaction, error) {
	return _Contract.Contract.SetOnceSignCode(&_Contract.TransactOpts, enable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// UpdateCostErc20 is a paid mutator transaction binding the contract method 0x0071ff6d.
//
// Solidity: function updateCostErc20(address costErc20, uint256 costAmount) returns()
func (_Contract *ContractTransactor) UpdateCostErc20(opts *bind.TransactOpts, costErc20 common.Address, costAmount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateCostErc20", costErc20, costAmount)
}

// UpdateCostErc20 is a paid mutator transaction binding the contract method 0x0071ff6d.
//
// Solidity: function updateCostErc20(address costErc20, uint256 costAmount) returns()
func (_Contract *ContractSession) UpdateCostErc20(costErc20 common.Address, costAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateCostErc20(&_Contract.TransactOpts, costErc20, costAmount)
}

// UpdateCostErc20 is a paid mutator transaction binding the contract method 0x0071ff6d.
//
// Solidity: function updateCostErc20(address costErc20, uint256 costAmount) returns()
func (_Contract *ContractTransactorSession) UpdateCostErc20(costErc20 common.Address, costAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateCostErc20(&_Contract.TransactOpts, costErc20, costAmount)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xa7ecd37e.
//
// Solidity: function updateSigner(address SIGNER) returns()
func (_Contract *ContractTransactor) UpdateSigner(opts *bind.TransactOpts, SIGNER common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateSigner", SIGNER)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xa7ecd37e.
//
// Solidity: function updateSigner(address SIGNER) returns()
func (_Contract *ContractSession) UpdateSigner(SIGNER common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSigner(&_Contract.TransactOpts, SIGNER)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xa7ecd37e.
//
// Solidity: function updateSigner(address SIGNER) returns()
func (_Contract *ContractTransactorSession) UpdateSigner(SIGNER common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSigner(&_Contract.TransactOpts, SIGNER)
}

// UpdateStakeErc20 is a paid mutator transaction binding the contract method 0x40f80ade.
//
// Solidity: function updateStakeErc20(address stakeErc20, uint256 stakeAmount) returns()
func (_Contract *ContractTransactor) UpdateStakeErc20(opts *bind.TransactOpts, stakeErc20 common.Address, stakeAmount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateStakeErc20", stakeErc20, stakeAmount)
}

// UpdateStakeErc20 is a paid mutator transaction binding the contract method 0x40f80ade.
//
// Solidity: function updateStakeErc20(address stakeErc20, uint256 stakeAmount) returns()
func (_Contract *ContractSession) UpdateStakeErc20(stakeErc20 common.Address, stakeAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateStakeErc20(&_Contract.TransactOpts, stakeErc20, stakeAmount)
}

// UpdateStakeErc20 is a paid mutator transaction binding the contract method 0x40f80ade.
//
// Solidity: function updateStakeErc20(address stakeErc20, uint256 stakeAmount) returns()
func (_Contract *ContractTransactorSession) UpdateStakeErc20(stakeErc20 common.Address, stakeAmount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateStakeErc20(&_Contract.TransactOpts, stakeErc20, stakeAmount)
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEUpdateSignerIterator is returned from FilterEUpdateSigner and is used to iterate over the raw logs and unpacked data for EUpdateSigner events raised by the Contract contract.
type ContractEUpdateSignerIterator struct {
	Event *ContractEUpdateSigner // Event containing the contract specifics and raw log

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
func (it *ContractEUpdateSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEUpdateSigner)
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
		it.Event = new(ContractEUpdateSigner)
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
func (it *ContractEUpdateSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEUpdateSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEUpdateSigner represents a EUpdateSigner event raised by the Contract contract.
type ContractEUpdateSigner struct {
	Signer   common.Address
	BlockNum *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEUpdateSigner is a free log retrieval operation binding the contract event 0x83fc816ecc6a4c6bad716c3c535b69cb1f49286557b74c0784cc742be8347075.
//
// Solidity: event eUpdateSigner(address signer, uint256 blockNum)
func (_Contract *ContractFilterer) FilterEUpdateSigner(opts *bind.FilterOpts) (*ContractEUpdateSignerIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "eUpdateSigner")
	if err != nil {
		return nil, err
	}
	return &ContractEUpdateSignerIterator{contract: _Contract.contract, event: "eUpdateSigner", logs: logs, sub: sub}, nil
}

// WatchEUpdateSigner is a free log subscription operation binding the contract event 0x83fc816ecc6a4c6bad716c3c535b69cb1f49286557b74c0784cc742be8347075.
//
// Solidity: event eUpdateSigner(address signer, uint256 blockNum)
func (_Contract *ContractFilterer) WatchEUpdateSigner(opts *bind.WatchOpts, sink chan<- *ContractEUpdateSigner) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "eUpdateSigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEUpdateSigner)
				if err := _Contract.contract.UnpackLog(event, "eUpdateSigner", log); err != nil {
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

// ParseEUpdateSigner is a log parse operation binding the contract event 0x83fc816ecc6a4c6bad716c3c535b69cb1f49286557b74c0784cc742be8347075.
//
// Solidity: event eUpdateSigner(address signer, uint256 blockNum)
func (_Contract *ContractFilterer) ParseEUpdateSigner(log types.Log) (*ContractEUpdateSigner, error) {
	event := new(ContractEUpdateSigner)
	if err := _Contract.contract.UnpackLog(event, "eUpdateSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
