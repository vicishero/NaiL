# 管理后台贴文和话题管理功能实现计划

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** 在管理后台实现贴文管理（p_post表，删改查）和话题管理（p_tag表，增删改查）功能

**Architecture:** 
- 后端：沿用现有Admin服务架构，在mirc定义路由，servants实现业务逻辑，DAO层复用现有Post和Tag模型
- 前端：参考现有h5Users.vue页面模式，创建posts.vue和tags.vue管理页面，添加路由和菜单

**Tech Stack:** Go + Gin + GORM, Vue 3 + TypeScript + Element Plus / Naive UI

---

## Task 1: 后端路由定义 (mirc/web/v1/admin.go)

**Files:**
- Modify: `mirc/web/v1/admin.go`
- Related: `internal/core/admin/h5.go` (结构体已定义)

- [ ] **Step 1: 添加贴文管理路由**

在`Admin`结构体中添加：
```go
// H5PostList 贴文列表
H5PostList func(Get, admin.H5PostListReq) admin.H5PostListResp `mir:"admin/post/list"`
// H5PostGet 获取单个贴文
H5PostGet func(Get, admin.H5PostGetReq) admin.H5PostItem `mir:"admin/post/get"`
// H5PostUpdate 更新贴文
H5PostUpdate func(Post, admin.H5PostUpdateReq) error `mir:"admin/post/update"`
// H5PostDelete 删除贴文
H5PostDelete func(Post, admin.H5PostDeleteReq) error `mir:"admin/post/delete"`
```

- [ ] **Step 2: 添加话题管理路由**

在`Admin`结构体中添加：
```go
// H5TagList 话题列表
H5TagList func(Get, admin.H5TagListReq) admin.H5TagListResp `mir:"admin/tag/list"`
// H5TagCreate 创建话题
H5TagCreate func(Post, admin.H5TagItem) error `mir:"admin/tag/create"`
// H5TagUpdate 更新话题
H5TagUpdate func(Post, admin.H5TagUpdateReq) error `mir:"admin/tag/update"`
// H5TagDelete 删除话题
H5TagDelete func(Post, admin.H5TagDeleteReq) error `mir:"admin/tag/delete"`
```

- [ ] **Step 3: 添加H5TagListResp结构体（如果缺失）**

在`internal/core/admin/h5.go`中添加（如果不存在）：
```go
// H5TagListResp 话题列表响应
type H5TagListResp struct {
	List  []H5TagItem `json:"list"`
	Total int64       `json:"total"`
}

// H5TagCreateReq 创建话题请求（如果需要）
type H5TagCreateReq struct {
	Tag      string `json:"tag"`
	QuoteNum int64  `json:"quoteNum"`
}
```

---

## Task 2: 后端服务实现 (internal/servants/web/admin.go)

**Files:**
- Modify: `internal/servants/web/admin.go`

- [ ] **Step 1: 实现贴文列表接口**

```go
import (
	// ... existing imports
	"github.com/rocboss/paopao-ce/internal/core"
)

func (s *adminSrv) H5PostList(req *core.H5PostListReq) (*core.H5PostListResp, error) {
	conditions := &dbr.ConditionsT{}
	if req.Keyword != "" {
		// 可以通过关联查询用户或内容来实现搜索
	}
	if req.UserID > 0 {
		(*conditions)["user_id"] = req.UserID
	}
	if req.Visibility != nil {
		(*conditions)["visibility"] = *req.Visibility
	}
	(*conditions)["ORDER"] = "id DESC"

	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	posts, err := s.Ds.ListPosts(conditions, offset, pageSize)
	if err != nil {
		return nil, err
	}

	total, err := s.Ds.CountPost(conditions)
	if err != nil {
		return nil, err
	}

	// 格式化数据
	list := make([]core.H5PostItem, 0, len(posts))
	for _, post := range posts {
		// 获取用户信息
		user, _ := s.Ds.GetUserByID(post.UserID)
		
		// 获取贴文内容
		postContents, _ := s.Ds.GetPostContents(post.Model.ID)
		
		contents := make([]core.H5PostContent, 0, len(postContents))
		for _, c := range postContents {
			contents = append(contents, core.H5PostContent{
				Type:    int(c.Type),
				Content: c.Content,
				Sort:    c.Sort,
			})
		}

		userItem := &core.H5UserItem{}
		if user != nil && user.Model != nil {
			userItem = &core.H5UserItem{
				ID:       user.ID,
				Nickname: user.Nickname,
				Username: user.Username,
				Avatar:   user.Avatar,
			}
		}

		list = append(list, core.H5PostItem{
			ID:             post.Model.ID,
			UserID:         post.UserID,
			User:           userItem,
			Contents:       contents,
			CommentCount:   post.CommentCount,
			UpvoteCount:    post.UpvoteCount,
			ShareCount:     post.ShareCount,
			CollectionCount: post.CollectionCount,
			Visibility:     int(post.Visibility),
			IsTop:          int8(post.IsTop),
			IsEssence:      int8(post.IsEssence),
			IsLock:         int8(post.IsLock),
			CreatedAt:      time.Unix(post.CreatedOn, 0).Format("2006-01-02 15:04:05"),
		})
	}

	return &core.H5PostListResp{
		List:  list,
		Total: total,
	}, nil
}
```

