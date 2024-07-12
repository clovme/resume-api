package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

// Intentions 技能特长
func (d *InitData) Intentions() {
	modelList := []models.Intentions{
		{BaseModelWithRIDUID: ridUID(), Content: "求职岗位：Python开发工程师"},
	}

	insertRecord[models.Intentions]("技能特长", modelList, func(model models.Intentions) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ?", model.RID, model.UID)
	})
}
