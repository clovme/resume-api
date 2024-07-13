package routers

import (
	"resume/controllers/honors"
)

func (router *Routers) HonorsGET() {
	router.Protected.GET("/honors", honors.Get)
}
