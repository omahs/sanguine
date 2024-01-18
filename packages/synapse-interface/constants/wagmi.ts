import '@/patch'

import {
  boba,
  cronos,
  dfk,
  dogechain,
  klaytn,
  metis,
  aurora,
  canto,
  base,
} from '@constants/extraWagmiChains'
import { createConfig, http } from 'wagmi'
import {
  arbitrum,
  avalanche,
  bsc,
  fantom,
  harmonyOne,
  mainnet,
  moonbeam,
  moonriver,
  optimism,
  polygon,
} from 'wagmi/chains'
import * as CHAINS from '@constants/chains/master'

import _TransactionsUpdater from '@/slices/_transactions/updater'
import { injected } from '@wagmi/core'

const rawChains = [
  mainnet,
  arbitrum,
  aurora,
  avalanche,
  base,
  bsc,
  canto,
  fantom,
  harmonyOne,
  metis,
  moonbeam,
  moonriver,
  optimism,
  polygon,
  klaytn,
  cronos,
  dfk,
  dogechain,
  boba,
]

// Add custom icons
const chains = []
for (const chain of rawChains) {
  const configChain = Object.values(CHAINS).filter(
    (chainObj) => chainObj.id === chain.id
  )[0]

  chains.push({
    ...chain,
    iconUrl: configChain.chainImg.src,
    configRpc: configChain.rpcUrls.primary,
    fallbackRpc: configChain.rpcUrls.fallback,
  })
}

export const connectors = [
  injected({
    shimDisconnect: true,
    unstable_shimAsyncInject: true,
    target: 'metaMask',
  }),
  injected({
    shimDisconnect: true,
    unstable_shimAsyncInject: true,
    target: 'rainbow',
  }),
]

export const wagmiConfig = createConfig({
  ssr: true,
  chains: [mainnet, optimism],
  transports: {
    [mainnet.id]: http(),
    [optimism.id]: http(),
  },
})

declare module 'wagmi' {
  interface Register {
    config: typeof wagmiConfig
  }
}
