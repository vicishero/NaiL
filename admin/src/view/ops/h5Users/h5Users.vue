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
        <el-table-column align="left" label="ID" width="170" prop="ID" />
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
                  <el-icon class="action-icon" @click="copyAddress(scope.row.walletAddress)"><CopyDocument /></el-icon>
                </el-tooltip>
                <el-tooltip content="显示二维码" placement="top">
                  <span class="action-icon qr-icon" @click="showQrCode(scope.row.walletAddress)">
                    <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                      <rect x="3" y="3" width="7" height="7" /><rect x="14" y="3" width="7" height="7" />
                      <rect x="3" y="14" width="7" height="7" /><rect x="14" y="14" width="3" height="3" />
                      <rect x="18" y="14" width="3" height="3" /><rect x="14" y="18" width="3" height="3" />
                    </svg>
                  </span>
                </el-tooltip>
              </div>
            </div>
            <span v-else class="no-address">-</span>
          </template>
        </el-table-column>
        <el-table-column align="left" label="简介" min-width="150" prop="bio" show-overflow-tooltip />
        <el-table-column align="center" label="状态" width="90">
          <template #default="scope">
            <el-switch v-model="scope.row.statusSwitch" :active-value="1" :inactive-value="2"
              active-text="正常" inline-prompt @change="(val) => updateField(scope.row, 'status', val)" />
          </template>
        </el-table-column>
        <el-table-column align="center" label="管理员" width="90">
          <template #default="scope">
            <el-switch v-model="scope.row.isAdmin" @change="(val) => updateField(scope.row, 'isAdmin', val)" />
          </template>
        </el-table-column>
        <el-table-column align="center" label="KOL" width="80">
          <template #default="scope">
            <el-switch v-model="scope.row.isKOL" @change="(val) => updateField(scope.row, 'isKOL', val)" />
          </template>
        </el-table-column>
        <el-table-column align="center" label="聊天" width="80">
          <template #default="scope">
            <el-switch v-model="scope.row.chatEnabled" @change="(val) => updateField(scope.row, 'chatEnabled', val)" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="注册时间" min-width="170" prop="CreatedAt">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column label="操作" min-width="150" fixed="right">
          <template #default="scope">
            <el-button v-if="scope.row.isKOL" type="primary" link icon="view" @click="viewKol(scope.row)">查看</el-button>
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
        <a :href="`https://bscscan.com/address/${currentWalletAddress}`" target="_blank" class="bsc-link">在 BSCScan 查看</a>
      </div>
      <template #footer>
        <el-button type="primary" @click="copyAddress(currentWalletAddress)">复制地址</el-button>
      </template>
    </el-dialog>

    <!-- KOL 查看/编辑面板 -->
    <el-drawer v-model="kolDrawerVisible" :before-close="closeKolDrawer" :show-close="false" size="500px">
      <template #header>
        <div class="flex justify-between items-center">
          <span class="text-lg">KOL 人物属性 - {{ kolForm.nickname }}</span>
          <div>
            <el-button @click="closeKolDrawer">取消</el-button>
            <el-button type="primary" @click="saveKolProfile">保存</el-button>
          </div>
        </div>
      </template>
      <el-form :model="kolForm" label-width="90px">
        <el-form-item label="用户ID">
          <el-input :value="kolForm.userId" disabled />
        </el-form-item>
        <el-form-item label="KOL分类">
          <el-select v-model="kolForm.categoryId" placeholder="选择分类" style="width:100%">
            <el-option v-for="c in kolCategoryList" :key="c.ID" :label="c.name" :value="c.ID" />
          </el-select>
        </el-form-item>
        <el-form-item label="身高">
          <el-input v-model="kolForm.height" />
        </el-form-item>
        <el-form-item label="体重">
          <el-input v-model="kolForm.weight" />
        </el-form-item>
        <el-form-item label="三围">
          <el-input v-model="kolForm.measurements" />
        </el-form-item>
        <el-form-item label="肤色">
          <el-input v-model="kolForm.skinTone" />
        </el-form-item>
        <el-form-item label="瞳色">
          <el-input v-model="kolForm.eyeColor" />
        </el-form-item>
        <el-form-item label="性向">
          <el-input v-model="kolForm.orientation" />
        </el-form-item>
        <el-form-item label="喜好">
          <el-input v-model="kolForm.preferences" type="textarea" :rows="2" />
        </el-form-item>
        <el-form-item label="喜欢食物">
          <el-input v-model="kolForm.favoriteFoods" type="textarea" :rows="2" />
        </el-form-item>
        <el-form-item label="服装风格">
          <el-input v-model="kolForm.clothingStyle" />
        </el-form-item>
        <el-form-item label="妆风格">
          <el-input v-model="kolForm.makeupStyle" />
        </el-form-item>
        <el-form-item label="系统提示词">
          <el-input v-model="kolForm.systemPrompt" type="textarea" :rows="6" placeholder="AI 对话的系统提示词..." />
        </el-form-item>
        <el-form-item label="API私钥">
          <el-input v-model="kolForm.apiKey" placeholder="Dify API Key" maxlength="64" show-word-limit />
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import { getH5UserList, updateH5User, deleteH5User, getKolProfile, saveKolProfileApi, getKolCategoryList } from '@/api/h5Admin'
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { formatDate } from '@/utils/format'
import { CopyDocument } from '@element-plus/icons-vue'
import QrcodeVue from 'qrcode.vue'

defineOptions({ name: 'h5Users' })

