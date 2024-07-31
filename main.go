package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"resume/database"
	"resume/libs"
	"resume/routers"
	"strings"
)

//go:embed public/*
var static embed.FS

func init() {
	libs.CheckChrome()

	// 旧的 PATH 环境变量
	oldPath := os.Getenv("PATH")

	// 追加路径到 PATH 环境变量
	newPath := fmt.Sprintf("%s;%s", oldPath, fmt.Sprintf("%s\\chrome", os.TempDir()))

	_ = os.Setenv("PATH", newPath)
}

func main() {
	var dsn gorm.Dialector

	libs.CreateDir("data")

	if len(os.Args) == 2 && strings.ToLower(os.Args[1]) == "init" {
		file, err := static.ReadFile("public/config.ini.tpl")
		if err != nil {
			return
		}
		os.WriteFile("data/config.ini", file, 0644)
		log.Println("配置文件已生成！")
		os.Exit(0)
		return
	}

	fmt.Println("============================================================================")
	fmt.Println(fmt.Sprintf("= %s init 生成配置文件。", filepath.Base(os.Args[0])))
	fmt.Println("============================================================================")

	libs.CreateDir("data/temp")

	// 读取 INI 文件
	cfg, err := ini.Load("data/config.ini")
	if err != nil {
		cfg = ini.Empty()
		log.Println("读取 INI 配置文件打开失败，使用默认配置，错误信息：", err)
	}

	if cfg.HasSection("mysql") {
		dsn = database.OpenMySQLDB(cfg)
	} else {
		dsn = sqlite.Open("data/resume.db")
	}

	if !cfg.HasSection("server") {
		cfg.Section("server").Key("host").SetValue("127.0.0.1")
		cfg.Section("server").Key("port").SetValue("8080")
	}

	// 初始化数据库
	db := database.AutoMigrate(dsn)

	router := gin.Default()
	// 视图路由配置
	routers.Initialization(router, db, &static)

	host := cfg.Section("server").Key("host").String()
	port := cfg.Section("server").Key("port").String()
	err = router.Run(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("WEB 服务器启动失败: %v", err)
	}
}
