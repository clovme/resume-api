package initdata

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"resume/models"
	"resume/types/enums"
)

type InitData struct {
	Db *gorm.DB
}

// insertRecord 插入数据
func insertRecord[T any](msg string, modelList []T, DbWhere func(model T) (db, where *gorm.DB)) {
	for _, model := range modelList {
		db, Where := DbWhere(model)
		if err := Where.First(&model).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 数据不存在，插入新记录
				if err := db.Create(&model).Error; err != nil {
					log.Println(fmt.Sprintf("[初始化数据] %s创建失败:", msg), err)
				} else {
					log.Println(fmt.Sprintf("[初始化数据] %s创建成功:", msg), model)
				}
			} else {
				log.Println(fmt.Sprintf("[初始化数据] %s查询失败:", msg), err)
			}
		}
	}
}

func ridUID() models.BaseModelWithRIDUID {
	return models.BaseModelWithRIDUID{RID: enums.Vx32, UID: enums.Vx32}
}
