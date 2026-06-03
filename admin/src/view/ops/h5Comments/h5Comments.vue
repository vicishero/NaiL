<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="关键词">
          <el-input v-model="searchInfo.keyword" placeholder="搜索内容" />
        </el-form-item>
        <el-form-item label="贴文ID">
          <el-input v-model="searchInfo.postId" placeholder="贴文ID" style="width:120px" />
        </el-form-item>
        <el-form-item label="用户ID">
          <el-input v-model="searchInfo.userId" placeholder="用户ID" style="width:120px" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSearch">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table :data="tableData" row-key="ID">
        <el-table-column align="left" label="ID" min-width="60" prop="ID" />
        <el-table-column align="left" label="用户" min-width="120">
          <template #default="scope">{{ scope.row.user?.nickname || `ID:${scope.row.userId}` }}</template>
        </el-table-column>
        <el-table-column align="left" label="贴文ID" min-width="80" prop="postId" />
        <el-table-column align="left" label="内容" min-width="250" prop="content" show-overflow-tooltip />
        <el-table-column align="left" label="点赞" min-width="70" prop="thumbsUpCount" />
        <el-table-column align="left" label="回复数" min-width="70" prop="replyCount" />
        <el-table-column align="left" label="IP" min-width="130" prop="ip" />
        <el-table-column align="left" label="时间" min-width="170">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column label="操作" min-width="100" fixed="right">
          <template #default="scope">
            <el-button type="primary" link icon="delete" @click="deleteFunc(scope.row)">删除</el-button>
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
  </div>
</template>

<script setup>
import { getH5CommentList, deleteH5Comment } from '@/api/h5Admin'
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate } from '@/utils/format'

defineOptions({ name: 'h5Comments' })

const page = ref(1), total = ref(0), pageSize = ref(10), tableData = ref([])
const searchInfo = reactive({ keyword: '', postId: '', userId: '' })

const getTableData = async () => {
  const params = { page: page.value, pageSize: pageSize.value }
  if (searchInfo.keyword) params.keyword = searchInfo.keyword
  if (searchInfo.postId) params.postId = Number(searchInfo.postId)
  if (searchInfo.userId) params.userId = Number(searchInfo.userId)
  const res = await getH5CommentList(params)
  if (res.code === 0) { tableData.value = res.data.list; total.value = res.data.total }
}
getTableData()

const onSearch = () => { page.value = 1; getTableData() }
const onReset = () => { searchInfo.keyword = ''; searchInfo.postId = ''; searchInfo.userId = ''; page.value = 1; getTableData() }
const handleSizeChange = (v) => { pageSize.value = v; getTableData() }
const handleCurrentChange = (v) => { page.value = v; getTableData() }

const deleteFunc = async (row) => {
  ElMessageBox.confirm('确定要删除该评论吗?', '提示', { confirmButtonText: '确定', cancelButtonText: '取消', type: 'warning' }).then(async () => {
    const res = await deleteH5Comment({ ID: row.ID })
    if (res.code === 0) { ElMessage.success('删除成功'); if (tableData.value.length === 1 && page.value > 1) page.value--; getTableData() }
  })
}
</script>
