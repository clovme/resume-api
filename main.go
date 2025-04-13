package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ncruces/zenity"
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

	libs.CreateDir(enums.TempPath)
	libs.CreateDir(enums.DataPath)

	log.Printf("数据文件夹，位置：%s。\n", enums.DataPath)
	log.Printf("临时数据文件夹，位置：%s。\n", enums.TempPath)
}

func setIniValue(cfg *ini.File, section, key, defaultValue, tip string) {
	fmt.Printf("　* %s>", tip)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n') // 读取整行，包括空格
	name = strings.TrimSpace(name)     // 去除首尾换行符

	if name == "" {
		name = defaultValue
	}
	cfg.Section(section).Key(key).SetValue(name)
}

func input(tip string) string {
	var name string
	fmt.Printf("　* %s>", tip)
	_, _ = fmt.Scanln(&name)
	return strings.TrimSpace(name)
}

func browserExePathSelect(cfg *ini.File, section, key, tip string) {
	name := input(tip)

	if name == "" || strings.ToLower(name) == "y" {
		brs := map[string]map[string]string{
			"Chrome": {
				"title": "Google Chrome文件路径选择",
				"name":  "chrome.exe",
				"path":  enums.BrowserExePath.Chrome,
			},
			"Edge": {
				"title": "Microsoft Edge文件路径选择",
				"name":  "msedge.exe",
				"path":  enums.BrowserExePath.Edge,
			},
		}

		_filepath, err := zenity.SelectFile(
			zenity.Filename(brs[key]["path"]),
			zenity.Title(brs[key]["title"]),
			zenity.FileFilters{
				{Name: brs[key]["name"], Patterns: []string{brs[key]["name"]}},
			})
		if err != nil {
			return
		}
		enums.BrowserExePath.Set(key, _filepath)
	}

	cfg.Section(section).Key(key).SetValue(enums.BrowserExePath.Get(key))
}

func fileExists(path string) bool {

	if _, err := os.Stat(path); err != nil {
		return false // 文件存在
	}
	return true
}

func BrowserCheck(path string) string {
	if fileExists(path) {
		return path
	} else {
		return fmt.Sprintf("%s(未配置)", path)
	}
}

func main() {
	// 读取 INI 文件
	cfg, err := ini.Load(enums.ConfigPath)
	if err != nil {
		cfg = ini.Empty()
		log.Println("没有配置文件，使用默认配置。")
	} else {
		log.Printf("读取配置文件：%s。\n", enums.ConfigPath)
		enums.BrowserExePath.Edge = cfg.Section("Browser").Key("Edge").String()
		enums.BrowserExePath.Chrome = cfg.Section("Browser").Key("Chrome").String()
	}

	log.Printf("Edge浏览器：%s。\n", BrowserCheck(enums.BrowserExePath.Edge))
	log.Printf("Chrome浏览器：%s。\n", BrowserCheck(enums.BrowserExePath.Chrome))

	if len(os.Args) == 2 && strings.ToLower(os.Args[1]) == "init" {
		cfg = ini.Empty()
		fmt.Println("数据库配置：")
		sel := input("数据库请选择: 1.SQLite3(默认) 2.MySQL")
		if sel == "1" || sel == "" {
			cfg.Section("SQLite3").Key("database").SetValue(enums.DatabasePath)
		} else {
			setIniValue(cfg, "MySQL", "host", "127.0.0.1", "MySQL主机地址(127.0.0.1)")
			setIniValue(cfg, "MySQL", "port", "3306", "MySQL端口(3306)")
			setIniValue(cfg, "MySQL", "username", "root", "MySQL用户名(root)")
			setIniValue(cfg, "MySQL", "password", "", "MySQL密码")
			setIniValue(cfg, "MySQL", "database", "resume", "MySQL数据库名称(resume)")
		}

		fmt.Println("WEB 服务配置：")
		setIniValue(cfg, "Server", "host", "localhost", "WEB地址(localhost)")
		setIniValue(cfg, "Server", "port", "8080", "WEB端口(8080)")
		setIniValue(cfg, "Server", "mode", "release", "日志级别debug/release/test(release)")

		fmt.Println("浏览器绝对路径：")
		browserExePathSelect(cfg, "Browser", "Edge", "选择(Edge)路径，Y(默认)/N")
		browserExePathSelect(cfg, "Browser", "Chrome", "选择(Chrome)路径，Y(默认)/N")

		fmt.Println("配置完成。")

		_ = cfg.SaveTo(enums.ConfigPath)
	}

	var dsn gorm.Dialector
	if cfg.HasSection("MySQL") {
		dsn = database.OpenMySQLDB(cfg)
	} else {
		dbname := cfg.Section("SQLite3").Key("database").String()
		if dbname == "" {
			dsn = sqlite.Open(enums.DatabasePath)
		} else {
			dsn = sqlite.Open(dbname)
		}
	}

	if !cfg.HasSection("Server") {
		cfg.Section("Server").Key("host").SetValue("localhost")
		cfg.Section("Server").Key("port").SetValue("8080")
		cfg.Section("Server").Key("mode").SetValue("release")
	}

	// 初始化数据库
	db := database.AutoMigrate(dsn)

	host := cfg.Section("Server").Key("host").String()
	port := cfg.Section("Server").Key("port").String()
	mode := cfg.Section("Server").Key("mode").String()

	log.Printf("日志级别：(%s)debug/release/test\n", mode)
	log.Println(fmt.Sprintf("地址端口：http://%s:%s", host, port))
	gin.SetMode(mode)
	router := gin.Default()
	// 视图路由配置
	routers.Initialization(router, db, &static)
	err = router.Run(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("WEB 服务器启动失败: %v", err)
	}
}
