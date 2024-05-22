package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 拿到请求方法，判断是否为options请求
		method := c.Request.Method
		// 拿到请求头
		requestHeader := c.Request.Header.Get("Access-Control-Request-Headers") //用户端请求，在下面全部允许
		// 拿到origin
		origin := c.Request.Header.Get("origin")
		//// 循环拿到请求头中元素
		//var headerkeys []string
		//for key, _ := range c.Request.Header {
		//	headerkeys = append(headerkeys, key)
		//}
		//// 分割
		//headerStr := strings.Join(headerkeys, ",")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			// 设置允许的请求方法
			c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS,DELETE,PUT")
			c.Header("Access-Control-Allow-Headers", requestHeader)
			c.Header("Access-Control-Max-Age", "172800")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request")
		}
		c.Next()
	}
}
