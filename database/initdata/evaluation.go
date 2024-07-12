package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

// Evaluation 自我评价
func (d *InitData) Evaluation() {
	modelList := []models.Evaluation{
		{BaseModelWithRIDUID: ridUID(), Content: "<ol><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>勤奋努力：</strong>我总是全力以赴，确保每项任务按时高质量完成，通过持续努力实现目标。</li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>创新思维：</strong>我具备创新思维，喜欢尝试新方法，提出创造性解决方案，在复杂情境中找到突破口。</li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>团队合作：</strong>我善于沟通和协作，倾听意见，分享经验，团队合作中不断学习和成长。</li></ol>"},
	}

	insertRecord[models.Evaluation]("自我评价", modelList, func(model models.Evaluation) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ?", model.RID, model.UID)
	})
}
