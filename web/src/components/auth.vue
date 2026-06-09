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
import { onMounted } from 'vue'
import { useStoreMain } from '@/store/main'
import { TOKEN_KEY, useStoreUser } from '@/store/user'
import { userInfo } from '@/api/auth'
import { storeToRefs } from 'pinia'
import WalletLogin from './WalletLogin.vue'

const storeMain = useStoreMain()
const storeUser = useStoreUser()
const { authModalShow } = storeToRefs(storeMain)

onMounted(() => {
  const token = localStorage.getItem(TOKEN_KEY) || ''
  if (token) {
    userInfo(token)
      .then((res) => {
        storeUser.updateUserinfo(res)
        storeMain.triggerAuth(false)
      })
      .catch((err) => {
        storeUser.userLogout()
      })
  } else {
    storeUser.userLogout()
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
