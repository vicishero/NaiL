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

// 缓存 provider 实例
let cachedProvider: BrowserProvider | null = null

// localStorage key
const WALLET_ADDRESS_KEY = 'PAOPAO_WALLET_ADDRESS'

// 登录成功时记录钱包地址（持久化到 localStorage）
export function recordWalletAddressOnLogin(address: string) {
  const normalized = address.toLowerCase()
  console.log('[Wallet] 登录成功，记录钱包地址:', normalized)
  localStorage.setItem(WALLET_ADDRESS_KEY, normalized)
}

// 获取已记录的钱包地址
export function getLoggedInWalletAddress(): string | null {
  return localStorage.getItem(WALLET_ADDRESS_KEY) || null
}

// 清除记录的钱包地址（退出登录时调用）
export function clearLoggedInWalletAddress() {
  console.log('[Wallet] 清除记录的钱包地址')
  localStorage.removeItem(WALLET_ADDRESS_KEY)
}

// 获取钱包提供者（兼容多种钱包）
export function getWalletProvider(): any {
  if (typeof window === 'undefined') return null

  console.log('[Wallet] 检测钱包 provider, window.ethereum:', !!window.ethereum)

  // TokenPocket 钱包可能注入在多个位置
  const tpProviders = [
    // TokenPocket 标准位置
    () => (window as any).ethereum,
    // TokenPocket 移动端可能的位置
    () => (window as any).tp,
    // TokenPocket DApp 浏览器
    () => (window as any).TokenPocket,
    // 旧版 web3 方式
    () => (window as any).web3?.currentProvider,
    // 通用以太坊 provider
    () => (window as any).web3Provider,
  ]

  for (const getProvider of tpProviders) {
    try {
      const provider = getProvider()
      if (provider) {
        console.log('[Wallet] 找到钱包 provider:', provider.isTokenPocket ? 'TokenPocket' : provider.constructor.name)
        return provider
      }
    } catch (e) {
      console.log('[Wallet] 获取 provider 失败:', e)
    }
  }

  console.log('[Wallet] 未找到任何钱包 provider')
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

// 强制刷新 provider（解决 MetaMask 缓存问题）
function forceRefreshProvider(): BrowserProvider | null {
  console.log('[Wallet] 强制刷新 provider 实例')
  const rawProvider = getWalletProvider()
  if (!rawProvider) return null

  // 断开旧 provider 的所有监听
  if (cachedProvider) {
    try {
      // @ts-ignore - 清理旧监听
      cachedProvider.removeAllListeners && cachedProvider.removeAllListeners()
    } catch (e) {
      console.log('[Wallet] 清理旧 provider 监听失败:', e)
    }
  }

  cachedProvider = new BrowserProvider(rawProvider)
  return cachedProvider
}

// 重置 provider 缓存（切换钱包时使用）
export function resetProviderCache() {
  console.log('[Wallet] 重置 provider 缓存')
  if (cachedProvider) {
    try {
      // @ts-ignore - 清理旧监听
      cachedProvider.removeAllListeners && cachedProvider.removeAllListeners()
    } catch (e) {
      console.log('[Wallet] 清理 provider 监听失败:', e)
    }
  }
  cachedProvider = null
}

// 连接钱包获取签名者（确保获取到当前选中的账户）
export async function connectWallet(): Promise<{
  signer: JsonRpcSigner
  address: string
} | null> {
  // 每次连接前强制刷新 provider，解决钱包缓存问题
  const provider = forceRefreshProvider()
  if (!provider) {
    console.log('[Wallet] 未检测到钱包提供者')
    return null
  }

  try {
    console.log('[Wallet] 开始请求账户授权...')

    let accounts: string[] = []
    const rawProvider = getWalletProvider()

    // TokenPocket 钱包特殊处理：先尝试直接调用
    if (rawProvider?.isTokenPocket) {
      console.log('[Wallet] 检测到 TokenPocket 钱包，使用兼容模式')
      try {
        // TokenPocket 可能需要先 enable
        if (typeof rawProvider.enable === 'function') {
          await rawProvider.enable()
        }
        // 尝试多种方式获取账户
        try {
          accounts = await provider.send('eth_requestAccounts', [])
        } catch {
          try {
            accounts = await provider.send('eth_accounts', [])
          } catch {
            // 最后尝试直接调用 requestAccounts
            accounts = await rawProvider.request({ method: 'eth_requestAccounts' })
          }
        }
      } catch (e) {
        console.error('[Wallet] TokenPocket 获取账户失败:', e)
        throw e
      }
    } else {
      // 标准钱包（MetaMask 等）
      accounts = await provider.send('eth_requestAccounts', [])
    }

    console.log('[Wallet] 钱包返回的账户列表:', accounts)

    if (!accounts || accounts.length === 0) {
      console.log('[Wallet] 没有可用的账户')
      return null
    }

    // 显式获取最新的签名者
    const signer = await provider.getSigner()
    const address = await signer.getAddress()
    console.log('[Wallet] 获取到的签名者地址:', address)

    return { signer, address }
  } catch (error) {
    console.error('[Wallet] 连接钱包失败:', error)
    // 出错时重置缓存，下次重试
    resetProviderCache()
    return null
  }
}

// 签名消息（兼容不同钱包）
export async function signMessage(signer: JsonRpcSigner, message: string): Promise<string | null> {
  try {
    const signature = await signer.signMessage(message)
    return signature
  } catch (error) {
    console.error('[Wallet] 签名消息失败:', error)
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
    console.error('[Wallet] 验证签名失败:', error)
    return false
  }
}
