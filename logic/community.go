package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"

	"go.uber.org/zap"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查数据库 查找到所有的community 并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}

func CommunityByName(communityName string) ([]*models.Post, error) {
	id := mysql.GetCommunityIDByName(communityName)
	post, err := mysql.GetPostListByCommunityIDs(id)
	if err != nil {
		zap.L().Error("CommunityByName mysql.GetPostListByCommunityIDs(id) ", zap.Error(err))
		return nil, err
	}
	return post, err
}

func GetComments(postID int64) ([]*models.Comment, error) {
	return redis.GetComments(postID)
}
