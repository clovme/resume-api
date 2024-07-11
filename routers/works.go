package routers

import "resume/controllers/works"

func (r *Routers) Works() {
	r.Protected.GET("/works", works.Get)
}
