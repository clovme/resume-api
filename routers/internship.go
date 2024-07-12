package routers

import "resume/controllers/internship"

func (router *Routers) Internship() {
	router.Protected.GET("/internship", internship.Get)
}
