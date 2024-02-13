// autogenerated file

package arbgasinfo

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// IArbGasInfoCaller ...
type IArbGasInfoCaller interface {
	// GetAmortizedCostCapBips is a free data retrieval call binding the contract method 0x7a7d6beb.
	//
	// Solidity: function getAmortizedCostCapBips() view returns(uint64)
	GetAmortizedCostCapBips(opts *bind.CallOpts) (uint64, error)
	// GetCurrentTxL1GasFees is a free data retrieval call binding the contract method 0xc6f7de0e.
	//
	// Solidity: function getCurrentTxL1GasFees() view returns(uint256)
	GetCurrentTxL1GasFees(opts *bind.CallOpts) (*big.Int, error)
	// GetGasAccountingParams is a free data retrieval call binding the contract method 0x612af178.
	//
	// Solidity: function getGasAccountingParams() view returns(uint256, uint256, uint256)
	GetGasAccountingParams(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, error)
	// GetGasBacklog is a free data retrieval call binding the contract method 0x1d5b5c20.
	//
	// Solidity: function getGasBacklog() view returns(uint64)
	GetGasBacklog(opts *bind.CallOpts) (uint64, error)
	// GetGasBacklogTolerance is a free data retrieval call binding the contract method 0x25754f91.
	//
	// Solidity: function getGasBacklogTolerance() view returns(uint64)
	GetGasBacklogTolerance(opts *bind.CallOpts) (uint64, error)
	// GetL1BaseFeeEstimate is a free data retrieval call binding the contract method 0xf5d6ded7.
	//
	// Solidity: function getL1BaseFeeEstimate() view returns(uint256)
	GetL1BaseFeeEstimate(opts *bind.CallOpts) (*big.Int, error)
	// GetL1BaseFeeEstimateInertia is a free data retrieval call binding the contract method 0x29eb31ee.
	//
	// Solidity: function getL1BaseFeeEstimateInertia() view returns(uint64)
	GetL1BaseFeeEstimateInertia(opts *bind.CallOpts) (uint64, error)
	// GetL1FeesAvailable is a free data retrieval call binding the contract method 0x5b39d23c.
	//
	// Solidity: function getL1FeesAvailable() view returns(uint256)
	GetL1FeesAvailable(opts *bind.CallOpts) (*big.Int, error)
	// GetL1GasPriceEstimate is a free data retrieval call binding the contract method 0x055f362f.
	//
	// Solidity: function getL1GasPriceEstimate() view returns(uint256)
	GetL1GasPriceEstimate(opts *bind.CallOpts) (*big.Int, error)
	// GetL1PricingSurplus is a free data retrieval call binding the contract method 0x520acdd7.
	//
	// Solidity: function getL1PricingSurplus() view returns(int256)
	GetL1PricingSurplus(opts *bind.CallOpts) (*big.Int, error)
	// GetMinimumGasPrice is a free data retrieval call binding the contract method 0xf918379a.
	//
	// Solidity: function getMinimumGasPrice() view returns(uint256)
	GetMinimumGasPrice(opts *bind.CallOpts) (*big.Int, error)
	// GetPerBatchGasCharge is a free data retrieval call binding the contract method 0x6ecca45a.
	//
	// Solidity: function getPerBatchGasCharge() view returns(int64)
	GetPerBatchGasCharge(opts *bind.CallOpts) (int64, error)
	// GetPricesInArbGas is a free data retrieval call binding the contract method 0x02199f34.
	//
	// Solidity: function getPricesInArbGas() view returns(uint256, uint256, uint256)
	GetPricesInArbGas(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, error)
	// GetPricesInArbGasWithAggregator is a free data retrieval call binding the contract method 0x7a1ea732.
	//
	// Solidity: function getPricesInArbGasWithAggregator(address aggregator) view returns(uint256, uint256, uint256)
	GetPricesInArbGasWithAggregator(opts *bind.CallOpts, aggregator common.Address) (*big.Int, *big.Int, *big.Int, error)
	// GetPricesInWei is a free data retrieval call binding the contract method 0x41b247a8.
	//
	// Solidity: function getPricesInWei() view returns(uint256, uint256, uint256, uint256, uint256, uint256)
	GetPricesInWei(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error)
	// GetPricesInWeiWithAggregator is a free data retrieval call binding the contract method 0xba9c916e.
	//
	// Solidity: function getPricesInWeiWithAggregator(address aggregator) view returns(uint256, uint256, uint256, uint256, uint256, uint256)
	GetPricesInWeiWithAggregator(opts *bind.CallOpts, aggregator common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error)
	// GetPricingInertia is a free data retrieval call binding the contract method 0x3dfb45b9.
	//
	// Solidity: function getPricingInertia() view returns(uint64)
	GetPricingInertia(opts *bind.CallOpts) (uint64, error)
}