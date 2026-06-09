<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true">
        <el-form-item>
          <el-button type="primary" icon="plus" @click="addCategory">新增分类</el-button>
          <el-button icon="refresh" @click="getTableData">刷新</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table :data="tableData" row-key="ID">
        <el-table-column align="left" label="ID" width="170" prop="ID" />
        <el-table-column align="left" label="分类名称" min-width="200" prop="name" />
        <el-table-column align="left" label="排序" width="100" prop="sort" />
        <el-table-column align="left" label="用户数" width="100" prop="userCount" />
        <el-table-column label="操作" min-width="220" fixed="right">
          <template #default="scope">
            <el-button type="primary" link icon="view" @click="viewKolUsers(scope.row)">查看KOL</el-button>
            <el-button type="primary" link icon="edit" @click="editCategory(scope.row)">编辑</el-button>
            <el-button type="primary" link icon="delete" @click="deleteCategoryFunc(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <el-dialog v-model="dialogVisible" :title="isEdit ? '编辑分类' : '新增分类'" width="400px">
      <el-form :model="form" label-width="80px">
        <el-form-item label="名称"><el-input v-model="form.name" /></el-form-item>
        <el-form-item label="排序"><el-input-number v-model="form.sort" :min="0" /></el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="saveCategory">确定</el-button>
      </template>
    </el-dialog>

    <!-- KOL用户列表弹窗 -->
    <el-dialog v-model="kolDrawerVisible" :title="'KOL列表 - ' + currentCategoryName" width="700px">
      <div v-loading="kolLoading">
        <el-table :data="kolUsers" row-key="ID" max-height="450">
          <el-table-column align="left" label="ID" width="170" prop="ID" />
          <el-table-column align="left" label="头像" width="65">
            <template #default="scope">
              <el-avatar v-if="scope.row.avatar" :src="scope.row.avatar" :size="36" />
              <el-avatar v-else :size="36">{{ scope.row.nickname?.charAt(0) || 'K' }}</el-avatar>
            </template>
          </el-table-column>
          <el-table-column align="left" label="昵称" min-width="100" prop="nickname" />
          <el-table-column align="left" label="用户名" min-width="100" prop="username" />
          <el-table-column align="left" label="钱包地址" min-width="180">
            <template #default="scope">
              <span v-if="scope.row.walletAddress" class="wallet-text">{{ formatAddressShort(scope.row.walletAddress) }}</span>
              <span v-else style="color:#c0c4cc">-</span>
            </template>
          </el-table-column>
        </el-table>
        <div v-if="kolUsers.length === 0 && !kolLoading" style="text-align:center;padding:40px;color:#909399">暂无KOL用户</div>
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import { getKolCategoryList, saveKolCategory, deleteKolCategory, getKolManageList } from '@/api/h5Admin'
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

defineOptions({ name: 'kolCategories' })

const tableData = ref([])
const dialogVisible = ref(false), isEdit = ref(false)
const form = reactive({ ID: 0, name: '', sort: 0 })

// KOL users panel
const kolDrawerVisible = ref(false), kolLoading = ref(false)
const kolUsers = ref([]), currentCategoryName = ref('')

const getTableData = async () => {
  const res = await getKolCategoryList()
  if (res.code === 0) tableData.value = res.data?.list || []
}
getTableData()

const addCategory = () => { isEdit.value = false; form.ID = 0; form.name = ''; form.sort = 0; dialogVisible.value = true }
const editCategory = (row) => { isEdit.value = true; form.ID = row.ID; form.name = row.name; form.sort = row.sort; dialogVisible.value = true }
const saveCategory = async () => {
  const res = await saveKolCategory({ ID: form.ID, name: form.name, sort: form.sort })
  if (res.code === 0) { ElMessage.success(isEdit.value ? '更新成功' : '新增成功'); dialogVisible.value = false; getTableData() }
}
const deleteCategoryFunc = async (row) => {
  ElMessageBox.confirm('确定要删除该分类吗?', '提示', { confirmButtonText: '确定', cancelButtonText: '取消', type: 'warning' }).then(async () => {
    const res = await deleteKolCategory({ ID: row.ID })
    if (res.code === 0) { ElMessage.success('删除成功'); getTableData() }
  })
}

const viewKolUsers = async (row) => {
  currentCategoryName.value = row.name
  kolDrawerVisible.value = true
  kolLoading.value = true
  try {
    const res = await getKolManageList({ categoryId: row.ID, page: 1, pageSize: 100 })
    if (res.code === 0) kolUsers.value = res.data?.list || []
  } catch { kolUsers.value = [] }
  finally { kolLoading.value = false }
}

const formatAddressShort = (addr) => {
  if (!addr) return ''
  return addr.length > 10 ? addr.substring(0, 6) + '...' + addr.substring(addr.length - 4) : addr
}
</script>

<style scoped>
.wallet-text { font-family: monospace; font-size: 13px; }
</style>
