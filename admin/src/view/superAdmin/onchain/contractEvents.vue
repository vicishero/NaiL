<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" @keyup.enter="onSearch">
        <el-form-item label="合约名称">
          <el-input v-model="searchInfo.name" placeholder="合约名称" />
        </el-form-item>
        <el-form-item label="合约地址">
          <el-input v-model="searchInfo.contractAddress" placeholder="合约地址" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSearch">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增合约</el-button>
      </div>
      <el-table
        style="width: 100%"
        :data="tableData"
        row-key="ID"
        @expand-change="onExpand"
      >
        <el-table-column type="expand">
          <template #default="scope">
            <el-table :data="scope.row.events || []" size="small" style="margin: 0 20px">
              <el-table-column align="left" label="事件ID" prop="ID" width="70" />
              <el-table-column align="left" label="事件名称" prop="eventName" />
              <el-table-column align="left" label="Topic Hash" prop="topicHash" show-overflow-tooltip />
              <el-table-column align="left" label="ABI签名" prop="abiSignature" show-overflow-tooltip />
              <el-table-column align="left" label="启用" prop="enabled" width="70">
                <template #default="s">
                  <el-tag :type="s.row.enabled ? 'success' : 'info'">{{ s.row.enabled ? '是' : '否' }}</el-tag>
                </template>
              </el-table-column>
            </el-table>
          </template>
        </el-table-column>
        <el-table-column align="left" label="合约名称" prop="name" width="150" />
        <el-table-column align="left" label="合约地址" prop="contractAddress" width="300" show-overflow-tooltip />
        <el-table-column align="left" label="链ID" prop="chainId" width="80" />
        <el-table-column align="left" label="启用" prop="enabled" width="70">
          <template #default="scope">
            <el-switch
              v-model="scope.row.enabled"
              @change="toggleEnabled(scope.row)"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="事件数" width="80">
          <template #default="scope">{{ (scope.row.events || []).length }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" min-width="160">
          <template #default="scope">
            <el-button type="primary" link icon="edit" @click="editContract(scope.row)">编辑</el-button>
            <el-button type="primary" link icon="delete" @click="deleteContract(scope.row)">删除</el-button>
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

    <el-dialog
      v-model="dialogVisible"
      :title="dialogTitle"
      width="700px"
      :before-close="closeDialog"
    >
      <el-form ref="formRef" :model="form" :rules="rules" label-width="110px">
        <el-form-item label="合约名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入合约名称" />
        </el-form-item>
        <el-form-item label="合约地址" prop="contractAddress">
          <el-input v-model="form.contractAddress" placeholder="0x..." />
        </el-form-item>
        <el-form-item label="链" prop="chainId">
          <el-select v-model="form.chainId" placeholder="选择链" style="width: 100%">
            <el-option
              v-for="item in chainList"
              :key="item.ID"
              :label="'链' + item.chainId + ' (' + item.rpcUrl + ')'"
              :value="item.chainId"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="启用同步">
          <el-switch v-model="form.enabled" />
        </el-form-item>

        <el-divider>事件配置</el-divider>
        <div v-for="(event, idx) in form.events" :key="idx" class="event-item">
          <el-row :gutter="10" align="middle" justify="space-between">
            <el-col v-if="isEdit" :span="4">
              <el-form-item label="事件ID">
                <span class="event-id-text">{{ event.ID }}</span>
              </el-form-item>
            </el-col>
            <el-col :span="isEdit ? 7 : 10">
              <el-form-item label="事件名" :prop="'events.' + idx + '.eventName'" :rules="[{required: true, message: '必填'}]">
                <el-input v-model="event.eventName" placeholder="Transfer" />
              </el-form-item>
            </el-col>
            <el-col :span="4">
              <el-form-item label="启用">
                <el-switch v-model="event.enabled" size="small" />
              </el-form-item>
            </el-col>
            <el-col :span="2" class="delete-col">
              <el-button type="danger" icon="delete" circle size="small" @click="removeEvent(idx)" />
            </el-col>
          </el-row>
          <el-form-item label="Topic Hash" :prop="'events.' + idx + '.topicHash'" :rules="[{required: true, message: '必填'}]">
            <el-input v-model="event.topicHash" placeholder="0x..." show-word-limit />
          </el-form-item>
          <el-form-item label="ABI签名">
            <el-input v-model="event.abiSignature" placeholder="event Transfer(address indexed from, address indexed to, uint256 value)" />
          </el-form-item>
        </div>
        <el-button type="primary" icon="plus" @click="addEvent">添加事件</el-button>
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
  createBscContractConfig,
  deleteBscContractConfig,
  updateBscContractConfig,
  getBscContractConfigList,
  getBscSyncInfoList
} from '@/api/bsc'

