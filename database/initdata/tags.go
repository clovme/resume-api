package initdata

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"resume/models"
	"resume/types/enums"
	"time"
)

func (d *InitData) Tags() {
	modelList := []models.Tags{
		{RID: 0, UID: 0, Type: enums.Skills, Name: "Office软件", IsChecked: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Type: enums.Skills, Name: "沟通能力", IsChecked: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Type: enums.Skills, Name: "口才", IsChecked: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Type: enums.Skills, Name: "文案编辑", IsChecked: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Type: enums.Skills, Name: "数据分析", IsChecked: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Type: enums.Skills, Name: "推广运营", IsChecked: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Type: enums.Skills, Name: "产品设计", IsChecked: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Type: enums.Skills, Name: "JavaScript", IsChecked: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Type: enums.Skills, Name: "Python", IsChecked: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Type: enums.Skills, Name: "PHP", IsChecked: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Type: enums.Skills, Name: "NodeJs", IsChecked: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{RID: 0, UID: 0, Type: enums.Skills, Name: "英语", IsChecked: false, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	for _, model := range modelList {
		if err := d.Db.Where("rid = ? and uid = ? and type = ? and name = ?", model.RID, model.UID, model.Type, model.Name).First(&models.Tags{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 数据不存在，插入新记录
				if err := d.Db.Create(&model).Error; err != nil {
					log.Println("[初始化]用户创建失败:", err)
				} else {
					log.Println("[初始化]用户创建创建成功:", model)
				}
			} else {
				log.Println("[初始化]用户查询失败:", err)
			}
		}
	}
}
