package routers

import (
	"resume/controllers/campus"
)

func (r *Routers) Campus() {
	r.Protected.GET("/campus", campus.Get)
}