defineOptions({
  name: 'ContractEvents'
})

const searchInfo = reactive({
  name: '',
  contractAddress: ''
})

const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const chainList = ref([])

const dialogVisible = ref(false)
const dialogTitle = ref('新增合约')
const formRef = ref(null)
const isEdit = ref(false)

const defaultForm = () => ({
  ID: 0,
  name: '',
  contractAddress: '',
  chainId: 56,
  enabled: true,
  events: []
})

const form = reactive(defaultForm())

const rules = {
  name: [{ required: true, message: '请输入合约名称', trigger: 'blur' }],
  contractAddress: [{ required: true, message: '请输入合约地址', trigger: 'blur' }],
  chainId: [{ required: true, message: '请选择链', trigger: 'change' }]
}

const fetchChainList = async () => {
  const res = await getBscSyncInfoList({ page: 1, pageSize: 1000 })
  if (res.code === 0) {
    chainList.value = res.data.list || []
  }
}

const fetchData = async () => {
  const res = await getBscContractConfigList({
    page: page.value,
    pageSize: pageSize.value,
    name: searchInfo.name,
    contractAddress: searchInfo.contractAddress
  })
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
  searchInfo.name = ''
  searchInfo.contractAddress = ''
  page.value = 1
  fetchData()
}

const openDialog = () => {
  isEdit.value = false
  dialogTitle.value = '新增合约'
  Object.assign(form, defaultForm())
  fetchChainList()
  dialogVisible.value = true
}

const editContract = (row) => {
  isEdit.value = true
  dialogTitle.value = '编辑合约'
  Object.assign(form, {
    ID: row.ID,
    name: row.name,
    contractAddress: row.contractAddress,
    chainId: row.chainId,
    enabled: row.enabled,
    events: (row.events || []).map(e => ({ ...e }))
  })
  fetchChainList()
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
  data.events = data.events.map(e => ({
    ID: e.ID || 0,
    eventName: e.eventName,
    topicHash: e.topicHash,
    abiSignature: e.abiSignature || '',
    enabled: e.enabled !== undefined ? e.enabled : true
  }))

  let res
  if (isEdit.value) {
    res = await updateBscContractConfig(data)
  } else {
    res = await createBscContractConfig(data)
  }
  if (res.code === 0) {
    ElMessage.success(isEdit.value ? '更新成功' : '创建成功')
    closeDialog()
    fetchData()
  }
}

const deleteContract = (row) => {
  ElMessageBox.confirm('确认删除该合约配置？相关事件也会被删除。', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteBscContractConfig({ ID: row.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      fetchData()
    }
  })
}

const toggleEnabled = async (row) => {
  const res = await updateBscContractConfig({
    ID: row.ID,
    name: row.name,
    contractAddress: row.contractAddress,
    chainId: row.chainId,
    enabled: row.enabled,
    events: (row.events || []).map(e => ({
      ID: e.ID || 0,
      eventName: e.eventName,
      topicHash: e.topicHash,
      abiSignature: e.abiSignature || '',
      enabled: e.enabled
    }))
  })
  if (res.code === 0) {
    ElMessage.success(row.enabled ? '同步已开启' : '同步已暂停')
  } else {
    row.enabled = !row.enabled
    fetchData()
  }
}

const addEvent = () => {
  form.events.push({
    ID: 0,
    eventName: '',
    topicHash: '',
    abiSignature: '',
    enabled: true
  })
}

const removeEvent = (idx) => {
  form.events.splice(idx, 1)
}

const onExpand = (row, expandedRows) => {
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

<style scoped>
.event-item {
  margin-bottom: 10px;
  padding: 10px;
  border: 1px solid #ebeef5;
  border-radius: 4px;
  background: #fafafa;
}
.event-id-text {
  display: inline-block;
  line-height: 32px;
  color: #606266;
}
.delete-col {
  display: flex;
  justify-content: flex-end;
}
</style>
