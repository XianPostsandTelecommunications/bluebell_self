package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"bluebell/setting"
	"strconv"

	"encoding/json"
	"errors"
	"fmt"

	"math/rand"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

const (
	mnsDomain = "1943695596114318.mns.cn-hangzhou.aliyuncs.com"
)

// SignUpHandler 用户注册接口
// @Summary 用户注册接口
// @Description 注册用户账户
// @Tags 用户相关接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param ParamSignUp body models.ParamSignUp true "用户注册参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseSuccess "成功响应"
// @Success 400 {object} _ResponseError "响应错误"
// @Success 500 {object} _ResponseError "服务器错误"
// @Router /community/:id [post]
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 用户登录接口
// @Summary 用户登录接口
// @Description 登录用户账户
// @Tags 用户相关接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param ParamLogin body models.ParamLogin true "用户登录参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseSuccess "成功响应"
// @Success 400 {object} _ResponseError "响应错误"
// @Success 500 {object} _ResponseError "服务器错误"
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	// 1.获取请求参数及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2.业务逻辑处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}

	// 3.返回响应
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID), // id值大于1<<53-1  int64类型的最大值是1<<63-1
		"user_name": user.Username,
		"token":     user.Token,
	})
}

// GetUserPage 用户界面
func GetUserPage(c *gin.Context) {
	userIDStr := c.Query("UserID")
	postIDStr := c.Query("ID")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		zap.L().Error("get userID detail with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		zap.L().Error("get postID detail with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	user, post := logic.GetUserPage(userID, postID)
	UserPage := models.UserPage{
		User: user,
		Post: post,
	}
	ResponseSuccess(c, UserPage)
}

// LoginSMSHandler 使用短信验证码登录
// @Summary 短信验证码接口
// @Description 使用阿里云短信服务SDK发送短信验证码
// @Tags 短信相关接口(api分组展示使用的)
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer JWT"
// @Param ParamLogin body models.ParamLogin true "用户登录参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponseSuccess "成功响应"
// @Success 400 {object} _ResponseError "响应错误"
// @Success 500 {object} _ResponseError "服务器错误"
// @Router /loginSMS [post]
func LoginSMSHandler(c *gin.Context) {
	// 获取参数
	phone := c.Query("phone")
	ms := setting.SMSConfig{}
	// 生成六位数验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	// 将验证码保存在数据库
	// 调用阿里云短信服务SDK
	//if err:=endpoints.AddEndpointMapping("cn-hangzhou", "Dybaseapi", "dybaseapi.aliyuncs.com");err!=nil {
	//	ResponseError(c,CodeServerBusy)
	//	zap.L().Error("endpoints.AddEndpointMapping error",zap.Error(err))
	//}
	// 阿里云账号AccessKey拥有所有API的访问权限，建议您使用RAM用户进行API访问或日常运维。
	// 强烈建议不要把AccessKey ID和AccessKey Secret保存到工程代码里，否则可能导致AccessKey泄露，威胁您账号下所有资源的安全。
	// 本示例以把AccessKey ID和AccessKey Secret保存在环境变量为例说明，来实现API访问的身份验证。
	// 创建client实例
	client, err := dysmsapi.NewClientWithAccessKey(
		ms.RegionID,  // 您的可用区ID
		ms.AppKey,    // 您的AccessKey ID
		ms.AppSecret) // 您的AccessKey Secret
	if err != nil {
		// 异常处理
		ResponseError(c, CodeServerBusy)
		panic(err)
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phone
	request.SignName = ms.SignName
	request.TemplateCode = ms.TemplateCode
	par, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(par)
	response, err := client.SendSms(request)
	fmt.Println(response)
	if err != nil {
		zap.L().Error("client.SendSms(request) error", zap.Error(err))
	}
}
