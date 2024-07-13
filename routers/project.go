package routers

import (
	"resume/controllers/project"
)

func (router *Routers) ProjectGET() {
	router.Protected.GET("/project", project.Get)
}