- [ ] **Step 2: 实现获取单个贴文接口**

```go
func (s *adminSrv) H5PostGet(req *core.H5PostGetReq) (*core.H5PostItem, error) {
	post, err := s.Ds.GetPostByID(req.ID)
	if err != nil || post.Model == nil {
		return nil, err
	}

	user, _ := s.Ds.GetUserByID(post.UserID)
	postContents, _ := s.Ds.GetPostContents(post.Model.ID)

	contents := make([]core.H5PostContent, 0, len(postContents))
	for _, c := range postContents {
		contents = append(contents, core.H5PostContent{
			Type:    int(c.Type),
			Content: c.Content,
			Sort:    c.Sort,
		})
	}

	userItem := &core.H5UserItem{}
	if user != nil && user.Model != nil {
		userItem = &core.H5UserItem{
			ID:       user.ID,
			Nickname: user.Nickname,
			Username: user.Username,
			Avatar:   user.Avatar,
		}
	}

	return &core.H5PostItem{
		ID:               post.Model.ID,
		UserID:           post.UserID,
		User:             userItem,
		Contents:         contents,
		CommentCount:     post.CommentCount,
		UpvoteCount:      post.UpvoteCount,
		ShareCount:       post.ShareCount,
		CollectionCount:  post.CollectionCount,
		Visibility:       int(post.Visibility),
		IsTop:            int8(post.IsTop),
		IsEssence:        int8(post.IsEssence),
		IsLock:           int8(post.IsLock),
		CreatedAt:        time.Unix(post.CreatedOn, 0).Format("2006-01-02 15:04:05"),
	}, nil
}
```

- [ ] **Step 3: 实现更新贴文接口**

```go
func (s *adminSrv) H5PostUpdate(req *core.H5PostUpdateReq) error {
	post, err := s.Ds.GetPostByID(req.ID)
	if err != nil || post.Model == nil {
		return err
	}

	post.Visibility = dbr.PostVisibleT(req.Visibility)
	post.IsTop = int(req.IsTop)
	post.IsEssence = int(req.IsEssence)
	post.IsLock = int(req.IsLock)

	return s.Ds.UpdatePost(post)
}
```

- [ ] **Step 4: 实现删除贴文接口**

```go
func (s *adminSrv) H5PostDelete(req *core.H5PostDeleteReq) error {
	post, err := s.Ds.GetPostByID(req.ID)
	if err != nil || post.Model == nil {
		return err
	}

	return s.Ds.DeletePost(post)
}
```

- [ ] **Step 5: 实现话题列表接口**

```go
func (s *adminSrv) H5TagList(req *core.H5TagListReq) (*core.H5TagListResp, error) {
	conditions := &dbr.ConditionsT{}
	if req.Keyword != "" {
		(*conditions)["tag LIKE ?"] = "%" + req.Keyword + "%"
	}
	(*conditions)["ORDER"] = "id DESC"

	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	tags, err := s.Ds.ListTags(conditions, offset, pageSize)
	if err != nil {
		return nil, err
	}

	total, err := s.Ds.CountTag(conditions)
	if err != nil {
		return nil, err
	}

	list := make([]core.H5TagItem, 0, len(tags))
	for _, tag := range tags {
		list = append(list, core.H5TagItem{
			ID:        tag.Model.ID,
			Tag:       tag.Tag,
			QuoteNum:  tag.QuoteNum,
			UserID:    tag.UserID,
			CreatedAt: time.Unix(tag.CreatedOn, 0).Format("2006-01-02 15:04:05"),
		})
	}

	return &core.H5TagListResp{
		List:  list,
		Total: total,
	}, nil
}
```

