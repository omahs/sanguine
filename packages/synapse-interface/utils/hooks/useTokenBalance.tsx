import { useBalance, useAccount } from 'wagmi'
import { Address } from 'viem'
import { Token } from '../types'

export const useTokenBalance = (token: Token, chainId: number) => {
  const { address } = useAccount()

  const tokenAddress = token && token.addresses[chainId]

  const balance = useBalance({
    address,
    token: tokenAddress as Address,
    chainId: chainId as any,
    // watch: true,
  })

  return balance
}
