package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

// Skills 技能特长
func (d *InitData) Skills() {
	modelList := []models.Skills{
		{BaseModelWithRIDUID: ridUID(), Content: "<ol><li data-list=\"bullet\"><span class=\"ql-ui\" contenteditable=\"false\"></span>熟练掌握Office软件</li><li data-list=\"bullet\"><span class=\"ql-ui\" contenteditable=\"false\"></span>如果不想用文字展示，可以用标签形式体现</li></ol>"},
	}

	insertRecord[models.Skills]("技能特长", modelList, func(model models.Skills) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ?", model.RID, model.UID)
	})
}
