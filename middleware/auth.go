package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resume/libs"
	"resume/models"
	"strings"
	"time"
)

// AuthMiddleware 是一个简单的认证中间件示例
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := libs.Context(c)
		var user models.Users

		authorization := c.GetHeader("Authorization")

		if !strings.Contains(authorization, "Bearer") {
			s.Msg(http.StatusUnauthorized, "登录授权过期或未登录!")
			c.Abort()
			return
		}

		token := strings.Replace(authorization, "Bearer ", "", 1)
		if err := s.First(&user, "token = ? AND expires_at >= ?", token, time.Now()).Error; err != nil {
			s.Msg(http.StatusUnauthorized, "登录授权过期或未登录!")
			c.Abort()
			return
		}
		rid := c.Query("rid")
		var resume models.Resumes
		if err := s.First(&resume, "id = ? AND uid = ?", rid, user.ID).Error; err != nil {
			if !strings.Contains(c.FullPath(), "resumes") {
				s.Msg(http.StatusNotFound, "当前简历不存在，请选择简历!")
				c.Abort()
				return
			}
		}

		c.Set("$", libs.HttpStatus{DB: s.DB, Context: c, User: user, Resume: resume, Embed: s.Embed})
		c.Next()
	}
}
