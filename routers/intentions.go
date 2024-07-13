package routers

import (
	"resume/controllers/intentions"
)

func (router *Routers) IntentionsGET() {
	router.Protected.GET("/intentions", intentions.Get)
}
