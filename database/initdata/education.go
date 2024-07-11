package initdata

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"resume/models"
	"time"
)

// Education 教育经历
func (d *InitData) Education() {
	modelList := []models.Education{
		{RID: 0, UID: 0, StartAt: "2017-06", EndAt: "2019-08", Name: "月球研发学院", Major: "月地开发技术", ToNow: false, Degree: "硕士", Sort: 0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	for _, model := range modelList {
		if err := d.Db.Where("rid = ? and uid = ? and name = ?", model.RID, model.UID, model.Name).First(&models.Education{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 数据不存在，插入新记录
				if err := d.Db.Create(&model).Error; err != nil {
					log.Fatalln("[初始化]教育经历创建失败:", err)
				} else {
					log.Println("[初始化]教育经历创建创建成功:", model)
				}
			} else {
				log.Fatalln("[初始化]教育经历查询失败:", err)
			}
		}
	}
}