- [ ] **Step 6: 实现创建话题接口**

```go
func (s *adminSrv) H5TagCreate(req *core.H5TagItem) error {
	// 检查tag是否已存在
	existingTag, err := s.Ds.GetTagByTag(req.Tag)
	if err == nil && existingTag != nil && existingTag.Model != nil {
		return xerror.New("话题已存在")
	}

	tag := &dbr.Tag{
		Model:    &dbr.Model{},
		UserID:   0, // 系统创建
		Tag:      req.Tag,
		QuoteNum: req.QuoteNum,
	}

	_, err = s.Ds.CreateTag(tag)
	return err
}
```

- [ ] **Step 7: 实现更新话题接口**

```go
func (s *adminSrv) H5TagUpdate(req *core.H5TagUpdateReq) error {
	tag, err := s.Ds.GetTagByID(req.ID)
	if err != nil || tag.Model == nil {
		return err
	}

	tag.Tag = req.Tag
	tag.QuoteNum = req.QuoteNum

	return s.Ds.UpdateTag(tag)
}
```

- [ ] **Step 8: 实现删除话题接口**

```go
func (s *adminSrv) H5TagDelete(req *core.H5TagDeleteReq) error {
	tag, err := s.Ds.GetTagByID(req.ID)
	if err != nil || tag.Model == nil {
		return err
	}

	// 注意：软删除
	return s.Ds.DeleteTag(tag)
}
```

---

## Task 3: DAO层方法验证 (internal/dao/jinzhu/)

**Files:**
- Verify: `internal/dao/jinzhu/posts.go`, `internal/dao/jinzhu/tags.go`

- [ ] **Step 1: 验证Posts DAO方法是否存在**

检查以下方法是否已实现：
- `ListPosts(conditions *dbr.ConditionsT, offset, limit int) ([]*dbr.Post, error)`
- `CountPost(conditions *dbr.ConditionsT) (int64, error)`
- `GetPostByID(id int64) (*dbr.Post, error)`
- `UpdatePost(post *dbr.Post) error`
- `DeletePost(post *dbr.Post) error`
- `GetPostContents(postID int64) ([]*dbr.PostContent, error)`

**如果缺失，参考dbr/post.go的List方法实现**

- [ ] **Step 2: 验证Tags DAO方法是否存在**

检查以下方法是否已实现：
- `ListTags(conditions *dbr.ConditionsT, offset, limit int) ([]*dbr.Tag, error)`
- `CountTag(conditions *dbr.ConditionsT) (int64, error)`
- `GetTagByID(id int64) (*dbr.Tag, error)`
- `GetTagByTag(tag string) (*dbr.Tag, error)`
- `CreateTag(tag *dbr.Tag) (*dbr.Tag, error)`
- `UpdateTag(tag *dbr.Tag) error`
- `DeleteTag(tag *dbr.Tag) error`

**如果缺失，参考dbr/topic.go的方法在jinzhu/tags.go中实现**

---

## Task 4: 前端API方法 (admin/src/api)

**Files:**
- Create/Modify: `admin/src/api/post.ts` or `admin/src/api/admin.ts`

- [ ] **Step 1: 创建贴文管理API**

```typescript
import request from '@/utils/request'

// 贴文列表
export function getH5PostList(params: any) {
  return request({
    url: '/v1/admin/post/list',
    method: 'get',
    params
  })
}

// 获取单个贴文
export function getH5Post(params: any) {
  return request({
    url: '/v1/admin/post/get',
    method: 'get',
    params
  })
}

// 更新贴文
export function updateH5Post(data: any) {
  return request({
    url: '/v1/admin/post/update',
    method: 'post',
    data
  })
}

// 删除贴文
export function deleteH5Post(data: any) {
  return request({
    url: '/v1/admin/post/delete',
    method: 'post',
    data
  })
}
```

- [ ] **Step 2: 创建话题管理API**

```typescript
// 话题列表
export function getH5TagList(params: any) {
  return request({
    url: '/v1/admin/tag/list',
    method: 'get',
    params
  })
}

// 创建话题
export function createH5Tag(data: any) {
  return request({
    url: '/v1/admin/tag/create',
    method: 'post',
    data
  })
}

// 更新话题
export function updateH5Tag(data: any) {
  return request({
    url: '/v1/admin/tag/update',
    method: 'post',
    data
  })
}

// 删除话题
export function deleteH5Tag(data: any) {
  return request({
    url: '/v1/admin/tag/delete',
    method: 'post',
    data
  })
}
```

