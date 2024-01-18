import '@styles/global.css'
import '@rainbow-me/rainbowkit/styles.css'
import type { AppProps } from 'next/app'
import Head from 'next/head'
import '@/patch'
import { Analytics } from '@vercel/analytics/react'
import { PersistGate } from 'redux-persist/integration/react'

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
import { RainbowKitProvider, darkTheme } from '@rainbow-me/rainbowkit'
import * as CHAINS from '@constants/chains/master'
import { SynapseProvider } from '@/utils/providers/SynapseProvider'
import CustomToaster from '@/components/toast'
import { SegmentAnalyticsProvider } from '@/contexts/SegmentAnalyticsProvider'

import { Provider } from 'react-redux'
import { store, persistor } from '@/store/store'
import { UserProvider } from '@/contexts/UserProvider'

import ApplicationUpdater from '@/slices/application/updater'
import BridgeUpdater from '@/slices/bridge/updater'
import PortfolioUpdater from '@/slices/portfolio/updater'
import TransactionsUpdater from '@/slices/transactions/updater'
import _TransactionsUpdater from '@/slices/_transactions/updater'
import { wagmiConfig } from '../constants/wagmi'
import { WagmiProvider } from 'wagmi'

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
const chainsMatured = []
for (const chain of rawChains) {
  const configChain = Object.values(CHAINS).filter(
    (chainObj) => chainObj.id === chain.id
  )[0]

  chainsMatured.push({
    ...chain,
    iconUrl: configChain.chainImg.src,
    configRpc: configChain.rpcUrls.primary,
    fallbackRpc: configChain.rpcUrls.fallback,
  })
}

function Updaters() {
  return (
    <>
      <ApplicationUpdater />
      <PortfolioUpdater />
      <TransactionsUpdater />
      <_TransactionsUpdater />
      <BridgeUpdater />
    </>
  )
}

const App = ({ Component, pageProps }: AppProps) => {
  return (
    <WagmiProvider config={wagmiConfig}>
      <RainbowKitProvider chains={chainsMatured} theme={darkTheme()}>
        <SynapseProvider chains={chainsMatured}>
          <Provider store={store}>
            <PersistGate loading={null} persistor={persistor}>
              <SegmentAnalyticsProvider>
                <UserProvider>
                  <Updaters />
                  <Component {...pageProps} />
                  <Analytics />
                  <CustomToaster />
                </UserProvider>
              </SegmentAnalyticsProvider>
            </PersistGate>
          </Provider>
        </SynapseProvider>
      </RainbowKitProvider>
    </WagmiProvider>
  )
}

export default App
