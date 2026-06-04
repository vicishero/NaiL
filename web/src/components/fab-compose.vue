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
    // 居中容器内的右边缘定位
    left: 50%;
    transform: translateX(174px);
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
        transform: translateX(174px) scale(1.05);
        box-shadow: 0 6px 16px rgba(24, 160, 88, 0.5);
    }

    &:active {
        transform: translateX(174px) scale(0.95);
    }
}
</style>
