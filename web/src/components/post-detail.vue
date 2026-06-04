<template>
    <article class="detail-item" @click="goPostDetail(post.id)">
        <!-- 删除确认 -->
        <n-modal
            v-model:show="showDelModal"
            :mask-closable="false"
            preset="dialog"
            title="提示"
            content="确定删除该泡泡动态吗？"
            positive-text="确认"
            negative-text="取消"
            @positive-click="execDelAction"
        />
        <!-- 锁定确认 -->
        <n-modal
            v-model:show="showLockModal"
            :mask-closable="false"
            preset="dialog"
            title="提示"
            :content="
                '确定' +
                (post.is_lock ? '解锁' : '锁定') +
                '该泡泡动态吗？'
            "
            positive-text="确认"
            negative-text="取消"
            @positive-click="execLockAction"
        />
        <!-- 置顶确认 -->
        <n-modal
            v-model:show="showStickModal"
            :mask-closable="false"
            preset="dialog"
            title="提示"
            :content="
                '确定' +
                (post.is_top ? '取消置顶' : '置顶') +
                '该泡泡动态吗？'
            "
            positive-text="确认"
            negative-text="取消"
            @positive-click="execStickAction"
        />
        <!-- 亮点确认 -->
        <n-modal
            v-model:show="showHighlightModal"
            :mask-closable="false"
            preset="dialog"
            title="提示"
            :content="
                '确定将该泡泡动态' +
                (post.is_essence ? '取消亮点' : '设为亮点') +
                '吗？'
            "
            positive-text="确认"
            negative-text="取消"
            @positive-click="execHighlightAction"
        />
        <!-- 修改可见度确认 -->
        <n-modal
            v-model:show="showVisibilityModal"
            :mask-closable="false"
            preset="dialog"
            title="提示"
            :content="
                '确定将该泡泡动态可见度修改为' +
                (tempVisibility == 0 ? '公开' : (tempVisibility == 1 ? '私密' : (tempVisibility == 2 ? '好友可见' : '关注可见'))) +
                '吗？'
            "
            positive-text="确认"
            negative-text="取消"
            @positive-click="execVisibilityAction"
        />
        <!-- 私信组件 -->
        <whisper :show="showWhisper" :user="whisperReceiver" @success="whisperSuccess" />

        <div class="post-layout">
            <!-- 左侧头像 -->
            <div class="post-avatar" @click.stop="goUserProfile">
                <n-avatar round :size="48" :src="post.user.avatar" />
            </div>

            <!-- 右侧内容 -->
            <div class="post-body">
                <!-- 头部 -->
                <div class="post-header">
                    <div class="post-header-left">
                        <span class="nickname" @click.stop="goUserProfile">
                            {{ post.user.nickname }}
                        </span>
                        <span class="username">@{{ post.user.username }}</span>
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
                            v-if="post.visibility == VisibilityEnum.PRIVATE"
                            class="badge-tag"
                            type="error"
                            size="tiny"
                            round
                        >
                            私密
                        </n-tag>
                        <n-tag
                            v-if="post.visibility == VisibilityEnum.FRIEND"
                            class="badge-tag"
                            type="info"
                            size="tiny"
                            round
                        >
                            好友可见
                        </n-tag>
                    </div>
                    <div class="options">
                        <n-dropdown
                            placement="bottom-end"
                            trigger="click"
                            size="small"
                            :options="adminOptions"
                            @select="handlePostAction"
                        >
                            <n-button quaternary circle size="tiny">
                                <template #icon>
                                    <n-icon :size="18"><more-horiz-filled /></n-icon>
                                </template>
                            </n-button>
                        </n-dropdown>
                    </div>
                </div>

                <!-- 正文内容 -->
                <div v-if="post.texts.length > 0" class="post-content">
                    <span
                        v-for="content in post.texts"
                        :key="content.id"
                        class="post-text"
                        @click.stop="doClickText($event, post.id)"
                        v-html="parsePostTag(content.content).content"
                    >
                    </span>
                </div>

                <!-- 附件/图片/视频/链接 -->
                <div class="post-media">
                    <post-attachment :attachments="post.attachments" />
                    <post-attachment
                        :attachments="post.charge_attachments"
                        :price="post.attachment_price"
                    />
                    <post-image :imgs="post.imgs" />
                    <post-video :videos="post.videos" :full="true" />
                    <post-link :links="post.links" />
                </div>

                <!-- 时间戳 -->
                <div class="timestamp">
                    发布于 {{ formatPrettyTime(post.created_on) }}
                    <span v-if="post.ip_loc">
                        <n-divider vertical />
                        {{ post.ip_loc }}
                    </span>
                    <span v-if="!collapsedLeft && post.created_on != post.latest_replied_on">
                        <n-divider vertical /> 最后回复
                        {{ formatPrettyTime(post.latest_replied_on) }}
                    </span>
                </div>

                <!-- 操作栏 -->
                <post-action-bar
                    :comment-count="post.comment_count"
                    :repost-count="post.share_count || 0"
                    :like-count="post.upvote_count"
                    :bookmark-count="post.collection_count"
                    :is-liked="hasStarred"
                    :is-bookmarked="hasCollected"
                    @comment="goPostDetail(post.id)"
                    @repost="handlePostShare"
                    @like="handlePostStar"
                    @bookmark="handlePostCollection"
                    @share="handlePostShare"
                />
            </div>
        </div>
    </article>
