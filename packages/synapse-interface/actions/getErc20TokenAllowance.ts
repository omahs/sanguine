import { wagmiConfig } from '@/constants/wagmi'
import { readContract } from '@wagmi/core'
import { Address, erc20Abi } from 'viem'

export const getErc20TokenAllowance = async ({
  address,
  chainId,
  tokenAddress,
  spender,
}: {
  address: Address
  chainId: number
  tokenAddress: Address
  spender: Address
}): Promise<bigint> => {
  const allowance = await readContract(wagmiConfig, {
    chainId: chainId as any,
    address: tokenAddress,
    abi: erc20Abi,
    functionName: 'allowance',
    args: [address, spender],
  })

  return allowance
}
