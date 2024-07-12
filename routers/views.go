package routers

import "resume/controllers"

func (router *Routers) Views() {
	router.Router.GET("", controllers.GetIndexView)
}
