package utils

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"net/http"
	"resume/libs"
	"resume/models"
)

func CopyMenus(tx *gorm.DB, s libs.HttpStatus, resume models.Resumes) {
	// 复制菜单
	var menus []models.Menus
	if err := tx.Find(&menus, "rid = 0").Error; err != nil {
		log.Println("系统菜单查询失败！", err.Error())
		s.Msg(http.StatusInternalServerError, "服务器异常，请重试！")
		return
	}
	// 创建简历菜单
	for _, menu := range menus {
		var data models.Menus

		if err := tx.Where("rid = ? and uid = ? and title = ?", resume.ID, s.User.ID, menu.Title).First(&data).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 数据不存在，插入新记录
				menu.ID = 0
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
