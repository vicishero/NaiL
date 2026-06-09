<template>
  <div class="wallet-login">
    <div v-if="!hasWallet" class="wallet-not-found">
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
import { ref, onMounted } from 'vue'
import { useStoreMain } from '@/store/main'
import { TOKEN_KEY, useStoreUser } from '@/store/user'
import { userInfo } from '@/api/auth'
import { hasWallet, connectWallet, signMessage } from '@/utils/web3'
import { Api } from '@/utils/request'

const storeMain = useStoreMain()
const storeUser = useStoreUser()

const connecting = ref(false)
const signing = ref(false)
const address = ref('')
const nonce = ref('')
const signMessageText = ref('')

// 检查是否安装了钱包
onMounted(() => {
  // 延迟一点检测，确保钱包已注入
  setTimeout(() => {
    // 触发响应式更新
  }, 100)
})

// 打开 MetaMask 官网
function openMetaMask() {
  window.open('https://metamask.io/download/', '_blank')
}

// 连接钱包
async function handleConnect() {
  connecting.value = true
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
    const result = await connectWallet()
    if (!result) {
      window.$message.error('钱包未连接')
      return
    }

    const signature = await signMessage(result.signer, signMessageText.value)
    if (!signature) {
      window.$message.error('签名失败')
      return
    }

    // 提交登录
    const loginResult = await Api.v1.auth.post.walletLogin({
      address: result.address,
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
        address.value = ''
        nonce.value = ''
        signMessageText.value = ''
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
