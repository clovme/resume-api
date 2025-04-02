package routers

import (
	"embed"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"html/template"
	"io/fs"
	"net/http"
	"reflect"
	"resume/middleware"
	"resume/types/enums"
)

type Routers struct {
	Router    *gin.Engine
	Public    *gin.RouterGroup
	Protected *gin.RouterGroup
}

// Initialization 初始化路由
// @params router web 路由
// @params db 数据库连接信息
// @params static embed.FS
func Initialization(router *gin.Engine, db *gorm.DB, static *embed.FS) {
	// 设置为信任所有代理
	err := router.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}

	// 使用数据库中间件
	router.Use(middleware.DBMiddleware(db, static))
	router.Use(middleware.CorsMiddleware())
	router.Use(middleware.NotFoundHandler())

	// 加载嵌入的模板文件
	tmpl := template.Must(template.New("").ParseFS(static, "public/*.html"))
	router.SetHTMLTemplate(tmpl)

	// 提供嵌入的静态文件
	staticFS, _ := fs.Sub(static, "public/assets")
	router.StaticFS("/assets", http.FS(staticFS))

	// 配置静态文件目录
	router.Static("/temp", enums.TempPath)

	public := router.Group("/api")
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())

	v := reflect.ValueOf(&Routers{Router: router, Public: public, Protected: protected})
	for i := 0; i < v.NumMethod(); i++ {
		v.Method(i).Call(nil)
	}
}
