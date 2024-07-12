package routers

import (
	"resume/controllers/honors"
)

func (router *Routers) Honors() {
	router.Protected.GET("/honors", honors.Get)
}
