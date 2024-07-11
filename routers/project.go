package routers

import (
	"resume/controllers/project"
)

func (r *Routers) Project() {
	r.Protected.GET("/project", project.Get)
}
