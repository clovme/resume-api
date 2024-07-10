package database

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"resume/models"
	"time"
)

// initMenusData 初始化菜单数据
func initMenusData(db *gorm.DB) {
	menus := []models.Menus{
		{RID: 0, UID: 0, Title: "基础信息", Name: "BasicInfo", IsActivate: true, IsOption: false, IsChecked: true, IsPage: true, Sort: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Title: "求职岗位", Name: "JobPosition", IsActivate: false, IsOption: true, IsChecked: true, IsPage: false, Sort: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Title: "技能特长", Name: "SkillsExpertise", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Title: "工作经历", Name: "WorkExperience", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Title: "项目经验", Name: "ProjectExperience", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Title: "教育经历", Name: "EducationExperience", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Title: "自我评价", Name: "SelfEvaluation", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Title: "实习经验", Name: "InternshipExperience", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Title: "报考信息", Name: "ApplicationInformation", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Title: "校园经历", Name: "CampusExperience", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Title: "荣誉证书", Name: "HonorsCertificates", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Title: "兴趣爱好", Name: "HobbiesInterests", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Title: "求职意向", Name: "JobIntentions", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	for _, menu := range menus {
		var data models.Menus

		if err := db.Where("title = ? and name = ? and rid = 0 and uid = 0", menu.Title, menu.Name).First(&data).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 数据不存在，插入新记录
				if err := db.Create(&menu).Error; err != nil {
					log.Println("[初始化]菜单创建失败:", err)
				} else {
					log.Println("[初始化]菜单创建创建成功:", data)
				}
			} else {
				log.Println("[初始化]菜单查询失败:", data)
			}
		}
	}
}
