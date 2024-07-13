package routers

import "resume/controllers/basicinfo"

func (router *Routers) BasicInfoGET() {
	router.Protected.GET("/basicinfo", basicinfo.Get)
}

func (router *Routers) BasicInfoPUT() {
	router.Protected.PUT("/basicinfo", basicinfo.Put)
}
