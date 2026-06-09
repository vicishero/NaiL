import { BrowserProvider, JsonRpcSigner, ethers } from 'ethers'

declare global {
  interface Window {
    ethereum?: {
      request: (args: { method: string; params?: any[] }) => Promise<any>
      on: (event: string, callback: (...args: any[]) => void) => void
      removeListener: (event: string, callback: (...args: any[]) => void) => void
    }
  }
}

// 检查是否已安装钱包
export function hasWallet(): boolean {
  return typeof window !== 'undefined' && window.ethereum !== undefined
}

// 获取钱包提供者
export function getProvider(): BrowserProvider | null {
  if (!window.ethereum) return null
  return new BrowserProvider(window.ethereum)
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

// 签名消息
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
