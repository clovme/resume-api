package routers

import (
	"resume/controllers/intentions"
)

func (r *Routers) Intentions() {
	r.Protected.GET("/intentions", intentions.Get)
}
