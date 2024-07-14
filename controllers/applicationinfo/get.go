package applicationinfo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resume/libs"
	"resume/models"
)

func Get(c *gin.Context) {
	s := libs.Context(c)

	var info models.ApplicationInfo
	if err := s.Find(&info, "rid = ? AND uid = ?", s.Resume.ID, s.User.ID).Error; err != nil {
		s.Msg(http.StatusNotFound, "没有查询到数据！")
		return
	}
	s.Json(http.StatusOK, "数据查询成功", info)
}

func GradeGet(c *gin.Context) {
	s := libs.Context(c)
	var grade []models.CourseGrade

	if err := s.Find(&grade, "rid = ? AND uid = ?", s.Resume.ID, s.User.ID).Error; err != nil {
		s.Msg(http.StatusNotFound, "没有查询到数据！")
		return
	}

	s.Json(http.StatusOK, "数据查询成功", grade)
}
