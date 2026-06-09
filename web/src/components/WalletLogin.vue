<template>
  <div class="wallet-login">
    <div v-if="!hasWalletInstalled" class="wallet-not-found">
      <n-empty description="未检测到钱包插件">
        <template #extra>
          <n-button type="primary" @click="openMetaMask">
            安装 MetaMask
          </n-button>
        </template>
      </n-empty>
    </div>

    <div v-else-if="!address" class="connect-wallet">
      <n-button
        type="primary"
        size="large"
        block
        :loading="connecting"
        @click="handleConnect"
      >
        <template #icon>
          <n-icon size="20">
            <svg viewBox="0 0 35 33" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path d="M32.9582 1L19.8241 10.8119L22.2561 1H32.9582Z" fill="#E17726"/>
              <path d="M2.66797 1L15.1829 10.8119L12.7509 1H2.66797Z" fill="#E27625"/>
              <path d="M32.958 1L22.2558 7.43864L19.8238 12.25L32.958 1Z" fill="#E27625"/>
              <path d="M2.66797 1L12.7509 12.25L10.3189 7.43864L2.66797 1Z" fill="#E27625"/>
              <path d="M19.8242 25.8864V21.2386L22.2562 18.5H32.9582V25.8864L28.0742 31L19.8242 25.8864Z" fill="#E27625"/>
              <path d="M12.75 18.5L15.182 21.2386V25.8864L6.93203 31L2.04797 25.8864V18.5H12.75Z" fill="#E27625"/>
            </svg>
          </n-icon>
        </template>
        连接钱包登录
      </n-button>
    </div>

    <div v-else class="sign-message">
      <n-alert type="info" title="签名验证">
        请在钱包弹窗中签名以完成登录
      </n-alert>
      <div class="mt-4">
        <n-button
          type="primary"
          size="large"
          block
          :loading="signing"
          @click="handleSign"
        >
          签名并登录
        </n-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch } from 'vue'
import { useStoreMain } from '@/store/main'
import { TOKEN_KEY, useStoreUser } from '@/store/user'
import { userInfo } from '@/api/auth'
import { connectWallet, signMessage, getWalletProvider } from '@/utils/web3'
import { Api } from '@/utils/request'

const storeMain = useStoreMain()
const storeUser = useStoreUser()

const connecting = ref(false)
const signing = ref(false)
const address = ref('')
const nonce = ref('')
const signMessageText = ref('')
// 是否安装了钱包
const hasWalletInstalled = ref(false)
// 重试计数
const retryCount = ref(0)

// SessionStorage key
const WALLET_LOGIN_STATE_KEY = 'paopao_wallet_login_state'

// 保存登录状态到 sessionStorage（处理钱包刷新页面问题）
function saveLoginState(state: any) {
  try {
    sessionStorage.setItem(WALLET_LOGIN_STATE_KEY, JSON.stringify(state))
  } catch (e) {
    console.error('保存登录状态失败:', e)
  }
}

// 从 sessionStorage 恢复登录状态
function restoreLoginState(): any | null {
  try {
    const state = sessionStorage.getItem(WALLET_LOGIN_STATE_KEY)
    if (state) {
      return JSON.parse(state)
    }
  } catch (e) {
    console.error('恢复登录状态失败:', e)
  }
  return null
}

// 清除登录状态
function clearLoginState() {
  try {
    sessionStorage.removeItem(WALLET_LOGIN_STATE_KEY)
  } catch (e) {
    console.error('清除登录状态失败:', e)
  }
}

// 重置组件状态
function resetState() {
  address.value = ''
  nonce.value = ''
  signMessageText.value = ''
  connecting.value = false
  signing.value = false
  retryCount.value = 0
  clearLoginState()
}

// 获取钱包地址并请求 nonce
async function fetchWalletAndNonce(): Promise<boolean> {
  try {
    const result = await connectWallet()
    if (result) {
      address.value = result.address

      // 获取 nonce
      const nonceResult = await Api.v1.auth.post.walletNonce({
        address: result.address,
      })
      nonce.value = nonceResult.nonce
      signMessageText.value = nonceResult.message

      return true
    }
  } catch (error: any) {
    console.error('获取钱包信息失败:', error)
    // 如果是用户拒绝请求，不显示错误提示
    if (error?.code !== 4001) {
      window.$message.error(error?.message || '连接钱包失败')
    }
  }
  return false
}

// 监听登录弹窗显示，自动连接钱包
watch(() => storeMain.authModalShow, async (show) => {
  if (show) {
    // 每次打开弹窗时重新检测钱包是否安装
    hasWalletInstalled.value = !!getWalletProvider()

    // 弹窗打开时，自动尝试连接钱包
    if (hasWalletInstalled.value) {
      // 先尝试恢复之前的登录状态（处理钱包刷新页面问题）
      const savedState = restoreLoginState()
      if (savedState?.address && savedState?.nonce) {
        address.value = savedState.address
        nonce.value = savedState.nonce
        signMessageText.value = savedState.signMessageText
        return
      }

      // 稍微延迟确保弹窗显示后再连接
      setTimeout(async () => {
        await fetchWalletAndNonce()
      }, 500)
    }
  } else {
    // 弹窗关闭时重置状态
    resetState()
  }
})

