package routers

import "resume/controllers/works"

func (router *Routers) Works() {
	router.Protected.GET("/works", works.Get)
}
