package routers

import (
	"resume/controllers/project"
)

func (router *Routers) ProjectGET() {
	router.Protected.GET("/project", project.Get)
}

func (router *Routers) ProjectPUT() {
	router.Protected.PUT("/project", project.Put)
	router.Protected.PUT("/project/sort", project.Sort)
}

func (router *Routers) ProjectPOST() {
	router.Protected.POST("/project", project.Post)
}

func (router *Routers) ProjectDELETE() {
	router.Protected.DELETE("/project", project.Delete)
}
