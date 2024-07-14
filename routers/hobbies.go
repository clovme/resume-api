package routers

import (
	"resume/controllers/hobbies"
)

func (router *Routers) HobbiesGET() {
	router.Protected.GET("/hobbies", hobbies.Get)
}

func (router *Routers) HobbiesPUT() {
	router.Protected.PUT("/hobbies", hobbies.Put)
}
