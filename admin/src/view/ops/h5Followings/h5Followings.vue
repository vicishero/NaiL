<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="关注者ID">
          <el-input v-model="searchInfo.userId" placeholder="用户ID" style="width:140px" />
        </el-form-item>
        <el-form-item label="被关注者ID">
          <el-input v-model="searchInfo.followId" placeholder="用户ID" style="width:140px" />
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
        <el-table-column align="left" label="关注者ID" min-width="120" prop="userId" />
        <el-table-column align="left" label="被关注者ID" min-width="120" prop="followId" />
        <el-table-column align="left" label="关注时间" min-width="170">
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
import { getH5FollowingList, deleteH5Following } from '@/api/h5Admin'
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate } from '@/utils/format'

defineOptions({ name: 'h5Followings' })

const page = ref(1), total = ref(0), pageSize = ref(10), tableData = ref([])
const searchInfo = reactive({ userId: '', followId: '' })

const getTableData = async () => {
  const params = { page: page.value, pageSize: pageSize.value }
  if (searchInfo.userId) params.userId = Number(searchInfo.userId)
  if (searchInfo.followId) params.followId = Number(searchInfo.followId)
  const res = await getH5FollowingList(params)
  if (res.code === 0) { tableData.value = res.data.list; total.value = res.data.total }
}
getTableData()

const onSearch = () => { page.value = 1; getTableData() }
const onReset = () => { searchInfo.userId = ''; searchInfo.followId = ''; page.value = 1; getTableData() }
const handleSizeChange = (v) => { pageSize.value = v; getTableData() }
const handleCurrentChange = (v) => { page.value = v; getTableData() }

const deleteFunc = async (row) => {
  ElMessageBox.confirm('确定要删除该关注关系吗?', '提示', { confirmButtonText: '确定', cancelButtonText: '取消', type: 'warning' }).then(async () => {
    const res = await deleteH5Following({ ID: row.ID })
    if (res.code === 0) { ElMessage.success('删除成功'); if (tableData.value.length === 1 && page.value > 1) page.value--; getTableData() }
  })
}
</script>
