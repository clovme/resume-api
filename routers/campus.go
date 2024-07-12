package routers

import (
	"resume/controllers/campus"
)

func (router *Routers) Campus() {
	router.Protected.GET("/campus", campus.Get)
}
