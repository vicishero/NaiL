<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="昵称">
          <el-input v-model="searchInfo.nickname" placeholder="昵称" />
        </el-form-item>
        <el-form-item label="用户名">
          <el-input v-model="searchInfo.username" placeholder="用户名" />
        </el-form-item>
        <el-form-item label="钱包地址">
          <el-input v-model="searchInfo.walletAddress" placeholder="钱包地址" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchInfo.status" placeholder="全部" clearable style="width:100px">
            <el-option label="正常" :value="1" />
            <el-option label="冻结" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSearch">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table :data="tableData" row-key="ID" :default-sort="{ prop: 'ID', order: 'descending' }">
        <el-table-column align="left" label="ID" min-width="60" prop="ID" />
        <el-table-column align="left" label="头像" min-width="75">
          <template #default="scope">
            <el-avatar v-if="scope.row.avatar" :src="scope.row.avatar" :size="40" />
            <el-avatar v-else :size="40">{{ scope.row.nickname?.charAt(0) || 'U' }}</el-avatar>
          </template>
        </el-table-column>
        <el-table-column align="left" label="昵称" min-width="120" prop="nickname" />
        <el-table-column align="left" label="用户名" min-width="120" prop="username" />
        <el-table-column align="left" label="钱包地址" min-width="200" prop="walletAddress" />
        <el-table-column align="left" label="简介" min-width="150" prop="bio" show-overflow-tooltip />
        <el-table-column align="left" label="关注数" min-width="80" prop="followingCount" />
        <el-table-column align="left" label="粉丝数" min-width="80" prop="followerCount" />
        <el-table-column align="left" label="发帖数" min-width="80" prop="postCount" />
        <el-table-column align="left" label="状态" min-width="80">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'">{{ scope.row.status === 1 ? '正常' : '冻结' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="注册时间" min-width="170" prop="CreatedAt">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column label="操作" min-width="180" fixed="right">
          <template #default="scope">
            <el-button type="primary" link icon="edit" @click="editUser(scope.row)">编辑</el-button>
            <el-button type="primary" link icon="delete" @click="deleteUserFunc(scope.row)">删除</el-button>
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
          <span class="text-lg">编辑用户</span>
          <div>
            <el-button @click="closeDrawer">取消</el-button>
            <el-button type="primary" @click="saveUser">确定</el-button>
          </div>
        </div>
      </template>
      <el-form :model="form" label-width="100px">
        <el-form-item label="昵称">
          <el-input v-model="form.nickname" />
        </el-form-item>
        <el-form-item label="简介">
          <el-input v-model="form.bio" type="textarea" :rows="3" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="form.status" style="width:100%">
            <el-option label="正常" :value="1" />
            <el-option label="冻结" :value="0" />
          </el-select>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import { getH5UserList, getH5User, updateH5User, deleteH5User } from '@/api/h5Admin'
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate } from '@/utils/format'

defineOptions({ name: 'h5Users' })

const page = ref(1), total = ref(0), pageSize = ref(10), tableData = ref([])
const searchInfo = reactive({ nickname: '', username: '', walletAddress: '', status: undefined })
const drawerVisible = ref(false)
const form = ref({ ID: 0, nickname: '', bio: '', status: 1 })

const getTableData = async () => {
  const params = { page: page.value, pageSize: pageSize.value, ...searchInfo }
  if (searchInfo.status === undefined || searchInfo.status === '') delete params.status
  const res = await getH5UserList(params)
  if (res.code === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
  }
}
getTableData()

const onSearch = () => { page.value = 1; getTableData() }
const onReset = () => {
  searchInfo.nickname = ''; searchInfo.username = ''; searchInfo.walletAddress = ''; searchInfo.status = undefined
  page.value = 1; getTableData()
}
const handleSizeChange = (v) => { pageSize.value = v; getTableData() }
const handleCurrentChange = (v) => { page.value = v; getTableData() }

const editUser = async (row) => {
  const res = await getH5User({ ID: row.ID })
  if (res.code === 0) {
    form.value = { ID: res.data.ID, nickname: res.data.nickname || '', bio: res.data.bio || '', status: res.data.status }
    drawerVisible.value = true
  }
}
const closeDrawer = () => { drawerVisible.value = false }
const saveUser = async () => {
  const res = await updateH5User(form.value)
  if (res.code === 0) { ElMessage.success('更新成功'); drawerVisible.value = false; getTableData() }
}
const deleteUserFunc = async (row) => {
  ElMessageBox.confirm('确定要删除该用户吗?', '提示', { confirmButtonText: '确定', cancelButtonText: '取消', type: 'warning' }).then(async () => {
    const res = await deleteH5User({ ID: row.ID })
    if (res.code === 0) { ElMessage.success('删除成功'); if (tableData.value.length === 1 && page.value > 1) page.value--; getTableData() }
  })
}
</script>
