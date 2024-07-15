package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

// Honors 荣誉证书
func (d *InitData) Honors() {
	modelList := []models.Honors{
		{BaseModelWithRIDUID: ridUID(), Content: "<ol><li data-list=\"bullet\"><span class=\"ql-ui\" contenteditable=\"false\"></span>英语四级，听说读写能力良好，能快速浏览英语专业文件及书籍；</li><li data-list=\"bullet\"><span class=\"ql-ui\" contenteditable=\"false\"></span>通过全国计算机二级考试，熟练运用office软件。</li></ol>"},
	}

	insertRecord[models.Honors]("荣誉证书", modelList, func(model models.Honors) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ?", model.RID, model.UID)
	})
}
