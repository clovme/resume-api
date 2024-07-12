package routers

import "resume/controllers/basicinfo"

func (router *Routers) BasicInfo() {
	router.Protected.GET("/basicinfo", basicinfo.Get)
}
