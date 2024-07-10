package resume

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"resume/libs"
	"resume/models"
	"strings"
)

func Delete(c *gin.Context) {
	s := libs.Context(c)
	rid := c.Query("rid")
	var resume models.Resumes

	if len(strings.Trim(rid, "")) <= 0 {
		s.Msg(http.StatusNotFound, "简历标识不能为空！")
		return
	}

	if err := s.First(&resume, "uid = ? and id = ?", s.User.ID, rid).Error; err != nil {
		s.Msg(http.StatusNotFound, "删除失败，简历名称不存在！")
		return
	}

	libs.Transaction(s, func(tx *gorm.DB) {
		if !libs.DeleteData(s, "简历", s.Where("id = ? and uid = ?", rid, s.User.ID).Delete(&models.Resumes{})) {
			return
		}

		if !libs.DeleteData(s, "菜单", s.Where("uid = ? and rid = ?", resume.UID, resume.ID).Delete(&models.Menus{})) {
			return
		}

		var resumes []models.Resumes
		if err := s.Find(&resumes, "uid = ?", s.User.ID).Error; err != nil {
			s.Msg(http.StatusNotFound, "暂无数据")
			return
		}
		s.Json(http.StatusOK, "数据删除成功", resumes)
	})
}
