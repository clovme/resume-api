package routers

import "resume/controllers/internship"

func (r *Routers) Internship() {
	r.Protected.GET("/internship", internship.Get)
}
