import _ from 'lodash'


import type { Token } from '@/utils/types'

import { PAUSED_TO_CHAIN_IDS } from '@/constants/chains'
import { findTokenByRouteSymbol } from '@/utils/tokens/findTokenByRouteSymbol'
import { flattenPausedTokens } from '@/utils/tokens/flattenPausedTokens'
import { getSymbol } from '@/utils/route/getSymbol'
import { getSwapFromChainIds } from './getSwapFromChainIds'
import { getSwapFromTokens } from './getSwapFromTokens'
import { getSwapToTokens } from './getSwapToTokens'

export const getSwapPossibilities = ({
  fromChainId,
  fromToken,
  toChainId,
  toToken,
}: {
  fromChainId?: number
  fromToken?: Token
  toChainId?: number
  toToken?: Token
}) => {
  const fromTokenRouteSymbol = fromToken && fromToken.routeSymbol
  const toTokenRouteSymbol = toToken && toToken.routeSymbol

  const fromChainIds: number[] = getSwapFromChainIds({
    fromChainId,
    fromTokenRouteSymbol,
    toChainId,
    toTokenRouteSymbol,
  })

  const fromTokens: Token[] = _(
    getSwapFromTokens({
      fromChainId,
      fromTokenRouteSymbol,
      toChainId,
      toTokenRouteSymbol,
    })
  )
    .difference(flattenPausedTokens())
    .map(getSymbol)
    .uniq()
    .map((symbol) => findTokenByRouteSymbol(symbol))
    .compact()
    .value()

  const toTokens: Token[] = _(
    getSwapToTokens({
      fromChainId,
      fromTokenRouteSymbol,
      toChainId,
      toTokenRouteSymbol,
    })
  )
    .difference(flattenPausedTokens())
    .filter((token) => {
      return !PAUSED_TO_CHAIN_IDS.some((value) => token.endsWith(`-${value}`))
    })
    .map(getSymbol)
    .uniq()
    .map((symbol) => findTokenByRouteSymbol(symbol))
    .compact()
    .value()

  return {
    fromChainId,
    fromToken,
    toChainId,
    toToken,
    fromChainIds,
    fromTokens,
    toTokens,
  }
}