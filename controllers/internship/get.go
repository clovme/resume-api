package internship

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resume/libs"
	"resume/models"
	"resume/types/enums"
)

func Get(c *gin.Context) {
	s := libs.Context(c)

	var internship []models.Internship
	if err := s.Find(&internship, "rid = ? AND uid = ?", s.Resume.ID, s.User.ID).Error; err != nil {
		s.Msg(http.StatusNotFound, "没有查询到数据！")
		return
	}
	if len(internship) <= 0 {
		if err := s.Find(&internship, "rid = ? AND uid = ?", enums.Vx32, enums.Vx32).Error; err != nil {
			s.Msg(http.StatusNotFound, "没有查询到数据！")
			return
		}
	}
	s.Json(http.StatusOK, "数据查询成功", internship)
}