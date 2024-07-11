package initdata

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"resume/models"
	"time"
)

// Evaluation 自我评价
func (d *InitData) Evaluation() {
	modelList := []models.Evaluation{
		{RID: 0, UID: 0, Content: "<ol><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>勤奋努力：</strong>我总是全力以赴，确保每项任务按时高质量完成，通过持续努力实现目标。</li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>创新思维：</strong>我具备创新思维，喜欢尝试新方法，提出创造性解决方案，在复杂情境中找到突破口。</li><li data-list=\"ordered\"><span class=\"ql-ui\" contenteditable=\"false\"></span><strong>团队合作：</strong>我善于沟通和协作，倾听意见，分享经验，团队合作中不断学习和成长。</li></ol>", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	for _, model := range modelList {
		if err := d.Db.Where("rid = ? and uid = ?", model.RID, model.UID).First(&models.Evaluation{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 数据不存在，插入新记录
				if err := d.Db.Create(&model).Error; err != nil {
					log.Fatalln("[初始化]自我评价创建失败:", err)
				} else {
					log.Println("[初始化]自我评价创建创建成功:", model)
				}
			} else {
				log.Fatalln("[初始化]自我评价查询失败:", err)
			}
		}
	}
}
