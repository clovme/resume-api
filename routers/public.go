package routers

import (
	"resume/controllers"
	"resume/controllers/users"
)

func (router *Routers) PublicGroup() {
	router.Public.GET("/icon", controllers.GetIcon)
	router.Public.POST("/login", users.SignIn)
}
