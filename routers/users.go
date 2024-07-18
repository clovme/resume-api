package routers

import (
	"resume/controllers/users"
)

func (router *Routers) UsersPUT() {
	router.Protected.PUT("/users", users.Put)
}

func (router *Routers) UsersPOST() {
	router.Public.POST("/login", users.SignIn)
	router.Public.POST("/regedit", users.Regedit)
}

func (router *Routers) UsersDELETE() {
	router.Public.DELETE("/logout", users.SignOut)
}
