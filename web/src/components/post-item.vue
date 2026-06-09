<template>
    <article class="post-item" @click="goPostDetail(post.id)">
        <div class="post-layout">
            <!-- 左侧头像 -->
            <div class="post-avatar" @click.stop="goUserProfile">
                <n-avatar round :size="48" :src="post.user.avatar" />
            </div>

            <!-- 右侧内容 -->
            <div class="post-body">
                <!-- 头部：昵称 + 用户名 + 时间 + 更多 -->
                <div class="post-header">
                    <div class="post-header-left">
                        <span class="nickname" @click.stop="goUserProfile">
                            {{ post.user.nickname }}
                        </span>
                        <span class="dot">·</span>
                        <span class="timestamp">{{ formatPrettyDate(post.created_on) }}</span>
                        <n-tag
                            v-if="post.is_top"
                            class="badge-tag"
                            type="warning"
                            size="tiny"
                            round
                        >
                            置顶
                        </n-tag>
                        <n-tag
                            v-if="post.visibility == 1"
                            class="badge-tag"
                            type="error"
                            size="tiny"
                            round
                        >
                            私密
                        </n-tag>
                        <n-tag
                            v-if="post.visibility == 2"
                            class="badge-tag"
                            type="info"
                            size="tiny"
                            round
                        >
                            好友可见
                        </n-tag>
                    </div>
                    <n-dropdown
                        placement="bottom-end"
                        trigger="click"
                        size="small"
                        :options="tweetOptions"
                        @select="handleTweetAction"
                    >
                        <n-button class="more-btn" quaternary circle size="tiny" @click.stop>
                            <template #icon>
                                <n-icon :size="18"><more-horiz-filled /></n-icon>
                            </template>
                        </n-button>
                    </n-dropdown>
                </div>

                <!-- 位置信息 -->
                <div v-if="post.ip_loc && !isMobile" class="post-location">
                    {{ post.ip_loc }}
                </div>

                <!-- 正文内容 -->
                <div v-if="post.texts.length > 0" class="post-content">
                    <span
                        v-for="content in post.texts"
                        :key="content.id"
                        class="post-text"
                        @click.stop="doClickText($event, post.id)"
                        v-html="preparePost(content.content, '展开', '收起', isMobile ? profile.tweetMobileEllipsisSize : profile.tweetWebEllipsisSize, inFoldStyle)"
                    ></span>
                </div>

                <!-- 附件/图片/视频/链接 -->
                <div class="post-media">
                    <post-attachment
                        v-if="post.attachments.length > 0"
                        :attachments="post.attachments" />
                    <post-attachment
                        v-if="post.charge_attachments.length > 0"
                        :attachments="post.charge_attachments"
                        :price="post.attachment_price"
                    />
                    <post-image
                        v-if="post.imgs.length > 0"
                        :imgs="post.imgs" />
                    <post-video
                        v-if="post.videos.length > 0"
                        :videos="post.videos" />
                    <post-link
                        v-if="post.links.length > 0"
                        :links="post.links" />
                </div>

                <!-- 操作栏 -->
                <post-action-bar
                    :comment-count="post.comment_count"
                    :repost-count="0"
                    :like-count="post.upvote_count"
                    :bookmark-count="post.collection_count"
                    :is-liked="false"
                    :is-bookmarked="false"
                    @comment="goPostDetail(post.id)"
                    @repost="handleRepost"
                    @like="handlePostStar"
                    @bookmark="handlePostCollection"
                    @share="handleShare"
                />
            </div>
        </div>
    </article>
</template>

