<template>
    <div v-if="slides.length > 0" class="banner-carousel">
        <div
            class="carousel-track"
            :style="{ transform: `translateX(-${currentIndex * 100}%)` }"
        >
            <div
                v-for="(slide, i) in slides"
                :key="i"
                class="carousel-slide"
                @click="goPost(slide.id)"
            >
                <div
                    class="slide-bg"
                    :style="{ backgroundImage: `url(${slide.cover})` }"
                />
                <div class="slide-overlay" />
                <div class="slide-content">
                    <span class="slide-tag" v-if="slide.tag">{{ slide.tag }}</span>
                    <h3 class="slide-title">{{ slide.title }}</h3>
                    <p class="slide-meta">{{ slide.user.nickname }} · {{ slide.upvote_count }} 赞</p>
                </div>
            </div>
        </div>
        <!-- 指示器 -->
        <div class="carousel-dots">
            <span
                v-for="(_, i) in slides"
                :key="i"
                class="dot"
                :class="{ active: i === currentIndex }"
                @click.stop="currentIndex = i"
            />
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, watch } from 'vue';
import { useRouter } from 'vue-router';
import { getPosts } from '@/api/post';

const router = useRouter();
const slides = ref<any[]>([]);
const currentIndex = ref(0);
let timer: ReturnType<typeof setInterval> | null = null;

function goPost(id: string) {
    router.push({ name: 'post', query: { id } });
}

function startAutoPlay() {
    stopAutoPlay();
    if (slides.value.length <= 1) return;
    timer = setInterval(() => {
        currentIndex.value = (currentIndex.value + 1) % slides.value.length;
    }, 4000);
}

function stopAutoPlay() {
    if (timer) {
        clearInterval(timer);
        timer = null;
    }
}

onMounted(async () => {
    try {
        const res = await getPosts({ style: 'hots', page: 1, page_size: 5 });
        const list = (res as any)?.list || [];
        slides.value = list
            .filter((p: any) => p.imgs?.length > 0 || p.texts?.length > 0)
            .slice(0, 5)
            .map((p: any) => {
                // 提取封面图
                const cover = p.imgs?.[0]?.url || p.imgs?.[0]?.content || '';
                // 提取文本摘要
                const title = p.texts?.[0]?.content?.replace(/<[^>]+>/g, '').slice(0, 40) || '查看详情';
                // 提取话题标签
                const tags = typeof p.tags === 'object' ? Object.keys(p.tags || {}) : [];
                return {
                    id: p.id,
                    cover,
                    title,
                    tag: tags.length > 0 ? `#${tags[0]}` : '',
                    user: p.user,
                    upvote_count: p.upvote_count,
                };
            });
        startAutoPlay();
    } catch (err) {
        console.log('Failed to load banner:', err);
    }
});

onUnmounted(() => {
    stopAutoPlay();
});
</script>

<style lang="less" scoped>
.banner-carousel {
    position: relative;
    width: 100%;
    aspect-ratio: 16 / 9;
    max-height: 240px;
    overflow: hidden;
    border-radius: 0;
}

.carousel-track {
    display: flex;
    height: 100%;
    transition: transform 0.5s ease-in-out;
}

.carousel-slide {
    min-width: 100%;
    height: 100%;
    position: relative;
    cursor: pointer;
    overflow: hidden;
}

.slide-bg {
    position: absolute;
    inset: 0;
    background-size: cover;
    background-position: center;
    background-color: #e0e0e0;
}

.slide-overlay {
    position: absolute;
    inset: 0;
    background: linear-gradient(to top, rgba(0,0,0,0.6) 0%, transparent 60%);
}

.slide-content {
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    padding: 16px;
    color: #fff;
}

.slide-tag {
    display: inline-block;
    background: rgba(24, 160, 88, 0.85);
    padding: 2px 10px;
    border-radius: 4px;
    font-size: 12px;
    margin-bottom: 6px;
}

.slide-title {
    font-size: 16px;
    font-weight: 700;
    margin: 0 0 4px;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
    text-shadow: 0 1px 3px rgba(0,0,0,0.5);
}

.slide-meta {
    font-size: 12px;
    opacity: 0.8;
    margin: 0;
}

.carousel-dots {
    position: absolute;
    bottom: 8px;
    right: 12px;
    display: flex;
    gap: 6px;
}

.dot {
    width: 6px;
    height: 6px;
    border-radius: 50%;
    background: rgba(255,255,255,0.5);
    cursor: pointer;
    transition: all 0.3s;

    &.active {
        background: #fff;
        width: 18px;
        border-radius: 3px;
    }
}
</style>
