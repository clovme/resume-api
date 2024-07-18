package users

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"resume/libs"
	"resume/models"
	"time"
)

// SignIn 用户登录
func SignIn(c *gin.Context) {
	s := libs.Context(c)

	var fUser FUsers
	if err := c.ShouldBind(&fUser); err != nil {
		s.Msg(http.StatusBadRequest, "登录失败，请重试！")
		return
	}

	var user models.Users
	if err := s.First(&user, "username = ? AND password = ?", fUser.Username, libs.MD5(fUser.Password)).Error; err != nil {
		s.Msg(http.StatusBadRequest, "登录失败，请重试！")
		return
	}

	if user.ExpiresAt.Before(time.Now().Add(12 * time.Hour)) {
		user.Token = libs.MD5(uuid.New().String())
		user.ExpiresAt = time.Now().Add(24 * time.Hour)
		user.UpdatedAt = time.Now()
		log.Println("更新", user.Username, "Token", user)
		s.Save(&user)
	}

	s.Json(http.StatusOK, "登录成功", user)
}

func SignOut(c *gin.Context) {
	s := libs.Context(c)

	if result := s.Model(&models.Users{}).Where("id = ? AND token = ?", s.User.ID, s.User.Token).Update("expires_at", time.Now()); result.Error != nil {
		log.Println("用户名称更新失败！", result.Error)
		s.Msg(http.StatusBadRequest, "注销失败！")
		return
	}
	s.Msg(http.StatusOK, "注销成功！")
}

func Regedit(c *gin.Context) {
	s := libs.Context(c)

	var fUser FUsers
	if err := c.ShouldBind(&fUser); err != nil {
		log.Println("用户注册->参数反序列化失败！")
		s.Msg(http.StatusBadRequest, "注册失败，请重试！")
		return
	}

	if fUser.Password != fUser.ConfirmPassword {
		s.Msg(http.StatusBadRequest, "两次输入密码不一致，请重试！")
		return
	}

	if result := s.Find(&models.Users{}, "username = ?", fUser.Username); result.Error != nil {
		log.Println("用户注册->用户查询:", result.Error)
		s.Msg(http.StatusBadRequest, "系统错误，请重试！")
		return
	} else if result.RowsAffected > 0 {
		s.Msg(http.StatusBadRequest, "用户名已存在，请重试！")
		return
	}

	user := models.Users{
		Nickname:  "个人简历",
		Username:  fUser.Username,
		Password:  libs.MD5(fUser.Password),
		Token:     libs.MD5(uuid.New().String()),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	if result := s.Create(&user); result.Error != nil {
		log.Println("用户注册->数据创建失败:", result.Error)
		s.Msg(http.StatusBadRequest, "注册失败，请重试！")
		return
	}
	s.Json(http.StatusOK, "注册成功，请重试！", user)
}
