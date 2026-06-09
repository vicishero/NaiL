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

                    <!-- 发帖 FAB -->
                    <fab-compose />

                    <!-- 底部 TabBar -->
                    <bottom-tab-bar />

                    <!-- 登录/注册公共组件 -->
                    <auth />

                    <!-- 发帖 Modal -->
                    <n-modal
                        v-model:show="composeModalShow"
                        preset="card"
                        title="发帖"
                        style="max-width: 600px;"
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
import { onMounted, computed, watch, ref, nextTick } from 'vue';
import { useRoute } from 'vue-router';
import { useStoreMain } from '@/store/main';
import { darkTheme } from 'naive-ui';
import { getSiteProfile } from '@/api/site';
import { useStoreProfile } from '@/store/profile';
import { useStoreUser } from '@/store/user';
import { storeToRefs } from 'pinia';
import { Api } from '@/utils/request';
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
            }, profile.value.defaultMsgLoopInterval || 5000);
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
});
</script>
