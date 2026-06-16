<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="用户ID">
          <el-input v-model="searchInfo.userId" placeholder="用户ID" style="width:120px" />
        </el-form-item>
        <el-form-item label="可见性">
          <el-select v-model="searchInfo.visibility" placeholder="全部" clearable style="width:100px">
            <el-option label="公开" :value="90" />
            <el-option label="私密" :value="0" />
            <el-option label="好友可见" :value="50" />
            <el-option label="关注可见" :value="60" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSearch">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button icon="connection" @click="onSyncIndex" :loading="syncing">同步索引</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table :data="tableData" row-key="ID" :default-sort="{ prop: 'ID', order: 'descending' }">
        <el-table-column align="left" label="ID" width="170" prop="ID" />
        <el-table-column align="left" label="用户" min-width="120">
          <template #default="scope">{{ scope.row.user?.nickname || `ID:${scope.row.userId}` }}</template>
        </el-table-column>
        <el-table-column align="left" label="内容" min-width="200" show-overflow-tooltip>
          <template #default="scope">
            <span v-if="scope.row.contents?.length">
              {{ scope.row.contents.filter(c => c.type === 2).map(c => c.content).join(' ').slice(0, 100) }}
            </span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="点赞" min-width="70" prop="upvoteCount" />
        <el-table-column align="left" label="评论" min-width="70" prop="commentCount" />
        <el-table-column align="left" label="收藏" min-width="70" prop="collectionCount" />
        <el-table-column align="left" label="可见性" min-width="80">
          <template #default="scope">{{ visibilityLabel(scope.row.visibility) }}</template>
        </el-table-column>
        <el-table-column align="left" label="置顶" min-width="70">
          <template #default="scope"><el-tag :type="scope.row.isTop ? 'warning' : 'info'" size="small">{{ scope.row.isTop ? '是' : '否' }}</el-tag></template>
        </el-table-column>
        <el-table-column align="left" label="锁定" min-width="70">
          <template #default="scope"><el-tag :type="scope.row.isLock ? 'danger' : 'info'" size="small">{{ scope.row.isLock ? '是' : '否' }}</el-tag></template>
        </el-table-column>
        <el-table-column align="left" label="发布时间" min-width="170" prop="CreatedAt">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column label="操作" min-width="180" fixed="right">
          <template #default="scope">
            <el-button type="primary" link icon="view" @click="viewPost(scope.row)">查看</el-button>
            <el-button type="primary" link icon="delete" @click="deletePostFunc(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]"
          :total="total" layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange" @size-change="handleSizeChange"
        />
      </div>
    </div>

    <el-drawer v-model="drawerVisible" :before-close="closeDrawer" :show-close="false" size="500px">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">贴文详情</span>
          <div>
            <el-button @click="closeDrawer">取消</el-button>
            <el-button type="primary" @click="savePost">确定</el-button>
          </div>
        </div>
      </template>
      <el-form :model="form" label-width="80px">
        <el-form-item label="可见性">
          <el-select v-model="form.visibility" style="width:100%">
            <el-option label="公开" :value="90" />
            <el-option label="私密" :value="0" />
            <el-option label="好友可见" :value="50" />
            <el-option label="关注可见" :value="60" />
          </el-select>
        </el-form-item>
        <el-form-item label="标记">
          <div style="display:flex;gap:24px">
            <span><el-switch v-model="form.isTop" /> 置顶</span>
            <span><el-switch v-model="form.isEssence" /> 精选</span>
            <span><el-switch v-model="form.isLock" /> 锁定</span>
          </div>
        </el-form-item>
        <el-form-item v-if="currentPost?.contents?.length" label="内容预览">
          <div style="max-height:400px;overflow-y:auto;width:100%">
            <div v-for="(c, i) in currentPost.contents" :key="i" style="margin-bottom:12px">
              <!-- 标题 -->
              <div v-if="c.type === 1" style="font-size:20px;font-weight:700;line-height:1.3;margin-bottom:4px">{{ c.content }}</div>
              <!-- 文本 -->
              <div v-else-if="c.type === 2" style="font-size:15px;line-height:1.6;white-space:pre-wrap;word-break:break-word;color:#0f1419">{{ c.content }}</div>
              <!-- 图片 -->
              <el-image v-else-if="c.type === 3" :src="c.content" style="max-width:100%;border-radius:12px" fit="contain" :preview-src-list="[c.content]" />
              <!-- 视频/音频/链接/附件 -->
              <div v-else style="padding:10px 14px;background:#f7f9f9;border-radius:10px;border:1px solid #eff3f4">
                <span style="font-size:13px;color:#536471">{{ ['','标题','文本','图片','视频','音频','链接','附件'][c.type] || `类型${c.type}` }}</span>
                <div style="font-size:14px;word-break:break-all;color:#1d9bf0;margin-top:2px">{{ c.content }}</div>
              </div>
            </div>
          </div>
        </el-form-item>
        <!-- 评论列表 -->
        <el-divider />
        <div style="font-size:15px;font-weight:600;margin-bottom:12px">评论 ({{ comments.length }})</div>
        <div v-loading="commentsLoading" style="max-height:300px;overflow-y:auto">
          <div v-if="comments.length === 0 && !commentsLoading" style="color:#909399;text-align:center;padding:20px">暂无评论</div>
          <div v-for="c in comments" :key="c.ID" style="padding:10px 0;border-bottom:1px solid #eff3f4">
            <div style="display:flex;justify-content:space-between;align-items:center;margin-bottom:4px">
              <span style="font-weight:600;font-size:13px">{{ c.nickname || `用户${c.userId}` }}</span>
              <span style="font-size:12px;color:#909399">{{ c.CreatedAt }}</span>
            </div>
            <div style="font-size:14px;line-height:1.5;white-space:pre-wrap;word-break:break-word;color:#0f1419">{{ c.content }}</div>
          </div>
        </div>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import { getH5PostList, getH5Post, updateH5Post, deleteH5Post, syncH5Index, getH5CommentList, deleteH5Comment } from '@/api/h5Admin'
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate } from '@/utils/format'

