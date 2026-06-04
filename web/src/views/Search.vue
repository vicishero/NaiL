<template>
    <div class="search-page">
        <!-- 搜索栏 -->
        <div class="search-bar">
            <div class="search-input-wrap">
                <n-input
                    v-model:value="keyword"
                    type="text"
                    placeholder="搜索泡泡内容..."
                    round
                    clearable
                    @keyup.enter="doSearch"
                >
                    <template #prefix>
                        <n-icon :component="SearchOutline" />
                    </template>
                </n-input>
            </div>
            <n-button
                type="primary"
                round
                size="small"
                :loading="searching"
                @click="doSearch"
            >
                搜索
            </n-button>
        </div>

        <!-- 搜索结果 -->
        <div v-if="searched" class="search-results">
            <div v-if="searching" class="empty-wrap">
                <n-spin :size="24" />
            </div>
            <div v-else-if="results.length === 0" class="empty-wrap">
                <n-empty description="未找到相关内容" />
            </div>
            <div v-else class="post-list">
                <div class="result-header">找到 {{ results.length }} 条结果</div>
                <post-item
                    v-for="p in results"
                    :key="p.id"
                    :post="p"
                    :is-owner="false"
                    :is-mobile="true"
                />
            </div>
        </div>

        <!-- 未搜索时的默认内容 -->
        <div v-else class="search-default">
            <!-- 热门话题 -->
            <div v-if="hotTags.length > 0" class="section">
                <div class="section-title">热门话题</div>
                <div class="tag-list">
                    <tag-item
                        v-for="tag in hotTags"
                        :key="tag.id"
                        :tag="tag"
                        :show-action="false"
                        :check-following="false"
                        :check-pin="false"
                    />
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { SearchOutline } from '@vicons/ionicons5';
import { getPosts, getTags } from '@/api/post';
import TagItem from '@/components/tag-item.vue';

const keyword = ref('');
const searching = ref(false);
const searched = ref(false);
const results = ref<any[]>([]);
const hotTags = ref<any[]>([]);

function doSearch() {
    const q = keyword.value.trim();
    if (!q) {
        window.$message?.warning('请输入搜索关键词');
        return;
    }
    searching.value = true;
    searched.value = true;
    getPosts({
        query: q,
        type: null,
        style: 'newest',
        page: 1,
        page_size: 50,
    })
        .then((rsp: any) => {
            results.value = rsp.list || [];
        })
        .catch((err: any) => {
            console.log(err);
            results.value = [];
        })
        .finally(() => {
            searching.value = false;
        });
}

onMounted(async () => {
    try {
        const res = await getTags({ type: 'hot', num: 10 });
        // 真实API返回结构是res.data.topics
        const rawList = (res as any)?.data?.topics || (res as any)?.topics || (res as any)?.list || (res as any)?.tags || [];
        hotTags.value = rawList.map((t: any) => ({
            ...t,
            id: t.id || t.tag,
            tag: t.tag || t.id,
            quote_num: t.quote_num || 0,
            is_following: t.is_following || 0,
            is_pin: t.is_pin || 0,
        }));
        console.log('Hot tags loaded:', hotTags.value); // 调试用
    } catch (err) {
        console.log('Failed to load hot tags:', err);
    }
});
</script>

<style lang="less" scoped>
.search-page {
    min-height: 60vh;
}

.search-bar {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 12px 16px;
    border-bottom: 1px solid var(--border-color, #eff3f4);
}

.search-input-wrap {
    flex: 1;
}

.search-results {
    .result-header {
        padding: 12px 16px;
        font-size: 14px;
        color: #888;
        border-bottom: 1px solid var(--border-color, #eff3f4);
    }
}

.search-default {
    padding: 16px;
}

.empty-wrap {
    display: flex;
    justify-content: center;
    padding: 40px 0;
}

.dark {
    .search-bar {
        border-bottom-color: #2f3336;
    }
}

.section {
    .section-title {
        font-size: 18px;
        font-weight: 700;
        padding: 12px 0 8px;
    }
}

.tag-list {
    display: flex;
    flex-wrap: wrap;
    gap: 8px;
}
</style>
