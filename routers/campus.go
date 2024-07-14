package routers

import (
	"resume/controllers/campus"
)

func (router *Routers) CampusGET() {
	router.Protected.GET("/campus", campus.Get)
}

func (router *Routers) CampusPUT() {
	router.Protected.PUT("/campus", campus.Put)
	router.Protected.PUT("/campus/sort", campus.Sort)
}

func (router *Routers) CampusPOST() {
	router.Protected.POST("/campus", campus.Post)
}

func (router *Routers) CampusDELETE() {
	router.Protected.DELETE("/campus", campus.Delete)
}
