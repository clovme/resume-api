package routers

import "resume/controllers/internship"

func (router *Routers) InternshipGET() {
	router.Protected.GET("/internship", internship.Get)
}
