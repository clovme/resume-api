package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

// Evaluation 自我评价
func (d *InitData) Evaluation() {
	modelList := []models.Evaluation{
		{BaseModelWithRIDUID: ridUID(), Content: "自我介绍篇幅不用太长，注意结合简历整体的美观度。自我评价应做到突出自身符合目标岗位要求的“卖点”，避免使用过多的形容词。"},
	}

	insertRecord[models.Evaluation]("自我评价", modelList, func(model models.Evaluation) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ?", model.RID, model.UID)
	})
}
