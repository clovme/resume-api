package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CorsMiddleware CORS 中间件
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许的来源，这里使用通配符允许所有来源访问，也可以指定具体的域名
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// 设置允许的请求方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// 设置允许的请求头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Browser-Type")

		// 如果请求方法是 OPTIONS，则直接返回 200 状态码
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// 继续处理请求
		c.Next()
	}
}
