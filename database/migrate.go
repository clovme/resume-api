package database

import (
	"gorm.io/gorm"
	"log"
	"resume/models"
)

func AutoMigrate(dbn gorm.Dialector) *gorm.DB {
	db, err := gorm.Open(dbn, &gorm.Config{})
	if err != nil {
		log.Fatal("[初始化]数据库拦截失败")
	}

	err = db.AutoMigrate(
		&models.Users{},
		&models.BasicInfo{},
		&models.Menus{},
		&models.Resumes{},
		&models.SkillsExpertise{},
		&models.SkillsExpertiseTags{},
		&models.SkillsExpertiseCheckedTags{},
	)

	if err != nil {
		log.Fatal("[初始化]数据库迁移失败：", err)
	}

	initUsersData(db)
	initMenusData(db)
	initBasicInfoData(db)

	return db
}
