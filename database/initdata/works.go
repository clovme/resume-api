package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

// WorkExperience 工作经历
func (d *InitData) WorkExperience() {
	modelList := []models.Works{
		{BaseModelWithRIDUID: ridUID()},
	}

	insertRecord[models.Works]("技能特长", modelList, func(model models.Works) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ? and name = ?", model.RID, model.UID, model.Name)
	})
}
