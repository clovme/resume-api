package resume

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"resume/controllers/resume/utils"
	"resume/libs"
	"resume/models"
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

	libs.Transaction(s, func(tx *gorm.DB) {
		// 创建简历模板
		if err := tx.Create(&resume).Error; err != nil {
			tx.Rollback()
			log.Println("简历模版创建失败", err.Error())
			s.Msg(http.StatusForbidden, "简历模版创建失败，请重试！")
			return
		}

		// 复制菜单
		utils.CopyMenus(tx, s, resume)
	})

	s.Json(http.StatusOK, "ok", resume)
}
