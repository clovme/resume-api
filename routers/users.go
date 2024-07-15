package routers

import (
	"resume/controllers/users"
)

func (router *Routers) UsersPUT() {
	router.Protected.PUT("/users", users.Put)
}

func (router *Routers) UsersPOST() {
	router.Public.POST("/login", users.SignIn)
}

func (router *Routers) UsersDELETE() {
	router.Protected.DELETE("/logout", users.SignOut)
}
