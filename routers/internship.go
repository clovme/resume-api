package routers

import (
	"resume/controllers/internship"
)

func (router *Routers) InternshipGET() {
	router.Protected.GET("/internship", internship.Get)
}

func (router *Routers) InternshipPUT() {
	router.Protected.PUT("/internship", internship.Put)
	router.Protected.PUT("/internship/sort", internship.Sort)
}

func (router *Routers) InternshipPOST() {
	router.Protected.POST("/internship", internship.Post)
}

func (router *Routers) InternshipDELETE() {
	router.Protected.DELETE("/internship", internship.Delete)
}