---

## Task 5: 贴文管理前端页面 (admin/src/view/ops/posts/posts.vue)

**Files:**
- Create: `admin/src/view/ops/posts/posts.vue`

- [ ] **Step 1: 创建贴文管理页面**

参考`h5Users.vue`的结构实现：

```vue
<template>
  <div class="post-management">
    <!-- 搜索区域 -->
    <el-form :inline="true" :model="searchInfo">
      <el-form-item label="关键词">
        <el-input v-model="searchInfo.keyword" placeholder="搜索" />
      </el-form-item>
      <el-form-item label="用户ID">
        <el-input v-model="searchInfo.userId" placeholder="用户ID" />
      </el-form-item>
      <el-form-item label="可见性">
        <el-select v-model="searchInfo.visibility" placeholder="全部" clearable>
          <el-option label="公开" :value="0" />
          <el-option label="私密" :value="1" />
          <el-option label="关注可见" :value="3" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSearch">搜索</el-button>
        <el-button @click="onReset">重置</el-button>
      </el-form-item>
    </el-form>

    <!-- 表格区域 -->
    <el-table :data="tableData" border>
      <el-table-column prop="ID" label="ID" width="80" />
      <el-table-column prop="user.nickname" label="用户" width="120">
        <template #default="{ row }">
          <span>{{ row.user?.nickname || '-' }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="contents" label="内容" min-width="200">
        <template #default="{ row }">
          <span class="content-preview">
            {{ getContentPreview(row.contents) }}
          </span>
        </template>
      </el-table-column>
      <el-table-column prop="commentCount" label="评论" width="80" />
      <el-table-column prop="upvoteCount" label="点赞" width="80" />
      <el-table-column prop="visibility" label="可见性" width="100">
        <template #default="{ row }">
          <el-tag :type="getVisibilityType(row.visibility)">
            {{ getVisibilityText(row.visibility) }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="isTop" label="置顶" width="80">
        <template #default="{ row }">
          <el-tag v-if="row.isTop === 1" type="danger">是</el-tag>
          <span v-else>否</span>
        </template>
      </el-table-column>
      <el-table-column prop="isEssence" label="精华" width="80">
        <template #default="{ row }">
          <el-tag v-if="row.isEssence === 1" type="success">是</el-tag>
          <span v-else>否</span>
        </template>
      </el-table-column>
      <el-table-column prop="isLock" label="锁定" width="80">
        <template #default="{ row }">
          <el-tag v-if="row.isLock === 1" type="warning">是</el-tag>
          <span v-else>否</span>
        </template>
      </el-table-column>
      <el-table-column prop="CreatedAt" label="创建时间" width="160" />
      <el-table-column label="操作" width="200" fixed="right">
        <template #default="{ row }">
          <el-button link type="primary" size="small" @click="editPost(row)">编辑</el-button>
          <el-button link type="danger" size="small" @click="deletePost(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 分页 -->
    <div class="pagination-container">
      <el-pagination
        v-model:current-page="page"
        v-model:page-size="pageSize"
        :page-sizes="[10, 20, 50, 100]"
        :total="total"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>

    <!-- 编辑抽屉 -->
    <el-drawer
      v-model="drawerVisible"
      title="编辑贴文"
      size="40%"
    >
      <el-form :model="form" label-width="100px">
        <el-form-item label="可见性">
          <el-select v-model="form.visibility">
            <el-option label="公开" :value="0" />
            <el-option label="私密" :value="1" />
            <el-option label="关注可见" :value="3" />
          </el-select>
        </el-form-item>
        <el-form-item label="置顶">
          <el-switch v-model="form.isTop" :active-value="1" :inactive-value="0" />
        </el-form-item>
        <el-form-item label="精华">
          <el-switch v-model="form.isEssence" :active-value="1" :inactive-value="0" />
        </el-form-item>
        <el-form-item label="锁定">
          <el-switch v-model="form.isLock" :active-value="1" :inactive-value="0" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="drawerVisible = false">取消</el-button>
        <el-button type="primary" @click="savePost">保存</el-button>
      </template>
    </el-drawer>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getH5PostList,
  getH5Post,
  updateH5Post,
  deleteH5Post,
} from '@/api/post'

defineOptions({ name: 'Posts' })

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const drawerVisible = ref(false)
const searchInfo = reactive({
  keyword: '',
  userId: '',
  visibility: undefined,
})
const form = ref({
  ID: 0,
  visibility: 0,
  isTop: 0,
  isEssence: 0,
  isLock: 0,
})

const getTableData = async () => {
  const params = {
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo,
    userId: searchInfo.userId ? parseInt(searchInfo.userId) : undefined,
  }
  const res = await getH5PostList(params)
  if (res.code === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
  }
}

getTableData()

const onSearch = () => {
  page.value = 1
  getTableData()
}

const onReset = () => {
  searchInfo.keyword = ''
  searchInfo.userId = ''
  searchInfo.visibility = undefined
  page.value = 1
  getTableData()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const getContentPreview = (contents) => {
  if (!contents || contents.length === 0) return '-'
  const textContent = contents.find(c => c.type === 0 || c.type === 1)
  if (textContent) {
    return textContent.content.length > 50 
      ? textContent.content.substring(0, 50) + '...' 
      : textContent.content
  }
  return '[媒体内容]'
}

const getVisibilityType = (visibility) => {
  const map = { 0: 'success', 1: 'info', 3: 'warning' }
  return map[visibility] || ''
}

const getVisibilityText = (visibility) => {
  const map = { 0: '公开', 1: '私密', 3: '关注可见' }
  return map[visibility] || '未知'
}

const editPost = async (row) => {
  const res = await getH5Post({ ID: row.ID })
  if (res.code === 0) {
    form.value = {
      ID: res.data.ID,
      visibility: res.data.visibility,
      isTop: res.data.isTop,
      isEssence: res.data.isEssence,
      isLock: res.data.isLock,
    }
    drawerVisible.value = true
  }
}

const savePost = async () => {
  const res = await updateH5Post(form.value)
  if (res.code === 0) {
    ElMessage.success('更新成功')
    drawerVisible.value = false
    getTableData()
  }
}

const deletePost = (row) => {
  ElMessageBox.confirm('确定要删除该贴文吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(async () => {
    const res = await deleteH5Post({ ID: row.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}
</script>

<style scoped>
.post-management {
  padding: 20px;
}
.pagination-container {
  margin-top: 20px;
  text-align: right;
}
.content-preview {
  color: #606266;
  font-size: 14px;
}
</style>
```

