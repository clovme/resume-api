package routers

import "resume/controllers"

func (router *Routers) Views() {
	router.Public.GET("/icon", controllers.GetIcon)
	router.Router.GET("", controllers.GetIndexView)
}
