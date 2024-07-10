package database

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"resume/libs"
	"resume/models"
	"time"
)

// 初始化用户表
func initUsersData(db *gorm.DB) {
	users := []models.Users{
		{Nickname: "我的", Username: "admin", Password: libs.MD5("admin"), CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	for _, user := range users {
		var data models.Users

		if err := db.Where("username = ?", user.Username).First(&data).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 数据不存在，插入新记录
				if err := db.Create(&user).Error; err != nil {
					log.Println("[初始化]用户创建失败:", err)
				} else {
					log.Println("[初始化]用户创建创建成功:", data)
				}
			} else {
				log.Println("[初始化]用户查询失败:", data)
			}
		}
	}
}
