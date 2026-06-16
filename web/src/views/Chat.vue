<template>
  <div class="chat-page">
    <!-- 消息列表 -->
    <div class="chat-messages" ref="messagesRef">
      <div v-for="(msg, i) in messages" :key="i" class="message-row" :class="msg.role">
        <div class="message-bubble" :class="msg.role">
          <template v-if="msg.role === 'user'">
            <div class="bubble-content" v-html="renderContent(msg.content)" />
            <div class="bubble-meta">{{ shortenAddr(walletAddr) }}</div>
          </template>
          <template v-else>
            <div class="ai-label">
              <n-avatar v-if="kolAvatar" :src="kolAvatar" :size="18" round />
              <span>{{ kolNickname || '🤖 AI' }}</span>
            </div>
            <div class="bubble-content" v-html="renderContent(msg.content)" />
            <div v-if="msg.streaming" class="typing-dot"><span>.</span><span>.</span><span>.</span></div>
          </template>
        </div>
      </div>
    </div>
    <!-- 底部工具栏 -->
    <div class="chat-toolbar">
      <n-button quaternary circle @click="triggerImage">
        <template #icon><n-icon size="22"><image-outline /></n-icon></template>
      </n-button>
      <n-button quaternary circle @click="triggerFile">
        <template #icon><n-icon size="22"><attach-outline /></n-icon></template>
      </n-button>
      <n-input class="chat-input" v-model:value="inputText" type="text" placeholder="输入消息..." @keydown.enter="sendMessage" />
      <n-button type="primary" :disabled="!inputText.trim()" @click="sendMessage">
        <template #icon><n-icon size="18"><send-outline /></n-icon></template>
      </n-button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, nextTick, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useStoreUser } from '@/store/user'
import { storeToRefs } from 'pinia'
import { TOKEN_KEY } from '@/store/user'
import { getLoggedInWalletAddress } from '@/utils/web3'
import { ImageOutline, AttachOutline, SendOutline } from '@vicons/ionicons5'

const route = useRoute()
const storeUser = useStoreUser()
const { userInfo } = storeToRefs(storeUser)
const messagesRef = ref<HTMLElement>()
const messages = ref<ChatMessage[]>([])
const inputText = ref('')
const loading = ref(false)
const walletAddr = ref('')
const conversationId = ref('')
const kolUserId = ref(Number(route.query.kol_user_id) || 0)
const kolNickname = ref((route.query.nickname as string) || '')
const kolAvatar = ref((route.query.avatar as string) || '')
interface ChatMessage { role: 'user' | 'assistant'; content: string; streaming?: boolean }

const uploadUrl = import.meta.env.VITE_HOST + '/v1/attachment'
const uploadHeaders = { Authorization: 'Bearer ' + localStorage.getItem(TOKEN_KEY) }
const DIFY_API = import.meta.env.VITE_HOST + '/v1/chat/dify'
const HISTORY_API = import.meta.env.VITE_HOST + '/v1/chat/history'

const shortenAddr = (addr: string) => addr && addr.length > 10 ? addr.slice(0, 6) + '...' + addr.slice(-4) : addr

const triggerImage = () => {
  const input = document.createElement('input')
  input.type = 'file'; input.accept = 'image/*';
  input.onchange = async () => {
    const file = input.files?.[0]; if (!file) return;
    const form = new FormData(); form.append('file', file); form.append('type', 'public/image');
    try {
      const res = await fetch(uploadUrl, { method: 'POST', headers: uploadHeaders, body: form })
      const data = await res.json()
      if (data.code === 0 && data.data?.content) inputText.value += `\n![](${data.data.content})\n`
    } catch { }
  }
  input.click()
}

const triggerFile = () => {
  const input = document.createElement('input')
  input.type = 'file'
  input.onchange = async () => {
    const file = input.files?.[0]; if (!file) return;
    const form = new FormData(); form.append('file', file); form.append('type', 'attachment');
    try {
      const res = await fetch(uploadUrl, { method: 'POST', headers: uploadHeaders, body: form })
      const data = await res.json()
      if (data.code === 0 && data.data?.content) inputText.value += `\n[文件](${data.data.content})\n`
    } catch { }
  }
  input.click()
}

