package routers

import (
	"resume/controllers/applicationinfo"
)

func (router *Routers) ApplicationInfoGET() {
	router.Protected.GET("/applicationinfo", applicationinfo.Get)
}
