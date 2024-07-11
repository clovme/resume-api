package routers

import "resume/controllers/education"

func (r *Routers) Education() {
	r.Protected.GET("/education", education.Get)
}
