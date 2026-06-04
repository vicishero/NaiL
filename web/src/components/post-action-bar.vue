<template>
    <div class="action-bar">
        <div class="action-item" @click.stop="$emit('comment')">
            <n-icon :size="18" :component="ChatboxOutline" />
            <span v-if="commentCount > 0" class="action-count">{{ commentCount }}</span>
        </div>
        <div class="action-item" @click.stop="$emit('repost')">
            <n-icon :size="18" :component="ShareSocialOutline" />
            <span v-if="repostCount > 0" class="action-count">{{ repostCount }}</span>
        </div>
        <div
            class="action-item"
            :class="{ liked: isLiked }"
            @click.stop="$emit('like')"
        >
            <n-icon :size="18" :component="isLiked ? HeartFilled : HeartOutline" />
            <span v-if="likeCount > 0" class="action-count" :class="{ liked: isLiked }">{{ likeCount }}</span>
        </div>
        <div
            class="action-item"
            :class="{ bookmarked: isBookmarked }"
            @click.stop="$emit('bookmark')"
        >
            <n-icon :size="18" :component="isBookmarked ? BookmarkFilled : BookmarkOutline" />
            <span v-if="bookmarkCount > 0" class="action-count" :class="{ bookmarked: isBookmarked }">{{ bookmarkCount }}</span>
        </div>
        <div class="action-item" @click.stop="$emit('share')">
            <n-icon :size="18" :component="ArrowUndoOutline" />
        </div>
    </div>
</template>

<script setup lang="ts">
import {
    HeartOutline,
    BookmarkOutline,
    ChatboxOutline,
    ShareSocialOutline,
    ArrowUndoOutline,
} from '@vicons/ionicons5';
import { Heart, Bookmark } from '@vicons/ionicons5';

// ionicons5 doesn't have filled variants for these, use the same with different color via CSS
const HeartFilled = Heart;
const BookmarkFilled = Bookmark;

withDefaults(defineProps<{
    commentCount?: number;
    repostCount?: number;
    likeCount?: number;
    bookmarkCount?: number;
    isLiked?: boolean;
    isBookmarked?: boolean;
}>(), {
    commentCount: 0,
    repostCount: 0,
    likeCount: 0,
    bookmarkCount: 0,
    isLiked: false,
    isBookmarked: false,
});

defineEmits<{
    comment: [];
    repost: [];
    like: [];
    bookmark: [];
    share: [];
}>();
</script>

<style lang="less" scoped>
.action-bar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    max-width: 400px;
    margin-top: 4px;
}

.action-item {
    display: flex;
    align-items: center;
    gap: 4px;
    padding: 6px 8px;
    border-radius: 999px;
    cursor: pointer;
    color: #888;
    transition: all 0.2s ease;

    &:hover {
        background: rgba(24, 160, 88, 0.08);
        color: #18a058;
    }

    &.liked {
        color: #e74c3c;

        &:hover {
            background: rgba(231, 76, 60, 0.08);
        }
    }

    &.bookmarked {
        color: #18a058;

        &:hover {
            background: rgba(24, 160, 88, 0.08);
        }
    }
}

.action-count {
    font-size: 13px;
}

.dark {
    .action-item {
        color: #999;

        &:hover {
            background: rgba(99, 226, 183, 0.12);
            color: #63e2b7;
        }

        &.liked {
            color: #e74c3c;
            &:hover {
                background: rgba(231, 76, 60, 0.12);
            }
        }

        &.bookmarked {
            color: #63e2b7;
            &:hover {
                background: rgba(99, 226, 183, 0.12);
            }
        }
    }
}
</style>
