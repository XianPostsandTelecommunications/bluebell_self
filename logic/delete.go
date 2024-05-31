package logic

import (
	"bluebell/dao/mysql"

	"go.uber.org/zap"
)

func DeletePost(postID int64) {
	// 把参数传递到dao层进行处理
	err := mysql.DeletePost(postID)
	if err != nil {
		zap.L().Error("mysql.DeletePost error", zap.Error(err))
		return
	}
}
