package routers

import "resume/controllers/skills"

func (router *Routers) SkillsGET() {
	router.Protected.GET("/skills", skills.Get)
}
