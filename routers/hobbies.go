package routers

import (
	"resume/controllers/hobbies"
)

func (router *Routers) Hobbies() {
	router.Protected.GET("/hobbies", hobbies.Get)
}
