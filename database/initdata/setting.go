package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

func (d *InitData) Setting() {
	modelList := []models.Setting{
		{BaseModelWithRIDUID: ridUID(), FontFamily: "微软雅黑", FontSize: 13, Module: 20, Lines: 0.7, Page: 30},
	}

	insertRecord[models.Setting]("简历设置", modelList, func(model models.Setting) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ?", model.RID, model.UID)
	})
}
