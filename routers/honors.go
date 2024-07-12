package routers

import (
	"resume/controllers/honors"
)

func (r *Routers) Honors() {
	r.Protected.GET("/honors", honors.Get)
}
