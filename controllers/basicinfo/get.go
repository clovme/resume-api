package basicinfo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resume/libs"
	"resume/models"
)

func Get(c *gin.Context) {
	s := libs.Context(c)

	var basic models.BasicInfo
	if err := s.Find(&basic, "rid = ? and uid = ?", s.Resume.ID, s.User.ID).Error; err != nil {
		s.Msg(http.StatusNotFound, "没有查询到基础信息！")
		return
	}
	if basic.ID == 0 {
		if err := s.Find(&basic, "rid = 0 and uid = 0").Error; err != nil {
			s.Msg(http.StatusNotFound, "没有查询到基础信息！")
			return
		}
	}
	s.Json(http.StatusOK, "数据查询成功", basic)
}
