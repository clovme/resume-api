package routers

import (
	"resume/controllers"
	"resume/controllers/users"
)

func (r *Routers) PublicGroup() {
	r.Public.GET("/icon", controllers.GetIcon)
	r.Public.POST("/login", users.SignIn)
}
