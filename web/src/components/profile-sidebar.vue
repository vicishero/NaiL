<template>
    <n-drawer
        v-model:show="showDrawer"
        placement="right"
        :width="320"
        :mask-closable="true"
        :native-scrollbar="false"
        title="个人中心"
    >
        <!-- 菜单项弹窗 -->
        <n-modal
            v-model:show="modalShow"
            preset="dialog"
            :title="modalTitle"
            positive-text="知道了"
            @positive-click="modalShow = false"
        >
            <div>{{ modalContent }}</div>
        </n-modal>

        <div class="sidebar-content">
                <!-- 第一组菜单 -->
                <div class="menu-group">
                    <n-list>
                        <n-list-item
                            class="menu-item"
                            @click="goToProfile"
                        >
                            <template #prefix>
                                <n-icon :size="20">
                                    <person-outline />
                                </n-icon>
                            </template>
                            个人资料
                        </n-list-item>
                        <n-list-item
                            class="menu-item"
                            @click="handleMenuAction('subscribes')"
                        >
                            <template #prefix>
                                <n-icon :size="20">
                                    <star-outline />
                                </n-icon>
                            </template>
                            我的订阅
                        </n-list-item>
                        <n-list-item
                            class="menu-item"
                            @click="handleMenuAction('consumption')"
                        >
                            <template #prefix>
                                <n-icon :size="20">
                                    <receipt-outline />
                                </n-icon>
                            </template>
                            消费查询
                        </n-list-item>
                    </n-list>
                </div>

                <!-- 第二组菜单 -->
                <div class="menu-group">
                    <n-list>
                        <n-list-item
                            class="menu-item"
                            @click="handleMenuAction('business')"
                        >
                            <template #prefix>
                                <n-icon :size="20">
                                    <briefcase-outline />
                                </n-icon>
                            </template>
                            商务合作
                        </n-list-item>
                        <n-list-item
                            class="menu-item"
                            @click="handleMenuAction('ai-creator')"
                        >
                            <template #prefix>
                                <n-icon :size="20">
                                    <sparkles-outline />
                                </n-icon>
                            </template>
                            AI创作者
                        </n-list-item>
                    </n-list>
                </div>

                <!-- 第三组菜单 -->
                <div class="menu-group">
                    <n-list>
                        <n-list-item class="menu-item">
                            <template #prefix>
                                <n-icon :size="20">
                                    <moon-outline />
                                </n-icon>
                            </template>
                            深色模式
                            <template #suffix>
                                <n-switch
                                    v-model:value="darkMode"
                                    size="small"
                                    @update:value="handleThemeChange"
                                />
                            </template>
                        </n-list-item>
                        <n-list-item
                            class="menu-item"
                            @click="toggleLanguage"
                        >
                            <template #prefix>
                                <n-icon :size="20">
                                    <language-outline />
                                </n-icon>
                            </template>
                            语言
                            <template #suffix>
                                <span class="suffix-text">{{ langLabel }}</span>
                            </template>
                        </n-list-item>
                        <n-list-item
                            class="menu-item"
                            @click="handleMenuAction('privacy')"
                        >
                            <template #prefix>
                                <n-icon :size="20">
                                    <shield-checkmark-outline />
                                </n-icon>
                            </template>
                            隐私设置
                        </n-list-item>
                        <n-list-item
                            class="menu-item"
                            @click="handleMenuAction('feedback')"
                        >
                            <template #prefix>
                                <n-icon :size="20">
                                    <chatbox-ellipses-outline />
                                </n-icon>
                            </template>
                            意见反馈
                        </n-list-item>
                        <n-list-item
                            class="menu-item"
                            @click="handleMenuAction('about')"
                        >
                            <template #prefix>
                                <n-icon :size="20">
                                    <information-circle-outline />
                                </n-icon>
                            </template>
                            关于平台
                        </n-list-item>
                        <n-list-item
                            class="menu-item logout-item"
                            @click="handleLogout"
                        >
                            <template #prefix>
                                <n-icon :size="20">
                                    <log-out-outline />
                                </n-icon>
                            </template>
                            退出登录
                        </n-list-item>
                    </n-list>
                </div>
            </div>
    </n-drawer>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import { useRouter } from 'vue-router';
import {
    PersonOutline,
    MoonOutline,
    LanguageOutline,
    BriefcaseOutline,
    SparklesOutline,
    ShieldCheckmarkOutline,
    ChatboxEllipsesOutline,
    InformationCircleOutline,
    StarOutline,
    ReceiptOutline,
    LogOutOutline,
} from '@vicons/ionicons5';
import { useStoreMain } from '@/store/main';
import { useStoreUser } from '@/store/user';
import { storeToRefs } from 'pinia';
import { useMessage, useDialog } from 'naive-ui';

