import {
  simulateContract,
  waitForTransactionReceipt,
  writeContract,
} from '@wagmi/core'
import { Address, TransactionReceipt } from 'viem'

import { MINICHEF_ABI } from '@/constants/abis/miniChef'
import { useConfig } from 'wagmi'

export const harvestLpPool = async ({
  address,
  chainId,
  poolId,
  lpAddress,
}: {
  address: Address
  chainId: number
  poolId: number
  lpAddress: Address
}) => {
  const config = useConfig()
  const { request } = await simulateContract(config, {
    chainId: chainId as any,
    address: lpAddress,
    abi: MINICHEF_ABI,
    functionName: 'harvest',
    args: [poolId, address],
  })

  const hash = await writeContract(config, request)
  const txReceipt = await waitForTransactionReceipt(config, { hash })

  return txReceipt
}
