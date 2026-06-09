<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true">
        <el-form-item>
          <el-button type="primary" icon="refresh" @click="onReset">刷新</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table :data="tableData" row-key="ID">
        <el-table-column align="left" label="ID" width="170" prop="ID" />
        <el-table-column align="left" label="话题" min-width="150" prop="tag">
          <template #default="scope">#{{ scope.row.tag }}</template>
        </el-table-column>
        <el-table-column align="left" label="引用次数" min-width="100" prop="quoteNum" />
        <el-table-column align="left" label="创建者ID" min-width="100" prop="userId" />
        <el-table-column align="left" label="创建时间" min-width="170">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column label="操作" min-width="150" fixed="right">
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

    <el-drawer v-model="drawerVisible" :before-close="closeDrawer" :show-close="false">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">编辑话题</span>
          <div>
            <el-button @click="closeDrawer">取消</el-button>
            <el-button type="primary" @click="saveTag">确定</el-button>
          </div>
        </div>
      </template>
      <el-form :model="form" label-width="100px">
        <el-form-item label="话题名">
          <el-input v-model="form.tag" />
        </el-form-item>
        <el-form-item label="引用次数">
          <el-input-number v-model="form.quoteNum" :min="0" />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import { getH5TagList, updateH5Tag, deleteH5Tag } from '@/api/h5Admin'
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate } from '@/utils/format'

defineOptions({ name: 'h5Tags' })

const page = ref(1), total = ref(0), pageSize = ref(10), tableData = ref([])
const drawerVisible = ref(false)
const form = ref({ ID: 0, tag: '', quoteNum: 0 })

const getTableData = async () => {
  const params = { page: page.value, pageSize: pageSize.value }
  try {
    const res = await getH5TagList(params)
    console.log('h5Tags getTableData response:', JSON.stringify(res))
    if (res && res.code === 0) {
      tableData.value = res.data?.list || []
      total.value = res.data?.total || 0
      console.log('h5Tags loaded:', total.value, 'items')
    } else {
      console.warn('h5Tags unexpected response:', res)
      tableData.value = []
      total.value = 0
    }
  } catch (err) {
    console.error('h5Tags getTableData error:', err)
    tableData.value = []
    total.value = 0
  }
}
getTableData()

const onReset = () => { page.value = 1; getTableData() }
const handleSizeChange = (v) => { pageSize.value = v; getTableData() }
const handleCurrentChange = (v) => { page.value = v; getTableData() }

const editTag = (row) => { form.value = { ID: row.ID, tag: row.tag, quoteNum: row.quoteNum }; drawerVisible.value = true }
const closeDrawer = () => { drawerVisible.value = false }
const saveTag = async () => {
  const res = await updateH5Tag(form.value)
  if (res.code === 0) { ElMessage.success('更新成功'); drawerVisible.value = false; getTableData() }
}
const deleteFunc = async (row) => {
  ElMessageBox.confirm('确定要删除该话题吗?', '提示', { confirmButtonText: '确定', cancelButtonText: '取消', type: 'warning' }).then(async () => {
    const res = await deleteH5Tag({ ID: row.ID })
    if (res.code === 0) { ElMessage.success('删除成功'); if (tableData.value.length === 1 && page.value > 1) page.value--; getTableData() }
  })
}
</script>
