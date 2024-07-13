package routers

import (
	"resume/controllers"
	"resume/controllers/users"
)

func (router *Routers) GET() {
	router.Public.GET("/icon", controllers.GetIcon)
}

func (router *Routers) POST() {
	router.Public.POST("/login", users.SignIn)
}
