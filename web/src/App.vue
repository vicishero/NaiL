<template>
    <n-config-provider :theme="iTheme">
        <n-message-provider>
            <n-dialog-provider>
                <div
                    class="app-container"
                    :class="{ dark: iTheme?.name === 'dark' }"
                >
                    <!-- 子页面返回栏 -->
                    <main-nav
                        v-if="showBack"
                        :title="pageTitle"
                    />

                    <!-- 内容区 -->
                    <div class="content-wrap">
                        <router-view
                            v-slot="{ Component }"
                        >
                            <keep-alive>
                                <component
                                    v-if="$route.meta.keepAlive"
                                    :is="Component"
                                />
                            </keep-alive>
                            <component
                                v-if="!$route.meta.keepAlive"
                                :is="Component"
                            />
                        </router-view>
                    </div>

                    <!-- 发帖 FAB（聊天页隐藏） -->
                    <fab-compose v-if="route.name !== 'chat'" />

                    <!-- 底部 TabBar（聊天页隐藏） -->
                    <bottom-tab-bar v-if="route.name !== 'chat'" />

                    <!-- 登录/注册公共组件 -->
                    <auth />

                    <!-- 发帖 Modal -->
                    <n-modal
                        v-model:show="composeModalShow"
                        preset="card"
                        title="发帖"
                        style="max-width: 460px;"
                        :mask-closable="true"
                    >
                        <compose
                            modal
                            @post-success="handlePostSuccess"
                        />
                    </n-modal>
                </div>
            </n-dialog-provider>
        </n-message-provider>
        <n-global-style />
    </n-config-provider>
</template>

<script setup lang="ts">
import { onMounted, computed, watch, ref, nextTick, onUnmounted } from 'vue';
import { useRoute } from 'vue-router';
import { useStoreMain } from '@/store/main';
import { darkTheme } from 'naive-ui';
import { getSiteProfile } from '@/api/site';
import { useStoreProfile } from '@/store/profile';
import { useStoreUser } from '@/store/user';
import { storeToRefs } from 'pinia';
import { Api } from '@/utils/request';
import { getWalletProvider, getLoggedInWalletAddress } from '@/utils/web3';
import BottomTabBar from '@/components/bottom-tab-bar.vue';
import FabCompose from '@/components/fab-compose.vue';
import Compose from '@/components/compose.vue';

// ========== 首屏性能日志 ==========
function perfLog(label: string) {
  const now = performance.now();
  const elapsed = (now - (window as any).__PERF_START__).toFixed(2);
  const info = `[${elapsed}ms] ${label}`;
  (window as any).__PERF_LOG__.push(info);
  console.log(`%c${info}`, 'color: #18a058; font-weight: bold;');
}

perfLog('App.vue setup 开始');

const storeMain = useStoreMain();
const storeProfile = useStoreProfile();
const storeUser = useStoreUser();
const { theme, composeModalShow } = storeToRefs(storeMain);
const { userInfo } = storeToRefs(storeUser);
const { profile } = storeToRefs(storeProfile);

perfLog('Store 初始化完成');

const route = useRoute();

const iTheme = computed(() => (theme.value === 'dark' ? darkTheme : null));

// 当前页面标题
const pageTitle = computed(() => (route.meta.title as string) || '');

// 是否显示返回按钮（非首页和非常规 Tab 页时显示）
const showBack = computed(() => {
    const tabRoutes = ['home', 'explore', 'assets', 'messages', 'profile'];
    // 关键：route.name 为 undefined 时也不显示（路由初始化中）
    return route.name && !tabRoutes.includes(route.name as string);
});

// 消息轮询（从 sidebar.vue 迁移过来）
const msgLoop = ref<ReturnType<typeof setInterval> | null>(null);

watch(() => [userInfo.value.id, profile.value.defaultMsgLoopInterval], () => {
    if (userInfo.value.id > 0) {
        if (!msgLoop.value) {
            Api.v1.user.get.msgcount.unread({})
                .then((res: any) => {
                    storeMain.updateUnreadMsgCount(res.count);
                })
                .catch((err: any) => {
                    console.log(err);
                });

            msgLoop.value = setInterval(() => {
                Api.v1.user.get.msgcount.unread({})
                    .then((res: any) => {
                        storeMain.updateUnreadMsgCount(res.count);
                    })
                    .catch((err: any) => {
                        console.log(err);
                    });
            }, profile.value.defaultMsgLoopInterval || 30000);
        }
    } else {
        if (msgLoop.value) {
            clearInterval(msgLoop.value);
            msgLoop.value = null;
        }
    }
});

