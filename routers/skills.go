package routers

import "resume/controllers/skills"

func (router *Routers) SkillsExpertise() {
	router.Protected.GET("/skills", skills.Get)
}
