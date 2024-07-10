package main

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"resume/database"
	"resume/libs"
	"resume/middleware"
	"resume/routers"
)

//go:embed public/*
var static embed.FS

func main() {
	var dsn gorm.Dialector

	libs.CreateDir("data")

	// 读取 INI 文件
	cfg, err := ini.Load("data/config.ini")
	if err != nil {
		cfg = ini.Empty()
		cfg.Section("server").Key("host").SetValue("127.0.0.1")
		cfg.Section("server").Key("port").SetValue("8080")
		dsn = sqlite.Open("data/resume.db")
		//dsn = sqlite.Open(":memory:")
		log.Println("读取 INI 配置文件打开失败，使用默认配置，错误信息：", err)
	} else {
		dsn = database.OpenDB(cfg)
	}

	// 初始化数据库
	DB := database.AutoMigrate(dsn)

	router := gin.Default()

	// 使用数据库中间件
	router.Use(middleware.DBMiddleware(DB))
	router.Use(middleware.CorsMiddleware())
	router.Use(middleware.NotFoundHandler())

	// 加载嵌入的模板文件
	tmpl := template.Must(template.New("").ParseFS(static, "public/*.html"))
	router.SetHTMLTemplate(tmpl)

	// 提供嵌入的静态文件
	staticFS, _ := fs.Sub(static, "public/assets")
	router.StaticFS("/assets", http.FS(staticFS))

	// 视图路由配置
	routers.Router(router)

	host := cfg.Section("server").Key("host").String()
	port := cfg.Section("server").Key("port").String()
	err = router.Run(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("WEB 服务器启动失败: %v", err)
	}
}
