package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

func (d *InitData) Slogan() {
	modelList := []models.Slogan{
		{BaseModelWithRIDUID: ridUID(), Title: "个人简历", Slogan: "----------------------------- 我命由我不由天 ---------------------------"},
	}

	insertRecord[models.Slogan]("简历标题", modelList, func(model models.Slogan) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ?", model.RID, model.UID)
	})
}
