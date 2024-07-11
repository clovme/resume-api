package database

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

// OpenMySQLDB MySQL数据库初始化 和 打开数据库
func OpenMySQLDB(cfg *ini.File) gorm.Dialector {
	var err error

	// 获取特定节和键的值
	host := cfg.Section("mysql").Key("host").String()
	port := cfg.Section("mysql").Key("port").String()
	username := cfg.Section("mysql").Key("username").String()
	password := cfg.Section("mysql").Key("password").String()
	database := cfg.Section("mysql").Key("database").String()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", username, password, host, port)
	conn, err := gorm.Open(mysql.Open(fmt.Sprintf(dsn)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("[初始化][创建数据库]无法连接到 MySQL 服务器: ", err)
	}
	// 检查数据库是否存在
	var result int
	conn.Raw("SELECT COUNT(*) FROM information_schema.schemata WHERE schema_name = ?", database).Scan(&result)
	if result == 0 {
		// 数据库不存在，创建数据库
		if err := conn.Exec(fmt.Sprintf("create schema `%s` collate utf8mb4_general_ci;", database)).Error; err != nil {
			log.Fatal("[初始化]无法创建数据库: ", err)
		}
	}
	return mysql.Open(fmt.Sprintf("%s%s?charset=utf8mb4&parseTime=True&loc=Local", dsn, database))
}