const renderContent = (text: string): string => {
  if (!text) return ''
  return text
    .replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;')
    .replace(/```(\w*)\n?([\s\S]*?)```/g, '<pre><code>$2</code></pre>')
    .replace(/`([^`]+)`/g, '<code>$1</code>')
    .replace(/!\[([^\]]*)\]\(([^)]+)\)/g, '<img src="$2" alt="$1" style="max-width:100%;border-radius:8px"/>')
    .replace(/\[([^\]]+)\]\(([^)]+)\)/g, '<a href="$2" target="_blank">$1</a>')
    .replace(/\*\*([^*]+)\*\*/g, '<strong>$1</strong>')
    .replace(/\n/g, '<br/>')
}

const sendMessage = async () => {
  const text = inputText.value.trim()
  if (!text || loading.value) return
  messages.value.push({ role: 'user', content: text }); inputText.value = ''; scrollToBottom()
  loading.value = true
  const aiMsg: ChatMessage = { role: 'assistant', content: '', streaming: true }; messages.value.push(aiMsg)
  const aiIdx = messages.value.length - 1
  try {
    const token = localStorage.getItem(TOKEN_KEY)
    const resp = await fetch(DIFY_API, {
      method: 'POST', headers: { 'Content-Type': 'application/json', 'Authorization': 'Bearer ' + token },
      body: JSON.stringify({ query: text, user: userInfo.value.id || 'anonymous', kol_user_id: kolUserId.value }),
    })
    if (!resp.ok) throw new Error('请求失败: ' + resp.status)
    const reader = resp.body?.getReader(); if (!reader) throw new Error('不支持')
    const decoder = new TextDecoder(); let buffer = ''
    while (true) {
      const { done, value } = await reader.read(); if (done) break
      buffer += decoder.decode(value, { stream: true })
      const lines = buffer.split('\n'); buffer = lines.pop() || ''
      for (const line of lines) {
        if (line.startsWith('data: ')) {
          try { const d = JSON.parse(line.slice(6)); if (d.answer) messages.value[aiIdx].content += d.answer } catch { }
        }
      }
      scrollToBottom()
    }
  } catch (err: any) {
    messages.value[aiIdx].content = '抱歉：' + (err.message || '未知错误')
  } finally { messages.value[aiIdx].streaming = false; loading.value = false }
}

const scrollToBottom = () => nextTick(() => { const el = messagesRef.value; if (el) el.scrollTop = el.scrollHeight })
onMounted(async () => {
  walletAddr.value = getLoggedInWalletAddress() || ''
  // 加载历史消息
  if (kolUserId.value > 0 || true) {
    try {
      const token = localStorage.getItem(TOKEN_KEY)
      const res = await fetch(HISTORY_API + '?kol_user_id=' + kolUserId.value, {
        headers: { 'Authorization': 'Bearer ' + token },
      })
      const data = await res.json()
      if (data.code === 0 && data.data) {
        conversationId.value = data.data.dify_conversation_id || ''
        const history = data.data.messages || []
        // Dify 每条记录含 query(用户)+answer(AI)，展开为两条消息
        const formatted: ChatMessage[] = []
        for (const m of history) {
          if (m.query) formatted.push({ role: 'user', content: m.query })
          if (m.answer) formatted.push({ role: 'assistant', content: m.answer })
        }
        messages.value = formatted
        scrollToBottom()
      }
    } catch { /* ignore */ }
  }
})
</script>

<style lang="less" scoped>
.chat-page { display:flex; flex-direction:column; overflow:hidden; background:var(--body-color,#f5f5f5); position:absolute; top:0; left:0; right:0; bottom:0; }
.chat-messages { flex:1; overflow-y:auto; padding:12px 16px;
  // 底部安全距离让最后一条消息不被工具栏遮挡
  &::after { content:''; display:block; height:8px; }
}

.message-row { display:flex; margin-bottom:14px;
  &.user { justify-content:flex-end; }
  &.assistant { justify-content:flex-start; }
}

.message-bubble { max-width:82%; padding:10px 14px; border-radius:14px; font-size:14px; line-height:1.55; word-break:break-word;
  &.user { background:#10b981; color:#fff; border-bottom-right-radius:4px;
    .bubble-meta { font-size:11px; opacity:.65; margin-top:2px; }
  }
  &.assistant { background:var(--card-color,#fff); border:1px solid var(--border-color,#e5e5e5); border-bottom-left-radius:4px;
    .ai-label { font-size:11px; color:#999; margin-bottom:3px; }
  }
}

.bubble-content :deep(pre) { background:rgba(0,0,0,.04); padding:8px 12px; border-radius:6px; overflow-x:auto; margin:6px 0; font-size:13px; white-space:pre-wrap; }
.bubble-content :deep(code) { font-family:monospace; font-size:13px; }
.bubble-content :deep(img) { max-width:100%; border-radius:8px; margin:6px 0; }
.bubble-content :deep(a) { color:var(--primary-color,#2080f0); }

.typing-dot span { display:inline-block; width:5px; height:5px; border-radius:50%; background:#aaa; margin:0 1px; animation:blink 1.4s infinite both; &:nth-child(2){animation-delay:.2s} &:nth-child(3){animation-delay:.4s} }
@keyframes blink { 0%,80%,100%{opacity:.2} 40%{opacity:1} }

.chat-toolbar { display:flex; align-items:center; gap:10px; padding:10px 14px; padding-bottom:max(10px, env(safe-area-inset-bottom)); background:var(--card-color,#fff); border-top:1px solid var(--border-color,#eee); flex-shrink:0;
  .chat-input { flex:1; }
}
</style>
