<template>
    <div>

        <div class="home-content">
            <!-- Banner 轮播 -->
            <banner-carousel />

            <!-- 标签切换: 推荐 / 关注 -->
            <div class="feed-tabs">
                <div
                    v-for="tab in feedTabs"
                    :key="tab.key"
                    class="feed-tab"
                    :class="{ active: activeFeed === tab.key }"
                    @click="switchFeed(tab.key)"
                >
                    {{ tab.label }}
                </div>
            </div>

            <!-- 搜索提示 -->
            <div v-if="route.query.q" class="search-hint">
                搜索: {{ decodeURIComponent(route.query.q as string) }}
                <n-button size="tiny" quaternary @click="clearSearch">清除</n-button>
            </div>

            <!-- 加载骨架 -->
            <div v-if="loading && list.length === 0" class="skeleton-wrap">
                <post-skeleton :num="pageSize" />
            </div>

            <!-- 空状态 -->
            <div v-if="!loading && list.length === 0" class="empty-wrap">
                <n-empty size="large" description="暂无数据" />
            </div>

            <!-- 帖子列表 -->
            <div class="post-list">
                <post-item
                    v-for="p in list"
                    :key="p.id"
                    :post="p"
                    :is-owner="userInfo.id == p.user_id"
                    :is-mobile="true"
                    add-follow-action
                    @send-whisper="onSendWhisper"
                    @post-follow-action="postFollowAction"
                    @handle-friend-action="onHandleFriendAction"
                />
            </div>

            <!-- 无限加载 -->
            <n-space v-if="totalPage > 0" justify="center">
                <InfiniteLoading
                    class="load-more"
                    :slots="{ complete: '没有更多泡泡了', error: '加载出错' }"
                    @infinite="handleNextPage"
                >
                    <template #spinner>
                        <div class="load-more-wrap">
                            <n-spin :size="14" v-if="!noMore" />
                            <span class="load-more-spinner">{{ noMore ? '没有更多泡泡了' : '加载更多' }}</span>
                        </div>
                    </template>
                </InfiniteLoading>
            </n-space>
        </div>

        <!-- 私信组件 -->
        <whisper :show="showWhisper" :user="whisperReceiver" @success="whisperSuccess" />
        <!-- 加好友组件 -->
        <whisper-add-friend :show="showAddFriendWhisper" :user="user" @success="addFriendWhisperSuccess" />
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, reactive, computed, watch } from 'vue';
import { useStoreMain } from '@/store/main';
import { useRoute, useRouter } from 'vue-router';
import { useDialog } from 'naive-ui';
import InfiniteLoading from 'v3-infinite-loading';
import { getPosts } from '@/api/post';
import { useStoreUser } from '@/store/user';
import { useStoreProfile } from '@/store/profile';
import { storeToRefs } from 'pinia';
import { Api } from '@/utils/request';
import { usePagination } from '@/composables/usePagination';
import UserAction from '@/composables/useUserAction';
import TweetComposerTrigger from '@/components/tweet-composer-trigger.vue';
import BannerCarousel from '@/components/banner-carousel.vue';

const storeMain = useStoreMain();
const storeUser = useStoreUser();
const storeProfile = useStoreProfile();
const { refresh } = storeToRefs(storeMain);
const { userInfo } = storeToRefs(storeUser);
const { profile } = storeToRefs(storeProfile);

const route = useRoute();
const router = useRouter();
const dialog = useDialog();

const activeFeed = ref<string>('recommend');

const feedTabs = [
    { key: 'recommend', label: '推荐' },
    { key: 'following', label: '关注' },
];

const title = ref<string>('泡泡广场');
const list = ref<any[]>([]);
const showAddFriendWhisper = ref(false);

// 使用 usePagination composable
const { loading, noMore, page, pageSize, totalPage, reset, nextPage } = usePagination(20);

// 使用 UserAction.useWhisper()
const { showWhisper, whisperReceiver, onSendWhisper, whisperSuccess } = UserAction.useWhisper();

const user = reactive<Item.UserInfo>({
  id: 0,
  avatar: '',
  username: '',
  nickname: '',
  is_admin: false,
  is_friend: false,
  is_following: false,
  created_on: 0,
  follows: 0,
  followings: 0,
  status: 1,
});
const inActionPost = ref<Item.PostProps | null>(null);

function openCompose() {
    storeMain.triggerCompose(true);
}

function switchFeed(key: string) {
    activeFeed.value = key;
    resetAll();
    loadPosts(key === 'following' ? 'following' : 'newest');
}

function clearSearch() {
    router.push({ name: 'home' });
}

