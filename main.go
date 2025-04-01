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
	"resume/types/enums"
	"strings"
)

//go:embed public/*
var static embed.FS

func init() {
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	fmt.Println(fmt.Sprintf("+ 使用 %s init 初始化配置文件。", filepath.Base(os.Args[0])))
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	//if err := libs.DownloadChromeFile(static); err != nil {
	//	fmt.Println("Chrome 浏览器下载失败，错误信息：", err)
	//} else {
	//	fmt.Printf("Chrome 浏览器位置：%s。\n", enums.ChromeExePath)
	//}
	fmt.Printf("数据文件夹，位置：%s。\n", enums.DataPath)
	fmt.Printf("临时数据文件夹，位置：%s。\n", filepath.Join(enums.DataPath, "temp"))
}

func setIniValue(cfg *ini.File, section, key, defaultValue, tip string) {
	name := input(tip)
	if name == "" {
		name = defaultValue
	}
	cfg.Section(section).Key(key).SetValue(name)
}

func input(tip string) string {
	var name string
	fmt.Printf("* %s>", tip)
	_, _ = fmt.Scanln(&name)
	return name
}

func main() {
	libs.CreateDir(filepath.Join(enums.DataPath, "temp"))

	// 读取 INI 文件
	cfg, err := ini.Load("data/config.ini")
	if err != nil {
		cfg = ini.Empty()
		log.Println("没有配置文件，使用默认配置。")
	}

	if len(os.Args) == 2 && strings.ToLower(os.Args[1]) == "init" {
		cfg = ini.Empty()
		fmt.Println("数据库配置：")
		sel := input("数据库请选择: 1.MySQL 2.sqlite3")
		if sel == "1" || sel == "" {
			setIniValue(cfg, "mysql", "host", "127.0.0.1", "MySQL主机地址(127.0.0.1)")
			setIniValue(cfg, "mysql", "port", "3306", "MySQL端口(3306)")
			setIniValue(cfg, "mysql", "username", "root", "MySQL用户名(root)")
			setIniValue(cfg, "mysql", "password", "", "MySQL密码")
			setIniValue(cfg, "mysql", "database", "resume", "MySQL数据库名称(resume)")
		} else {
			cfg.Section("sqlite3").Key("database").SetValue("data/resume.db")
		}

		fmt.Println("WEB 服务配置：")
		setIniValue(cfg, "server", "host", "localhost", "WEB地址(localhost)")
		setIniValue(cfg, "server", "port", "8080", "WEB端口(8080)")
		setIniValue(cfg, "server", "mode", "release", "日志级别debug/release/test(release)")
		fmt.Println("配置完成。")

		_ = cfg.SaveTo("data/config.ini")
	}

	var dsn gorm.Dialector
	if cfg.HasSection("mysql") {
		dsn = database.OpenMySQLDB(cfg)
	} else {
		dbname := cfg.Section("sqlite3").Key("database").String()
		if dbname == "" {
			dsn = sqlite.Open("data/resume.db")
		} else {
			dsn = sqlite.Open(dbname)
		}
	}

	if !cfg.HasSection("server") {
		cfg.Section("server").Key("host").SetValue("localhost")
		cfg.Section("server").Key("port").SetValue("8080")
		cfg.Section("server").Key("mode").SetValue("release")
	}

	// 初始化数据库
	db := database.AutoMigrate(dsn)

	host := cfg.Section("server").Key("host").String()
	port := cfg.Section("server").Key("port").String()
	mode := cfg.Section("server").Key("mode").String()

	fmt.Printf("日志级别：(%s)debug/release/test\n", mode)
	fmt.Println(fmt.Sprintf("地址端口：http://%s:%s", host, port))
	gin.SetMode(mode)
	router := gin.Default()
	// 视图路由配置
	routers.Initialization(router, db, &static)
	err = router.Run(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("WEB 服务器启动失败: %v", err)
	}
}
