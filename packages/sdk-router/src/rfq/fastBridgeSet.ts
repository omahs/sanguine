import { Provider } from '@ethersproject/abstract-provider'

import {
  BridgeQuote,
  BridgeRoute,
  SynapseModule,
  SynapseModuleSet,
} from '../module'
import { ChainProvider } from '../router'
import { FastBridge } from './fastBridge'

export class FastBridgeSet extends SynapseModuleSet {
  public readonly bridgeModuleName = 'SynapseRFQ'
  public readonly allEvents = ['BridgeRequestedEvent', 'BridgeRelayedEvent']

  public fastBridges: {
    [chainId: number]: FastBridge
  }
  public providers: {
    [chainId: number]: Provider
  }

  constructor(chains: ChainProvider[], addressMap: AddressMap) {
    super()
    this.fastBridges = {}
    this.providers = {}
    chains.forEach(({ chainId, provider }) => {
      const address = addressMap[chainId]
      // Skip chains without a FastBridge address
      if (address) {
        this.fastBridges[chainId] = new FastBridge(chainId, provider, address)
        this.providers[chainId] = provider
      }
    })
  }

  /**
   * @inheritdoc SynapseModuleSet.getModule
   */
  public getModule(chainId: number): SynapseModule | undefined {
    return this.fastBridges[chainId]
  }

  /**
   * @inheritdoc RouterSet.getOriginAmountOut
   */
  public getEstimatedTime(chainId: number): number {
    // TODO: figure out estimated time for FastBridge
  }

  /**
   * @inheritdoc SynapseModuleSet.getBridgeRoutes
   */
  public async getBridgeRoutes(
    originChainId: number,
    destChainId: number,
    tokenIn: string,
    tokenOut: string,
    amountIn: BigintIsh
  ): Promise<BridgeRoute[]> {
    // TODO
  }

  /**
   * @inheritdoc SynapseModuleSet.finalizeBridgeRoute
   */
  public async finalizeBridgeRoute(
    bridgeRoute: BridgeRoute,
    deadline?: BigNumber
  ): Promise<BridgeQuote> {
    // TODO
  }
}
