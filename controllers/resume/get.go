package resume

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resume/libs"
	"resume/models"
	"strings"
)

// Get 获取简历列表
func Get(c *gin.Context) {
	s := libs.Context(c)

	var resumes []models.Resumes
	if err := s.Find(&resumes, "uid = ?", s.User.ID).Error; err != nil {
		s.Msg(http.StatusNotFound, "没有查询得到数据")
		return
	}

	s.Json(http.StatusOK, "数据查询成功！", resumes)
}

// GetResumeId 判断简历是否存在
func GetResumeId(c *gin.Context) {
	s := libs.Context(c)

	rid := c.Query("rid")

	var resume models.Resumes
	if err := s.First(&resume, "id = ? AND uid = ?", rid, s.User.ID).Error; err != nil {
		s.Msg(http.StatusNotFound, "简历id不存在")
		return
	}
	s.Json(http.StatusOK, "数据查询成功", resume)
}

// CheckResume 检测简历名称是否存在
func CheckResume(c *gin.Context) {
	s := libs.Context(c)
	name := c.Query("name")

	if len(strings.Trim(name, "")) > 0 {
		var resume models.Resumes
		if err := s.First(&resume, "uid = ? AND name = ?", s.User.ID, name).Error; err == nil {
			s.Json(http.StatusBadRequest, "简历名称已存在", false)
		} else {
			s.Json(http.StatusOK, "简历名称不存在", true)
		}
		return
	}
	s.Json(http.StatusBadRequest, "简历名称不能为空", false)
}
