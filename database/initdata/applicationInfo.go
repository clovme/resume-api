package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

// ApplicationInfo 报考专业
func (d *InitData) ApplicationInfo() {
	modelList := []models.ApplicationInfo{
		{BaseModelWithRIDUID: ridUID(), Name: "哈尔滨佛学院", Major: "哈佛", CName: "上学期"},
	}
	insertRecord[models.ApplicationInfo]("报考专业", modelList, func(model models.ApplicationInfo) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? AND uid = ?", model.RID, model.UID)
	})
}

// CourseGrade 报考专业分数列表
func (d *InitData) CourseGrade() {
	modelList := []models.CourseGrade{
		{BaseModelWithRIDUID: ridUID(), Name: "语文", Score: "80"},
		{BaseModelWithRIDUID: ridUID(), Name: "数学", Score: "80"},
		{BaseModelWithRIDUID: ridUID(), Name: "英文", Score: "80"},
		{BaseModelWithRIDUID: ridUID(), Name: "物理", Score: "80"},
		{BaseModelWithRIDUID: ridUID(), Name: "化学", Score: "80"},
		{BaseModelWithRIDUID: ridUID(), Name: "历史", Score: "80"},
	}
	insertRecord[models.CourseGrade]("报考专业", modelList, func(model models.CourseGrade) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? AND uid = ? AND name = ?", model.RID, model.UID, model.Name)
	})
}
