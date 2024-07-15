package routers

import (
	"resume/controllers/slogan"
)

func (router *Routers) SloganGET() {
	router.Protected.GET("/slogan", slogan.Get)
}

func (router *Routers) SloganPUT() {
	router.Protected.PUT("/slogan", slogan.Put)
}
