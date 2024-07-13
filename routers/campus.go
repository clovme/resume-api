package routers

import (
	"resume/controllers/campus"
)

func (router *Routers) CampusGET() {
	router.Protected.GET("/campus", campus.Get)
}
