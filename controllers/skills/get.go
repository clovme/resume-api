package skills

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resume/libs"
	"resume/models"
	"resume/types"
)

func Get(c *gin.Context) {
	s := libs.Context(c)

	var skills models.Skills
	if err := s.Find(&skills, "rid = ? AND uid = ?", s.Resume.ID, s.User.ID).Error; err != nil {
		s.Msg(http.StatusNotFound, "没有查询到技能特长数据！")
		return
	}
	if skills.ID == "" {
		if err := s.Find(&skills, "rid = ? AND uid = ?", types.Ox320, types.Ox320).Error; err != nil {
			s.Msg(http.StatusNotFound, "没有查询到技能特长数据！")
			return
		}
	}
	s.Json(http.StatusOK, "数据查询成功", skills)
}
