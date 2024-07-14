package routers

import (
	"resume/controllers/applicationinfo"
)

func (router *Routers) ApplicationInfoGET() {
	router.Protected.GET("/applicationinfo", applicationinfo.Get)
	router.Protected.GET("/applicationinfo/grade", applicationinfo.GradeGet)
}

func (router *Routers) ApplicationInfoPUT() {
	router.Protected.PUT("/applicationinfo", applicationinfo.Put)
	router.Protected.PUT("/applicationinfo/grade", applicationinfo.GradePut)
}

func (router *Routers) ApplicationInfoPOST() {
	router.Protected.POST("/applicationinfo/grade", applicationinfo.GradePost)
}

func (router *Routers) ApplicationInfoDELETE() {
	router.Protected.DELETE("/applicationinfo/grade", applicationinfo.GradeDelete)
}
