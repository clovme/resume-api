package routers

import (
	"resume/controllers/menus"
)

func (router *Routers) Menus() {
	router.Protected.GET("/menus", menus.Get)
	router.Protected.PUT("/menus", menus.Put)
}
