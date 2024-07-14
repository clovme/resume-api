package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

// Hobbies 兴趣爱好
func (d *InitData) Hobbies() {
	modelList := []models.Hobbies{
		{BaseModelWithRIDUID: ridUID()},
	}

	insertRecord[models.Hobbies]("兴趣爱好", modelList, func(model models.Hobbies) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ?", model.RID, model.UID)
	})
}
