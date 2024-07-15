package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

// WorkExperience 工作经历
func (d *InitData) WorkExperience() {
	modelList := []models.Works{
		{
			BaseModelWithRIDUID: ridUID(), StartAt: "2019-08", EndAt: "2021-03", Name: "XXX科技有限公司", Title: "UI设计师", Content: "<ol><li data-list=\"bullet\"><span class=\"ql-ui\" contenteditable=\"false\"></span>GPA：X.X / 4.0（专业前X%）</li><li data-list=\"bullet\"><span class=\"ql-ui\" contenteditable=\"false\"></span>获得学校一等奖学金</li></ol>", ToNow: false, Sort: 0},
	}

	insertRecord[models.Works]("技能特长", modelList, func(model models.Works) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ? and name = ?", model.RID, model.UID, model.Name)
	})
}