</template>

<script setup lang="ts">
import { h, ref, onMounted, computed } from 'vue';
import type { Component } from 'vue';
import { NIcon, useDialog } from 'naive-ui';
import { useStoreMain } from '@/store/main';
import { useRouter } from 'vue-router';
import { formatPrettyTime } from '@/utils/formatTime';
import { parsePostTag } from '@/utils/content';
import {
  PaperPlaneOutline,
  Heart,
  HeartOutline,
  Bookmark,
  BookmarkOutline,
  ShareSocialOutline,
  ChatboxOutline,
  PushOutline,
  TrashOutline,
  LockClosedOutline,
  LockOpenOutline,
  EyeOutline,
  EyeOffOutline,
  BodyOutline,
  WalkOutline,
  PersonOutline,
  FlameOutline,
} from '@vicons/ionicons5';
import { MoreHorizFilled } from '@vicons/material';
import {
  getPostStar,
  postStar,
  getPostCollection,
  postCollection,
  deletePost,
  lockPost,
  stickPost,
  highlightPost,
  visibilityPost,
} from '@/api/post';
import type { DropdownOption } from 'naive-ui';
import { VisibilityEnum } from '@/utils/IEnum';
import copy from 'copy-to-clipboard';
import { storeToRefs } from 'pinia';
import { useStoreUser } from '@/store/user';
import { Api } from '@/utils/request';
import UserAction from '@/composables/useUserAction';
import { usePostContent } from '@/composables/usePostContent';
import PostActionBar from '@/components/post-action-bar.vue';

const useFriendship =
  import.meta.env.VITE_USE_FRIENDSHIP.toLowerCase() === 'true';

const storeMain = useStoreMain();
const storeUser = useStoreUser();
const { collapsedLeft } = storeToRefs(storeMain);
const { userInfo } = storeToRefs(storeUser);

const router = useRouter();
const dialog = useDialog();
const hasStarred = ref(false);
const hasCollected = ref(false);
const props = withDefaults(
  defineProps<{
    post: Item.PostProps;
  }>(),
  {},
);
const showDelModal = ref(false);
const showLockModal = ref(false);
const showStickModal = ref(false);
const showHighlightModal = ref(false);
const showVisibilityModal = ref(false);
const loading = ref(false);
const tempVisibility = ref<VisibilityEnum>(VisibilityEnum.PUBLIC);
const showWhisper = ref(false);
const whisperReceiver = ref<Item.UserInfo>({
  id: 0,
  avatar: '',
  username: '',
  nickname: '',
  is_admin: false,
  is_friend: true,
  is_following: false,
  created_on: 0,
  follows: 0,
  followings: 0,
  status: 1,
});

const onSendWhisper = (user: Item.UserInfo) => {
  whisperReceiver.value = user;
  showWhisper.value = true;
};

const whisperSuccess = () => {
  showWhisper.value = false;
};

const emit = defineEmits<{
  (e: 'reload', post_id: string): void;
}>();

// 使用 usePostContent composable (包含额外字段)
const post = usePostContent(props.post, true);

const renderIcon = (icon: Component) => {
  return () => {
    return h(NIcon, null, {
      default: () => h(icon),
    });
  };
};

