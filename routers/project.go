package routers

import (
	"resume/controllers/project"
)

func (router *Routers) Project() {
	router.Protected.GET("/project", project.Get)
}
