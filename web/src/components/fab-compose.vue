<template>
    <div v-if="showFab" class="fab-btn" @click="openCompose">
        <n-icon :size="24" color="#fff">
            <create-outline />
        </n-icon>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useRoute } from 'vue-router';
import { useStoreMain } from '@/store/main';
import { CreateOutline } from '@vicons/ionicons5';

const route = useRoute();
const storeMain = useStoreMain();

const showFab = computed(() => {
    return route.name === 'home';
});

function openCompose() {
    storeMain.triggerCompose(true);
}
</script>

<style lang="less" scoped>
.fab-btn {
    position: fixed;
    bottom: 80px;
    // 使用 right 定位，确保在所有屏幕尺寸下都在可视范围内
    right: 16px;
    width: 56px;
    height: 56px;
    border-radius: 50%;
    background: #18a058;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    z-index: 99;
    box-shadow: 0 4px 12px rgba(24, 160, 88, 0.4);
    transition: all 0.2s ease;

    &:hover {
        background: #16914f;
        transform: scale(1.05);
        box-shadow: 0 6px 16px rgba(24, 160, 88, 0.5);
    }

    &:active {
        transform: scale(0.95);
    }

    // 移动端适配：距离底部 TabBar 稍远一些
    @media (max-width: 768px) {
        bottom: 90px;
        right: 16px;
    }

    // 极小屏幕适配
    @media (max-width: 375px) {
        bottom: 85px;
        right: 12px;
        width: 52px;
        height: 52px;
    }
}
</style>
