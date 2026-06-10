// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

// H5UserListReq 运维用户列表请求
type H5UserListReq struct {
	Page          int    `form:"page" json:"page"`
	PageSize      int    `form:"pageSize" json:"pageSize"`
	Nickname      string `form:"nickname" json:"nickname"`
	Username      string `form:"username" json:"username"`
	WalletAddress string `form:"walletAddress" json:"walletAddress"`
	Status        *int   `form:"status" json:"status"`
}

// H5UserListResp 运维用户列表响应
type H5UserListResp struct {
	List  []H5UserItem `json:"list"`
	Total int64        `json:"total"`
}

// H5UserItem 运维用户列表项
type H5UserItem struct {
	ID             int64  `json:"ID,string"`
	Nickname       string `json:"nickname"`
	Username       string `json:"username"`
	Phone          string `json:"phone"`
	WalletAddress  string `json:"walletAddress"`
	Bio            string `json:"bio"`
	Avatar         string `json:"avatar"`
	Status         int    `json:"status"`
	IsAdmin        bool   `json:"isAdmin"`
	IsKOL          bool   `json:"isKOL"`
	CreatedAt      string `json:"CreatedAt"`
	FollowingCount int64  `json:"followingCount"`
	FollowerCount  int64  `json:"followerCount"`
}

// H5UserGetReq 获取单个运维用户
type H5UserGetReq struct {
	ID int64 `form:"ID" json:"ID"`
}

// H5UserUpdateReq 更新运维用户
type H5UserUpdateReq struct {
	ID       int64  `json:"ID,string"`
	Nickname string `json:"nickname"`
	Bio      string `json:"bio"`
	Status   *int   `json:"status"`
	IsAdmin  *bool  `json:"isAdmin"`
	IsKOL    *bool  `json:"isKOL"`
}

// H5UserDeleteReq 删除运维用户
type H5UserDeleteReq struct {
	ID int64 `json:"ID,string"`
}

// H5PostListReq 贴文列表请求
type H5PostListReq struct {
	Page       int    `form:"page" json:"page"`
	PageSize   int    `form:"pageSize" json:"pageSize"`
	Keyword    string `form:"keyword" json:"keyword"`
	UserID     int64  `form:"userId" json:"userId"`
	Visibility *int   `form:"visibility" json:"visibility"`
}

// H5PostItem 贴文列表项
type H5PostItem struct {
	ID             int64       `json:"ID,string"`
	UserID          int64       `json:"userId,string"`
	User            *H5UserItem `json:"user"`
	Contents        []H5PostContent `json:"contents"`
	CommentCount    int64       `json:"commentCount"`
	CollectionCount int64       `json:"collectionCount"`
	UpvoteCount    int64       `json:"upvoteCount"`
	ShareCount     int64       `json:"shareCount"`
	Visibility     int         `json:"visibility"`
	IsTop          int8        `json:"isTop"`
	IsEssence      int8        `json:"isEssence"`
	IsLock         int8        `json:"isLock"`
	CreatedAt      string      `json:"CreatedAt"`
}

// H5PostContent 贴文内容项
type H5PostContent struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
	Sort   int    `json:"sort"`
}

// H5PostGetReq 获取单个贴文请求
type H5PostGetReq struct {
	ID int64 `form:"ID" json:"ID"`
}

// H5PostUpdateReq 更新贴文请求
type H5PostUpdateReq struct {
	ID         int64 `json:"ID,string"`
	Visibility int   `json:"visibility"`
	IsTop      int8  `json:"isTop"`
	IsEssence  int8  `json:"isEssence"`
	IsLock     int8  `json:"isLock"`
}

// H5PostDeleteReq 删除贴文请求
type H5PostDeleteReq struct {
	ID int64 `json:"ID,string"`
}

// H5TagListReq 话题列表请求
type H5TagListReq struct {
	Page     int    `form:"page" json:"page"`
	PageSize int    `form:"pageSize" json:"pageSize"`
	Keyword  string `form:"keyword" json:"keyword"`
}

// H5TagItem 话题列表项
type H5TagItem struct {
	ID        int64  `json:"ID,string"`
	Tag       string `json:"tag"`
	QuoteNum int64  `json:"quoteNum"`
	UserID    int64  `json:"userId,string"`
	CreatedAt string `json:"CreatedAt"`
}

// H5TagUpdateReq 更新话题请求
type H5TagUpdateReq struct {
	ID       int64  `json:"ID,string"`
	Tag      string `json:"tag"`
	QuoteNum int64  `json:"quoteNum"`
}

// H5TagDeleteReq 删除话题请求
type H5TagDeleteReq struct {
	ID int64 `json:"ID,string"`
}

