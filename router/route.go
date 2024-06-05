package router

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"net/http"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "bluebell/docs"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// SetupRouter 路由
func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	//r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.RateLimitMiddleware(2*time.Second, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.Cors())

	r.LoadHTMLFiles("./templates/index.html")
	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		// 注册
		v1.POST("/signup", controller.SignUpHandler)
		// 登录
		v1.GET("/login", controller.LoginHandler)
		// 短信验证码登录
		v1.POST("/loginSMS", controller.LoginSMSHandler)
		// 根据时间或分数获取帖子列表
		v1.GET("/posts2", controller.GetPostListHandler2)
		v1.GET("/posts", controller.GetPostListHandler)
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/id/:id", controller.CommunityDetailHandler)
		v1.GET("/community/name/:name", controller.CommunityByName)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/select", controller.GetPostBySelect)
	}

	auth := v1.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())
	{
		// 帖子评论
		// auth.POST("/comment", controller.PostComment)
		// 用户头像上传
		auth.POST("/user/:user_id/avatar", controller.PostAvatar)
		// 发布帖子
		auth.POST("/post", controller.CreatePostHandler)
		// 投票
		auth.POST("/vote", controller.PostVoteController)
		// 个人页面
		auth.GET("/userPage", controller.GetUserPage)
		// 删除帖子
		auth.DELETE("/deleteV1", controller.DeletePost)
	}

	manager := r.Group("/manager", middlewares.JWTAuthMiddleware(), middlewares.AuthManager())
	{
		// 删除帖子
		manager.DELETE("/deleteRoot", controller.DeletePost)
		// 置顶帖子
		//manager.POST("/postTop", controller.PostTop)
		// 删除用户头像
		//manager.DELETE("/deleteAvatar", controller.DeleteAvatar)
	}
	pprof.Register(r) // 注册pprof相关路由

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
