<template>
    <div class="composer-trigger" @click="$emit('click')">
        <n-avatar
            v-if="userInfo.id > 0"
            round
            :size="40"
            :src="userInfo.avatar"
            class="trigger-avatar"
        />
        <n-avatar
            v-else
            round
            :size="40"
            class="trigger-avatar"
        >
            <n-icon :size="24"><person-outline /></n-icon>
        </n-avatar>
        <span class="trigger-placeholder">说说您的新鲜事...</span>
    </div>
</template>

<script setup lang="ts">
import { PersonOutline } from '@vicons/ionicons5';
import { useStoreUser } from '@/store/user';
import { storeToRefs } from 'pinia';

const storeUser = useStoreUser();
const { userInfo } = storeToRefs(storeUser);

defineEmits<{
    click: [];
}>();
</script>

<style lang="less" scoped>
.composer-trigger {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 12px 16px;
    border-bottom: 1px solid var(--border-color, #eff3f4);
    cursor: pointer;
    transition: background-color 0.2s;

    &:hover {
        background: #f7f9f9;
    }
}

.trigger-avatar {
    flex-shrink: 0;
}

.trigger-placeholder {
    font-size: 15px;
    color: #888;
    flex: 1;
}

.dark {
    .composer-trigger {
        border-bottom-color: #2f3336;

        &:hover {
            background: #080808;
        }
    }
}
</style>
