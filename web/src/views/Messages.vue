<template>
    <div>
        <!-- 私信组件 -->
        <whisper :show="showWhisper" :user="whisperReceiver" @success="whisperSuccess" />

        <!-- 三Tab切换 -->
        <div class="msg-tabs">
            <div
                v-for="tab in tabs"
                :key="tab.key"
                class="msg-tab"
                :class="{ active: activeTab === tab.key }"
                @click="switchTab(tab.key)"
            >
                {{ tab.label }}
                <span
                    v-if="(tab.key === 'dm' || tab.key === 'system') && getTabUnread(tab.key) > 0"
                    class="tab-unread"
                >{{ getTabUnread(tab.key) > 99 ? '99+' : getTabUnread(tab.key) }}</span>
            </div>
        </div>

        <!-- 聊天 Tab: Mock对话列表 -->
        <div v-if="activeTab === 'chat'" class="chat-list">
            <div
                v-for="conv in chatList"
                :key="conv.id"
                class="chat-item"
                @click="openChat(conv)"
            >
                <n-badge :value="conv.unread" :max="99" :show="conv.unread > 0">
                    <n-avatar round :size="48" :src="conv.avatar" />
                </n-badge>
                <div class="chat-info">
                    <div class="chat-top">
                        <span class="chat-name">{{ conv.name }}</span>
                        <span class="chat-time">{{ conv.time }}</span>
                    </div>
                    <div class="chat-bottom">
                        <span class="chat-last">{{ conv.lastMsg }}</span>
                    </div>
                </div>
            </div>
            <div v-if="chatList.length === 0" class="empty-wrap">
                <n-empty description="暂无聊天" />
            </div>
        </div>

        <!-- 私信 Tab -->
        <div v-if="activeTab === 'dm'">
            <div v-if="loading && list.length === 0" class="skeleton-wrap">
                <message-skeleton :num="pageSize" />
            </div>
            <div v-else>
                <div class="empty-wrap" v-if="list.length === 0">
                    <n-empty size="large" description="暂无消息" />
                </div>
                <div v-else class="msg-list">
                    <div v-for="m in list" :key="m.id" class="msg-card">
                        <message-item :message="m" @send-whisper="onSendWhisper" @reload="reloadMessages" @read-message="onReadMessage" />
                    </div>
                </div>
            </div>
            <infinite-load-more
                :total-page="totalPage"
                :no-more="noMore"
                complete-text="没有更多消息了"
                @load-more="nextPage"
            />
        </div>

        <!-- 系统消息 Tab -->
        <div v-if="activeTab === 'system'">
            <div v-if="sysLoading && sysList.length === 0" class="skeleton-wrap">
                <message-skeleton :num="sysPageSize" />
            </div>
            <div v-else>
                <div class="empty-wrap" v-if="sysList.length === 0">
                    <n-empty size="large" description="暂无系统消息" />
                </div>
                <div v-else class="msg-list">
                    <div v-for="m in sysList" :key="m.id" class="msg-card">
                        <message-item :message="m" @send-whisper="onSendWhisper" @reload="reloadSysMessages" @read-message="onReadMessage" />
                    </div>
                </div>
            </div>
            <infinite-load-more
                :total-page="sysTotalPage"
                :no-more="sysNoMore"
                complete-text="没有更多系统消息了"
                @load-more="sysNextPage"
            />
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useStoreMain } from '@/store/main';
import { storeToRefs } from 'pinia';
import { Api } from '@/utils/request';
import InfiniteLoadMore from '@/components/infinite-load-more.vue';

const route = useRoute();
const router = useRouter();
const storeMain = useStoreMain();
const { unreadMsgCount } = storeToRefs(storeMain);

// Tab state
type TabKey = 'chat' | 'dm' | 'system';
const activeTab = ref<TabKey>('chat');
const tabs = [
    { key: 'chat' as const, label: '聊天' },
    { key: 'dm' as const, label: '私信' },
    { key: 'system' as const, label: '系统消息' },
];