---

## Task 6: 话题管理前端页面 (admin/src/view/ops/tags/tags.vue)

**Files:**
- Create: `admin/src/view/ops/tags/tags.vue`

- [ ] **Step 1: 创建话题管理页面**

```vue
<template>
  <div class="tag-management">
    <!-- 搜索和操作区域 -->
    <el-form :inline="true" :model="searchInfo">
      <el-form-item label="关键词">
        <el-input v-model="searchInfo.keyword" placeholder="搜索话题" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSearch">搜索</el-button>
        <el-button @click="onReset">重置</el-button>
        <el-button type="success" @click="createTag">新建话题</el-button>
      </el-form-item>
    </el-form>

    <!-- 表格区域 -->
    <el-table :data="tableData" border>
      <el-table-column prop="ID" label="ID" width="80" />
      <el-table-column prop="tag" label="话题名称" min-width="150">
        <template #default="{ row }">
          <el-tag type="success">#{{ row.tag }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="quoteNum" label="引用数" width="120" />
      <el-table-column prop="CreatedAt" label="创建时间" width="160" />
      <el-table-column label="操作" width="200" fixed="right">
        <template #default="{ row }">
          <el-button link type="primary" size="small" @click="editTag(row)">编辑</el-button>
          <el-button link type="danger" size="small" @click="deleteTag(row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 分页 -->
    <div class="pagination-container">
      <el-pagination
        v-model:current-page="page"
        v-model:page-size="pageSize"
        :page-sizes="[10, 20, 50, 100]"
        :total="total"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>

    <!-- 编辑/创建对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="isEditMode ? '编辑话题' : '新建话题'"
      width="400px"
    >
      <el-form :model="form" label-width="80px">
        <el-form-item label="话题名称">
          <el-input v-model="form.tag" placeholder="请输入话题名称" />
        </el-form-item>
        <el-form-item label="引用数">
          <el-input-number v-model="form.quoteNum" :min="0" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="saveTag">保存</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import {
  getH5TagList,
  createH5Tag,
  updateH5Tag,
  deleteH5Tag,
} from '@/api/post'

defineOptions({ name: 'Tags' })

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const dialogVisible = ref(false)
const isEditMode = ref(false)
const searchInfo = reactive({
  keyword: '',
})
const form = ref({
  ID: 0,
  tag: '',
  quoteNum: 0,
})

const getTableData = async () => {
  const params = {
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo,
  }
  const res = await getH5TagList(params)
  if (res.code === 0) {
    tableData.value = res.data.list
    total.value = res.data.total
  }
}

getTableData()

const onSearch = () => {
  page.value = 1
  getTableData()
}

const onReset = () => {
  searchInfo.keyword = ''
  page.value = 1
  getTableData()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const createTag = () => {
  isEditMode.value = false
  form.value = {
    ID: 0,
    tag: '',
    quoteNum: 0,
  }
  dialogVisible.value = true
}

const editTag = (row) => {
  isEditMode.value = true
  form.value = {
    ID: row.ID,
    tag: row.tag,
    quoteNum: row.quoteNum,
  }
  dialogVisible.value = true
}

const saveTag = async () => {
  if (!form.value.tag.trim()) {
    ElMessage.warning('请输入话题名称')
    return
  }

  let res
  if (isEditMode.value) {
    res = await updateH5Tag(form.value)
  } else {
    res = await createH5Tag(form.value)
  }

  if (res.code === 0) {
    ElMessage.success(isEditMode.value ? '更新成功' : '创建成功')
    dialogVisible.value = false
    getTableData()
  }
}

const deleteTag = (row) => {
  ElMessageBox.confirm('确定要删除该话题吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(async () => {
    const res = await deleteH5Tag({ ID: row.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      if (tableData.value.length === 1 && page.value > 1) {
        page.value--
      }
      getTableData()
    }
  })
}
</script>

<style scoped>
.tag-management {
  padding: 20px;
}
.pagination-container {
  margin-top: 20px;
  text-align: right;
}
</style>
```

