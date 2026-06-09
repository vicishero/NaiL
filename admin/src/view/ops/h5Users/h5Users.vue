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
            <el-option label="冻结" :value="2" />
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
        <el-table-column align="left" label="ID" min-width="170" prop="ID" />
        <el-table-column align="left" label="头像" min-width="75">
          <template #default="scope">
            <el-avatar v-if="scope.row.avatar" :src="scope.row.avatar" :size="40" />
            <el-avatar v-else :size="40">{{ scope.row.nickname?.charAt(0) || 'U' }}</el-avatar>
          </template>
        </el-table-column>
        <el-table-column align="left" label="昵称" min-width="120" prop="nickname" />
        <el-table-column align="left" label="用户名" min-width="120" prop="username" />
        <el-table-column align="left" label="钱包地址" min-width="280" prop="walletAddress">
          <template #default="scope">
            <div v-if="scope.row.walletAddress" class="wallet-address-cell">
              <el-tooltip :content="scope.row.walletAddress" placement="top">
                <a :href="`https://bscscan.com/address/${scope.row.walletAddress}`" target="_blank" class="address-link">
                  {{ formatAddressShort(scope.row.walletAddress) }}
                </a>
              </el-tooltip>
              <div class="address-actions">
                <el-tooltip content="复制地址" placement="top">
                  <el-icon class="action-icon" @click="copyAddress(scope.row.walletAddress)">
                    <CopyDocument />
                  </el-icon>
                </el-tooltip>
                <el-tooltip content="显示二维码" placement="top">
                  <span class="action-icon qr-icon" @click="showQrCode(scope.row.walletAddress)">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <rect x="3" y="3" width="7" height="7" />
                      <rect x="14" y="3" width="7" height="7" />
                      <rect x="3" y="14" width="7" height="7" />
                      <rect x="14" y="14" width="3" height="3" />
                      <rect x="18" y="14" width="3" height="3" />
                      <rect x="14" y="18" width="3" height="3" />
                    </svg>
                  </span>
                </el-tooltip>
              </div>
            </div>
            <span v-else class="no-address">-</span>
          </template>
        </el-table-column>
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

    <!-- 钱包地址二维码弹窗 -->
    <el-dialog v-model="qrCodeVisible" title="BSC 钱包地址" width="400px" center>
      <div class="qr-code-container">
        <qrcode-vue :value="currentWalletAddress" :size="250" level="H" />
        <p class="address-text">{{ currentWalletAddress }}</p>
        <a :href="`https://bscscan.com/address/${currentWalletAddress}`" target="_blank" class="bsc-link">
          在 BSCScan 查看
        </a>
      </div>
      <template #footer>
        <el-button type="primary" @click="copyAddress(currentWalletAddress)">复制地址</el-button>
      </template>
    </el-dialog>

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
            <el-option label="冻结" :value="2" />
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
import { CopyDocument } from '@element-plus/icons-vue'
import QrcodeVue from 'qrcode.vue'

defineOptions({ name: 'h5Users' })

const page = ref(1), total = ref(0), pageSize = ref(10), tableData = ref([])
const searchInfo = reactive({ nickname: '', username: '', walletAddress: '', status: undefined })
const drawerVisible = ref(false)
const form = ref({ ID: 0, nickname: '', bio: '', status: 1 })
const qrCodeVisible = ref(false)
const currentWalletAddress = ref('')

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

// 复制钱包地址到剪贴板
const copyAddress = async (address) => {
  try {
    await navigator.clipboard.writeText(address)
    ElMessage.success('地址已复制到剪贴板')
  } catch (err) {
    // 降级方案
    const textArea = document.createElement('textarea')
    textArea.value = address
    document.body.appendChild(textArea)
    textArea.select()
    document.execCommand('copy')
    document.body.removeChild(textArea)
    ElMessage.success('地址已复制到剪贴板')
  }
}

// 显示二维码
const showQrCode = (address) => {
  currentWalletAddress.value = address
  qrCodeVisible.value = true
}

// 钱包地址缩写格式：0x1234...123a
const formatAddressShort = (address) => {
  if (!address || address.length < 10) return address
  return address.substring(0, 6) + '...' + address.substring(address.length - 4)
}
</script>

<style scoped>
.wallet-address-cell {
  display: flex;
  align-items: center;
  gap: 8px;
}

.address-link {
  color: #409eff;
  text-decoration: none;
  font-family: monospace;
  font-size: 14px;
}

.address-link:hover {
  text-decoration: underline;
}

.address-actions {
  display: flex;
  gap: 6px;
}

.action-icon {
  cursor: pointer;
  color: #909399;
  font-size: 16px;
  transition: color 0.2s;
  display: inline-flex;
  align-items: center;
  justify-content: center;
}

.action-icon:hover {
  color: #409eff;
}

.qr-icon svg {
  width: 16px;
  height: 16px;
}

.no-address {
  color: #c0c4cc;
}

.qr-code-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16px;
  padding: 20px 0;
}

.address-text {
  margin: 0;
  font-family: monospace;
  font-size: 14px;
  color: #606266;
  word-break: break-all;
  text-align: center;
  max-width: 250px;
}

.bsc-link {
  color: #409eff;
  text-decoration: none;
  font-size: 13px;
}

.bsc-link:hover {
  text-decoration: underline;
}
</style>