defineOptions({ name: 'h5Posts' })

const visibilityLabel = (v) => {
  const map = { 90: '公开', 0: '私密', 50: '好友可见', 60: '关注可见' }
  return map[v] ?? v
}

const page = ref(1), total = ref(0), pageSize = ref(10), tableData = ref([])
const searchInfo = reactive({ userId: '', visibility: undefined })
const drawerVisible = ref(false), currentPost = ref(null), syncing = ref(false)
const form = ref({ ID: 0, visibility: 90, isTop: false, isEssence: false, isLock: false })
const comments = ref([]), commentsLoading = ref(false)

const getTableData = async () => {
  const params = { page: page.value, pageSize: pageSize.value }
  if (searchInfo.userId) params.userId = Number(searchInfo.userId)
  if (searchInfo.visibility !== undefined && searchInfo.visibility !== '') params.visibility = searchInfo.visibility
  try {
    const res = await getH5PostList(params)
    console.log('h5Posts getTableData response:', JSON.stringify(res))
    if (res && res.code === 0) {
      tableData.value = res.data?.list || []
      total.value = res.data?.total || 0
      console.log('h5Posts loaded:', total.value, 'items')
    } else {
      console.warn('h5Posts unexpected response:', res)
      tableData.value = []
      total.value = 0
    }
  } catch (err) {
    console.error('h5Posts getTableData error:', err)
    tableData.value = []
    total.value = 0
  }
}
getTableData()

const onSearch = () => { page.value = 1; getTableData() }
const onReset = () => { searchInfo.userId = ''; searchInfo.visibility = undefined; page.value = 1; getTableData() }
const onSyncIndex = async () => { syncing.value = true; try { await syncH5Index(); ElMessage.success('索引同步已启动'); } catch { ElMessage.error('同步失败'); } finally { syncing.value = false } }
const handleSizeChange = (v) => { pageSize.value = v; getTableData() }
const handleCurrentChange = (v) => { page.value = v; getTableData() }

const viewPost = async (row) => {
  const res = await getH5Post({ ID: row.ID })
  if (res.code === 0) {
    currentPost.value = res.data
    form.value = { ID: res.data.ID, visibility: res.data.visibility, isTop: res.data.isTop, isEssence: res.data.isEssence, isLock: res.data.isLock }
    drawerVisible.value = true
    loadComments(res.data.ID)
  }
}
const loadComments = async (postId) => {
  commentsLoading.value = true
  try {
    const res = await getH5CommentList({ postId, page: 1, pageSize: 50 })
    if (res.code === 0) { comments.value = res.data?.list || [] }
  } catch { comments.value = [] }
  finally { commentsLoading.value = false }
}
const closeDrawer = () => { drawerVisible.value = false; currentPost.value = null; comments.value = [] }
const savePost = async () => {
  const res = await updateH5Post(form.value)
  if (res.code === 0) { ElMessage.success('更新成功'); drawerVisible.value = false; getTableData() }
}
const deletePostFunc = async (row) => {
  ElMessageBox.confirm('确定要删除该贴文吗?', '提示', { confirmButtonText: '确定', cancelButtonText: '取消', type: 'warning' }).then(async () => {
    const res = await deleteH5Post({ ID: row.ID })
    if (res.code === 0) { ElMessage.success('删除成功'); if (tableData.value.length === 1 && page.value > 1) page.value--; getTableData() }
  })
}
</script>
