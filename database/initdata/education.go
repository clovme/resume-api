package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

// Education 教育经历
func (d *InitData) Education() {
	modelList := []models.Education{
		{BaseModelWithRIDUID: ridUID()},
	}

	insertRecord[models.Education]("教育经历", modelList, func(model models.Education) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ? and name = ?", model.RID, model.UID, model.Name)
	})
}
