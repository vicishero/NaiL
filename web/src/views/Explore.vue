<template>
    <div class="explore-page">
        <!-- 搜索框：点击跳转搜索子页面 -->
        <div class="search-bar" @click="goSearch">
            <n-input
                type="text"
                placeholder="搜索泡泡内容..."
                round
                readonly
            >
                <template #prefix>
                    <n-icon :component="SearchOutline" />
                </template>
            </n-input>
        </div>

        <!-- Trending 推荐用户 -->
        <div v-if="trendingUsers.length > 0" class="section">
            <div class="section-header">
                <span class="section-title">Trending</span>
                <span class="section-more" @click="goSearch">查看全部</span>
            </div>
            <div class="trending-list">
                <div
                    v-for="user in trendingUsers"
                    :key="user.id"
                    class="trending-user"
                    @click="goUser(user.username)"
                >
                    <n-avatar
                        round
                        :size="56"
                        :src="user.avatar"
                        :fallback-src="defaultAvatar"
                    />
                    <span class="trending-nickname">{{ user.nickname }}</span>
                </div>
            </div>
        </div>

        <!-- 分类用户列表 -->
        <div v-for="cat in categories" :key="cat.id" class="section">
            <div class="section-header">
                <span class="section-title">{{ cat.name }}</span>
                <span class="section-more">查看全部</span>
            </div>
            <div class="category-grid">
                <div
                    v-for="user in cat.users"
                    :key="user.id"
                    class="category-user"
                    @click="goUser(user.username)"
                >
                    <div class="category-avatar">
                        <img v-if="user.cover_image" :src="user.cover_image" class="cover-img" />
                        <n-avatar v-else :size="80" :src="user.avatar" :fallback-src="defaultAvatar" shape="square" />
                    </div>
                    <span class="category-nickname">{{ user.nickname }}</span>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { SearchOutline } from '@vicons/ionicons5';
import { getIndexTrends } from '@/api/post';

const router = useRouter();
const trendingUsers = ref<any[]>([]);
const defaultAvatar = '';

interface Category {
    id: number;
    name: string;
    users: Array<{
        id: number;
        nickname: string;
        avatar: string;
        username: string;
    }>;
}

// KOL分类数据
const categories = ref<Category[]>([]);

function goSearch() {
    router.push({ name: 'search' });
}

function goUser(username: string) {
    router.push({ name: 'user', query: { s: username } });
}

onMounted(async () => {
    // 加载 Trends 用户
    try {
        const res = await getIndexTrends({ page: 1, page_size: 5 });
        const data = (res as any)?.data || res;
        trendingUsers.value = data?.list || (res as any)?.list || [];
    } catch (err) {
        console.log('Failed to load trends:', err);
    }

    // 加载KOL分类数据
    try {
        const apiRes = await fetch(import.meta.env.VITE_HOST + '/v1/explore/kolCategories');
        const json = await apiRes.json();
        if (json.code === 0 && json.data?.categories) {
            categories.value = json.data.categories.map((c: any) => ({
                id: Number(c.id),
                name: c.name,
                users: (c.users || []).map((u: any) => ({
                    id: Number(u.id),
                    nickname: u.nickname,
                    avatar: u.avatar || '',
                    cover_image: u.cover_image || '',
                    username: u.username,
                })),
            }));
        } else {
            categories.value = [];
        }
    } catch {
        categories.value = [];
    }
});
</script>

<style lang="less" scoped>
.explore-page {
    padding-bottom: 16px;
}

.search-bar {
    padding: 12px 16px;
    cursor: pointer;
}

.section {
    padding: 0 16px 16px;
}

.section-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 12px 0 8px;
}

.section-title {
    font-size: 18px;
    font-weight: 700;
}

.section-more {
    font-size: 13px;
    color: #18a058;
    cursor: pointer;

    &:hover {
        text-decoration: underline;
    }
}

.trending-list {
    display: flex;
    gap: 16px;
    overflow-x: auto;
    padding: 4px 0;

    &::-webkit-scrollbar {
        display: none;
    }
}

.trending-user {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 6px;
    cursor: pointer;
    flex-shrink: 0;
    width: 72px;

    &:hover {
        opacity: 0.8;
    }
}

.trending-nickname {
    font-size: 12px;
    color: #666;
    text-align: center;
    max-width: 72px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
}

.category-grid {
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    gap: 12px;
}

.category-user {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 6px;
    cursor: pointer;

    &:hover {
        opacity: 0.85;
    }
}

.category-avatar {
    width: 100%;
    aspect-ratio: 3/4;
    overflow: hidden;
    border-radius: 8px;
    background: #f0f0f0;

	.cover-img { width: 100%; height: 100%; object-fit: cover; }
    :deep(.n-avatar) {
        width: 100%;
        height: 100%;
    }
}

.category-nickname {
    font-size: 13px;
    color: #333;
    text-align: center;
}

.dark {
    .category-avatar {
        background: #222;
    }

    .trending-nickname {
        color: #999;
    }

    .category-nickname {
        color: #ddd;
    }
}
</style>
