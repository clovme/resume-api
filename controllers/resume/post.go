package resume

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
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

	status := true

	libs.Transaction(s, func(tx *gorm.DB) {
		// 创建简历模板
		if err := tx.Create(&resume).Error; err != nil {
			tx.Rollback()
			log.Println("简历模版创建失败", err.Error())
			s.Msg(http.StatusForbidden, "简历模版创建失败，请重试！")
			return
		}

		// 复制菜单模板
		status = CopyData[models.Menus]("复制菜单模板", tx, []models.Menus{}, func(m *models.Menus) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ? AND title = ?", resume.ID, s.User.ID, m.Title)
		})

		// 复制基础信息模板
		status = CopyData[models.Basicinfo]("复制基础信息模板", tx, []models.Basicinfo{}, func(m *models.Basicinfo) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			m.Name = resume.Name

			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 复制技能特长模板
		status = CopyData[models.Skills]("复制技能特长模板", tx, []models.Skills{}, func(m *models.Skills) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 复制标签模板
		status = CopyData[models.Tags]("复制标签模板", tx, []models.Tags{}, func(m *models.Tags) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ? AND name = ?", resume.ID, s.User.ID, m.Name)
		})

		// 复制工作经历模板
		status = CopyData[models.Works]("复制工作经历模板", tx, []models.Works{}, func(m *models.Works) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})
	})
	if status {
		s.Json(http.StatusOK, "简历创建成功！", resume)
		return
	}
	s.Msg(http.StatusForbidden, "简历创建失败！")
}
