package controller

import (
	"bluebell/logic"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// DeletePost 删除帖子功能
// @Summary 删除帖子功能
// @Description 删除帖子在数据库中数据
// @Tags 帖子相关接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param ID query int true "帖子ID"
// @Security ApiKeyAuth
// @Success 200 {object} models.ResponseSuccess "成功响应"
// @Success 400 {object} models.ResponseError "响应错误"
// @Success 500 {object} models.ResponseError "服务器错误"
// @Router /deleteV1 [delete]
func DeletePost(c *gin.Context) {
	// 获取参数与参数校验
	poststr := c.Query("ID")
	postID, err := strconv.ParseInt(poststr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		zap.L().Error("delete post invalid param", zap.Error(err))
		return
	}
	// 传入参数到logic层
	logic.DeletePost(postID)
	// 响应参数
	ResponseSuccess(c, CodeSuccess)

}

func DeleteAvatar(c *gin.Context) {

}
