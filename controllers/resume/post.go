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

		// 01 复制菜单模板
		status = CopyData[models.Menus]("复制菜单模板", tx, []models.Menus{}, func(m *models.Menus) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ? AND title = ?", resume.ID, s.User.ID, m.Title)
		})

		// 02 复制标签模板
		status = CopyData[models.Tags]("复制标签模板", tx, []models.Tags{}, func(m *models.Tags) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ? AND name = ?", resume.ID, s.User.ID, m.Name)
		})

		// 03 复制基础信息模板
		status = CopyData[models.Basicinfo]("复制基础信息模板", tx, []models.Basicinfo{}, func(m *models.Basicinfo) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			m.Name = resume.Name

			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 04 复制技能特长模板
		status = CopyData[models.Skills]("复制技能特长模板", tx, []models.Skills{}, func(m *models.Skills) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("uid = ? AND rid = ?", resume.UID, resume.ID)
		})

		// 05 复制工作经历模板
		status = CopyData[models.Works]("复制工作经历模板", tx, []models.Works{}, func(m *models.Works) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 06 复制项目经验模板
		status = CopyData[models.Project]("复制项目经验模板", tx, []models.Project{}, func(m *models.Project) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 07 复制教育经历模板
		status = CopyData[models.Education]("复制教育经历模板", tx, []models.Education{}, func(m *models.Education) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 08 复制自我评价模板
		status = CopyData[models.Evaluation]("复制自我评价模板", tx, []models.Evaluation{}, func(m *models.Evaluation) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 09 复制实习经验模板
		status = CopyData[models.Internship]("复制实习经验模板", tx, []models.Internship{}, func(m *models.Internship) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 10 复制报考信息模板
		status = CopyData[models.ApplicationInfo]("复制报考信息模板", tx, []models.ApplicationInfo{}, func(m *models.ApplicationInfo) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 11 复制校园经历模板
		status = CopyData[models.Campus]("复制校园经历模板", tx, []models.Campus{}, func(m *models.Campus) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 12 复制荣誉证书模板
		status = CopyData[models.Honors]("复制荣誉证书模板", tx, []models.Honors{}, func(m *models.Honors) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 13 复制兴趣爱好模板
		status = CopyData[models.Hobbies]("复制兴趣爱好模板", tx, []models.Hobbies{}, func(m *models.Hobbies) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 14 复制求职意向模板
		status = CopyData[models.Intentions]("复制求职意向模板", tx, []models.Intentions{}, func(m *models.Intentions) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 15 复制课程成绩模板
		status = CopyData[models.CourseGrade]("复制课程成绩模板", tx, []models.CourseGrade{}, func(m *models.CourseGrade) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 16 复制简历标题模板
		status = CopyData[models.Slogan]("复制简历标题模板", tx, []models.Slogan{}, func(m *models.Slogan) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 17 复制简历标题模板
		status = CopyData[models.Setting]("复制简历设置模板", tx, []models.Setting{}, func(m *models.Setting) (string, *gorm.DB) {
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
