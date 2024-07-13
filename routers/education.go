package routers

import "resume/controllers/education"

func (router *Routers) EducationGET() {
	router.Protected.GET("/education", education.Get)
}