// 检查是否安装了钱包（多次尝试，兼容移动端钱包注入慢）
function checkWalletInstalled() {
  const maxRetries = 5
  let retry = 0

  const check = () => {
    hasWalletInstalled.value = !!getWalletProvider()
    if (!hasWalletInstalled.value && retry < maxRetries) {
      retry++
      setTimeout(check, 200)
    }
  }
  check()
}

onMounted(() => {
  // 多次检测，确保钱包已注入（兼容移动端钱包）
  checkWalletInstalled()

  // 监听钱包账户和网络变化
  const provider = getWalletProvider()
  if (provider && provider.on) {
    provider.on('accountsChanged', handleAccountsChanged)
    provider.on('chainChanged', handleChainChanged)
  }

  // 页面隐藏/可见时重新检测（处理移动端切换 App 后返回的情况）
  document.addEventListener('visibilitychange', () => {
    if (!document.hidden) {
      checkWalletInstalled()
      // 如果弹窗打开着，尝试重新连接
      if (storeMain.authModalShow && !address.value) {
        setTimeout(async () => {
          await fetchWalletAndNonce()
        }, 300)
      }
    }
  })
})

// 组件卸载时移除监听器
onUnmounted(() => {
  const provider = getWalletProvider()
  if (provider && provider.removeListener) {
    provider.removeListener('accountsChanged', handleAccountsChanged)
    provider.removeListener('chainChanged', handleChainChanged)
  }
})

// 处理网络切换
function handleChainChanged(chainId: string) {
  // 有些钱包切换网络会刷新页面，这里先保存状态
  if (storeMain.authModalShow && address.value && nonce.value) {
    saveLoginState({
      address: address.value,
      nonce: nonce.value,
      signMessageText: signMessageText.value,
    })
  }

  if (storeMain.authModalShow) {
    resetState()
    // 网络切换后重新获取 nonce
    setTimeout(async () => {
      await fetchWalletAndNonce()
    }, 500)
  }
}

// 监听钱包账户变化（组件内）
function handleAccountsChanged(accounts: string[]) {
  // 重新检测钱包状态
  hasWalletInstalled.value = !!getWalletProvider() && accounts.length > 0

  if (storeMain.authModalShow) {
    // 有些钱包切换账户会刷新页面，这里先保存状态
    if (address.value && nonce.value) {
      saveLoginState({
        address: address.value,
        nonce: nonce.value,
        signMessageText: signMessageText.value,
      })
    }
    resetState()
    if (accounts.length > 0) {
      // 延迟一下确保钱包连接稳定后再获取 nonce
      setTimeout(async () => {
        await fetchWalletAndNonce()
      }, 500)
    }
  }
}

// 打开 MetaMask 官网
function openMetaMask() {
  window.open('https://metamask.io/download/', '_blank')
}

// 连接钱包
async function handleConnect() {
  connecting.value = true
  try {
    const success = await fetchWalletAndNonce()
    if (success) {
      window.$message.success('钱包连接成功，请签名完成登录')
    } else {
      window.$message.error('连接钱包失败')
    }
  } catch (error: any) {
    console.error('连接钱包失败:', error)
    window.$message.error(error?.message || '连接钱包失败')
  } finally {
    connecting.value = false
  }
}

// 签名并登录
async function handleSign() {
  signing.value = true
  try {
    // 重新连接确保获取正确的签名者对象
    const result = await connectWallet()
    if (!result || !address.value) {
      window.$message.error('钱包未连接')
      return
    }

    const signature = await signMessage(result.signer, signMessageText.value)
    if (!signature) {
      window.$message.error('签名失败')
      return
    }

    // 提交登录（使用已获取的地址）
    const loginResult = await Api.v1.auth.post.walletLogin({
      address: address.value,
      signature,
      nonce: nonce.value,
    })

    // 写入用户信息
    const token = loginResult.token
    localStorage.setItem(TOKEN_KEY, token)

    await userInfo(token)
      .then((res) => {
        storeUser.updateUserinfo(res)
        storeMain.triggerAuth(false)

        if (loginResult.is_new_user) {
          window.$message.success('欢迎加入 PaoPao！')
        } else {
          window.$message.success('登录成功')
        }

        // 重置状态
        resetState()
      })
      .catch((err) => {
        console.error('获取用户信息失败:', err)
        window.$message.error('登录失败')
      })
  } catch (error: any) {
    console.error('签名登录失败:', error)
    window.$message.error(error?.message || '签名验证失败')
  } finally {
    signing.value = false
  }
}
</script>

<style lang="less" scoped>
.wallet-login {
  padding: 16px 0;
}

.connect-wallet {
  display: flex;
  justify-content: center;
}

.sign-message {
  .mt-4 {
    margin-top: 16px;
  }
}

.wallet-not-found {
  text-align: center;
}
</style>
