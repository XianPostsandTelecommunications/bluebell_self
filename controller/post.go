package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"bluebell/pkg/badword"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	// swagger 嵌入文件
)

// CreatePostHandler 创建帖子的处理函数
// @Summary 创建帖子的处理函数
// @Description 创建帖子
// @Tags 帖子相关接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object body models.Post true "帖子信息"
// @Security ApiKeyAuth
// @Success 200 {object} models.ResponseSuccess "成功响应"
// @Success 400 {object} models.ResponseError "响应错误"
// @Success 500 {object} models.ResponseError "服务器错误"
// @Router /post [post]
func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数及参数的校验
	//c.ShouldBindJSON()  // validator --> binding tag
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) error", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 初始化敏感词树
	SensitiveTree := badword.NewTrieV1()
	// 初始化敏感词库
_:
	SensitiveTree.LoadWordsFromFile("./pkg/sensitivewords/广告.txt ")
_:
	SensitiveTree.LoadWordsFromFile("./pkg/sensitivewords/政治类.txt ")
_:
	SensitiveTree.LoadWordsFromFile("./pkg/sensitivewords/涉枪涉爆违法信息关键词.txt ")
_:
	SensitiveTree.LoadWordsFromFile("./pkg/sensitivewords/网址.txt ")
_:
	SensitiveTree.LoadWordsFromFile("./pkg/sensitivewords/色情类.txt ")
	if err := SensitiveTree.Contains(p.Content); err != false {
		ResponseError(c, CodeServerBusy)
		return
	}
	// 从 c 取到当前发请求的用户的ID
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 2. 创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. 返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获取帖子详情的处理函数
// @Summary 获取帖子详情的处理函数
// @Description 用id得到帖子
// @Tags 帖子相关接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param id path int true "帖子ID"
// @Security ApiKeyAuth
// @Success 200 {object} models.ResponseSuccess "成功响应"
// @Success 400 {object} models.ResponseError "响应错误"
// @Success 500 {object} models.ResponseError "服务器错误"
// @Router /post/:id [get]
func GetPostDetailHandler(c *gin.Context) {
	// 1. 获取参数（从URL中获取帖子的id）
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 2. 根据id取出帖子数据（查数据库）
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler 获取帖子列表的处理函数
func GetPostListHandler(c *gin.Context) {
	// 获取分页参数
	page, size := getPageInfo(c)
	// 获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
	// 返回响应
}

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} models.ResponseSuccess
// @Success 400 {object} models.ResponseError "响应错误"
// @Success 500 {object} models.ResponseError "服务器错误"
// @Router /posts2 [get]
func GetPostListHandler2(c *gin.Context) {
	// GET请求参数(query string)：/api/v1/posts2?page=1&size=10&order=time
	// 初始化结构体时指定初始参数
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime, // magic string
	}
	//c.ShouldBind()  根据请求的数据类型选择相应的方法去获取数据
	//c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostListNew(p) // 更新：合二为一
	fmt.Println(data)
	// 获取数据
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
	// 返回响应
}

// PostAvatar 提交头像的处理函数
// @Summary 提交头像的处理函数
// @Description 提交头像地址到数据库
// @Tags 用户相关接口(api分组展示使用的)
// @Accept multipart/form-data
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param user_id formData int true "用户ID"
// @Param avatar formData file true "头像文件"
// @Security ApiKeyAuth
// @Success 200 {object} models.ResponseSuccess "成功响应"
// @Success 400 {object} models.ResponseError "响应错误"
// @Success 500 {object} models.ResponseError "服务器错误"
// @Router /user/:user_id/avatar [post]
func PostAvatar(c *gin.Context) {
	// 获取参数，id与file
	userID := c.PostForm("user_id")
	file, err := c.FormFile("avatar")
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		zap.Error(err)
	}
	// 传递参数到logic层
	logic.PostAvatar(c, userID, file)
}

// PostTop 置顶帖子功能实现
func PostTop(c *gin.Context) {

}

// GetPostBySelect 使用查询拿到帖子
// @Summary 使用查询拿到帖子
// @Description 用帖子标题得到帖子
// @Tags 帖子相关接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param title path string true "帖子标题"
// @Security ApiKeyAuth
// @Success 200 {object} models.ResponseSuccess "成功响应"
// @Success 400 {object} models.ResponseError "响应错误"
// @Success 500 {object} models.ResponseError "服务器错误"
// @Router /select [get]
func GetPostBySelect(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		ResponseError(c, CodeInvalidParam)
		return
	}
	post, err := logic.GetPostByTitle(title)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	ResponseSuccess(c, post)
}

// PostComment 发布评论
func PostComment(c *gin.Context) {
	comment := &models.Comment{
		Time: time.Now().Unix(),
	}
	if err := c.ShouldBindJSON(comment); err != nil {
		zap.L().Error("PostComment with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	logic.PostComment(comment)
}

//// 根据社区去查询帖子列表
//func GetCommunityPostListHandler(c *gin.Context) {
//	// 初始化结构体时指定初始参数
//	p := &models.ParamCommunityPostList{
//		ParamPostList: &models.ParamPostList{
//			Page:  1,
//			Size:  10,
//			Order: models.OrderTime,
//		},
//	}
//	//c.ShouldBind()  根据请求的数据类型选择相应的方法去获取数据
//	//c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据
//	if err := c.ShouldBindQuery(p); err != nil {
//		zap.L().Error("GetCommunityPostListHandler with invalid params", zap.Error(err))
//		ResponseError(c, CodeInvalidParam)
//		return
//	}
//
//	// 获取数据
//	data, err := logic.GetCommunityPostList(p)
//	if err != nil {
//		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//	ResponseSuccess(c, data)
//}