const adminOptions = computed(() => {
  let options: DropdownOption[] = [];
  if (
    !userInfo.value.is_admin &&
    userInfo.value.id != props.post.user.id
  ) {
    options.push({
      label: '私信 @' + props.post.user.username,
      key: 'whisper',
      icon: renderIcon(PaperPlaneOutline),
    });
    if (props.post.user.is_following) {
      options.push({
        label: '取消关注 @' + props.post.user.username,
        key: 'unfollow',
        icon: renderIcon(WalkOutline),
      });
    } else {
      options.push({
        label: '关注 @' + props.post.user.username,
        key: 'follow',
        icon: renderIcon(BodyOutline),
      });
    }
    return options;
  }
  options.push({
    label: '删除',
    key: 'delete',
    icon: renderIcon(TrashOutline),
  });
  if (post.value.is_lock === 0) {
    options.push({
      label: '锁定',
      key: 'lock',
      icon: renderIcon(LockClosedOutline),
    });
  } else {
    options.push({
      label: '解锁',
      key: 'unlock',
      icon: renderIcon(LockOpenOutline),
    });
  }
  if (userInfo.value.is_admin) {
    if (post.value.is_top === 0) {
      options.push({
        label: '置顶',
        key: 'stick',
        icon: renderIcon(PushOutline),
      });
    } else {
      options.push({
        label: '取消置顶',
        key: 'unstick',
        icon: renderIcon(PushOutline),
      });
    }
  }
  if (post.value.is_essence === 0) {
    options.push({
      label: '设为亮点',
      key: 'highlight',
      icon: renderIcon(FlameOutline),
    });
  } else {
    options.push({
      label: '取消亮点',
      key: 'unhighlight',
      icon: renderIcon(FlameOutline),
    });
  }
  let visitMenu: DropdownOption;
  if (post.value.visibility === VisibilityEnum.PUBLIC) {
    visitMenu = {
      label: '公开',
      key: 'vpublic',
      icon: renderIcon(EyeOutline),
      children: [
        { label: '私密', key: 'vprivate', icon: renderIcon(EyeOffOutline) },
        { label: '关注可见', key: 'vfollowing', icon: renderIcon(BodyOutline) },
      ],
    };
  } else if (post.value.visibility === VisibilityEnum.PRIVATE) {
    visitMenu = {
      label: '私密',
      key: 'vprivate',
      icon: renderIcon(EyeOffOutline),
      children: [
        { label: '公开', key: 'vpublic', icon: renderIcon(EyeOutline) },
        { label: '关注可见', key: 'vfollowing', icon: renderIcon(BodyOutline) },
      ],
    };
  } else if (useFriendship && post.value.visibility === VisibilityEnum.FRIEND) {
    visitMenu = {
      label: '好友可见',
      key: 'vfriend',
      icon: renderIcon(PersonOutline),
      children: [
        { label: '公开', key: 'vpublic', icon: renderIcon(EyeOutline) },
        { label: '私密', key: 'vprivate', icon: renderIcon(EyeOffOutline) },
        { label: '关注可见', key: 'vfollowing', icon: renderIcon(BodyOutline) },
      ],
    };
  } else {
    visitMenu = {
      label: '关注可见',
      key: 'vfollowing',
      icon: renderIcon(BodyOutline),
      children: [
        { label: '公开', key: 'vpublic', icon: renderIcon(EyeOutline) },
        { label: '私密', key: 'vprivate', icon: renderIcon(EyeOffOutline) },
      ],
    };
  }
  if (useFriendship && post.value.visibility !== VisibilityEnum.FRIEND) {
    visitMenu.children?.push({
      label: '好友可见',
      key: 'vfriend',
      icon: renderIcon(PersonOutline),
    });
  }
  options.push(visitMenu);
  return options;
});