<script setup lang="ts">
import { h, ref, computed } from 'vue';
import { useStoreMain } from '@/store/main';
import { useRouter } from 'vue-router';
import { NIcon, useDialog } from 'naive-ui';
import type { Component } from 'vue';
import type { DropdownOption } from 'naive-ui';
import { formatPrettyDate } from '@/utils/formatTime';
import { preparePost } from '@/utils/content';
import { postStar, postCollection } from '@/api/post';
import {
  PaperPlaneOutline,
  HeartOutline,
  BookmarkOutline,
  ChatboxOutline,
  ShareSocialOutline,
  PersonAddOutline,
  PersonRemoveOutline,
  BodyOutline,
  WalkOutline,
} from '@vicons/ionicons5';
import { MoreHorizFilled } from '@vicons/material';
import copy from 'copy-to-clipboard';
import { useStoreProfile } from '@/store/profile';
import { storeToRefs } from 'pinia';
import { Api } from '@/utils/request';
import UserAction from '@/composables/useUserAction';
import { usePostContent } from '@/composables/usePostContent';
import PostActionBar from '@/components/post-action-bar.vue';

const router = useRouter();

const storeMain = useStoreMain();
const storeProfile = useStoreProfile();
const { profile } = storeToRefs(storeProfile);

const dialog = useDialog();

const inFoldStyle = ref<boolean>(true);
const props = withDefaults(defineProps<{
    post: Item.PostProps;
    isOwner: boolean;
    addFriendAction?: boolean;
    addFollowAction?: boolean;
    isMobile?: boolean;
}>(), {
    addFollowAction: false,
    addFriendAction: false,
    isMobile: false,
});

const emit = defineEmits<{
  (e: 'send-whisper', user: Item.UserInfo): void;
  (e: 'handle-follow-action', user: Item.PostProps): void;
  (e: 'handle-friend-action', user: Item.PostProps): void;
  (e: 'post-follow-action', user_id: string, is_following: boolean): void;
}>();

const renderIcon = (icon: Component) => {
  return () => {
    return h(NIcon, null, {
      default: () => h(icon),
    });
  };
};

const tweetOptions = computed(() => {
  let options: DropdownOption[] = [];
  if (!props.isOwner) {
    options.push({
      label: '私信 ' + props.post.user.nickname,
      key: 'whisper',
      icon: renderIcon(PaperPlaneOutline),
    });
  }
  if (!props.isOwner && props.addFollowAction) {
    if (props.post.user.is_following) {
      options.push({
        label: '取消关注 ' + props.post.user.nickname,
        key: 'unfollow',
        icon: renderIcon(WalkOutline),
      });
    } else {
      options.push({
        label: '关注 ' + props.post.user.nickname,
        key: 'follow',
        icon: renderIcon(BodyOutline),
      });
    }
  }
  if (!props.isOwner && props.addFriendAction) {
    if (props.post.user.is_friend) {
      options.push({
        label: '删除好友 ' + props.post.user.nickname,
        key: 'delete',
        icon: renderIcon(PersonRemoveOutline),
      });
    } else {
      options.push({
        label: '添加朋友 ' + props.post.user.nickname,
        key: 'requesting',
        icon: renderIcon(PersonAddOutline),
      });
    }
  }
  options.push({
    label: '复制链接',
    key: 'copyTweetLink',
    icon: renderIcon(ShareSocialOutline),
  });
  return options;
});

const handleTweetAction = async (
  item:
    | 'copyTweetLink'
    | 'whisper'
    | 'follow'
    | 'unfollow'
    | 'delete'
    | 'requesting',
) => {
  switch (item) {
    case 'copyTweetLink':
      copy(
        `${window.location.origin}/#/post?id=${post.value.id}&share=copy_link&t=${new Date().getTime()}`,
      );
      window.$message.success('链接已复制到剪贴板');
      break;
    case 'whisper':
      emit('send-whisper', props.post.user);
      break;
    case 'delete':
    case 'requesting':
      emit('handle-friend-action', props.post);
      break;
    case 'follow':
    case 'unfollow':
      UserAction.followAction(dialog, props.post.user.id, props.post.user.username, props.post.user.is_following)
        .then(_action => {
          emit('post-follow-action', props.post.user.id, _action);
        });
      emit('handle-follow-action', props.post);
      break;
    default:
        break;
  }
};

