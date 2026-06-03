<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" @keyup.enter="onSearch">
        <el-form-item label="链ID">
          <el-input-number v-model="searchInfo.chainId" :min="1" placeholder="链ID" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSearch">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增链配置</el-button>
      </div>
      <el-table style="width: 100%" :data="tableData" row-key="ID">
        <el-table-column align="left" label="链ID" prop="chainId" width="80" />
        <el-table-column align="left" label="RPC节点" prop="rpcUrl" min-width="250" show-overflow-tooltip />
        <el-table-column align="left" label="起始区块" prop="startBlock" width="120" />
        <el-table-column align="left" label="最后同步区块" prop="lastSyncedBlock" width="140" />
        <el-table-column align="left" label="区块步长" prop="blockStep" width="100" />
        <el-table-column align="left" label="同步间隔(秒)" prop="syncIntervalSeconds" width="120" />
        <el-table-column align="left" label="启用" prop="enabled" width="70">
          <template #default="scope">
            <el-switch v-model="scope.row.enabled" @change="toggleEnabled(scope.row)" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="160">
          <template #default="scope">
            <el-button type="primary" link icon="edit" @click="editInfo(scope.row)">编辑</el-button>
            <el-button type="primary" link icon="delete" @click="deleteInfo(scope.row)">删除</el-button>
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

    <el-dialog v-model="dialogVisible" :title="dialogTitle" width="550px" :before-close="closeDialog">
      <el-form ref="formRef" :model="form" :rules="rules" label-width="130px">
        <el-form-item label="链ID" prop="chainId">
          <el-input-number v-model="form.chainId" :min="1" style="width: 100%" />
        </el-form-item>
        <el-form-item label="RPC节点地址" prop="rpcUrl">
          <el-input v-model="form.rpcUrl" placeholder="https://bsc-dataseed.binance.org" />
        </el-form-item>
        <el-form-item label="起始区块号">
          <el-input-number v-model="form.startBlock" :min="0" style="width: 100%" />
        </el-form-item>
        <el-form-item label="区块步长">
          <el-input-number v-model="form.blockStep" :min="1" :max="100000" style="width: 100%" />
        </el-form-item>
        <el-form-item label="同步间隔(秒)">
          <el-input-number v-model="form.syncIntervalSeconds" :min="1" :max="3600" style="width: 100%" />
        </el-form-item>
        <el-form-item label="启用同步">
          <el-switch v-model="form.enabled" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="closeDialog">取消</el-button>
        <el-button type="primary" @click="submitForm">确认</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  createBscSyncInfo,
  deleteBscSyncInfo,
  updateBscSyncInfo,
  getBscSyncInfoList
} from '@/api/bsc'

defineOptions({
  name: 'SyncInfo'
})

const searchInfo = reactive({
  chainId: null
})

const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)

const dialogVisible = ref(false)
const dialogTitle = ref('新增链配置')
const formRef = ref(null)
const isEdit = ref(false)

const defaultForm = () => ({
  ID: 0,
  chainId: 56,
  rpcUrl: '',
  startBlock: 0,
  lastSyncedBlock: 0,
  blockStep: 5000,
  syncIntervalSeconds: 1,
  enabled: true
})

const form = reactive(defaultForm())

const rules = {
  chainId: [{ required: true, message: '请输入链ID', trigger: 'blur' }],
  rpcUrl: [{ required: true, message: '请输入RPC节点地址', trigger: 'blur' }]
}

const fetchData = async () => {
  const params = { page: page.value, pageSize: pageSize.value }
  if (searchInfo.chainId) {
    params.chainId = searchInfo.chainId
  }
  const res = await getBscSyncInfoList(params)
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
  searchInfo.chainId = null
  page.value = 1
  fetchData()
}

const openDialog = () => {
  isEdit.value = false
  dialogTitle.value = '新增链配置'
  Object.assign(form, defaultForm())
  dialogVisible.value = true
}

const editInfo = (row) => {
  isEdit.value = true
  dialogTitle.value = '编辑链配置'
  Object.assign(form, {
    ID: row.ID,
    chainId: row.chainId,
    rpcUrl: row.rpcUrl || '',
    startBlock: row.startBlock || 0,
    lastSyncedBlock: row.lastSyncedBlock || 0,
    blockStep: row.blockStep || 5000,
    syncIntervalSeconds: row.syncIntervalSeconds || 1,
    enabled: row.enabled
  })
  dialogVisible.value = true
}

const closeDialog = () => {
  dialogVisible.value = false
  formRef.value?.resetFields()
}

const submitForm = async () => {
  const valid = await formRef.value.validate().catch(() => false)
  if (!valid) return

  const data = { ...form }

  let res
  if (isEdit.value) {
    res = await updateBscSyncInfo(data)
  } else {
    res = await createBscSyncInfo(data)
  }
  if (res.code === 0) {
    ElMessage.success(isEdit.value ? '更新成功' : '创建成功')
    closeDialog()
    fetchData()
  }
}

const deleteInfo = (row) => {
  ElMessageBox.confirm('确认删除该链配置？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteBscSyncInfo({ ID: row.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      fetchData()
    }
  })
}

const toggleEnabled = async (row) => {
  const res = await updateBscSyncInfo({
    ID: row.ID,
    chainId: row.chainId,
    rpcUrl: row.rpcUrl || '',
    startBlock: row.startBlock || 0,
    lastSyncedBlock: row.lastSyncedBlock || 0,
    blockStep: row.blockStep || 5000,
    syncIntervalSeconds: row.syncIntervalSeconds || 1,
    enabled: row.enabled
  })
  if (res.code === 0) {
    ElMessage.success(row.enabled ? '同步已开启' : '同步已暂停')
  } else {
    row.enabled = !row.enabled
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