---

## Task 7: 路由和菜单配置

**Files:**
- Modify: `admin/src/router/index.ts` (或对应的路由配置文件)
- Modify: 菜单配置文件（通常在router或pinia中）

- [ ] **Step 1: 添加路由配置**

在路由文件中添加：

```typescript
// 贴文管理
{
  path: '/ops/posts',
  name: 'Posts',
  component: () => import('@/view/ops/posts/posts.vue'),
  meta: {
    title: '贴文管理',
    icon: 'Document',
  },
},
// 话题管理
{
  path: '/ops/tags',
  name: 'Tags',
  component: () => import('@/view/ops/tags/tags.vue'),
  meta: {
    title: '话题管理',
    icon: 'PriceTag',
  },
},
```

- [ ] **Step 2: 添加菜单配置**

在菜单配置中找到"运维管理"或"内容管理"分组，添加这两个菜单项。

---

## Task 8: 代码生成和编译测试

**Files:**
- Generated: `auto/api/v1/admin.go` (通过 `go generate`)

- [ ] **Step 1: 运行go generate生成API代码**

```bash
cd /home/v3/workspace/NaiL
go generate ./mirc/web/v1/...
```

- [ ] **Step 2: 编译后端代码**

```bash
go build -v ./cmd/paopao
```

- [ ] **Step 3: 编译前端代码**

```bash
cd admin
npm run build
```

---

## Task 9: 测试验证

- [ ] **Step 1: 启动服务并登录管理后台**
- [ ] **Step 2: 验证贴文管理功能**
  - 列表展示正确
  - 搜索功能正常
  - 编辑贴文状态正常
  - 删除贴文正常
- [ ] **Step 3: 验证话题管理功能**
  - 列表展示正确
  - 搜索功能正常
  - 创建新话题正常
  - 编辑话题正常
  - 删除话题正常

---

## Summary

完成以上所有任务后，管理后台将具备：
1. **贴文管理模块**：支持贴文列表查询、按条件搜索、编辑贴文状态（可见性/置顶/精华/锁定）、删除贴文
2. **话题管理模块**：支持话题列表查询、搜索、创建新话题、编辑话题、删除话题

两个模块都遵循现有代码架构和风格，使用统一的API结构和前端组件模式。
