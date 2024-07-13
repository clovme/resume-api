package routers

import "resume/controllers/tags"

func (router *Routers) TagsGET() {
	router.Protected.GET("/tags", tags.Get)
}
