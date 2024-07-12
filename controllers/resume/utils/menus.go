package utils

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"resume/libs"
	"resume/models"
	"resume/types/enums"
)

func CopyMenus(tx *gorm.DB, s libs.HttpStatus, resume models.Resumes) {
	// 复制菜单
	var menus []models.Menus
	if err := tx.Find(&menus, "rid = ? AND uid = ?", enums.Vx32, enums.Vx32).Error; err != nil {
		log.Println("系统菜单查询失败！", err.Error())
		return
	}
	// 创建简历菜单
	for _, menu := range menus {
		if err := tx.Where("rid = ? AND uid = ? AND title = ?", resume.ID, s.User.ID, menu.Title).First(&models.Menus{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 数据不存在，插入新记录
				menu.UID = s.User.ID
				menu.RID = resume.ID
				if err := tx.Create(&menu).Error; err != nil {
					tx.Rollback()
					log.Println("["+resume.Name+"]菜单创建失败:", err)
				} else {
					log.Println("["+resume.Name+"]菜单创建创建成功！", menu)
				}
			} else {
				tx.Rollback()
				log.Println("["+resume.Name+"]菜单查询失败:", err.Error())
			}
		}
	}
}
