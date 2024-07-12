package initdata

import (
	"gorm.io/gorm"
	"resume/models"
)

// Hobbies 兴趣爱好
func (d *InitData) Hobbies() {
	modelList := []models.Hobbies{
		{BaseModelWithRIDUID: ridUID(), Content: "<ol><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>阅读与写作:</strong>喜欢阅读各类书籍，尤其喜欢文学作品，并享受写作带来的创作乐趣。</li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>旅行与摄影:</strong>热爱旅行，喜欢探索不同的地方和文化，同时用摄影记录旅途中的美好瞬间。</li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>音乐与乐器:</strong>对音乐充满热情，喜欢听各种风格的音乐，同时擅长弹奏吉他和钢琴，享受音乐带来的愉悦和放松。</li></ol>"},
	}

	insertRecord[models.Hobbies]("兴趣爱好", modelList, func(model models.Hobbies) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ?", model.RID, model.UID)
	})
}
