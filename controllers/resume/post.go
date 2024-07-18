package resume

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"resume/libs"
	"resume/models"
	"resume/types/enums"
	"strings"
)

func Post(c *gin.Context) {
	s := libs.Context(c)
	var resume models.Resumes

	if err := c.ShouldBind(&resume); err != nil {
		s.Msg(http.StatusInternalServerError, "服务器异常，请重试！")
		return
	}

	if len(strings.Trim(resume.Name, "")) <= 0 {
		s.Msg(http.StatusForbidden, "简历名称不能为空！")
		return
	}

	if err := s.First(&resume, "uid = ? AND name = ?", s.User.ID, resume.Name).Error; err == nil {
		s.Json(http.StatusForbidden, "简历名称已存在", false)
		return
	}
	resume.UID = s.User.ID

	vx32 := enums.Vx32

	if CopyData(s, &resume, vx32, vx32) {
		s.Json(http.StatusOK, "简历创建成功！", resume)
		return
	}
	s.Msg(http.StatusForbidden, "简历创建失败！")
}

func Copy(c *gin.Context) {
	s := libs.Context(c)
	var resume models.Resumes

	if err := c.ShouldBind(&resume); err != nil {
		s.Msg(http.StatusInternalServerError, "服务器异常，请重试！")
		return
	}

	if len(strings.Trim(resume.Name, "")) <= 0 {
		s.Msg(http.StatusForbidden, "简历名称不能为空！")
		return
	}
	resume.Name = fmt.Sprintf("%s_副本", resume.Name)

	if err := s.First(&resume, "uid = ? AND name = ?", s.User.ID, resume.Name).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			s.Msg(http.StatusForbidden, "简历复制失败！")
			return
		}
	}
	resume.UID = s.User.ID

	if CopyData(s, &resume, resume.ID, s.User.ID) {
		s.Json(http.StatusOK, "简历复制成功！", resume)
		return
	}
	s.Msg(http.StatusForbidden, "简历复制失败！")
}
