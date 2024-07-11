package routers

import (
	"resume/controllers"
)

func (r *Routers) Menus() {
	r.Protected.GET("/menus", controllers.GetMenus)
}
