package routers

import (
	"resume/controllers/education"
)

func (router *Routers) EducationGET() {
	router.Protected.GET("/education", education.Get)
}

func (router *Routers) EducationPUT() {
	router.Protected.PUT("/education", education.Put)
	router.Protected.PUT("/education/sort", education.Sort)
}

func (router *Routers) EducationPOST() {
	router.Protected.POST("/education", education.Post)
}

func (router *Routers) EducationDELETE() {
	router.Protected.DELETE("/education", education.Delete)
}
