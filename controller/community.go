package controller

import (
	"bluebell/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ---- 跟社区相关的 ----

// CommunityHandler 查询社区
// @Summary 查询社区
// @Description 查询所有社区
// @Tags 社区相关接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Security ApiKeyAuth
// @Success 200 {object} models.ResponseSuccess "成功响应"
// @Success 400 {object} models.ResponseError "响应错误"
// @Success 500 {object} models.ResponseError "服务器错误"
// @Router /community [get]
func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区（community_id, community_name) 以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 社区详情
// @Summary 社区详情
// @Description 用id得到社区详情
// @Tags 社区相关接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param id path int true "社区ID"
// @Security ApiKeyAuth
// @Success 200 {object} models.ResponseSuccess "成功响应"
// @Success 400 {object} models.ResponseError "响应错误"
// @Success 500 {object} models.ResponseError "服务器错误"
// @Router /community/:id [get]
func CommunityDetailHandler(c *gin.Context) {
	// 1. 获取社区id
	idStr := c.Param("id") // 获取URL参数
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 根据id获取社区详情
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}

func CommunityByName(c *gin.Context) {
	communityName := c.Param("name")
	// 2. 根据id获取社区详情
	data, err := logic.CommunityByName(communityName)
	if err != nil {
		zap.L().Error("logic.CommunityByName(id) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}
