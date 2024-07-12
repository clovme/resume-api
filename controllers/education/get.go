package education

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resume/libs"
	"resume/models"
	"resume/types"
)

func Get(c *gin.Context) {
	s := libs.Context(c)

	var edu []models.Education
	if err := s.Find(&edu, "rid = ? AND uid = ?", s.Resume.ID, s.User.ID).Error; err != nil {
		s.Msg(http.StatusNotFound, "没有查询到教育经历数据！")
		return
	}
	if len(edu) <= 0 {
		if err := s.Find(&edu, "rid = ? AND uid = ?", types.Ox320, types.Ox320).Error; err != nil {
			s.Msg(http.StatusNotFound, "没有查询到教育经历数据！")
			return
		}
	}
	s.Json(http.StatusOK, "数据查询成功", edu)
}
