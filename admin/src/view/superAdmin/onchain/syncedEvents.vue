<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" @keyup.enter="onSearch">
        <el-form-item label="合约ID">
          <el-input-number v-model="searchInfo.contractConfigId" :min="1" placeholder="合约ID" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchInfo.status" placeholder="全部" clearable style="width: 120px">
            <el-option label="未处理" :value="0" />
            <el-option label="已处理" :value="1" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSearch">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table style="width: 100%" :data="tableData" row-key="ID">
        <el-table-column align="left" label="事件ID" prop="ID" width="80" />
        <el-table-column align="left" label="合约ID" prop="contractConfigId" width="80" />
        <el-table-column align="left" label="事件定义ID" prop="eventId" width="100" />
        <el-table-column align="left" label="区块号" prop="blockNumber" width="120" />
        <el-table-column align="left" label="Tx Hash" prop="txHash" width="200" show-overflow-tooltip />
        <el-table-column align="left" label="Log Index" prop="logIndex" width="80" />
        <el-table-column align="left" label="Topics" prop="topics" min-width="200" show-overflow-tooltip />
        <el-table-column align="left" label="同步时间" prop="syncedAt" width="160" />
        <el-table-column align="left" label="状态" width="90">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'info'">
              {{ scope.row.status === 1 ? '已处理' : '未处理' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" width="100">
          <template #default="scope">
            <el-button
              v-if="scope.row.status === 0"
              type="primary"
              link
              @click="markProcessed(scope.row)"
            >标记已处理</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage } from 'element-plus'
import { getBscSyncedEventList, updateBscSyncedEventStatus } from '@/api/bsc'

defineOptions({
  name: 'SyncedEvents'
})

const searchInfo = reactive({
  contractConfigId: null,
  status: null
})

const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)

const fetchData = async () => {
  const params = { page: page.value, pageSize: pageSize.value }
  if (searchInfo.contractConfigId) {
    params.contractConfigId = searchInfo.contractConfigId
  }
  if (searchInfo.status !== null && searchInfo.status !== undefined && searchInfo.status !== '') {
    params.status = searchInfo.status
  }
  const res = await getBscSyncedEventList(params)
  if (res.code === 0) {
    tableData.value = res.data.list || []
    total.value = res.data.total || 0
  }
}

const onSearch = () => {
  page.value = 1
  fetchData()
}

const onReset = () => {
  searchInfo.contractConfigId = null
  searchInfo.status = null
  page.value = 1
  fetchData()
}

const markProcessed = async (row) => {
  const res = await updateBscSyncedEventStatus({ ID: row.ID, status: 1 })
  if (res.code === 0) {
    ElMessage.success('已标记为已处理')
    fetchData()
  }
}

const handleCurrentChange = (val) => {
  page.value = val
  fetchData()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  page.value = 1
  fetchData()
}

onMounted(() => {
  fetchData()
})
</script>
