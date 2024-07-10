package routers

import (
	"github.com/gin-gonic/gin"
	"resume/controllers"
	"resume/controllers/basicinfo"
	"resume/controllers/resume"
	"resume/controllers/users"
	"resume/middleware"
)

func Router(router *gin.Engine) {
	router.GET("", controllers.GetIndexView)

	public := router.Group("/api")
	{
		public.GET("/icon", controllers.GetIcon)
		public.POST("/login", users.SignIn)
	}

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.PUT("/users", users.Put)
		protected.GET("/menus", controllers.GetMenus)

		protected.GET("/resumes", resume.Get)
		protected.GET("/resumes/id", resume.GetResumeId)
		protected.GET("/resumes/check", resume.CheckResume)
		protected.PUT("/resumes", resume.Put)
		protected.POST("/resumes", resume.Post)
		protected.DELETE("/resumes", resume.Delete)

		protected.GET("/basicinfo", basicinfo.Get)
	}
}
