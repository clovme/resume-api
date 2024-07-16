package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

// BasicInfo 初始化基础信息表
func (d *InitData) BasicInfo() {
	modelList := []models.Basicinfo{
		{BaseModelWithRIDUID: ridUID(), IsAge: true, Gender: "不填", MaritalStatus: "不填", PoliticalStatus: "不填", WorkExperienceYears: "不填", StartDate: "不填"},
	}
	insertRecord[models.Basicinfo]("基础信息", modelList, func(model models.Basicinfo) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ?", model.RID, model.UID)
	})
}