const dmUnread = ref(0);
const sysUnread = ref(0);
function getTabUnread(key: string) {
    if (key === 'dm') return dmUnread.value;
    if (key === 'system') return sysUnread.value;
    return 0;
}
function switchTab(key: TabKey) {
    activeTab.value = key;
}

// === 聊天 Mock数据 ===
interface ChatConv {
    id: number;
    name: string;
    avatar: string;
    lastMsg: string;
    time: string;
    unread: number;
}

const chatList = ref<ChatConv[]>([
    { id: 1, name: '张经理', avatar: '', lastMsg: '好的，明天会议上讨论这个方案', time: '10:32', unread: 2 },
    { id: 2, name: '李总监', avatar: '', lastMsg: '文件已经发你邮箱了，请查收', time: '09:15', unread: 0 },
    { id: 3, name: '王主管', avatar: '', lastMsg: '收到，我马上安排', time: '昨天', unread: 1 },
    { id: 4, name: '赵工程师', avatar: '', lastMsg: '这个bug已经修好了，你测试一下', time: '昨天', unread: 0 },
    { id: 5, name: '钱设计师', avatar: '', lastMsg: '新版的UI稿你看一下，有问题随时沟通', time: '周三', unread: 0 },
    { id: 6, name: '孙分析师', avatar: '', lastMsg: '数据报告已经完成，下午汇报', time: '周三', unread: 3 },
    { id: 7, name: '周秘书', avatar: '', lastMsg: '明天的会议改到下午3点', time: '周一', unread: 0 },
]);

function openChat(_conv: ChatConv) {
    router.push({ name: 'chat' });
}

// === 私信 (whisper + requesting) ===
const list = ref<Item.MessageProps[]>([]);
const loading = ref(false);
const noMore = ref(false);
const page = ref(1);
const pageSize = ref(20);
const totalPage = ref(0);
const showWhisper = ref(false);
const whisperReceiver = ref<Item.UserInfo>({
    id: 0, avatar: '', username: '', nickname: '',
    is_admin: false, is_friend: true, is_following: false,
    created_on: 0, follows: 0, followings: 0, status: 1,
});

function reset() {
    page.value = 1;
    list.value = [];
}

function reloadMessages() {
    reset();
    loadMessages();
}

function loadMessages() {
    loading.value = true;
    // 合并 whispers 和 friend requests
    Promise.all([
        Api.v1.user.get.messages({ style: 'whisper', page: page.value, page_size: pageSize.value }),
        Api.v1.user.get.messages({ style: 'requesting', page: 1, page_size: 50 }),
    ])
        .then(([whisperRes, friendRes]: any[]) => {
            loading.value = false;
            const combined = [
                ...(friendRes?.list || []),
                ...(whisperRes?.list || []),
            ].sort((a: any, b: any) => b.created_on - a.created_on);
            dmUnread.value = combined.filter((m: any) => !m.is_read).length;
            if (combined.length === 0 && page.value === 1) {
                noMore.value = true;
            }
            list.value = combined;
            totalPage.value = Math.ceil((whisperRes?.pager?.total_rows || 0) / pageSize.value);
        })
        .catch((_err: any) => {
            loading.value = false;
        });
}

function nextPage() {
    if (page.value < totalPage.value || totalPage.value === 0) {
        noMore.value = false;
        page.value++;
        loadMessages();
    } else {
        noMore.value = true;
    }
}

// === 系统消息 ===
const sysList = ref<Item.MessageProps[]>([]);
const sysLoading = ref(false);
const sysNoMore = ref(false);
const sysPage = ref(1);
const sysPageSize = ref(20);
const sysTotalPage = ref(0);

function sysReset() {
    sysPage.value = 1;
    sysList.value = [];
}

function reloadSysMessages() {
    sysReset();
    loadSysMessages();
}

// 处理消息已读 - 立即更新对应Tab的未读计数
function onReadMessage(messageType: number) {
    // type = 99 是系统消息，其他（1-5）是私信相关
    if (messageType === 99) {
        // 系统消息未读减1
        if (sysUnread.value > 0) {
            sysUnread.value--;
        }
    } else {
        // 私信未读减1
        if (dmUnread.value > 0) {
            dmUnread.value--;
        }
    }
}

