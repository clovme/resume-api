package routers

import (
	"resume/controllers/intentions"
)

func (router *Routers) Intentions() {
	router.Protected.GET("/intentions", intentions.Get)
}
