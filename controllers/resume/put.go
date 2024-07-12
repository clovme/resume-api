package resume

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resume/libs"
	"resume/models"
	"strings"
	"time"
)

// Put 更新简历
func Put(c *gin.Context) {
	s := libs.Context(c)

	var params models.Resumes
	if err := c.ShouldBind(&params); err != nil {
		s.Msg(http.StatusInternalServerError, "服务器异常，请重试！")
		return
	}

	if len(strings.Trim(params.Name, "")) <= 0 {
		s.Msg(http.StatusNotFound, "简历名称不能为空！")
		return
	}

	var resume models.Resumes
	if err := s.First(&resume, "uid = ? AND id = ?", s.User.ID, params.ID).Error; err != nil {
		s.Msg(http.StatusNotFound, "更新失败，更新的数据不存在！")
		return
	}
	resume.Name = params.Name
	resume.UpdatedAt = time.Now()
	result := s.Save(&resume)
	if result.Error != nil {
		s.Msg(http.StatusNotFound, "更新失败，更新的不存在！")
		return
	} else if result.RowsAffected == 0 {
		s.Msg(http.StatusNotFound, "没有找到符合条件的记录！")
		return
	}

	var resumes []models.Resumes
	if err := s.Find(&resumes, "uid = ?", s.User.ID).Error; err != nil {
		s.Msg(http.StatusNotFound, "暂无数据")
		return
	}
	s.Json(http.StatusOK, "数据更新成功", resumes)
}
