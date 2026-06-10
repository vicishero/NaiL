<template>
  <n-modal
    v-model:show="authModalShow"
    class="auth-card"
    preset="card"
    size="small"
    :mask-closable="false"
    :bordered="false"
    :style="{
      width: '360px',
    }"
  >
    <div class="auth-wrap">
      <n-card :bordered="false">
        <n-space justify="center" class="mb-4">
          <n-h3>
            <n-text type="success">连接钱包登录</n-text>
          </n-h3>
        </n-space>

        <WalletLogin />

        <n-alert type="info" class="mt-4">
          <template #icon>
            <n-icon>
              <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <circle cx="12" cy="12" r="10"></circle>
                <path d="M12 16v-4"></path>
                <path d="M12 8h.01"></path>
              </svg>
            </n-icon>
          </template>
          <template #header>使用说明</template>
          <ul class="mt-2 text-sm">
            <li>• 请先安装 MetaMask 或其他以太坊钱包插件</li>
            <li>• 点击按钮连接钱包，签名后自动完成登录</li>
            <li>• 首次登录将自动创建账户，无需注册</li>
          </ul>
        </n-alert>
      </n-card>
    </div>
  </n-modal>
</template>

<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue'
import { useStoreMain } from '@/store/main'
import { TOKEN_KEY, useStoreUser } from '@/store/user'
import { userInfo } from '@/api/auth'
import { storeToRefs } from 'pinia'
import WalletLogin from './WalletLogin.vue'
import { getWalletProvider, getLoggedInWalletAddress, recordWalletAddressOnLogin } from '@/utils/web3'

const storeMain = useStoreMain()
const storeUser = useStoreUser()
const { authModalShow } = storeToRefs(storeMain)

// 钱包账户变化时自动退出登录
function handleAccountsChanged(accounts: string[]) {
  console.log('[Auth] 钱包账户变化（全局监听）:', accounts)
  if (storeUser.userLogined) {
    console.log('[Auth] 检测到钱包账户变化，自动退出登录')
    storeUser.userLogout()
    window.$message?.info?.('检测到钱包账户变化，已自动退出登录')
  }
}

onMounted(() => {
  const token = localStorage.getItem(TOKEN_KEY) || ''
  if (token) {
    userInfo(token)
      .then((res) => {
        storeUser.updateUserinfo(res)
        storeMain.triggerAuth(false)

        // 页面刷新后，如果已经登录但没有记录钱包地址，获取并记录
        const savedAddress = getLoggedInWalletAddress()
        if (!savedAddress) {
          const provider = getWalletProvider()
          if (provider) {
            provider.request?.({ method: 'eth_accounts' })
              .then((accounts: string[]) => {
                if (accounts?.[0]) {
                  console.log('[Auth] 页面刷新后记录钱包地址:', accounts[0])
                  recordWalletAddressOnLogin(accounts[0])
                }
              })
              .catch((err: any) => {
                console.log('[Auth] 获取钱包地址失败:', err)
              })
          }
        }
      })
      .catch((err) => {
        storeUser.userLogout()
      })
  } else {
    storeUser.userLogout()
  }

  // 全局注册钱包账户变化监听
  setTimeout(() => {
    const provider = getWalletProvider()
    if (provider && provider.on) {
      console.log('[Auth] 注册全局钱包账户变化监听')
      provider.on('accountsChanged', handleAccountsChanged)
    }
  }, 500)
})

// 组件卸载时移除监听器
onUnmounted(() => {
  const provider = getWalletProvider()
  if (provider && provider.removeListener) {
    provider.removeListener('accountsChanged', handleAccountsChanged)
  }
})
</script>

<style lang="less" scoped>
.auth-wrap {
  margin-top: -30px;

  .mb-4 {
    margin-bottom: 16px;
  }

  .mt-4 {
    margin-top: 16px;
  }

  ul {
    margin: 0;
    padding-left: 0;
    list-style: none;

    li {
      line-height: 1.8;
    }
  }
}
.dark {
  .auth-wrap {
    background-color: rgba(16, 16, 20, 0.75);
  }
}
</style>
