package initdata

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"resume/models"
	"resume/types"
)

func (d *InitData) Menus() {
	modelList := []models.Menus{
		{BaseModelWithRIDUID: d.ridUID(), Title: "基础信息", Name: "BasicInfo", IsActivate: true, IsOption: false, IsChecked: true, IsPage: true, Sort: 0},
		{BaseModelWithRIDUID: d.ridUID(), Title: "求职岗位", Name: "JobPosition", IsActivate: false, IsOption: true, IsChecked: true, IsPage: false, Sort: 1},
		{BaseModelWithRIDUID: d.ridUID(), Title: "技能特长", Name: "SkillsExpertise", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 2},
		{BaseModelWithRIDUID: d.ridUID(), Title: "工作经历", Name: "WorkExperience", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 3},
		{BaseModelWithRIDUID: d.ridUID(), Title: "项目经验", Name: "ProjectExperience", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 4},
		{BaseModelWithRIDUID: d.ridUID(), Title: "教育经历", Name: "EducationExperience", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 5},
		{BaseModelWithRIDUID: d.ridUID(), Title: "自我评价", Name: "SelfEvaluation", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 6},
		{BaseModelWithRIDUID: d.ridUID(), Title: "实习经验", Name: "InternshipExperience", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 7},
		{BaseModelWithRIDUID: d.ridUID(), Title: "报考信息", Name: "ApplicationInformation", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 8},
		{BaseModelWithRIDUID: d.ridUID(), Title: "校园经历", Name: "CampusExperience", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 9},
		{BaseModelWithRIDUID: d.ridUID(), Title: "荣誉证书", Name: "HonorsCertificates", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 10},
		{BaseModelWithRIDUID: d.ridUID(), Title: "兴趣爱好", Name: "HobbiesInterests", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 11},
		{BaseModelWithRIDUID: d.ridUID(), Title: "求职意向", Name: "JobIntentions", IsActivate: false, IsOption: true, IsChecked: true, IsPage: true, Sort: 12},
	}

	for _, model := range modelList {
		if err := d.Db.Where("title = ? and name = ? and rid = ? and uid = ?", model.Title, model.Name, types.Ox320, types.Ox320).First(&models.Menus{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 数据不存在，插入新记录
				if err := d.Db.Create(&model).Error; err != nil {
					log.Fatalln("[初始化]菜单创建失败:", err)
				} else {
					log.Println("[初始化]菜单创建创建成功:", model)
				}
			} else {
				log.Fatalln("[初始化]菜单查询失败:", err)
			}
		}
	}
}
