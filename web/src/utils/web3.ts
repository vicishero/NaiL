import { BrowserProvider, JsonRpcSigner, ethers } from 'ethers'

declare global {
  interface Window {
    ethereum?: {
      request: (args: { method: string; params?: any[] }) => Promise<any>
      on: (event: string, callback: (...args: any[]) => void) => void
      removeListener: (event: string, callback: (...args: any[]) => void) => void
      isMetaMask?: boolean
      isTokenPocket?: boolean
      chainId?: string
    }
    web3?: {
      currentProvider?: any
    }
  }
}

// 获取钱包提供者（兼容多种钱包）
export function getWalletProvider(): any {
  if (typeof window === 'undefined') return null

  // 优先使用 ethereum
  if (window.ethereum) {
    return window.ethereum
  }

  // 兼容旧版 web3 方式
  if (window.web3?.currentProvider) {
    return window.web3.currentProvider
  }

  return null
}

// 检查是否已安装钱包（兼容多种钱包）
export function hasWallet(): boolean {
  return !!getWalletProvider()
}

// 获取钱包名称
export function getWalletName(): string {
  const provider = getWalletProvider()
  if (!provider) return 'Unknown'
  if (provider.isMetaMask) return 'MetaMask'
  if (provider.isTokenPocket) return 'TokenPocket'
  return 'Wallet'
}

// 获取钱包提供者
export function getProvider(): BrowserProvider | null {
  const provider = getWalletProvider()
  if (!provider) return null
  return new BrowserProvider(provider)
}

// 连接钱包获取签名者
export async function connectWallet(): Promise<{
  signer: JsonRpcSigner
  address: string
} | null> {
  const provider = getProvider()
  if (!provider) return null

  try {
    // 请求账户授权
    await provider.send('eth_requestAccounts', [])
    const signer = await provider.getSigner()
    const address = await signer.getAddress()

    return { signer, address }
  } catch (error) {
    console.error('连接钱包失败:', error)
    return null
  }
}

// 签名消息（兼容不同钱包）
export async function signMessage(signer: JsonRpcSigner, message: string): Promise<string | null> {
  try {
    const signature = await signer.signMessage(message)
    return signature
  } catch (error) {
    console.error('签名消息失败:', error)
    return null
  }
}

// 验证签名（客户端预览验证，实际在服务端验证）
export function verifySignature(
  message: string,
  signature: string,
  expectedAddress: string
): boolean {
  try {
    const recoveredAddress = ethers.verifyMessage(message, signature)
    return recoveredAddress.toLowerCase() === expectedAddress.toLowerCase()
  } catch (error) {
    console.error('验证签名失败:', error)
    return false
  }
}
