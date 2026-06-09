// Copyright 2024 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/core/admin"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

// GetKolProfile 获取KOL人物属性
func (s *AuthServant) GetKolProfile(c *gin.Context) {
	var req admin.H5KolProfileGetReq
	if err := c.ShouldBindQuery(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if req.UserID <= 0 {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails("userId无效"))
		return
	}
	profile, err := s.service.GetKolProfile(c.Request.Context(), req.UserID)
	if err != nil {
		// 没有记录时返回默认值，方便前端编辑
		profile = &admin.H5KolProfileItem{
			UserID: req.UserID,
			Height: "160cm", Weight: "44kg", Measurements: "84/58/84",
			SkinTone: "冷白病态肌", EyeColor: "酒红", Orientation: "偏双性恋（情感依赖向）",
			Preferences: "独占欲、暗调氛围、偏执温柔", FavoriteFoods: "黑森林、红酒、冷食",
			ClothingStyle: "黑裙、蕾丝、丝带、哥特风", MakeupStyle: "苍白底妆、下垂眼、暗红眼影、冷唇",
		}
	}
	app.NewResponse(c).ToResponse(profile)
}

// SaveKolProfile 保存KOL人物属性
func (s *AuthServant) SaveKolProfile(c *gin.Context) {
	var req admin.H5KolProfileSaveReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if req.UserID <= 0 {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails("userId无效"))
		return
	}
	if err := s.service.SaveKolProfile(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(nil)
}

// GetKolCategoryList 获取KOL分类列表
func (s *AuthServant) GetKolCategoryList(c *gin.Context) {
	list, err := s.service.GetKolCategoryList(c.Request.Context())
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(gin.H{"list": list})
}

// SaveKolCategory 保存KOL分类
func (s *AuthServant) SaveKolCategory(c *gin.Context) {
	var req admin.H5KolCategorySaveReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.SaveKolCategory(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(nil)
}

// DeleteKolCategory 删除KOL分类
func (s *AuthServant) DeleteKolCategory(c *gin.Context) {
	var req admin.H5KolCategoryDeleteReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.DeleteKolCategory(c.Request.Context(), req.ID); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(nil)
}

// GetKolManageList 获取KOL管理列表
func (s *AuthServant) GetKolManageList(c *gin.Context) {
	var req admin.H5KolManageListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	total, list, err := s.service.GetKolManageList(c.Request.Context(), &req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	if list == nil { list = []admin.H5KolManageItem{} }
	app.NewResponse(c).ToResponse(gin.H{"list": list, "total": total})
}

// AssignKolCategory 分配KOL分类
func (s *AuthServant) AssignKolCategory(c *gin.Context) {
	var req admin.H5KolAssignCategoryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.InvalidParams.WithDetails(err.Error()))
		return
	}
	if err := s.service.AssignKolCategory(c.Request.Context(), &req); err != nil {
		app.NewResponse(c).ToErrorResponse(xerror.ServerError.WithDetails(err.Error()))
		return
	}
	app.NewResponse(c).ToResponse(nil)
}
