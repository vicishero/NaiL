<template>
    <div class="assets-page">
        <div class="assets-content">
            <n-icon :size="64" color="#ccc">
                <wallet-outline />
            </n-icon>
            <h2>资产管理</h2>
            <p class="assets-desc">资产管理功能即将上线，敬请期待</p>
        </div>
    </div>
</template>

<script setup lang="ts">
// KeepAlive 需要组件 name 来匹配 include
defineOptions({
  name: 'assets',
});
import { WalletOutline } from '@vicons/ionicons5';
</script>

<style lang="less" scoped>
.assets-page {
    min-height: 60vh;
    display: flex;
    align-items: center;
    justify-content: center;
}

.assets-content {
    text-align: center;
    color: #999;

    h2 {
        margin: 16px 0 8px;
        font-size: 20px;
        color: #666;
    }

    .assets-desc {
        font-size: 14px;
    }
}
</style>
