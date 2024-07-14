package routers

import "resume/controllers/skills"

func (router *Routers) SkillsGET() {
	router.Protected.GET("/skills", skills.Get)
}

func (router *Routers) SkillsPUT() {
	router.Protected.PUT("/skills", skills.Put)
}