// H5CommentListReq 评论列表请求
type H5CommentListReq struct {
	PostID int64 `form:"postId" json:"postId,string"`
	Page   int   `form:"page" json:"page"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

// H5CommentItem 评论列表项
type H5CommentItem struct {
	ID        int64  `json:"ID,string"`
	PostID    int64  `json:"postId,string"`
	UserID    int64  `json:"userId,string"`
	Nickname  string `json:"nickname"`
	Content   string `json:"content"`
	CreatedAt string `json:"CreatedAt"`
}

// H5CommentDeleteReq 删除评论请求
type H5CommentDeleteReq struct {
	ID int64 `json:"ID,string"`
}

// H5KolProfileItem KOL人物属性
type H5KolProfileItem struct {
	UserID        int64  `json:"userId,string"`
	Height        string `json:"height"`
	Weight        string `json:"weight"`
	Measurements  string `json:"measurements"`
	SkinTone      string `json:"skinTone"`
	EyeColor      string `json:"eyeColor"`
	Orientation   string `json:"orientation"`
	Preferences   string `json:"preferences"`
	FavoriteFoods string `json:"favoriteFoods"`
	ClothingStyle string `json:"clothingStyle"`
	MakeupStyle   string `json:"makeupStyle"`
	CategoryID    int64  `json:"categoryId,string"`
}

// H5KolProfileGetReq 获取KOL属性请求
type H5KolProfileGetReq struct {
	UserID int64 `form:"userId" json:"userId,string"`
}

// H5KolProfileSaveReq 保存KOL属性请求
type H5KolProfileSaveReq struct {
	UserID        int64  `json:"userId,string"`
	Height        string `json:"height"`
	Weight        string `json:"weight"`
	Measurements  string `json:"measurements"`
	SkinTone      string `json:"skinTone"`
	EyeColor      string `json:"eyeColor"`
	Orientation   string `json:"orientation"`
	Preferences   string `json:"preferences"`
	FavoriteFoods string `json:"favoriteFoods"`
	ClothingStyle string `json:"clothingStyle"`
	MakeupStyle   string `json:"makeupStyle"`
	CategoryID    int64  `json:"categoryId,string"`
}

// H5KolCategoryItem KOL分类
type H5KolCategoryItem struct {
	ID        int64  `json:"ID,string"`
	Name      string `json:"name"`
	Sort      int    `json:"sort"`
	UserCount int64  `json:"userCount"`
}

// H5KolCategorySaveReq 保存KOL分类
type H5KolCategorySaveReq struct {
	ID   int64  `json:"ID,string"`
	Name string `json:"name"`
	Sort int    `json:"sort"`
}

// H5KolCategoryDeleteReq 删除KOL分类
type H5KolCategoryDeleteReq struct {
	ID int64 `json:"ID,string"`
}

// H5KolManageListReq KOL管理列表请求
type H5KolManageListReq struct {
	Page       int   `form:"page" json:"page"`
	PageSize   int   `form:"pageSize" json:"pageSize"`
	CategoryID int64 `form:"categoryId" json:"categoryId,string"`
	Keyword    string `form:"keyword" json:"keyword"`
}

// H5KolManageItem KOL管理列表项
type H5KolManageItem struct {
	ID           int64  `json:"ID,string"`
	Nickname     string `json:"nickname"`
	Username     string `json:"username"`
	Avatar       string `json:"avatar"`
	CategoryID   int64  `json:"categoryId,string"`
	CategoryName string `json:"categoryName"`
	Height       string `json:"height"`
	Weight        string `json:"weight"`
	Measurements  string `json:"measurements"`
	WalletAddress string `json:"walletAddress"`
	CreatedAt     string `json:"CreatedAt"`
}

// H5KolAssignCategoryReq 分配KOL分类
type H5KolAssignCategoryReq struct {
	UserID     int64 `json:"userId,string"`
	CategoryID int64 `json:"categoryId,string"`
}

// ExploreKolCategoryResp 探索页KOL分类响应
type ExploreKolCategoryResp struct {
	Categories []ExploreKolCategory `json:"categories"`
}

type ExploreKolCategory struct {
	ID    int64            `json:"id,string"`
	Name  string           `json:"name"`
	Users []ExploreKolUser `json:"users"`
}

type ExploreKolUser struct {
	ID       int64  `json:"id,string"`
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

// H5SysMsgListReq 系统消息列表请求
type H5SysMsgListReq struct {
	Page     int `form:"page" json:"page"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

// H5SysMsgItem 系统消息列表项
type H5SysMsgItem struct {
	ID         int64  `json:"ID,string"`
	SenderID   int64  `json:"senderId,string"`
	ReceiverID int64  `json:"receiverId,string"`
	Brief      string `json:"brief"`
	Content    string `json:"content"`
	IsRead     int8   `json:"isRead"`
	CreatedAt  string `json:"CreatedAt"`
}

// H5SysMsgCreateReq 创建系统消息
type H5SysMsgCreateReq struct {
	ReceiverID string `json:"receiverId"` // 字符串类型，避免大ID精度丢失，"0"表示全员
	Brief      string `json:"brief"`
	Content    string `json:"content"`
}

// H5SysMsgDeleteReq 删除系统消息
type H5SysMsgDeleteReq struct {
	ID int64 `json:"ID,string"`
}
