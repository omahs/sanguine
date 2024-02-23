package testutil

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/contracts"
	"github.com/synapsecns/sanguine/ethergo/deployer"
	"github.com/synapsecns/sanguine/ethergo/manager"
	"github.com/synapsecns/sanguine/sin-executor/contracts/interchainclient"
	"github.com/synapsecns/sanguine/sin-executor/contracts/interchaindb"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/executionfeesmock"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/executionservicemock"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/interchainappmock"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/interchainmodulemock"
	"github.com/synapsecns/sanguine/sin-executor/contracts/mocks/optionslibexport"
	"testing"
)

// DeployManager wraps DeployManager and allows typed contract handles to be returned.
type DeployManager struct {
	*manager.DeployerManager
}

// NewDeployManager creates a new DeployManager.
func NewDeployManager(t *testing.T) *DeployManager {
	t.Helper()

	return &DeployManager{manager.NewDeployerManager(t, interchainClientDeployer, interchainDBDeployer, interchainModuleMockDeployer, interchainAppMockDeployer, optionsLibMock, executionServiceMock, executionFeesMock)}
}

var (
	// TODO: this looks horrible, rip out the functional deployer.
	interchainClientDeployer = deployer.NewFunctionalDeployer(InterchainClient,
		func(ctx context.Context, helpers deployer.IFunctionalDeployer,
			transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {

			deployAddress, deployTx, deployIface, err := interchainclient.DeployInterchainClientV1(transactOps, backend)
			if err != nil {
				return common.Address{}, nil, nil, fmt.Errorf("could not deploy interchain client: %w", err)
			}

			helpers.Backend().WaitForConfirmation(ctx, deployTx)

			// set the interchain db
			idb := helpers.Registry().Get(ctx, InterchainDB)
			transactor := helpers.Backend().GetTxContext(ctx, &transactOps.From)
			idbSet, err := deployIface.SetInterchainDB(transactor.TransactOpts, idb.Address())
			if err != nil {
				return common.Address{}, nil, nil, fmt.Errorf("could not set interchain db: %w", err)
			}
			helpers.Backend().WaitForConfirmation(ctx, idbSet)

			// set the execution service
			em := helpers.Registry().Get(ctx, ExecutionFeesMock)
			emSet, err := deployIface.SetExecutionFees(transactor.TransactOpts, em.Address())
			if err != nil {
				return common.Address{}, nil, nil, fmt.Errorf("could not set execution fees: %w", err)
			}
			helpers.Backend().WaitForConfirmation(ctx, emSet)

			return deployAddress, deployTx, deployIface, nil
		}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
			return interchainclient.NewInterchainClientRef(address, backend)
		}, []contracts.ContractType{InterchainDB, ExecutionFeesMock})

	interchainDBDeployer = deployer.NewFunctionalDeployer(InterchainDB,
		func(ctx context.Context, helpers deployer.IFunctionalDeployer, transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
			deployAddress, deployTx, deployIface, err := interchaindb.DeployInterchainDB(transactOps, backend)
			if err != nil {
				return common.Address{}, nil, nil, fmt.Errorf("could not deploy interchain db: %w", err)
			}
			helpers.Backend().WaitForConfirmation(ctx, deployTx)

			return deployAddress, deployTx, deployIface, nil
		}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
			return interchaindb.NewInterchainDBRef(address, backend)
		}, []contracts.ContractType{})

	interchainModuleMockDeployer = deployer.NewFunctionalDeployer(InterchainModuleMock, func(ctx context.Context, helpers deployer.IFunctionalDeployer, transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return interchainmodulemock.DeployInterchainModuleMock(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return interchainmodulemock.NewInterchainModuleMockRef(address, backend)
	}, []contracts.ContractType{})

	interchainAppMockDeployer = deployer.NewFunctionalDeployer(InterchainAppMock, func(ctx context.Context, helpers deployer.IFunctionalDeployer, transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		appAddress, appTx, appMock, err := interchainappmock.DeployInterchainAppMock(transactOps, backend)
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not deploy interchain app mock: %w", err)
		}
		helpers.Backend().WaitForConfirmation(ctx, appTx)

		txOpts := helpers.Backend().GetTxContext(ctx, &transactOps.From)
		setTx, err := appMock.SetReceivingModule(txOpts.TransactOpts, helpers.Registry().Get(ctx, InterchainModuleMock).Address())
		if err != nil {
			return common.Address{}, nil, nil, fmt.Errorf("could not set receiving module: %w", err)
		}
		helpers.Backend().WaitForConfirmation(ctx, setTx)

		return appAddress, appTx, appMock, nil
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return interchainappmock.NewInterchainAppMockRef(address, backend)
	}, []contracts.ContractType{InterchainModuleMock})

	optionsLibMock = deployer.NewFunctionalDeployer(OptionsLib, func(ctx context.Context, helpers deployer.IFunctionalDeployer, transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return optionslibexport.DeployOptionsLibMocks(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return optionslibexport.NewOptionsLibExportRef(address, backend)
	}, []contracts.ContractType{})

	executionServiceMock = deployer.NewFunctionalDeployer(ExecutionServiceMock, func(ctx context.Context, helpers deployer.IFunctionalDeployer, transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return executionservicemock.DeployExecutionServiceMock(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return executionservicemock.NewExecutionserviceMockRef(address, backend)
	}, []contracts.ContractType{})

	executionFeesMock = deployer.NewFunctionalDeployer(ExecutionFeesMock, func(ctx context.Context, helpers deployer.IFunctionalDeployer, transactOps *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, interface{}, error) {
		return executionfeesmock.DeployExecutionFeesMock(transactOps, backend)
	}, func(address common.Address, backend bind.ContractBackend) (interface{}, error) {
		return executionfeesmock.NewExecutionfeesMockRef(address, backend)
	}, []contracts.ContractType{})
)