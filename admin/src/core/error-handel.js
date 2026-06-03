import { createSysError } from '@/api/system/sysError'
import { useUserStore } from '@/pinia/modules/user'

// 防止错误上报的无限循环
let isReportingError = false

function sendErrorTip(errorInfo) {
  // 如果正在上报错误，或者没有token，就不上报了，避免无限循环
  const userStore = useUserStore()
  if (isReportingError || !userStore.token) return

  isReportingError = true
  setTimeout(() => {
    const errorData = {
      form: errorInfo.type,
      info: `${errorInfo.message}\nStack: ${errorInfo.stack}${errorInfo.component ? `\nComponent: ${errorInfo.component.name || 'Unknown'}` : ''}${errorInfo.vueInfo ? `\nVue Info: ${errorInfo.vueInfo}` : ''}${errorInfo.source ? `\nSource: ${errorInfo.source}:${errorInfo.lineno}:${errorInfo.colno}` : ''}`,
      level: 'error',
      solution: null
    }

    createSysError(errorData).catch(apiErr => {
      console.error('Failed to create error record:', apiErr)
    }).finally(() => {
      isReportingError = false
    })
  }, 0)
}

  window.addEventListener('unhandledrejection', (event) => {
    sendErrorTip({
      type: '前端',
      message: `错误信息: ${event.reason}`,
      stack: `调用栈: ${event.reason?.stack || '没有调用栈信息'}`,
    });
  });
