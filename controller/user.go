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

// SignUpHandler 处理注册请求的函数
// @Summary 账户注册业务
// @Description 若账户未注册，则注册账户
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param
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

// LoginHandler 登录
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

	//queueName := "MyQueue"
	//messageType := "SMS"
	//cstLoc, _ := time.LoadLocation("Asia/Shanghai")
	//var token *dybaseapi.MessageTokenDTO
	//var expireTime time.Time
	//for {
	//	if token == nil || expireTime.Unix()-time.Now().Unix() < 2*60 {
	//		// 创建 API 请求并设置参数
	//		request := dybaseapi.CreateQueryTokenForMnsQueueRequest()
	//		request.MessageType = messageType
	//		request.QueueName = queueName
	//		// 发起请求并处理异常
	//		response, err := client.QueryTokenForMnsQueue(request)
	//		if err != nil {
	//			panic(err)
	//		}
	//
	//		token = &response.MessageTokenDTO
	//	}
	//	expireTime, err = time.ParseInLocation("2006-01-02 15:04:05", token.ExpireTime, cstLoc)
	//	if err != nil {
	//		panic(err)
	//	}
	//	mnsClient, err := mns.NewClientWithStsToken(
	//		"cn-hangzhou",
	//		token.AccessKeyId,
	//		token.AccessKeySecret,
	//		token.SecurityToken,
	//	)
	//	if err != nil {
	//		panic(err)
	//	}
	//	mnsRequest := mns.CreateBatchReceiveMessageRequest()
	//	mnsRequest.Domain = mnsDomain
	//	mnsRequest.QueueName = queueName
	//	mnsRequest.NumOfMessages = "10"
	//	mnsRequest.WaitSeconds = "5"
	//	mnsResponse, err := mnsClient.BatchReceiveMessage(mnsRequest)
	//	if err != nil {
	//		panic(err)
	//	}
	//	// fmt.Println(mnsResponse)
	//	receiptHandles := make([]string, len(mnsResponse.Message))
	//	for i, message := range mnsResponse.Message {
	//		messageBody, decodeErr := base64.StdEncoding.DecodeString(message.MessageBody)
	//		if decodeErr != nil {
	//			panic(decodeErr)
	//		}
	//		fmt.Println(string(messageBody))
	//		receiptHandles[i] = message.ReceiptHandle
	//	}
	//	if len(receiptHandles) > 0 {
	//		mnsDeleteRequest := mns.CreateBatchDeleteMessageRequest()
	//		mnsDeleteRequest.Domain = mnsDomain
	//		mnsDeleteRequest.QueueName = queueName
	//		mnsDeleteRequest.SetReceiptHandles(receiptHandles)
	//		//_, err = mnsClient.BatchDeleteMessage(mnsDeleteRequest) //取消注释将删除队列中的消息
	//		if err != nil {
	//			panic(err)
	//		}
	//	}
	//}
}
