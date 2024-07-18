package resume

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"resume/libs"
	"resume/models"
)

func CopyDataItem[T any](msg string, tx *gorm.DB, models []T, rid string, uid string, setValues func(m *T) (string, *gorm.DB)) bool {
	// 判断默认数据是否存在
	if err := tx.Find(&models, "rid = ? AND uid = ?", rid, uid).Error; err != nil {
		log.Println(fmt.Sprintf("%s查询失败！", msg), err.Error())
		return false
	}

	flag := true
	// 添加数据
	for _, model := range models {
		// 设置uid和rid以及其他参数
		RName, Where := setValues(&model)

		if err := Where.First(&model).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 数据不存在，插入新记录
				if err := tx.Create(&model).Error; err != nil {
					tx.Rollback()
					log.Println(fmt.Sprintf("简历名称[%s]->[%s]复制失败:", RName, msg), err)
					flag = false
				} else {
					log.Println(fmt.Sprintf("简历名称[%s]->[%s]复制创建成功！", RName, msg))
				}
			} else {
				tx.Rollback()
				log.Println(fmt.Sprintf("简历名称[%s]->[%s]查询失败:", RName, msg), err.Error())
				flag = false
			}
		}
	}
	return flag
}

func CopyData(s libs.HttpStatus, resume *models.Resumes, rid, uid string) (status bool) {
	status = true
	libs.Transaction(s, func(tx *gorm.DB) {
		// 创建简历模板
		if err := tx.Create(&resume).Error; err != nil {
			tx.Rollback()
			log.Println("简历模版创建失败", err.Error())
			s.Msg(http.StatusForbidden, "简历模版创建失败，请重试！")
			return
		}

		// 01 复制菜单模板
		status = CopyDataItem[models.Menus]("复制菜单模板", tx, []models.Menus{}, rid, uid, func(m *models.Menus) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ? AND title = ?", resume.ID, s.User.ID, m.Title)
		})

		// 02 复制标签模板
		status = CopyDataItem[models.Tags]("复制标签模板", tx, []models.Tags{}, rid, uid, func(m *models.Tags) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ? AND name = ?", resume.ID, s.User.ID, m.Name)
		})

		// 03 复制基础信息模板
		status = CopyDataItem[models.Basicinfo]("复制基础信息模板", tx, []models.Basicinfo{}, rid, uid, func(m *models.Basicinfo) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			m.Name = resume.Name

			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 04 复制技能特长模板
		status = CopyDataItem[models.Skills]("复制技能特长模板", tx, []models.Skills{}, rid, uid, func(m *models.Skills) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("uid = ? AND rid = ?", resume.UID, resume.ID)
		})

		// 05 复制工作经历模板
		status = CopyDataItem[models.Works]("复制工作经历模板", tx, []models.Works{}, rid, uid, func(m *models.Works) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 06 复制项目经验模板
		status = CopyDataItem[models.Project]("复制项目经验模板", tx, []models.Project{}, rid, uid, func(m *models.Project) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 07 复制教育经历模板
		status = CopyDataItem[models.Education]("复制教育经历模板", tx, []models.Education{}, rid, uid, func(m *models.Education) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 08 复制自我评价模板
		status = CopyDataItem[models.Evaluation]("复制自我评价模板", tx, []models.Evaluation{}, rid, uid, func(m *models.Evaluation) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 09 复制实习经验模板
		status = CopyDataItem[models.Internship]("复制实习经验模板", tx, []models.Internship{}, rid, uid, func(m *models.Internship) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 10 复制报考信息模板
		status = CopyDataItem[models.ApplicationInfo]("复制报考信息模板", tx, []models.ApplicationInfo{}, rid, uid, func(m *models.ApplicationInfo) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 11 复制校园经历模板
		status = CopyDataItem[models.Campus]("复制校园经历模板", tx, []models.Campus{}, rid, uid, func(m *models.Campus) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 12 复制荣誉证书模板
		status = CopyDataItem[models.Honors]("复制荣誉证书模板", tx, []models.Honors{}, rid, uid, func(m *models.Honors) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 13 复制兴趣爱好模板
		status = CopyDataItem[models.Hobbies]("复制兴趣爱好模板", tx, []models.Hobbies{}, rid, uid, func(m *models.Hobbies) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 14 复制求职意向模板
		status = CopyDataItem[models.Intentions]("复制求职意向模板", tx, []models.Intentions{}, rid, uid, func(m *models.Intentions) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 15 复制课程成绩模板
		status = CopyDataItem[models.CourseGrade]("复制课程成绩模板", tx, []models.CourseGrade{}, rid, uid, func(m *models.CourseGrade) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 16 复制简历标题模板
		status = CopyDataItem[models.Slogan]("复制简历标题模板", tx, []models.Slogan{}, rid, uid, func(m *models.Slogan) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})

		// 17 复制简历标题模板
		status = CopyDataItem[models.Setting]("复制简历设置模板", tx, []models.Setting{}, rid, uid, func(m *models.Setting) (string, *gorm.DB) {
			m.RID = resume.ID
			m.UID = s.User.ID
			return resume.Name, tx.Where("rid = ? AND uid = ?", resume.ID, s.User.ID)
		})
	})
	return
}
