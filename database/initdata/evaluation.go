package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

// Evaluation 自我评价
func (d *InitData) Evaluation() {
	modelList := []models.Evaluation{
		{BaseModelWithRIDUID: ridUID()},
	}

	insertRecord[models.Evaluation]("自我评价", modelList, func(model models.Evaluation) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ?", model.RID, model.UID)
	})
}
