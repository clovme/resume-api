package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

// ProjectExperience 项目经验
func (d *InitData) ProjectExperience() {
	modelList := []models.Project{
		{BaseModelWithRIDUID: ridUID()},
	}

	insertRecord[models.Project]("项目经验", modelList, func(model models.Project) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ? and name = ?", model.RID, model.UID, model.Name)
	})
}
