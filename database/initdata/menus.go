package initdata

import (
	"gorm.io/gorm"
	"resume/models"
	"resume/types/enums"
)

// Menus 操作菜单
func (d *InitData) Menus() {
	modelList := []models.Menus{
		{BaseModelWithRIDUID: ridUID(), Title: "基础信息", Name: "BasicInfo", IsActivate: true, IsOption: false, IsChecked: true, IsPage: true, Sort: 0},
		{BaseModelWithRIDUID: ridUID(), Title: "求职岗位", Name: "JobPosition", IsActivate: false, IsOption: true, IsChecked: true, IsPage: false, Sort: 1},
		{BaseModelWithRIDUID: ridUID(), Title: "技能特长", Name: "SkillsExpertise", IsActivate: false, IsOption: true, IsChecked: false, IsPage: true, Sort: 2},
		{BaseModelWithRIDUID: ridUID(), Title: "工作经历", Name: "WorkExperience", IsActivate: false, IsOption: true, IsChecked: false, IsPage: true, Sort: 3},
		{BaseModelWithRIDUID: ridUID(), Title: "项目经验", Name: "ProjectExperience", IsActivate: false, IsOption: true, IsChecked: false, IsPage: true, Sort: 4},
		{BaseModelWithRIDUID: ridUID(), Title: "教育经历", Name: "EducationExperience", IsActivate: false, IsOption: true, IsChecked: false, IsPage: true, Sort: 5},
		{BaseModelWithRIDUID: ridUID(), Title: "自我评价", Name: "SelfEvaluation", IsActivate: false, IsOption: true, IsChecked: false, IsPage: true, Sort: 6},
		{BaseModelWithRIDUID: ridUID(), Title: "实习经验", Name: "InternshipExperience", IsActivate: false, IsOption: true, IsChecked: false, IsPage: true, Sort: 7},
		{BaseModelWithRIDUID: ridUID(), Title: "报考信息", Name: "ApplicationInformation", IsActivate: false, IsOption: true, IsChecked: false, IsPage: true, Sort: 8},
		{BaseModelWithRIDUID: ridUID(), Title: "校园经历", Name: "CampusExperience", IsActivate: false, IsOption: true, IsChecked: false, IsPage: true, Sort: 9},
		{BaseModelWithRIDUID: ridUID(), Title: "荣誉证书", Name: "HonorsCertificates", IsActivate: false, IsOption: true, IsChecked: false, IsPage: true, Sort: 10},
		{BaseModelWithRIDUID: ridUID(), Title: "兴趣爱好", Name: "HobbiesInterests", IsActivate: false, IsOption: true, IsChecked: false, IsPage: true, Sort: 11},
		{BaseModelWithRIDUID: ridUID(), Title: "求职意向", Name: "JobIntentions", IsActivate: false, IsOption: true, IsChecked: false, IsPage: true, Sort: 12},
	}

	insertRecord[models.Menus]("操作菜单", modelList, func(model models.Menus) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("title = ? and name = ? and rid = ? and uid = ?", model.Title, model.Name, enums.Vx32, enums.Vx32)
	})
}
