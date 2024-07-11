package works

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resume/libs"
	"resume/models"
)

func Get(c *gin.Context) {
	s := libs.Context(c)

	var work []models.Works
	if err := s.Find(&work, "rid = ? and uid = ?", s.Resume.ID, s.User.ID).Error; err != nil {
		s.Msg(http.StatusNotFound, "没有查询到基础信息！")
		return
	}
	if len(work) <= 0 {
		if err := s.Find(&work, "rid = 0 and uid = 0").Error; err != nil {
			s.Msg(http.StatusNotFound, "没有查询到基础信息！")
			return
		}
	}
	s.Json(http.StatusOK, "数据查询成功", work)
}
