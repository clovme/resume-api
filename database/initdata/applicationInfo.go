package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

// ApplicationInfo 报考专业
func (d *InitData) ApplicationInfo() {
	modelList := []models.ApplicationInfo{
		{BaseModelWithRIDUID: ridUID()},
	}
	insertRecord[models.ApplicationInfo]("报考专业", modelList, func(model models.ApplicationInfo) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? AND uid = ?", model.RID, model.UID)
	})
}
