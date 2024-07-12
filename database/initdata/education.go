package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

// Education 教育经历
func (d *InitData) Education() {
	modelList := []models.Education{
		{BaseModelWithRIDUID: ridUID(), StartAt: "2017-06", EndAt: "2019-08", Name: "月球研发学院", Major: "月地开发技术", ToNow: false, Degree: "硕士", Sort: 0},
	}

	insertRecord[models.Education]("教育经历", modelList, func(model models.Education) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ? and name = ?", model.RID, model.UID, model.Name)
	})
}
