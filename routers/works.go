package routers

import "resume/controllers/works"

func (router *Routers) WorksGET() {
	router.Protected.GET("/works", works.Get)
}
