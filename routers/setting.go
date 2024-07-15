package routers

import (
	"resume/controllers/setting"
)

func (router *Routers) SettingGET() {
	router.Protected.GET("/setting", setting.Get)
}

func (router *Routers) SettingPUT() {
	router.Protected.PUT("/setting", setting.Put)
}