const page = ref(1), total = ref(0), pageSize = ref(10), tableData = ref([])
const searchInfo = reactive({ nickname: '', username: '', walletAddress: '', status: undefined })
const qrCodeVisible = ref(false)
const currentWalletAddress = ref('')

// KOL drawer
const kolDrawerVisible = ref(false)
const kolCategoryList = ref([])
const kolForm = reactive({
  userId: 0, nickname: '',
  height: '160cm', weight: '44kg', measurements: '84/58/84',
  skinTone: '冷白病态肌', eyeColor: '酒红', orientation: '偏双性恋（情感依赖向）',
  preferences: '独占欲、暗调氛围、偏执温柔', favoriteFoods: '黑森林、红酒、冷食',
  clothingStyle: '黑裙、蕾丝、丝带、哥特风', makeupStyle: '苍白底妆、下垂眼、暗红眼影、冷唇',
  systemPrompt: '',
  apiKey: "",
  categoryId: 0
})

const getTableData = async () => {
  const params = { page: page.value, pageSize: pageSize.value, ...searchInfo }
  if (searchInfo.status === undefined || searchInfo.status === '') delete params.status
  const res = await getH5UserList(params)
  if (res.code === 0) {
    tableData.value = (res.data.list || []).map(u => ({
      ...u,
      statusSwitch: u.status // for el-switch two-value binding
    }))
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

const updateField = async (row, field, val) => {
  const payload = { ID: row.ID }
  if (field === 'status') payload.status = val
  else if (field === 'isAdmin') payload.isAdmin = val
  else if (field === 'isKOL') payload.isKOL = val
  else if (field === 'chatEnabled') payload.chatEnabled = val
  const res = await updateH5User(payload)
  if (res.code === 0) {
    ElMessage.success('更新成功')
    if (field === 'isKOL' && !val) row.isKOL = false
  } else {
    // revert on failure
    if (field === 'isAdmin') row.isAdmin = row.isAdmin
    else if (field === 'isKOL') row.isKOL = !val
    else if (field === 'status') row.statusSwitch = val === 1 ? 2 : 1
  }
}

const deleteUserFunc = async (row) => {
  ElMessageBox.confirm('确定要删除该用户吗?', '提示', { confirmButtonText: '确定', cancelButtonText: '取消', type: 'warning' }).then(async () => {
    const res = await deleteH5User({ ID: row.ID })
    if (res.code === 0) { ElMessage.success('删除成功'); if (tableData.value.length === 1 && page.value > 1) page.value--; getTableData() }
  })
}

// KOL profile
const viewKol = async (row) => {
  // 加载分类列表
  try { const r = await getKolCategoryList(); if (r.code === 0) kolCategoryList.value = r.data?.list || [] } catch { kolCategoryList.value = [] }
  Object.assign(kolForm, {
    userId: row.ID, nickname: row.nickname,
    height: '160cm', weight: '44kg', measurements: '84/58/84',
    skinTone: '冷白病态肌', eyeColor: '酒红', orientation: '偏双性恋（情感依赖向）',
    preferences: '独占欲、暗调氛围、偏执温柔', favoriteFoods: '黑森林、红酒、冷食',
    clothingStyle: '黑裙、蕾丝、丝带、哥特风', makeupStyle: '苍白底妆、下垂眼、暗红眼影、冷唇',
  systemPrompt: '',
  apiKey: "",
    categoryId: 0
  })
  try {
    const res = await getKolProfile({ userId: row.ID })
    if (res.code === 0 && res.data) {
      Object.assign(kolForm, res.data)
    }
  } catch { /* use defaults */ }
  kolDrawerVisible.value = true
}
const closeKolDrawer = () => { kolDrawerVisible.value = false }
const saveKolProfile = async () => {
  const res = await saveKolProfileApi(kolForm)
  if (res.code === 0) { ElMessage.success('KOL属性保存成功'); kolDrawerVisible.value = false }
}

// wallet helpers
const copyAddress = async (address) => {
  try { await navigator.clipboard.writeText(address); ElMessage.success('地址已复制到剪贴板') }
  catch {
    const ta = document.createElement('textarea'); ta.value = address; document.body.appendChild(ta); ta.select()
    document.execCommand('copy'); document.body.removeChild(ta); ElMessage.success('地址已复制到剪贴板')
  }
}
const showQrCode = (address) => { currentWalletAddress.value = address; qrCodeVisible.value = true }
const formatAddressShort = (address) => {
  if (!address || address.length < 10) return address
  return address.substring(0, 6) + '...' + address.substring(address.length - 4)
}
</script>

<style scoped>
.wallet-address-cell { display: flex; align-items: center; gap: 8px; }
.address-link { color: #409eff; text-decoration: none; font-family: monospace; font-size: 14px; }
.address-link:hover { text-decoration: underline; }
.address-actions { display: flex; gap: 6px; }
.action-icon { cursor: pointer; color: #909399; font-size: 16px; transition: color 0.2s; display: inline-flex; align-items: center; justify-content: center; }
.action-icon:hover { color: #409eff; }
.qr-icon svg { width: 16px; height: 16px; }
.no-address { color: #c0c4cc; }
.qr-code-container { display: flex; flex-direction: column; align-items: center; gap: 16px; padding: 20px 0; }
.address-text { margin: 0; font-family: monospace; font-size: 14px; color: #606266; word-break: break-all; text-align: center; max-width: 250px; }
.bsc-link { color: #409eff; text-decoration: none; font-size: 13px; }
.bsc-link:hover { text-decoration: underline; }
</style>
