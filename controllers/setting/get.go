package setting

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"resume/libs"
	"resume/models"
)

func Get(c *gin.Context) {
	s := libs.Context(c)

	var setting models.Setting

	if err := s.Where("rid = ? AND uid = ?", s.Resume.ID, s.User.ID).First(&setting).Error; err != nil {
		log.Println("没有查询到数据，并且查询异常:", err.Error())
		s.Msg(http.StatusNotFound, "没有查询到数据！")
		return
	}

	s.Json(http.StatusOK, "数据查询成功", setting)
}
