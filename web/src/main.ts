import { createApp } from 'vue';
import { createPinia } from 'pinia';
import router from './router';
import App from './App.vue';
import '@/assets/css/main.less';

import type { MessageApiInjection } from 'naive-ui/lib/message/src/MessageProvider';

// 通用字体
import 'vfonts/Lato.css';
// 等宽字体
import 'vfonts/FiraCode.css';

// ========== 首屏性能日志 ==========
function perfLog(label) {
  const now = performance.now();
  const elapsed = (now - window.__PERF_START__).toFixed(2);
  const info = `[${elapsed}ms] ${label}`;
  window.__PERF_LOG__.push(info);
  console.log(`%c${info}`, 'color: #18a058; font-weight: bold;');
}

perfLog('main.ts 开始执行');

const pinia = createPinia();
perfLog('Pinia 创建完成');

const app = createApp(App);
perfLog('Vue App 创建完成');

app.use(router);
perfLog('Router 注册完成');

app.use(pinia);
perfLog('Pinia 注册完成');

// 监控路由加载时间
router.beforeEach((to, from, next) => {
  if (from.name === undefined) {
    perfLog(`首路由开始加载: ${to.name?.toString() || to.path}`);
    to.meta.__startTime = performance.now();
  }
  next();
});

router.afterEach((to) => {
  if (to.meta.__startTime) {
    const elapsed = (performance.now() - to.meta.__startTime).toFixed(2);
    perfLog(`首路由加载完成, 耗时: ${elapsed}ms`);
  }
});

app.mount('#app');
perfLog('App 挂载到 DOM 完成');

// 隐藏启动页
if (typeof window !== 'undefined' && (window as any).hideSplashScreen) {
  // 延迟一帧确保 DOM 渲染完成
  requestAnimationFrame(() => {
    setTimeout(() => {
      (window as any).hideSplashScreen();
    }, 0);
  });
}

declare global {
  interface Window {
    $message: MessageApiInjection;
  }
}
