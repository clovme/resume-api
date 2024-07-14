package resume

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"resume/types/enums"
)

func CopyData[T any](msg string, tx *gorm.DB, models []T, setValues func(m *T) (string, *gorm.DB)) bool {
	// 判断默认数据是否存在
	if err := tx.Find(&models, "rid = ? AND uid = ?", enums.Vx32, enums.Vx32).Error; err != nil {
		log.Println(fmt.Sprintf("%s查询失败！", msg), err.Error())
		return false
	}

	flag := true
	// 添加数据
	for _, model := range models {
		// 设置uid和rid以及其他参数
		RName, Where := setValues(&model)

		if err := Where.First(&model).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 数据不存在，插入新记录
				if err := tx.Create(&model).Error; err != nil {
					tx.Rollback()
					log.Println(fmt.Sprintf("简历名称[%s]->[%s]复制失败:", RName, msg), err)
					flag = false
				} else {
					log.Println(fmt.Sprintf("简历名称[%s]->[%s]复制创建成功！", RName, msg))
				}
			} else {
				tx.Rollback()
				log.Println(fmt.Sprintf("简历名称[%s]->[%s]查询失败:", RName, msg), err.Error())
				flag = false
			}
		}
	}
	return flag
}
