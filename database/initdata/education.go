package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

// Education 教育经历
func (d *InitData) Education() {
	var modelList = []models.Education{
		{BaseModelWithRIDUID: ridUID(), StartAt: "2015-09", EndAt: "2019-07", Name: "师范大学", Major: "计算机科学与技术", Content: "<ol><li data-list=\"bullet\"><span class=\"ql-ui\" contenteditable=\"false\"></span>GPA：X.X / 4.0（专业前X%）</li><li data-list=\"bullet\"><span class=\"ql-ui\" contenteditable=\"false\"></span>获得学校一等奖学金</li></ol>", ToNow: false, Degree: "不填", Sort: 0},
	}
	insertRecord[models.Education]("教育经历", modelList, func(model models.Education) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ? and name = ?", model.RID, model.UID, model.Name)
	})
}