const storeMain = useStoreMain();
const storeUser = useStoreUser();
const { theme } = storeToRefs(storeMain);
const router = useRouter();
const message = useMessage();
const dialog = useDialog();

const emit = defineEmits<{
    close: [];
}>();

const props = defineProps<{
    show: boolean;
}>();

const showDrawer = computed({
    get: () => props.show,
    set: (val) => {
        if (!val) emit('close');
    },
});

const darkMode = computed({
    get: () => theme.value === 'dark',
    set: (val) => val,
});

// 切换主题
function handleThemeChange(val: boolean) {
    if (val) {
        localStorage.setItem('PAOPAO_THEME', 'dark');
        storeMain.triggerTheme('dark');
    } else {
        localStorage.setItem('PAOPAO_THEME', 'light');
        storeMain.triggerTheme('light');
    }
}

// 跳转到个人资料页面
function goToProfile() {
    handleClose();
    router.push({ name: 'setting' });
}

// 弹窗状态
const modalShow = ref(false);
const modalTitle = ref('');
const modalContent = ref('');

// 语言切换
const lang = ref<'cn' | 'en'>((localStorage.getItem('PAOPAO_LANG') as 'cn' | 'en') || 'cn');
const langLabel = computed(() => (lang.value === 'cn' ? '简体中文' : 'English'));

function toggleLanguage() {
    lang.value = lang.value === 'cn' ? 'en' : 'cn';
    localStorage.setItem('PAOPAO_LANG', lang.value);
    message.success(lang.value === 'cn' ? '已切换为简体中文' : 'Switched to English');
}

// 处理菜单项点击
function handleMenuAction(action: string) {
    switch (action) {
        case 'subscribes':
            modalTitle.value = '我的订阅';
            modalContent.value = '我的订阅功能开发中，敬请期待。';
            break;
        case 'consumption':
            modalTitle.value = '消费查询';
            modalContent.value = '消费查询功能开发中，敬请期待。';
            break;
        case 'business':
            modalTitle.value = '商务合作';
            modalContent.value = '如需商务合作，请联系：business@paopao.info';
            break;
        case 'ai-creator':
            modalTitle.value = 'AI创作者';
            modalContent.value = 'AI创作者平台即将上线，为创作者提供AI辅助内容生成、智能推荐等能力。敬请期待！';
            break;
        case 'privacy':
            modalTitle.value = '隐私设置';
            modalContent.value = '隐私设置功能开发中。我们将严格保护您的个人信息安全，遵循相关法律法规。';
            break;
        case 'feedback':
            modalTitle.value = '意见反馈';
            modalContent.value = '欢迎通过邮件 feedback@paopao.info 提交您的宝贵意见和建议，我们会认真对待每一条反馈。';
            break;
        case 'about':
            modalTitle.value = '关于泡泡';
            modalContent.value = '泡泡（PaoPao）是一个清新文艺的微社区平台。版本：v0.6-alpha';
            break;
    }
    modalShow.value = true;
}

// 关闭侧边栏
function handleClose() {
    emit('close');
}

// 退出登录
function handleLogout() {
    dialog.warning({
        title: '确认退出',
        content: '确定要退出当前账号吗？',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: () => {
            storeUser.userLogout();
            handleClose();
            message.success('已退出登录');
            router.push({ name: 'home' });
        },
    });
}
</script>

<style lang="less" scoped>
.sidebar-content {
    padding: 16px 0;
    overflow-y: auto;
    height: 100%;
}

.menu-group {
    margin-bottom: 16px;

    &:last-child {
        margin-bottom: 0;
    }
}

.menu-item {
    cursor: pointer;
    transition: background-color 0.2s ease;

    &:hover {
        background-color: rgba(24, 160, 88, 0.05);
    }

    .suffix-text {
        font-size: 13px;
        color: #999;
        white-space: nowrap;
        min-width: 56px;
        text-align: right;
    }

    &.logout-item {
        color: #d03050;

        &:hover {
            background-color: rgba(208, 48, 80, 0.05);
        }
    }
}

:deep(.n-list) {
    --n-border-radius: 12px;
}

// 去掉边框 + 左侧10px
:deep(.n-list) {
    border: none !important;
    border-radius: 0 !important;
    box-shadow: none !important;
}

:deep(.n-list-item) {
    border-bottom: none !important;
    padding-left: 10px !important;
    padding-right: 10px !important;
}

:deep(.n-list-item__divider) {
    display: none !important;
}

:deep(.n-list-item__main) {
    align-items: center;
}

.dark {
    .menu-item:hover {
        background-color: rgba(99, 226, 183, 0.05);
    }

    .suffix-text {
        color: #666;
    }

    .logout-item {
        color: #e53935;

        &:hover {
            background-color: rgba(229, 57, 53, 0.1);
        }
    }
}

</style>