// 使用 usePostContent composable
const post = usePostContent(props.post);
const handlePostStar = () => {
  postStar({
    id: post.value.id,
  })
    .then((res) => {
      if (res.status) {
        post.value = {
          ...post.value,
          upvote_count: post.value.upvote_count + 1,
        };
      } else {
        post.value = {
          ...post.value,
          upvote_count:
            post.value.upvote_count > 0 ? post.value.upvote_count - 1 : 0,
        };
      }
    })
    .catch((err) => {
      console.log(err);
    });
};
const handlePostCollection = () => {
  postCollection({
    id: post.value.id,
  })
    .then((res) => {
      if (res.status) {
        post.value = {
          ...post.value,
          collection_count: post.value.collection_count + 1,
        };
      } else {
        post.value = {
          ...post.value,
          collection_count:
            post.value.collection_count > 0
              ? post.value.collection_count - 1
              : 0,
        };
      }
    })
    .catch((err) => {
      console.log(err);
    });
};
const handleRepost = () => {
  copy(
    `${window.location.origin}/#/post?id=${post.value.id}&share=copy_link&t=${new Date().getTime()}`,
  );
  window.$message.success('链接已复制到剪贴板');
};
const handleShare = () => {
  copy(
    `${window.location.origin}/#/post?id=${post.value.id}&share=copy_link&t=${new Date().getTime()}`,
  );
  window.$message.success('链接已复制到剪贴板');
};
const goPostDetail = (id: string) => {
  router.push({
    name: 'post',
    query: {
      id,
    },
  });
};
const goUserProfile = () => {
  router.push({
    name: 'user',
    query: {
      s: post.value.user.username,
    },
  });
};
const doClickText = (e: MouseEvent, id: string) => {
  const detail = (e.target as any).dataset.detail;
  if (detail && detail !== 'post') {
    const d = detail.split(':');
    if (d.length === 2) {
      storeMain.doRefresh();
      if (d[0] === 'tag') {
        router.push({
          name: 'home',
          query: {
            q: d[1],
            t: 'tag',
          },
        });
      } else {
        router.push({
          name: 'user',
          query: {
            s: d[1],
          },
        });
      }
    }
  } else if (detail && detail === 'post') {
    inFoldStyle.value = !inFoldStyle.value;
  } else {
    goPostDetail(id);
  }
};
</script>

<style lang="less" scoped>
.post-item {
    padding: 12px 16px;
    border-bottom: 1px solid var(--border-color, #eff3f4);
    cursor: pointer;
    transition: background-color 0.2s ease;

    &:hover {
        background: #f7f9f9;
    }
}

.post-layout {
    display: flex;
    gap: 12px;
}

.post-avatar {
    flex-shrink: 0;
    cursor: pointer;

    &:hover {
        opacity: 0.85;
    }
}

.post-body {
    flex: 1;
    min-width: 0;
}

.post-header {
    display: flex;
    align-items: flex-start;
    justify-content: space-between;
    gap: 8px;
}

.post-header-left {
    display: flex;
    align-items: center;
    flex-wrap: wrap;
    gap: 4px;
    font-size: 14px;
    line-height: 20px;
}

.nickname {
    font-weight: 700;
    color: inherit;
    cursor: pointer;

    &:hover {
        text-decoration: underline;
    }
}

.username {
    color: #888;
    font-size: 14px;
}

.dot {
    color: #888;
}

.timestamp {
    color: #888;
    font-size: 14px;
}

.badge-tag {
    transform: scale(0.75);
    transform-origin: left center;
}

.more-btn {
    flex-shrink: 0;
    color: #888;
}

.post-location {
    font-size: 13px;
    color: #888;
    margin-top: 2px;
}

.post-content {
    margin-top: 4px;
}

.post-text {
    text-align: justify;
    overflow: hidden;
    white-space: pre-wrap;
    word-break: break-all;
    font-size: 15px;
    line-height: 1.5;
}

.post-media {
    margin-top: 8px;
}

.dark {
    .post-item {
        border-bottom-color: #2f3336;

        &:hover {
            background: #080808;
        }
    }

    .username,
    .dot,
    .timestamp,
    .more-btn,
    .post-location {
        color: #999;
    }
}
</style>