function handlePostSuccess() {
    storeMain.triggerCompose(false);
    storeMain.doRefresh();
}

function loadSiteProfile() {
    storeProfile.loadDefaultSiteProfile();
    if (import.meta.env.VITE_USE_WEB_PROFILE.toLowerCase() === 'true') {
        getSiteProfile()
            .then((res) => {
                storeProfile.updateSiteProfile(res);
            }).catch((err) => {
                console.log(err);
            });
    }
}

// 钱包账户变化时自动退出登录（全局监听）
function handleAccountsChanged(accounts: string[]) {
    console.log('[App] 钱包账户变化（事件监听）:', accounts)
    checkAndHandleWalletChange(accounts?.[0] || null)
}

// 检查并处理钱包变化
function checkAndHandleWalletChange(newAddress: string | null) {
    const loggedInAddress = getLoggedInWalletAddress()
    const normalizedNewAddress = newAddress?.toLowerCase() || null
    console.log('[App] 检查钱包变化 - 登录状态:', storeUser.userLogined, '已记录地址:', loggedInAddress, '新地址:', normalizedNewAddress)

    // 如果用户已登录，且钱包地址发生了变化
    if (storeUser.userLogined && loggedInAddress && normalizedNewAddress &&
        loggedInAddress !== normalizedNewAddress) {
        console.log('[App] 检测到钱包账户变化，自动退出登录')
        storeUser.userLogout()
        window.$message?.info?.('检测到钱包账户变化，已自动退出登录')
        // 自动打开登录弹窗，方便用户用新账户登录
        setTimeout(() => {
            storeMain.triggerAuth(true)
        }, 500)
    } else {
        console.log('[App] 钱包地址未变化或未登录，不处理')
    }
}

// 轮询检测钱包账户变化（作为事件监听的 fallback）
async function pollWalletAddress() {
    if (!storeUser.userLogined) {
        // 未登录时只记录日志，不处理
        return
    }

    try {
        const provider = getWalletProvider()
        if (provider) {
            const accounts = await provider.request?.({ method: 'eth_accounts' })
            checkAndHandleWalletChange(accounts?.[0] || null)
        }
    } catch (e) {
        console.log('[App] 轮询钱包地址失败:', e)
    }
}

// 注册钱包监听，兼容多种钱包 API
function setupWalletListeners() {
    const provider = getWalletProvider()
    if (!provider) {
        console.log('[App] 未找到钱包 provider，稍后重试')
        return false
    }

    console.log('[App] 尝试注册钱包监听...')

    // 尝试多种监听 API
    const methods = ['on', 'addListener', 'addEventListener']
    let registered = false

    for (const method of methods) {
        if (typeof provider[method] === 'function') {
            try {
                provider[method]('accountsChanged', handleAccountsChanged)
                console.log(`[App] 使用 ${method} 注册监听成功`)
                registered = true
                break
            } catch (e) {
                console.log(`[App] 使用 ${method} 注册监听失败:`, e)
            }
        }
    }

    // 无论事件监听是否成功，都启动轮询作为 fallback
    console.log('[App] 启动钱包地址轮询（每 2 秒）')
    setInterval(pollWalletAddress, 2000)

    // 初始化时获取一次地址
    pollWalletAddress()

    return registered
}

onMounted(async () => {
    perfLog('App.vue onMounted 触发');

    await nextTick();
    perfLog('首次 nextTick 完成 (DOM 渲染完成)');

    // 延迟执行非关键操作
    setTimeout(() => {
        perfLog('开始加载站点配置');
        loadSiteProfile();
        perfLog('站点配置加载完成');
    }, 0);

    // 多次尝试注册钱包监听，适配注入延迟
    let retryCount = 0
    const trySetup = () => {
        if (setupWalletListeners()) return
        if (retryCount < 10) {
            retryCount++
            setTimeout(trySetup, 500)
        }
    }
    trySetup()
});

onUnmounted(() => {
    const provider = getWalletProvider()
    if (provider) {
        const methods = ['removeListener', 'off', 'removeEventListener']
        for (const method of methods) {
            if (typeof provider[method] === 'function') {
                try {
                    provider[method]('accountsChanged', handleAccountsChanged)
                } catch (e) { /* ignore */ }
            }
        }
    }
});
</script>
