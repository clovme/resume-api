package routers

import (
	"resume/controllers/applicationinfo"
)

func (router *Routers) ApplicationInfo() {
	router.Protected.GET("/applicationinfo", applicationinfo.Get)
}
