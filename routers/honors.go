package routers

import (
	"resume/controllers/honors"
)

func (router *Routers) HonorsGET() {
	router.Protected.GET("/honors", honors.Get)
}

func (router *Routers) HonorsPUT() {
	router.Protected.PUT("/honors", honors.Put)
}
