package routers

import "resume/controllers/tags"

func (r *Routers) Tags() {
	r.Protected.GET("/tags", tags.Get)
}
