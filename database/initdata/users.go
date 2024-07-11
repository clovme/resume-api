package initdata

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"resume/libs"
	"resume/models"
	"time"
)

func (d *InitData) Users() {
	users := []models.Users{
		{Nickname: "我的", Username: "admin", Password: libs.MD5("admin"), Token: libs.MD5("admin"), ExpiresAt: time.Now(), CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	for _, user := range users {
		if err := d.Db.Where("username = ?", user.Username).First(&models.Users{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 数据不存在，插入新记录
				if err := d.Db.Create(&user).Error; err != nil {
					log.Println("[初始化]用户创建失败:", err)
				} else {
					log.Println("[初始化]用户创建创建成功:", user)
				}
			} else {
				log.Println("[初始化]用户查询失败:", err)
			}
		}
	}
}
