package routers

import "resume/controllers"

func (router *Routers) SettingPOST() {
	router.Protected.POST("/pdf", controllers.Html2PDF)
}
