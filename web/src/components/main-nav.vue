<template>
    <!-- 只有 title 有值时才显示，避免闪烁 -->
    <div v-if="props.title" class="back-header">
        <n-button
            class="back-btn"
            @click="goBack"
            quaternary
            circle
            size="small"
        >
            <template #icon>
                <n-icon><chevron-left-round /></n-icon>
            </template>
        </n-button>
        <span class="back-title">{{ props.title }}</span>
    </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { ChevronLeftRound } from '@vicons/material';

const router = useRouter();

const props = withDefaults(
  defineProps<{
    title: string;
  }>(),
  {
    title: '',
  },
);

const goBack = () => {
  if (window.history.length <= 1) {
    router.push({ path: '/' });
  } else {
    router.go(-1);
  }
};
</script>

<style lang="less" scoped>
.back-header {
    display: flex;
    align-items: center;
    gap: 4px;
    padding: 12px 16px;
    background: rgba(255, 255, 255, 0.85);
    backdrop-filter: blur(12px);
    position: sticky;
    top: 0;
    z-index: 99;

    .back-btn {
        flex-shrink: 0;
    }

    .back-title {
        font-size: 18px;
        font-weight: 700;
    }
}

.dark {
    .back-header {
        background: rgba(16, 16, 20, 0.85);
    }
}
</style>
