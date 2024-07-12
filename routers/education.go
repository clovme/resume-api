package routers

import "resume/controllers/education"

func (router *Routers) Education() {
	router.Protected.GET("/education", education.Get)
}
