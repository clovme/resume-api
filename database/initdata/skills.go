package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

// Skills 技能特长
func (d *InitData) Skills() {
	modelList := []models.Skills{
		{BaseModelWithRIDUID: ridUID()},
	}

	insertRecord[models.Skills]("技能特长", modelList, func(model models.Skills) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ?", model.RID, model.UID)
	})
}
