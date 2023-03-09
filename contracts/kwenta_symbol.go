// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// IPerpsV2MarketBaseTypesPosition is an auto generated low-level Go binding around an user-defined struct.
type IPerpsV2MarketBaseTypesPosition struct {
	Id               uint64
	LastFundingIndex uint64
	Margin           *big.Int
	LastPrice        *big.Int
	Size             *big.Int
}

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketState\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"CacheUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerNominated\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accessibleMargin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"marginAccessible\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"invalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accruedFunding\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"funding\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"invalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"assetPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"invalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseAsset\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"canLiquidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentFundingRate\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentFundingVelocity\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"int256\",\"name\":\"sizeDelta\",\"type\":\"int256\"}],\"name\":\"fillPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"invalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fundingLastRecomputed\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"fundingSequence\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fundingSequenceLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isResolverCached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"liquidationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"liquidationPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"invalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"marketDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"invalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"marketKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"marketSize\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"marketSizes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"long\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"short\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"marketSkew\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"marketState\",\"outputs\":[{\"internalType\":\"contractIPerpsV2MarketState\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"nominateNewOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nominatedOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"notionalValue\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"invalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"int256\",\"name\":\"sizeDelta\",\"type\":\"int256\"},{\"internalType\":\"enumIPerpsV2MarketBaseTypes.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"}],\"name\":\"orderFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"invalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"positions\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"lastFundingIndex\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"margin\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"lastPrice\",\"type\":\"uint128\"},{\"internalType\":\"int128\",\"name\":\"size\",\"type\":\"int128\"}],\"internalType\":\"structIPerpsV2MarketBaseTypes.Position\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"int256\",\"name\":\"sizeDelta\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"tradePrice\",\"type\":\"uint256\"},{\"internalType\":\"enumIPerpsV2MarketBaseTypes.OrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"postTradeDetails\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"margin\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"size\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liqPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"enumIPerpsV2MarketBaseTypes.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"profitLoss\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"pnl\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"invalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"rebuildCache\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"remainingMargin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"marginRemaining\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"invalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"contractAddressResolver\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"resolverAddressesRequired\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"addresses\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unrecordedFunding\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"funding\",\"type\":\"int256\"},{\"internalType\":\"bool\",\"name\":\"invalid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// AccessibleMargin is a free data retrieval call binding the contract method 0x785cdeec.
//
// Solidity: function accessibleMargin(address account) view returns(uint256 marginAccessible, bool invalid)
func (_Contracts *ContractsCaller) AccessibleMargin(opts *bind.CallOpts, account common.Address) (struct {
	MarginAccessible *big.Int
	Invalid          bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "accessibleMargin", account)

	outstruct := new(struct {
		MarginAccessible *big.Int
		Invalid          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MarginAccessible = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Invalid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// AccessibleMargin is a free data retrieval call binding the contract method 0x785cdeec.
//
// Solidity: function accessibleMargin(address account) view returns(uint256 marginAccessible, bool invalid)
func (_Contracts *ContractsSession) AccessibleMargin(account common.Address) (struct {
	MarginAccessible *big.Int
	Invalid          bool
}, error) {
	return _Contracts.Contract.AccessibleMargin(&_Contracts.CallOpts, account)
}

// AccessibleMargin is a free data retrieval call binding the contract method 0x785cdeec.
//
// Solidity: function accessibleMargin(address account) view returns(uint256 marginAccessible, bool invalid)
func (_Contracts *ContractsCallerSession) AccessibleMargin(account common.Address) (struct {
	MarginAccessible *big.Int
	Invalid          bool
}, error) {
	return _Contracts.Contract.AccessibleMargin(&_Contracts.CallOpts, account)
}

// AccruedFunding is a free data retrieval call binding the contract method 0x1bf556d0.
//
// Solidity: function accruedFunding(address account) view returns(int256 funding, bool invalid)
func (_Contracts *ContractsCaller) AccruedFunding(opts *bind.CallOpts, account common.Address) (struct {
	Funding *big.Int
	Invalid bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "accruedFunding", account)

	outstruct := new(struct {
		Funding *big.Int
		Invalid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Funding = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Invalid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// AccruedFunding is a free data retrieval call binding the contract method 0x1bf556d0.
//
// Solidity: function accruedFunding(address account) view returns(int256 funding, bool invalid)
func (_Contracts *ContractsSession) AccruedFunding(account common.Address) (struct {
	Funding *big.Int
	Invalid bool
}, error) {
	return _Contracts.Contract.AccruedFunding(&_Contracts.CallOpts, account)
}

// AccruedFunding is a free data retrieval call binding the contract method 0x1bf556d0.
//
// Solidity: function accruedFunding(address account) view returns(int256 funding, bool invalid)
func (_Contracts *ContractsCallerSession) AccruedFunding(account common.Address) (struct {
	Funding *big.Int
	Invalid bool
}, error) {
	return _Contracts.Contract.AccruedFunding(&_Contracts.CallOpts, account)
}

// AssetPrice is a free data retrieval call binding the contract method 0xd24378eb.
//
// Solidity: function assetPrice() view returns(uint256 price, bool invalid)
func (_Contracts *ContractsCaller) AssetPrice(opts *bind.CallOpts) (struct {
	Price   *big.Int
	Invalid bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "assetPrice")

	outstruct := new(struct {
		Price   *big.Int
		Invalid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Invalid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// AssetPrice is a free data retrieval call binding the contract method 0xd24378eb.
//
// Solidity: function assetPrice() view returns(uint256 price, bool invalid)
func (_Contracts *ContractsSession) AssetPrice() (struct {
	Price   *big.Int
	Invalid bool
}, error) {
	return _Contracts.Contract.AssetPrice(&_Contracts.CallOpts)
}

// AssetPrice is a free data retrieval call binding the contract method 0xd24378eb.
//
// Solidity: function assetPrice() view returns(uint256 price, bool invalid)
func (_Contracts *ContractsCallerSession) AssetPrice() (struct {
	Price   *big.Int
	Invalid bool
}, error) {
	return _Contracts.Contract.AssetPrice(&_Contracts.CallOpts)
}

// BaseAsset is a free data retrieval call binding the contract method 0xcdf456e1.
//
// Solidity: function baseAsset() view returns(bytes32 key)
func (_Contracts *ContractsCaller) BaseAsset(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "baseAsset")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BaseAsset is a free data retrieval call binding the contract method 0xcdf456e1.
//
// Solidity: function baseAsset() view returns(bytes32 key)
func (_Contracts *ContractsSession) BaseAsset() ([32]byte, error) {
	return _Contracts.Contract.BaseAsset(&_Contracts.CallOpts)
}

// BaseAsset is a free data retrieval call binding the contract method 0xcdf456e1.
//
// Solidity: function baseAsset() view returns(bytes32 key)
func (_Contracts *ContractsCallerSession) BaseAsset() ([32]byte, error) {
	return _Contracts.Contract.BaseAsset(&_Contracts.CallOpts)
}

// CanLiquidate is a free data retrieval call binding the contract method 0xb9f4ff55.
//
// Solidity: function canLiquidate(address account) view returns(bool)
func (_Contracts *ContractsCaller) CanLiquidate(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "canLiquidate", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanLiquidate is a free data retrieval call binding the contract method 0xb9f4ff55.
//
// Solidity: function canLiquidate(address account) view returns(bool)
func (_Contracts *ContractsSession) CanLiquidate(account common.Address) (bool, error) {
	return _Contracts.Contract.CanLiquidate(&_Contracts.CallOpts, account)
}

// CanLiquidate is a free data retrieval call binding the contract method 0xb9f4ff55.
//
// Solidity: function canLiquidate(address account) view returns(bool)
func (_Contracts *ContractsCallerSession) CanLiquidate(account common.Address) (bool, error) {
	return _Contracts.Contract.CanLiquidate(&_Contracts.CallOpts, account)
}

// CurrentFundingRate is a free data retrieval call binding the contract method 0x3aef4d0b.
//
// Solidity: function currentFundingRate() view returns(int256)
func (_Contracts *ContractsCaller) CurrentFundingRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "currentFundingRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentFundingRate is a free data retrieval call binding the contract method 0x3aef4d0b.
//
// Solidity: function currentFundingRate() view returns(int256)
func (_Contracts *ContractsSession) CurrentFundingRate() (*big.Int, error) {
	return _Contracts.Contract.CurrentFundingRate(&_Contracts.CallOpts)
}

// CurrentFundingRate is a free data retrieval call binding the contract method 0x3aef4d0b.
//
// Solidity: function currentFundingRate() view returns(int256)
func (_Contracts *ContractsCallerSession) CurrentFundingRate() (*big.Int, error) {
	return _Contracts.Contract.CurrentFundingRate(&_Contracts.CallOpts)
}

// CurrentFundingVelocity is a free data retrieval call binding the contract method 0xb74e3806.
//
// Solidity: function currentFundingVelocity() view returns(int256)
func (_Contracts *ContractsCaller) CurrentFundingVelocity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "currentFundingVelocity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentFundingVelocity is a free data retrieval call binding the contract method 0xb74e3806.
//
// Solidity: function currentFundingVelocity() view returns(int256)
func (_Contracts *ContractsSession) CurrentFundingVelocity() (*big.Int, error) {
	return _Contracts.Contract.CurrentFundingVelocity(&_Contracts.CallOpts)
}

// CurrentFundingVelocity is a free data retrieval call binding the contract method 0xb74e3806.
//
// Solidity: function currentFundingVelocity() view returns(int256)
func (_Contracts *ContractsCallerSession) CurrentFundingVelocity() (*big.Int, error) {
	return _Contracts.Contract.CurrentFundingVelocity(&_Contracts.CallOpts)
}

// FillPrice is a free data retrieval call binding the contract method 0xea9f9aa7.
//
// Solidity: function fillPrice(int256 sizeDelta) view returns(uint256 price, bool invalid)
func (_Contracts *ContractsCaller) FillPrice(opts *bind.CallOpts, sizeDelta *big.Int) (struct {
	Price   *big.Int
	Invalid bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "fillPrice", sizeDelta)

	outstruct := new(struct {
		Price   *big.Int
		Invalid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Invalid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// FillPrice is a free data retrieval call binding the contract method 0xea9f9aa7.
//
// Solidity: function fillPrice(int256 sizeDelta) view returns(uint256 price, bool invalid)
func (_Contracts *ContractsSession) FillPrice(sizeDelta *big.Int) (struct {
	Price   *big.Int
	Invalid bool
}, error) {
	return _Contracts.Contract.FillPrice(&_Contracts.CallOpts, sizeDelta)
}

// FillPrice is a free data retrieval call binding the contract method 0xea9f9aa7.
//
// Solidity: function fillPrice(int256 sizeDelta) view returns(uint256 price, bool invalid)
func (_Contracts *ContractsCallerSession) FillPrice(sizeDelta *big.Int) (struct {
	Price   *big.Int
	Invalid bool
}, error) {
	return _Contracts.Contract.FillPrice(&_Contracts.CallOpts, sizeDelta)
}

// FundingLastRecomputed is a free data retrieval call binding the contract method 0x27b9a236.
//
// Solidity: function fundingLastRecomputed() view returns(uint32)
func (_Contracts *ContractsCaller) FundingLastRecomputed(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "fundingLastRecomputed")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FundingLastRecomputed is a free data retrieval call binding the contract method 0x27b9a236.
//
// Solidity: function fundingLastRecomputed() view returns(uint32)
func (_Contracts *ContractsSession) FundingLastRecomputed() (uint32, error) {
	return _Contracts.Contract.FundingLastRecomputed(&_Contracts.CallOpts)
}

// FundingLastRecomputed is a free data retrieval call binding the contract method 0x27b9a236.
//
// Solidity: function fundingLastRecomputed() view returns(uint32)
func (_Contracts *ContractsCallerSession) FundingLastRecomputed() (uint32, error) {
	return _Contracts.Contract.FundingLastRecomputed(&_Contracts.CallOpts)
}

// FundingSequence is a free data retrieval call binding the contract method 0x41108cf2.
//
// Solidity: function fundingSequence(uint256 index) view returns(int128)
func (_Contracts *ContractsCaller) FundingSequence(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "fundingSequence", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingSequence is a free data retrieval call binding the contract method 0x41108cf2.
//
// Solidity: function fundingSequence(uint256 index) view returns(int128)
func (_Contracts *ContractsSession) FundingSequence(index *big.Int) (*big.Int, error) {
	return _Contracts.Contract.FundingSequence(&_Contracts.CallOpts, index)
}

// FundingSequence is a free data retrieval call binding the contract method 0x41108cf2.
//
// Solidity: function fundingSequence(uint256 index) view returns(int128)
func (_Contracts *ContractsCallerSession) FundingSequence(index *big.Int) (*big.Int, error) {
	return _Contracts.Contract.FundingSequence(&_Contracts.CallOpts, index)
}

// FundingSequenceLength is a free data retrieval call binding the contract method 0xcded0cea.
//
// Solidity: function fundingSequenceLength() view returns(uint256)
func (_Contracts *ContractsCaller) FundingSequenceLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "fundingSequenceLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingSequenceLength is a free data retrieval call binding the contract method 0xcded0cea.
//
// Solidity: function fundingSequenceLength() view returns(uint256)
func (_Contracts *ContractsSession) FundingSequenceLength() (*big.Int, error) {
	return _Contracts.Contract.FundingSequenceLength(&_Contracts.CallOpts)
}

// FundingSequenceLength is a free data retrieval call binding the contract method 0xcded0cea.
//
// Solidity: function fundingSequenceLength() view returns(uint256)
func (_Contracts *ContractsCallerSession) FundingSequenceLength() (*big.Int, error) {
	return _Contracts.Contract.FundingSequenceLength(&_Contracts.CallOpts)
}

// IsResolverCached is a free data retrieval call binding the contract method 0x2af64bd3.
//
// Solidity: function isResolverCached() view returns(bool)
func (_Contracts *ContractsCaller) IsResolverCached(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isResolverCached")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsResolverCached is a free data retrieval call binding the contract method 0x2af64bd3.
//
// Solidity: function isResolverCached() view returns(bool)
func (_Contracts *ContractsSession) IsResolverCached() (bool, error) {
	return _Contracts.Contract.IsResolverCached(&_Contracts.CallOpts)
}

// IsResolverCached is a free data retrieval call binding the contract method 0x2af64bd3.
//
// Solidity: function isResolverCached() view returns(bool)
func (_Contracts *ContractsCallerSession) IsResolverCached() (bool, error) {
	return _Contracts.Contract.IsResolverCached(&_Contracts.CallOpts)
}

// LiquidationFee is a free data retrieval call binding the contract method 0xc8023af4.
//
// Solidity: function liquidationFee(address account) view returns(uint256)
func (_Contracts *ContractsCaller) LiquidationFee(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "liquidationFee", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationFee is a free data retrieval call binding the contract method 0xc8023af4.
//
// Solidity: function liquidationFee(address account) view returns(uint256)
func (_Contracts *ContractsSession) LiquidationFee(account common.Address) (*big.Int, error) {
	return _Contracts.Contract.LiquidationFee(&_Contracts.CallOpts, account)
}

// LiquidationFee is a free data retrieval call binding the contract method 0xc8023af4.
//
// Solidity: function liquidationFee(address account) view returns(uint256)
func (_Contracts *ContractsCallerSession) LiquidationFee(account common.Address) (*big.Int, error) {
	return _Contracts.Contract.LiquidationFee(&_Contracts.CallOpts, account)
}

// LiquidationPrice is a free data retrieval call binding the contract method 0x964db90c.
//
// Solidity: function liquidationPrice(address account) view returns(uint256 price, bool invalid)
func (_Contracts *ContractsCaller) LiquidationPrice(opts *bind.CallOpts, account common.Address) (struct {
	Price   *big.Int
	Invalid bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "liquidationPrice", account)

	outstruct := new(struct {
		Price   *big.Int
		Invalid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Invalid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// LiquidationPrice is a free data retrieval call binding the contract method 0x964db90c.
//
// Solidity: function liquidationPrice(address account) view returns(uint256 price, bool invalid)
func (_Contracts *ContractsSession) LiquidationPrice(account common.Address) (struct {
	Price   *big.Int
	Invalid bool
}, error) {
	return _Contracts.Contract.LiquidationPrice(&_Contracts.CallOpts, account)
}

// LiquidationPrice is a free data retrieval call binding the contract method 0x964db90c.
//
// Solidity: function liquidationPrice(address account) view returns(uint256 price, bool invalid)
func (_Contracts *ContractsCallerSession) LiquidationPrice(account common.Address) (struct {
	Price   *big.Int
	Invalid bool
}, error) {
	return _Contracts.Contract.LiquidationPrice(&_Contracts.CallOpts, account)
}

// MarketDebt is a free data retrieval call binding the contract method 0xe8c63470.
//
// Solidity: function marketDebt() view returns(uint256 debt, bool invalid)
func (_Contracts *ContractsCaller) MarketDebt(opts *bind.CallOpts) (struct {
	Debt    *big.Int
	Invalid bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "marketDebt")

	outstruct := new(struct {
		Debt    *big.Int
		Invalid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Debt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Invalid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// MarketDebt is a free data retrieval call binding the contract method 0xe8c63470.
//
// Solidity: function marketDebt() view returns(uint256 debt, bool invalid)
func (_Contracts *ContractsSession) MarketDebt() (struct {
	Debt    *big.Int
	Invalid bool
}, error) {
	return _Contracts.Contract.MarketDebt(&_Contracts.CallOpts)
}

// MarketDebt is a free data retrieval call binding the contract method 0xe8c63470.
//
// Solidity: function marketDebt() view returns(uint256 debt, bool invalid)
func (_Contracts *ContractsCallerSession) MarketDebt() (struct {
	Debt    *big.Int
	Invalid bool
}, error) {
	return _Contracts.Contract.MarketDebt(&_Contracts.CallOpts)
}

// MarketKey is a free data retrieval call binding the contract method 0xd7103a46.
//
// Solidity: function marketKey() view returns(bytes32 key)
func (_Contracts *ContractsCaller) MarketKey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "marketKey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MarketKey is a free data retrieval call binding the contract method 0xd7103a46.
//
// Solidity: function marketKey() view returns(bytes32 key)
func (_Contracts *ContractsSession) MarketKey() ([32]byte, error) {
	return _Contracts.Contract.MarketKey(&_Contracts.CallOpts)
}

// MarketKey is a free data retrieval call binding the contract method 0xd7103a46.
//
// Solidity: function marketKey() view returns(bytes32 key)
func (_Contracts *ContractsCallerSession) MarketKey() ([32]byte, error) {
	return _Contracts.Contract.MarketKey(&_Contracts.CallOpts)
}

// MarketSize is a free data retrieval call binding the contract method 0xeb56105d.
//
// Solidity: function marketSize() view returns(uint128)
func (_Contracts *ContractsCaller) MarketSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "marketSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketSize is a free data retrieval call binding the contract method 0xeb56105d.
//
// Solidity: function marketSize() view returns(uint128)
func (_Contracts *ContractsSession) MarketSize() (*big.Int, error) {
	return _Contracts.Contract.MarketSize(&_Contracts.CallOpts)
}

// MarketSize is a free data retrieval call binding the contract method 0xeb56105d.
//
// Solidity: function marketSize() view returns(uint128)
func (_Contracts *ContractsCallerSession) MarketSize() (*big.Int, error) {
	return _Contracts.Contract.MarketSize(&_Contracts.CallOpts)
}

// MarketSizes is a free data retrieval call binding the contract method 0x5fc890c2.
//
// Solidity: function marketSizes() view returns(uint256 long, uint256 short)
func (_Contracts *ContractsCaller) MarketSizes(opts *bind.CallOpts) (struct {
	Long  *big.Int
	Short *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "marketSizes")

	outstruct := new(struct {
		Long  *big.Int
		Short *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Long = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Short = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MarketSizes is a free data retrieval call binding the contract method 0x5fc890c2.
//
// Solidity: function marketSizes() view returns(uint256 long, uint256 short)
func (_Contracts *ContractsSession) MarketSizes() (struct {
	Long  *big.Int
	Short *big.Int
}, error) {
	return _Contracts.Contract.MarketSizes(&_Contracts.CallOpts)
}

// MarketSizes is a free data retrieval call binding the contract method 0x5fc890c2.
//
// Solidity: function marketSizes() view returns(uint256 long, uint256 short)
func (_Contracts *ContractsCallerSession) MarketSizes() (struct {
	Long  *big.Int
	Short *big.Int
}, error) {
	return _Contracts.Contract.MarketSizes(&_Contracts.CallOpts)
}

// MarketSkew is a free data retrieval call binding the contract method 0x2b58ecef.
//
// Solidity: function marketSkew() view returns(int128)
func (_Contracts *ContractsCaller) MarketSkew(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "marketSkew")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketSkew is a free data retrieval call binding the contract method 0x2b58ecef.
//
// Solidity: function marketSkew() view returns(int128)
func (_Contracts *ContractsSession) MarketSkew() (*big.Int, error) {
	return _Contracts.Contract.MarketSkew(&_Contracts.CallOpts)
}

// MarketSkew is a free data retrieval call binding the contract method 0x2b58ecef.
//
// Solidity: function marketSkew() view returns(int128)
func (_Contracts *ContractsCallerSession) MarketSkew() (*big.Int, error) {
	return _Contracts.Contract.MarketSkew(&_Contracts.CallOpts)
}

// MarketState is a free data retrieval call binding the contract method 0x08fb1b77.
//
// Solidity: function marketState() view returns(address)
func (_Contracts *ContractsCaller) MarketState(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "marketState")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketState is a free data retrieval call binding the contract method 0x08fb1b77.
//
// Solidity: function marketState() view returns(address)
func (_Contracts *ContractsSession) MarketState() (common.Address, error) {
	return _Contracts.Contract.MarketState(&_Contracts.CallOpts)
}

// MarketState is a free data retrieval call binding the contract method 0x08fb1b77.
//
// Solidity: function marketState() view returns(address)
func (_Contracts *ContractsCallerSession) MarketState() (common.Address, error) {
	return _Contracts.Contract.MarketState(&_Contracts.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Contracts *ContractsCaller) NominatedOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "nominatedOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Contracts *ContractsSession) NominatedOwner() (common.Address, error) {
	return _Contracts.Contract.NominatedOwner(&_Contracts.CallOpts)
}

// NominatedOwner is a free data retrieval call binding the contract method 0x53a47bb7.
//
// Solidity: function nominatedOwner() view returns(address)
func (_Contracts *ContractsCallerSession) NominatedOwner() (common.Address, error) {
	return _Contracts.Contract.NominatedOwner(&_Contracts.CallOpts)
}

// NotionalValue is a free data retrieval call binding the contract method 0xb895daab.
//
// Solidity: function notionalValue(address account) view returns(int256 value, bool invalid)
func (_Contracts *ContractsCaller) NotionalValue(opts *bind.CallOpts, account common.Address) (struct {
	Value   *big.Int
	Invalid bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "notionalValue", account)

	outstruct := new(struct {
		Value   *big.Int
		Invalid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Value = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Invalid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// NotionalValue is a free data retrieval call binding the contract method 0xb895daab.
//
// Solidity: function notionalValue(address account) view returns(int256 value, bool invalid)
func (_Contracts *ContractsSession) NotionalValue(account common.Address) (struct {
	Value   *big.Int
	Invalid bool
}, error) {
	return _Contracts.Contract.NotionalValue(&_Contracts.CallOpts, account)
}

// NotionalValue is a free data retrieval call binding the contract method 0xb895daab.
//
// Solidity: function notionalValue(address account) view returns(int256 value, bool invalid)
func (_Contracts *ContractsCallerSession) NotionalValue(account common.Address) (struct {
	Value   *big.Int
	Invalid bool
}, error) {
	return _Contracts.Contract.NotionalValue(&_Contracts.CallOpts, account)
}

// OrderFee is a free data retrieval call binding the contract method 0x4dd9d7e9.
//
// Solidity: function orderFee(int256 sizeDelta, uint8 orderType) view returns(uint256 fee, bool invalid)
func (_Contracts *ContractsCaller) OrderFee(opts *bind.CallOpts, sizeDelta *big.Int, orderType uint8) (struct {
	Fee     *big.Int
	Invalid bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "orderFee", sizeDelta, orderType)

	outstruct := new(struct {
		Fee     *big.Int
		Invalid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Invalid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// OrderFee is a free data retrieval call binding the contract method 0x4dd9d7e9.
//
// Solidity: function orderFee(int256 sizeDelta, uint8 orderType) view returns(uint256 fee, bool invalid)
func (_Contracts *ContractsSession) OrderFee(sizeDelta *big.Int, orderType uint8) (struct {
	Fee     *big.Int
	Invalid bool
}, error) {
	return _Contracts.Contract.OrderFee(&_Contracts.CallOpts, sizeDelta, orderType)
}

// OrderFee is a free data retrieval call binding the contract method 0x4dd9d7e9.
//
// Solidity: function orderFee(int256 sizeDelta, uint8 orderType) view returns(uint256 fee, bool invalid)
func (_Contracts *ContractsCallerSession) OrderFee(sizeDelta *big.Int, orderType uint8) (struct {
	Fee     *big.Int
	Invalid bool
}, error) {
	return _Contracts.Contract.OrderFee(&_Contracts.CallOpts, sizeDelta, orderType)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Positions is a free data retrieval call binding the contract method 0x55f57510.
//
// Solidity: function positions(address account) view returns((uint64,uint64,uint128,uint128,int128))
func (_Contracts *ContractsCaller) Positions(opts *bind.CallOpts, account common.Address) (IPerpsV2MarketBaseTypesPosition, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "positions", account)

	if err != nil {
		return *new(IPerpsV2MarketBaseTypesPosition), err
	}

	out0 := *abi.ConvertType(out[0], new(IPerpsV2MarketBaseTypesPosition)).(*IPerpsV2MarketBaseTypesPosition)

	return out0, err

}

// Positions is a free data retrieval call binding the contract method 0x55f57510.
//
// Solidity: function positions(address account) view returns((uint64,uint64,uint128,uint128,int128))
func (_Contracts *ContractsSession) Positions(account common.Address) (IPerpsV2MarketBaseTypesPosition, error) {
	return _Contracts.Contract.Positions(&_Contracts.CallOpts, account)
}

// Positions is a free data retrieval call binding the contract method 0x55f57510.
//
// Solidity: function positions(address account) view returns((uint64,uint64,uint128,uint128,int128))
func (_Contracts *ContractsCallerSession) Positions(account common.Address) (IPerpsV2MarketBaseTypesPosition, error) {
	return _Contracts.Contract.Positions(&_Contracts.CallOpts, account)
}

// PostTradeDetails is a free data retrieval call binding the contract method 0xea1d5478.
//
// Solidity: function postTradeDetails(int256 sizeDelta, uint256 tradePrice, uint8 orderType, address sender) view returns(uint256 margin, int256 size, uint256 price, uint256 liqPrice, uint256 fee, uint8 status)
func (_Contracts *ContractsCaller) PostTradeDetails(opts *bind.CallOpts, sizeDelta *big.Int, tradePrice *big.Int, orderType uint8, sender common.Address) (struct {
	Margin   *big.Int
	Size     *big.Int
	Price    *big.Int
	LiqPrice *big.Int
	Fee      *big.Int
	Status   uint8
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "postTradeDetails", sizeDelta, tradePrice, orderType, sender)

	outstruct := new(struct {
		Margin   *big.Int
		Size     *big.Int
		Price    *big.Int
		LiqPrice *big.Int
		Fee      *big.Int
		Status   uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Margin = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Size = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LiqPrice = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// PostTradeDetails is a free data retrieval call binding the contract method 0xea1d5478.
//
// Solidity: function postTradeDetails(int256 sizeDelta, uint256 tradePrice, uint8 orderType, address sender) view returns(uint256 margin, int256 size, uint256 price, uint256 liqPrice, uint256 fee, uint8 status)
func (_Contracts *ContractsSession) PostTradeDetails(sizeDelta *big.Int, tradePrice *big.Int, orderType uint8, sender common.Address) (struct {
	Margin   *big.Int
	Size     *big.Int
	Price    *big.Int
	LiqPrice *big.Int
	Fee      *big.Int
	Status   uint8
}, error) {
	return _Contracts.Contract.PostTradeDetails(&_Contracts.CallOpts, sizeDelta, tradePrice, orderType, sender)
}

// PostTradeDetails is a free data retrieval call binding the contract method 0xea1d5478.
//
// Solidity: function postTradeDetails(int256 sizeDelta, uint256 tradePrice, uint8 orderType, address sender) view returns(uint256 margin, int256 size, uint256 price, uint256 liqPrice, uint256 fee, uint8 status)
func (_Contracts *ContractsCallerSession) PostTradeDetails(sizeDelta *big.Int, tradePrice *big.Int, orderType uint8, sender common.Address) (struct {
	Margin   *big.Int
	Size     *big.Int
	Price    *big.Int
	LiqPrice *big.Int
	Fee      *big.Int
	Status   uint8
}, error) {
	return _Contracts.Contract.PostTradeDetails(&_Contracts.CallOpts, sizeDelta, tradePrice, orderType, sender)
}

// ProfitLoss is a free data retrieval call binding the contract method 0xb111dfac.
//
// Solidity: function profitLoss(address account) view returns(int256 pnl, bool invalid)
func (_Contracts *ContractsCaller) ProfitLoss(opts *bind.CallOpts, account common.Address) (struct {
	Pnl     *big.Int
	Invalid bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "profitLoss", account)

	outstruct := new(struct {
		Pnl     *big.Int
		Invalid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pnl = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Invalid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// ProfitLoss is a free data retrieval call binding the contract method 0xb111dfac.
//
// Solidity: function profitLoss(address account) view returns(int256 pnl, bool invalid)
func (_Contracts *ContractsSession) ProfitLoss(account common.Address) (struct {
	Pnl     *big.Int
	Invalid bool
}, error) {
	return _Contracts.Contract.ProfitLoss(&_Contracts.CallOpts, account)
}

// ProfitLoss is a free data retrieval call binding the contract method 0xb111dfac.
//
// Solidity: function profitLoss(address account) view returns(int256 pnl, bool invalid)
func (_Contracts *ContractsCallerSession) ProfitLoss(account common.Address) (struct {
	Pnl     *big.Int
	Invalid bool
}, error) {
	return _Contracts.Contract.ProfitLoss(&_Contracts.CallOpts, account)
}

// RemainingMargin is a free data retrieval call binding the contract method 0x9cfbf4e4.
//
// Solidity: function remainingMargin(address account) view returns(uint256 marginRemaining, bool invalid)
func (_Contracts *ContractsCaller) RemainingMargin(opts *bind.CallOpts, account common.Address) (struct {
	MarginRemaining *big.Int
	Invalid         bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "remainingMargin", account)

	outstruct := new(struct {
		MarginRemaining *big.Int
		Invalid         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MarginRemaining = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Invalid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// RemainingMargin is a free data retrieval call binding the contract method 0x9cfbf4e4.
//
// Solidity: function remainingMargin(address account) view returns(uint256 marginRemaining, bool invalid)
func (_Contracts *ContractsSession) RemainingMargin(account common.Address) (struct {
	MarginRemaining *big.Int
	Invalid         bool
}, error) {
	return _Contracts.Contract.RemainingMargin(&_Contracts.CallOpts, account)
}

// RemainingMargin is a free data retrieval call binding the contract method 0x9cfbf4e4.
//
// Solidity: function remainingMargin(address account) view returns(uint256 marginRemaining, bool invalid)
func (_Contracts *ContractsCallerSession) RemainingMargin(account common.Address) (struct {
	MarginRemaining *big.Int
	Invalid         bool
}, error) {
	return _Contracts.Contract.RemainingMargin(&_Contracts.CallOpts, account)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Contracts *ContractsCaller) Resolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "resolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Contracts *ContractsSession) Resolver() (common.Address, error) {
	return _Contracts.Contract.Resolver(&_Contracts.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Contracts *ContractsCallerSession) Resolver() (common.Address, error) {
	return _Contracts.Contract.Resolver(&_Contracts.CallOpts)
}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0x899ffef4.
//
// Solidity: function resolverAddressesRequired() view returns(bytes32[] addresses)
func (_Contracts *ContractsCaller) ResolverAddressesRequired(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "resolverAddressesRequired")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0x899ffef4.
//
// Solidity: function resolverAddressesRequired() view returns(bytes32[] addresses)
func (_Contracts *ContractsSession) ResolverAddressesRequired() ([][32]byte, error) {
	return _Contracts.Contract.ResolverAddressesRequired(&_Contracts.CallOpts)
}

// ResolverAddressesRequired is a free data retrieval call binding the contract method 0x899ffef4.
//
// Solidity: function resolverAddressesRequired() view returns(bytes32[] addresses)
func (_Contracts *ContractsCallerSession) ResolverAddressesRequired() ([][32]byte, error) {
	return _Contracts.Contract.ResolverAddressesRequired(&_Contracts.CallOpts)
}

// UnrecordedFunding is a free data retrieval call binding the contract method 0x917e77f5.
//
// Solidity: function unrecordedFunding() view returns(int256 funding, bool invalid)
func (_Contracts *ContractsCaller) UnrecordedFunding(opts *bind.CallOpts) (struct {
	Funding *big.Int
	Invalid bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "unrecordedFunding")

	outstruct := new(struct {
		Funding *big.Int
		Invalid bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Funding = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Invalid = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// UnrecordedFunding is a free data retrieval call binding the contract method 0x917e77f5.
//
// Solidity: function unrecordedFunding() view returns(int256 funding, bool invalid)
func (_Contracts *ContractsSession) UnrecordedFunding() (struct {
	Funding *big.Int
	Invalid bool
}, error) {
	return _Contracts.Contract.UnrecordedFunding(&_Contracts.CallOpts)
}

// UnrecordedFunding is a free data retrieval call binding the contract method 0x917e77f5.
//
// Solidity: function unrecordedFunding() view returns(int256 funding, bool invalid)
func (_Contracts *ContractsCallerSession) UnrecordedFunding() (struct {
	Funding *big.Int
	Invalid bool
}, error) {
	return _Contracts.Contract.UnrecordedFunding(&_Contracts.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Contracts *ContractsTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Contracts *ContractsSession) AcceptOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.AcceptOwnership(&_Contracts.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Contracts *ContractsTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Contracts.Contract.AcceptOwnership(&_Contracts.TransactOpts)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Contracts *ContractsTransactor) NominateNewOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "nominateNewOwner", _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Contracts *ContractsSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.NominateNewOwner(&_Contracts.TransactOpts, _owner)
}

// NominateNewOwner is a paid mutator transaction binding the contract method 0x1627540c.
//
// Solidity: function nominateNewOwner(address _owner) returns()
func (_Contracts *ContractsTransactorSession) NominateNewOwner(_owner common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.NominateNewOwner(&_Contracts.TransactOpts, _owner)
}

// RebuildCache is a paid mutator transaction binding the contract method 0x74185360.
//
// Solidity: function rebuildCache() returns()
func (_Contracts *ContractsTransactor) RebuildCache(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "rebuildCache")
}

// RebuildCache is a paid mutator transaction binding the contract method 0x74185360.
//
// Solidity: function rebuildCache() returns()
func (_Contracts *ContractsSession) RebuildCache() (*types.Transaction, error) {
	return _Contracts.Contract.RebuildCache(&_Contracts.TransactOpts)
}

// RebuildCache is a paid mutator transaction binding the contract method 0x74185360.
//
// Solidity: function rebuildCache() returns()
func (_Contracts *ContractsTransactorSession) RebuildCache() (*types.Transaction, error) {
	return _Contracts.Contract.RebuildCache(&_Contracts.TransactOpts)
}

// ContractsCacheUpdatedIterator is returned from FilterCacheUpdated and is used to iterate over the raw logs and unpacked data for CacheUpdated events raised by the Contracts contract.
type ContractsCacheUpdatedIterator struct {
	Event *ContractsCacheUpdated // Event containing the contract specifics and raw log

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
func (it *ContractsCacheUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCacheUpdated)
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
		it.Event = new(ContractsCacheUpdated)
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
func (it *ContractsCacheUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCacheUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCacheUpdated represents a CacheUpdated event raised by the Contracts contract.
type ContractsCacheUpdated struct {
	Name        [32]byte
	Destination common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCacheUpdated is a free log retrieval operation binding the contract event 0x88a93678a3692f6789d9546fc621bf7234b101ddb7d4fe479455112831b8aa68.
//
// Solidity: event CacheUpdated(bytes32 name, address destination)
func (_Contracts *ContractsFilterer) FilterCacheUpdated(opts *bind.FilterOpts) (*ContractsCacheUpdatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CacheUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractsCacheUpdatedIterator{contract: _Contracts.contract, event: "CacheUpdated", logs: logs, sub: sub}, nil
}

// WatchCacheUpdated is a free log subscription operation binding the contract event 0x88a93678a3692f6789d9546fc621bf7234b101ddb7d4fe479455112831b8aa68.
//
// Solidity: event CacheUpdated(bytes32 name, address destination)
func (_Contracts *ContractsFilterer) WatchCacheUpdated(opts *bind.WatchOpts, sink chan<- *ContractsCacheUpdated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CacheUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCacheUpdated)
				if err := _Contracts.contract.UnpackLog(event, "CacheUpdated", log); err != nil {
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

// ParseCacheUpdated is a log parse operation binding the contract event 0x88a93678a3692f6789d9546fc621bf7234b101ddb7d4fe479455112831b8aa68.
//
// Solidity: event CacheUpdated(bytes32 name, address destination)
func (_Contracts *ContractsFilterer) ParseCacheUpdated(log types.Log) (*ContractsCacheUpdated, error) {
	event := new(ContractsCacheUpdated)
	if err := _Contracts.contract.UnpackLog(event, "CacheUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Contracts contract.
type ContractsOwnerChangedIterator struct {
	Event *ContractsOwnerChanged // Event containing the contract specifics and raw log

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
func (it *ContractsOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnerChanged)
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
		it.Event = new(ContractsOwnerChanged)
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
func (it *ContractsOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnerChanged represents a OwnerChanged event raised by the Contracts contract.
type ContractsOwnerChanged struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Contracts *ContractsFilterer) FilterOwnerChanged(opts *bind.FilterOpts) (*ContractsOwnerChangedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return &ContractsOwnerChangedIterator{contract: _Contracts.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Contracts *ContractsFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *ContractsOwnerChanged) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnerChanged)
				if err := _Contracts.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address oldOwner, address newOwner)
func (_Contracts *ContractsFilterer) ParseOwnerChanged(log types.Log) (*ContractsOwnerChanged, error) {
	event := new(ContractsOwnerChanged)
	if err := _Contracts.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsOwnerNominatedIterator is returned from FilterOwnerNominated and is used to iterate over the raw logs and unpacked data for OwnerNominated events raised by the Contracts contract.
type ContractsOwnerNominatedIterator struct {
	Event *ContractsOwnerNominated // Event containing the contract specifics and raw log

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
func (it *ContractsOwnerNominatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsOwnerNominated)
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
		it.Event = new(ContractsOwnerNominated)
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
func (it *ContractsOwnerNominatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsOwnerNominatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsOwnerNominated represents a OwnerNominated event raised by the Contracts contract.
type ContractsOwnerNominated struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerNominated is a free log retrieval operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Contracts *ContractsFilterer) FilterOwnerNominated(opts *bind.FilterOpts) (*ContractsOwnerNominatedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return &ContractsOwnerNominatedIterator{contract: _Contracts.contract, event: "OwnerNominated", logs: logs, sub: sub}, nil
}

// WatchOwnerNominated is a free log subscription operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Contracts *ContractsFilterer) WatchOwnerNominated(opts *bind.WatchOpts, sink chan<- *ContractsOwnerNominated) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "OwnerNominated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsOwnerNominated)
				if err := _Contracts.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
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

// ParseOwnerNominated is a log parse operation binding the contract event 0x906a1c6bd7e3091ea86693dd029a831c19049ce77f1dce2ce0bab1cacbabce22.
//
// Solidity: event OwnerNominated(address newOwner)
func (_Contracts *ContractsFilterer) ParseOwnerNominated(log types.Log) (*ContractsOwnerNominated, error) {
	event := new(ContractsOwnerNominated)
	if err := _Contracts.contract.UnpackLog(event, "OwnerNominated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
