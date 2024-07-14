package routers

import "resume/controllers/works"

func (router *Routers) WorksGET() {
	router.Protected.GET("/works", works.Get)
}

func (router *Routers) WorksPUT() {
	router.Protected.PUT("/works", works.Put)
	router.Protected.PUT("/works/sort", works.Sort)
}

func (router *Routers) WorksPOST() {
	router.Protected.POST("/works", works.Post)
}

func (router *Routers) WorksDELETE() {
	router.Protected.DELETE("/works", works.Delete)
}
