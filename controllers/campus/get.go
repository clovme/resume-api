package campus

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resume/libs"
	"resume/models"
	"resume/types"
)

func Get(c *gin.Context) {
	s := libs.Context(c)

	var campus []models.Campus
	if err := s.Find(&campus, "rid = ? AND uid = ?", s.Resume.ID, s.User.ID).Error; err != nil {
		s.Msg(http.StatusNotFound, "没有查询到校园经历数据！")
		return
	}
	if len(campus) <= 0 {
		if err := s.Find(&campus, "rid = ? AND uid = ?", types.Ox320, types.Ox320).Error; err != nil {
			s.Msg(http.StatusNotFound, "没有查询到校园经历数据！")
			return
		}
	}
	s.Json(http.StatusOK, "数据查询成功", campus)
}
