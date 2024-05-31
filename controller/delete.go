package controller

import (
	"bluebell/logic"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

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
