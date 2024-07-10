package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resume/libs"
)

// NotFoundHandler is a middleware to handle 404 errors
func NotFoundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process request

		if c.Writer.Status() == http.StatusNotFound {
			s := libs.Context(c)
			s.Msg(http.StatusNotFound, "请输入正确的请求地址!")
			c.Abort()
		}
	}
}
