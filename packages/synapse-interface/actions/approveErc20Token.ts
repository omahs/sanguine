import {
  simulateContract,
  waitForTransaction,
  waitForTransactionReceipt,
  writeContract,
} from '@wagmi/core'
import { TransactionReceipt, Address, erc20Abi } from 'viem'

import { MAX_UINT256 } from '@/constants'
import { USDT } from '@/constants/tokens/bridgeable'
import { USDT_ABI } from '@/constants/abis/usdtAbi'
import { ETH as Ethereum } from '@/constants/chains/master'
import { wagmiConfig } from '@/constants/wagmi'

export const approveErc20Token = async ({
  chainId,
  amount,
  spender,
  tokenAddress,
}: {
  chainId: number
  amount: bigint
  spender: Address
  tokenAddress: Address
}) => {
  let abi

  if (tokenAddress === USDT.addresses[Ethereum.id]) {
    abi = USDT_ABI
  } else {
    abi = erc20Abi
  }

  const { request } = await simulateContract(wagmiConfig, {
    chainId: chainId as any,
    address: tokenAddress,
    abi,
    functionName: 'approve',
    args: [spender, amount ?? MAX_UINT256],
  })

  const hash = await writeContract(wagmiConfig, request)

  const txReceipt = await waitForTransactionReceipt(wagmiConfig, { hash })

  return txReceipt
}