function postFollowAction(userId: string, isFollowing: boolean) {
  for (let index in list.value) {
    if (list.value[index].user_id == userId) {
      list.value[index].user.is_following = isFollowing;
    }
  }
  if (activeFeed.value === 'following' && !isFollowing) {
    resetAll();
    loadPosts('following');
  }
}

const updateTitle = () => {
  title.value = '泡泡广场';
  if (route.query && route.query.q) {
    if (route.query.t && route.query.t === 'tag') {
      title.value = '#' + decodeURIComponent(route.query.q as string);
    } else {
      title.value = '搜索: ' + decodeURIComponent(route.query.q as string);
    }
  }
};

const resetAll = () => {
  reset();
  list.value = [];
};

function loadPosts(style: 'newest' | 'hots' | 'following' | 'search') {
  loading.value = true;
  getPosts({
    query: route.query.q ? decodeURIComponent(route.query.q as string) : null,
    type: route.query.t as string,
    style: style,
    page: page.value,
    page_size: pageSize.value,
  })
    .then((rsp) => {
      loading.value = false;
      if (rsp.list.length === 0) {
        noMore.value = true;
      }
      if (page.value > 1) {
        list.value = list.value.concat(rsp.list);
      } else {
        list.value = rsp.list;
        window.scrollTo(0, 0);
      }
      totalPage.value = Math.ceil(rsp.pager.total_rows / pageSize.value);
    })
    .catch((err) => {
      loading.value = false;
      if (page.value > 1) {
        page.value--;
      }
    });
}

function handleNextPage() {
  const style = activeFeed.value === 'following' ? 'following' : 'newest';
  nextPage(() => loadPosts(style));
}

const addFriendWhisperSuccess = () => {
  showAddFriendWhisper.value = false;
  inActionPost.value = null;
};

const onHandleFriendAction = (post: Item.PostProps) => {
  inActionPost.value = post;
  user.id = post.user.id;
  user.username = post.user.username;
  user.nickname = post.user.nickname;
  if (post.user.is_friend) {
    dialog.warning({
      title: '删除好友',
      content: '将好友 "' + post.user.nickname + '" 删除',
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        Api.v1.friend.post.delete({ user_id: user.id })
          .then((_res) => {
            window.$message.success('操作成功');
            post.user.is_friend = false;
          })
          .catch((_err) => {});
      },
    });
  } else {
    showAddFriendWhisper.value = true;
  }
};

onMounted(() => {
  resetAll();
  loadPosts('newest');
});

watch(
  () => ({
    path: route.path,
    query: route.query,
    refresh: refresh.value,
  }),
  (to, from) => {
    updateTitle();
    if (to.refresh !== from.refresh) {
      resetAll();
      setTimeout(() => {
        const style = activeFeed.value === 'following' ? 'following' : 'newest';
        loadPosts(style);
      }, 0);
      return;
    }
    if (from.path !== '/post' && to.path === '/') {
      resetAll();
      setTimeout(() => {
        const style = activeFeed.value === 'following' ? 'following' : 'newest';
        loadPosts(style);
      }, 0);
    }
  },
);
</script>

<style lang="less" scoped>
.home-content {
    min-height: 60vh;
}

.feed-tabs {
    display: flex;
    border-bottom: 1px solid var(--border-color, #eff3f4);
    background: rgba(255, 255, 255, 0.85);
    backdrop-filter: blur(12px);
    position: sticky;
    top: 0;
    z-index: 10;
}

.feed-tab {
    padding: 14px 20px;
    font-size: 15px;
    font-weight: 500;
    color: #666;
    cursor: pointer;
    position: relative;
    transition: color 0.2s;

    &:hover {
        color: #18a058;
    }

    &.active {
        color: #18a058;
        font-weight: 700;

        &::after {
            content: '';
            position: absolute;
            bottom: 0;
            left: 20px;
            right: 20px;
            height: 4px;
            background: #18a058;
            border-radius: 2px;
        }
    }
}

.search-hint {
    padding: 12px 16px;
    font-size: 14px;
    color: #888;
    display: flex;
    align-items: center;
    justify-content: space-between;
    border-bottom: 1px solid var(--border-color, #eff3f4);
}

.skeleton-wrap {
    padding: 8px 16px;
}

.post-list {
    // Posts render with their own border-bottom
}

.load-more {
    margin: 20px;

    .load-more-wrap {
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;
        gap: 14px;

        .load-more-spinner {
            font-size: 14px;
            opacity: 0.65;
        }
    }
}

.dark {
    .feed-tabs {
        background: rgba(16, 16, 20, 0.85);
        border-bottom-color: #2f3336;
    }

    .search-hint {
        border-bottom-color: #2f3336;
    }
}
</style>