function loadSysMessages() {
    sysLoading.value = true;
    Api.v1.user.get.messages({ style: 'system', page: sysPage.value, page_size: sysPageSize.value })
        .then((res: any) => {
            sysLoading.value = false;
            // 统计未读系统消息
            sysUnread.value = (res.list || []).filter((m: any) => !m.is_read).length;
            if (res.list.length === 0) {
                sysNoMore.value = true;
            }
            if (sysPage.value > 1) {
                sysList.value = sysList.value.concat(res.list);
            } else {
                sysList.value = res.list;
                window.scrollTo(0, 0);
            }
            sysTotalPage.value = Math.ceil(res.pager.total_rows / sysPageSize.value);
        })
        .catch((_err: any) => {
            sysLoading.value = false;
        });
}

function sysNextPage() {
    if (sysPage.value < sysTotalPage.value || sysTotalPage.value === 0) {
        sysNoMore.value = false;
        sysPage.value++;
        loadSysMessages();
    } else {
        sysNoMore.value = true;
    }
}

// === Whisper ===
const onSendWhisper = (user: Item.UserInfo) => {
    whisperReceiver.value = user;
    showWhisper.value = true;
};
const whisperSuccess = () => {
    showWhisper.value = false;
};

onMounted(() => {
    const tab = route.query.tab as string;
    if (tab === 'chat' || tab === 'dm' || tab === 'system') {
        activeTab.value = tab;
    }
    loadMessages();
    loadSysMessages();
});
</script>

<style lang="less" scoped>
// Tab栏
.msg-tabs {
    display: flex;
    border-bottom: 1px solid var(--border-color, #eff3f4);
    background: rgba(255, 255, 255, 0.85);
    backdrop-filter: blur(12px);
    position: sticky;
    top: 0;
    z-index: 10;
}

.msg-tab {
    flex: 1;
    text-align: center;
    padding: 14px 0;
    font-size: 15px;
    font-weight: 500;
    color: #666;
    cursor: pointer;
    position: relative;
    transition: color 0.2s;

    &:hover { color: #18a058; }

    &.active {
        color: #18a058;
        font-weight: 700;
        &::after {
            content: '';
            position: absolute;
            bottom: 0;
            left: 50%;
            transform: translateX(-50%);
            width: 36px;
            height: 3px;
            background: #18a058;
            border-radius: 2px;
        }
    }
}

.tab-unread {
    position: absolute;
    top: 6px;
    right: 20px;
    min-width: 18px;
    height: 18px;
    line-height: 18px;
    padding: 0 5px;
    font-size: 11px;
    font-weight: 600;
    color: #fff;
    background: #d93025;
    border-radius: 9px;
    text-align: center;
    z-index: 1;
}

// 聊天列表
.chat-list {
    min-height: 60vh;
}

.chat-item {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 14px 16px;
    border-bottom: 1px solid var(--border-color, #eff3f4);
    cursor: pointer;
    transition: background 0.15s;

    &:hover { background: #f7f9f9; }
}

.chat-info {
    flex: 1;
    min-width: 0;
}

.chat-top {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 4px;
}

.chat-name {
    font-size: 16px;
    font-weight: 600;
}

.chat-time {
    font-size: 12px;
    color: #999;
    flex-shrink: 0;
}

.chat-last {
    font-size: 13px;
    color: #999;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    display: block;
}

// 消息列表
.msg-list {
    min-height: 60vh;
}

.msg-card {
    border-bottom: 1px solid var(--border-color, #eff3f4);
}

.empty-wrap {
    min-height: 300px;
    display: flex;
    align-items: center;
    justify-content: center;
}

.skeleton-wrap {
    padding: 12px 16px;
}

.dark {
    .msg-tabs {
        background: rgba(16, 16, 20, 0.85);
        border-bottom-color: #2f3336;
    }
    .chat-item {
        border-bottom-color: #2f3336;
        &:hover { background: #080808; }
    }
    .msg-card {
        border-bottom-color: #2f3336;
    }
}
</style>
