<template>
    <div class="tab-bar-wrap">
        <div
            v-for="tab in tabs"
            :key="tab.key"
            class="tab-item"
            :class="{ active: activeTab === tab.key }"
            @click="goTab(tab)"
        >
            <n-badge
                v-if="tab.key === 'messages'"
                :dot="unreadMsgCount > 0"
                :show="unreadMsgCount > 0"
                processing
            >
                <n-icon :size="24" :component="tab.icon" />
            </n-badge>
            <n-icon v-else :size="24" :component="tab.icon" />
            <span class="tab-label">{{ tab.label }}</span>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { NIcon, NBadge, useMessage } from 'naive-ui';
import {
    HomeOutline,
    ChatbubblesOutline,
    WalletOutline,
    PersonOutline,
} from '@vicons/ionicons5';
import { CompassOutline } from '@vicons/ionicons5';
import { useStoreMain } from '@/store/main';
import { storeToRefs } from 'pinia';

const storeMain = useStoreMain();
const { unreadMsgCount } = storeToRefs(storeMain);

// 初始化全局 message（必须在 n-message-provider 内部调用）
window.$message = useMessage();

const route = useRoute();
const router = useRouter();

const tabs = [
    { key: 'home', label: '首页', icon: HomeOutline, route: '/' },
    { key: 'explore', label: '探索', icon: CompassOutline, route: '/explore' },
    { key: 'assets', label: '资产', icon: WalletOutline, route: '/assets' },
    { key: 'messages', label: '消息', icon: ChatbubblesOutline, route: '/messages' },
    { key: 'profile', label: '个人', icon: PersonOutline, route: '/profile' },
];

const activeTab = computed(() => {
    const name = route.name as string;
    if (name === 'home') return 'home';
    if (name === 'explore') return 'explore';
    if (name === 'assets') return 'assets';
    if (name === 'messages') return 'messages';
    if (name === 'profile') return 'profile';
    return '';
});

function goTab(tab: { key: string; route: string }) {
    if (activeTab.value === tab.key) {
        // 已经在当前页，触发刷新
        if (tab.key === 'home') {
            storeMain.doRefresh();
        }
        return;
    }
    router.push(tab.route);
}
</script>

<style lang="less" scoped>
.tab-bar-wrap {
    position: fixed;
    bottom: 0;
    left: 50%;
    transform: translateX(-50%);
    width: 100%;
    max-width: 500px;
    height: 56px;
    display: flex;
    justify-content: space-around;
    align-items: center;
    background-color: rgba(255, 255, 255, 0.92);
    backdrop-filter: blur(12px);
    border-top: 1px solid var(--border-color, #eff3f4);
    z-index: 100;
    padding-bottom: env(safe-area-inset-bottom, 0);
}

.tab-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 2px;
    cursor: pointer;
    color: #888;
    padding: 4px 12px;
    border-radius: 8px;
    transition: color 0.2s ease;

    &:hover {
        color: #18a058;
    }

    &.active {
        color: #18a058;
    }
}

.tab-label {
    font-size: 10px;
    line-height: 1;
}

.dark {
    .tab-bar-wrap {
        background-color: rgba(16, 16, 20, 0.92);
        border-top-color: #2f3336;
    }
}
</style>
