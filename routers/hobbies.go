package routers

import (
	"resume/controllers/hobbies"
)

func (r *Routers) Hobbies() {
	r.Protected.GET("/hobbies", hobbies.Get)
}
