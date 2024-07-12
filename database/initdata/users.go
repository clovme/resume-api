package initdata

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"resume/libs"
	"resume/models"
	"time"
)

// Users 初始化用户
func (d *InitData) Users() {
	modelList := []models.Users{
		{Nickname: "我的", Username: "admin", Password: libs.MD5("admin"), Token: libs.MD5(uuid.New().String()), ExpiresAt: time.Now()},
		{Nickname: "Admin", Username: "admin1", Password: libs.MD5("admin"), Token: libs.MD5(uuid.New().String()), ExpiresAt: time.Now()},
	}

	insertRecord[models.Users]("技能特长", modelList, func(model models.Users) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("username = ?", model.Username)
	})
}
