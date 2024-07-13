package routers

import (
	"resume/controllers/menus"
)

func (router *Routers) MenusGET() {
	router.Protected.GET("/menus", menus.Get)
}

func (router *Routers) MenusPUT() {
	router.Protected.PUT("/menus/sort", menus.Sort)
	router.Protected.PUT("/menus/edit/name", menus.EditName)
	router.Protected.PUT("/menus/switch/status", menus.SwitchStatus)
}
