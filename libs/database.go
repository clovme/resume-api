package libs

import (
	"gorm.io/gorm"
	"log"
	"net/http"
	"resume/types/enums"
)

// Transaction 开启事务
func Transaction(s HttpStatus, callback func(tx *gorm.DB)) {
	// 开始事务
	tx := s.Begin()
	// 事务操作
	if err := tx.Error; err != nil {
		log.Println("事务开启失败！", err.Error())
		s.Msg(http.StatusForbidden, "创建简历模版创建失败，请重试！")
		return
	}

	callback(tx)

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Println("创建简历模版事务提交失败！", err.Error())
		s.Msg(http.StatusForbidden, "创建简历模版创建失败，请重试！")
		return
	}
}

// UpdateNotCreatedAt 更新数据或者创建数据
func UpdateNotCreatedAt[T any](data *T, db HttpStatus) bool {
	if result := db.Where("rid != ? AND uid != ?", enums.Vx32, enums.Vx32).Omit("id", "rid", "uid", "CreatedAt").Save(&data); result.Error != nil {
		log.Println("UpdateSave数据查询异常:", result.Error)
		return false
	}
	return true
}