const onHandleFollowAction = (post: Item.PostProps) => {
	UserAction.followAction(dialog, post.user.id, post.user.username, post.user.is_following)
		.then(_action => {
			post.user.is_following = _action;
		})
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
  if ((e.target as any).dataset.detail) {
    const d = (e.target as any).dataset.detail.split(':');
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
      return;
    }
  }
  goPostDetail(id);
};
const handlePostAction = (
  item:
    | 'whisper'
    | 'follow'
    | 'unfollow'
    | 'delete'
    | 'lock'
    | 'unlock'
    | 'stick'
    | 'unstick'
    | 'highlight'
    | 'unhighlight'
    | 'vpublic'
    | 'vprivate'
    | 'vfriend'
    | 'vfollowing',
) => {
  switch (item) {
    case 'whisper':
      onSendWhisper(props.post.user);
      break;
    case 'follow':
    case 'unfollow':
      onHandleFollowAction(props.post);
      break;
    case 'delete':
      showDelModal.value = true;
      break;
    case 'lock':
    case 'unlock':
      showLockModal.value = true;
      break;
    case 'stick':
    case 'unstick':
      showStickModal.value = true;
      break;
    case 'highlight':
    case 'unhighlight':
      showHighlightModal.value = true;
      break;
    case 'vpublic':
      tempVisibility.value = 0;
      showVisibilityModal.value = true;
      break;
    case 'vprivate':
      tempVisibility.value = 1;
      showVisibilityModal.value = true;
      break;
    case 'vfriend':
      tempVisibility.value = 2;
      showVisibilityModal.value = true;
      break;
    case 'vfollowing':
      tempVisibility.value = 3;
      showVisibilityModal.value = true;
      break;
    default:
      break;
  }
};
const execDelAction = () => {
  deletePost({
    id: post.value.id,
  })
    .then((_res) => {
      window.$message.success('删除成功');
      router.replace('/');

      setTimeout(() => {
        storeMain.doRefresh();
      }, 50);
    })
    .catch((_err) => {
      loading.value = false;
    });
};
const execLockAction = () => {
  lockPost({
    id: post.value.id,
  })
    .then((res) => {
      emit('reload', post.value.id);
      if (res.lock_status === 1) {
        window.$message.success('锁定成功');
      } else {
        window.$message.success('解锁成功');
      }
    })
    .catch((_err) => {
      loading.value = false;
    });
};
const execStickAction = () => {
  stickPost({
    id: post.value.id,
  })
    .then((res) => {
      emit('reload', post.value.id);
      if (res.top_status === 1) {
        window.$message.success('置顶成功');
      } else {
        window.$message.success('取消置顶成功');
      }
    })
    .catch((_err) => {
      loading.value = false;
    });
};
const execHighlightAction = () => {
  highlightPost({
    id: post.value.id,
  })
    .then((res) => {
      post.value = {
        ...post.value,
        is_essence: res.highlight_status,
      };
      if (res.highlight_status === 1) {
        window.$message.success('设为亮点成功');
      } else {
        window.$message.success('取消亮点成功');
      }
    })
    .catch((_err) => {
      loading.value = false;
    });
};
const execVisibilityAction = () => {
  visibilityPost({
    id: post.value.id,
    visibility: tempVisibility.value,
  })
    .then((_res) => {
      emit('reload', post.value.id);
      window.$message.success('修改可见性成功');
    })
    .catch((_err) => {
      loading.value = false;
    });
};
const handlePostStar = () => {
  postStar({
    id: post.value.id,
  })
    .then((res) => {
      hasStarred.value = res.status;
      if (res.status) {
        post.value = {
          ...post.value,
          upvote_count: post.value.upvote_count + 1,
        };
      } else {
        post.value = {
          ...post.value,
          upvote_count: post.value.upvote_count - 1,
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
      hasCollected.value = res.status;
      if (res.status) {
        post.value = {
          ...post.value,
          collection_count: post.value.collection_count + 1,
        };
      } else {
        post.value = {
          ...post.value,
          collection_count: post.value.collection_count - 1,
        };
      }
    })
    .catch((err) => {
      console.log(err);
    });
};
const handlePostShare = () => {
  copy(
    `${window.location.origin}/#/post?id=${post.value.id}&share=copy_link&t=${new Date().getTime()}`,
  );
  window.$message.success('链接已复制到剪贴板');
};

onMounted(() => {
  if (userInfo.value.id > 0) {
    getPostStar({
      id: post.value.id,
    })
      .then((res) => {
        hasStarred.value = res.status;
      })
      .catch((err) => {
        console.log(err);
      });

    getPostCollection({
      id: post.value.id,
    })
      .then((res) => {
        hasCollected.value = res.status;
      })
      .catch((err) => {
        console.log(err);
      });
  }
});
</script>

<style lang="less" scoped>
.detail-item {
    width: 100%;
    padding: 16px;
    box-sizing: border-box;
    border-bottom: 1px solid var(--border-color, #eff3f4);
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
    font-size: 15px;
    line-height: 20px;
}

.nickname {
    font-weight: 700;
    cursor: pointer;

    &:hover {
        text-decoration: underline;
    }
}

.username {
    color: #888;
    font-size: 15px;
}

.badge-tag {
    transform: scale(0.75);
    transform-origin: left center;
}

.options {
    flex-shrink: 0;
    color: #888;
}

.post-content {
    margin-top: 8px;
}

.post-text {
    font-size: 17px;
    text-align: justify;
    overflow: hidden;
    white-space: pre-wrap;
    word-break: break-all;
    line-height: 1.6;
}

.post-media {
    margin-top: 12px;
}

.timestamp {
    opacity: 0.6;
    font-size: 13px;
    margin-top: 12px;
    padding-bottom: 8px;
    border-bottom: 1px solid var(--border-color, #eff3f4);
}

.dark {
    .detail-item {
        border-bottom-color: #2f3336;
    }

    .username {
        color: #999;
    }

    .timestamp {
        border-bottom-color: #2f3336;
    }
}
</style>