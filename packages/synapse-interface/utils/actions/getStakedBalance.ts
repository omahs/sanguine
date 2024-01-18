import { readContracts } from '@wagmi/core'
import { Address } from 'viem'
import { MINICHEF_ABI } from '@abis/miniChef'
import { Token } from '@types'
import { wagmiConfig } from '@/constants/wagmi'

type UserInfoResult = {
  result: [amount: bigint, rewardDebt: bigint]
  status: string
}

type PendingSynapseResult = {
  result: bigint
  status: string
}

export const getStakedBalance = async (
  address: Address,
  chainId: number,
  poolId: number,
  pool: Token
) => {
  return {
    amount: 0n,
    reward: 0n,
  }
  // const miniChefContractAddress: Address = pool.miniChefAddress as Address
  // try {
  //   const data = await readContracts(wagmiConfig, {
  //     contracts: [
  //       {
  //         address: miniChefContractAddress,
  //         abi: MINICHEF_ABI,
  //         functionName: 'userInfo',
  //         chainId,
  //         args: [BigInt(poolId), address],
  //       },
  //       {
  //         address: miniChefContractAddress,
  //         abi: MINICHEF_ABI,
  //         functionName: 'pendingSynapse',
  //         chainId,
  //         args: [BigInt(poolId), address],
  //       },
  //     ],
  //   })

  //   const userInfo: UserInfoResult = data[0]
  //   const pendingSynapse: PendingSynapseResult = data[1]

  //   if (userInfo.status === 'success') {
  //     return {
  //       amount: userInfo?.result[0] ?? 0n,
  //       reward: pendingSynapse?.result ?? 0n,
  //     }
  //   } else {
  //     return {
  //       amount: 0n,
  //       reward: 0n,
  //     }
  //   }
  // } catch (error) {
  //   console.error('Error from getStakedBalance: ', error)
  //   return { amount: 0n, reward: 0n }
  // }
}
