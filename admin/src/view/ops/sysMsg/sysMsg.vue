<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true">
        <el-form-item>
          <el-button type="primary" icon="plus" @click="addMsg">发送系统消息</el-button>
          <el-button icon="refresh" @click="getTableData">刷新</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table :data="tableData" row-key="ID">
        <el-table-column align="left" label="ID" width="170" prop="ID" />
        <el-table-column align="left" label="接收者ID" width="170" prop="receiverId" />
        <el-table-column align="left" label="摘要" min-width="150" prop="brief" show-overflow-tooltip />
        <el-table-column align="left" label="内容" min-width="200" prop="content" show-overflow-tooltip />
        <el-table-column align="left" label="已读" width="70">
          <template #default="scope">{{ scope.row.isRead ? '是' : '否' }}</template>
        </el-table-column>
        <el-table-column align="left" label="发送时间" min-width="170" prop="CreatedAt">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column label="操作" width="80" fixed="right">
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

    <el-dialog v-model="dialogVisible" title="发送系统消息" width="500px">
      <el-form :model="form" label-width="80px">
        <el-form-item label="接收者ID">
          <el-input v-model="form.receiverId" placeholder="用户ID，0表示全员" />
        </el-form-item>
        <el-form-item label="摘要">
          <el-input v-model="form.brief" />
        </el-form-item>
        <el-form-item label="内容">
          <el-input v-model="form.content" type="textarea" :rows="4" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="sendMsg">发送</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { getSysMsgList, createSysMsg, deleteSysMsg } from '@/api/h5Admin'
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate } from '@/utils/format'

defineOptions({ name: 'sysMsg' })

const page = ref(1), total = ref(0), pageSize = ref(10), tableData = ref([])
const dialogVisible = ref(false)
const form = reactive({ receiverId: '0', brief: '', content: '' })

const getTableData = async () => {
  const res = await getSysMsgList({ page: page.value, pageSize: pageSize.value })
  if (res.code === 0) { tableData.value = res.data?.list || []; total.value = res.data?.total || 0 }
}
getTableData()

const handleSizeChange = (v) => { pageSize.value = v; getTableData() }
const handleCurrentChange = (v) => { page.value = v; getTableData() }

const addMsg = () => { form.receiverId = '0'; form.brief = ''; form.content = ''; dialogVisible.value = true }
const sendMsg = async () => {
  // 发送字符串格式，跟用户ID格式保持一致，"0"表示全员
  const receiverIdVal = form.receiverId?.trim() || '0'
  const res = await createSysMsg({ receiverId: receiverIdVal, brief: form.brief, content: form.content })
  if (res.code === 0) { ElMessage.success('发送成功'); dialogVisible.value = false; getTableData() }
}
const deleteFunc = async (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', { confirmButtonText: '确定', cancelButtonText: '取消', type: 'warning' }).then(async () => {
    const res = await deleteSysMsg({ ID: row.ID })
    if (res.code === 0) { ElMessage.success('删除成功'); getTableData() }
  })
}
</script>
