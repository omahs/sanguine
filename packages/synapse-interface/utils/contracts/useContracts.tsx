import { useAccount } from 'wagmi'
import { Address, erc20Abi, getContract } from 'viem'
import { Token } from '../types'
import { wagmiConfig } from '@/constants/wagmi'

/**
 * @param {Token} token token the token used
 */
export function useTokenContract({ token }: { token: Token }) {
  // const { chain } = useAccount()

  // const contract = getContract({
  //   address: token ? (token.addresses[chain.id] as Address) : `0x`,
  //   abi: erc20Abi,
  // })

  return null
  // return contract
}
