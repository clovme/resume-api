package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

// Honors 荣誉证书
func (d *InitData) Honors() {
	modelList := []models.Honors{
		{BaseModelWithRIDUID: ridUID()},
	}

	insertRecord[models.Honors]("荣誉证书", modelList, func(model models.Honors) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ?", model.RID, model.UID)
	})
}
