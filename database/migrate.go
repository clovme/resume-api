package database

import (
	"gorm.io/gorm"
	"log"
	"reflect"
	"resume/database/initdata"
	"resume/models"
)

func AutoMigrate(dbn gorm.Dialector) *gorm.DB {
	db, err := gorm.Open(dbn, &gorm.Config{})
	if err != nil {
		log.Fatalln("[初始化]数据库打开失败！")
	}

	err = db.AutoMigrate(
		&models.Tags{},
		&models.Users{},
		&models.Menus{},
		&models.Resumes{},
		&models.BasicInfo{},
		&models.Skills{},
	)

	if err != nil {
		log.Fatalln("[初始化]数据库迁移失败：", err)
	}

	v := reflect.ValueOf(&initdata.InitData{Db: db})

	for i := 0; i < v.NumMethod(); i++ {
		v.Method(i).Call(nil)
	}

	return db
}
