package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"resume/libs"
	"resume/models"
	"time"
)

// SignIn 用户登录
func SignIn(c *gin.Context) {
	s := libs.Context(c)

	var user models.Users
	if err := c.ShouldBind(&user); err != nil {
		s.Msg(http.StatusBadRequest, "登录失败，请重试！")
		return
	}

	if err := s.First(&user, "username = ? and password = ?", user.Username, libs.MD5(user.Password)).Error; err != nil {
		s.Msg(http.StatusBadRequest, "登录失败，请重试！")
		return
	}

	if user.ExpiresAt.Before(time.Now()) {
		user.Token = libs.MD5(fmt.Sprintf("%s%d", user.Password, time.Now().Nanosecond()))
		user.ExpiresAt = time.Now().Add(24 * time.Hour)
		user.UpdatedAt = time.Now()
		log.Println("更新", user.Username, "Token", user)
		s.Save(&user)
	}

	s.Json(http.StatusOK, "登录成功", user)
}
