package routers

import (
	"resume/controllers/menus"
)

func (r *Routers) Menus() {
	r.Protected.GET("/menus", menus.Get)
}
