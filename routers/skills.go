package routers

import "resume/controllers/skills"

func (r *Routers) SkillsExpertise() {
	r.Protected.GET("/skills", skills.Get)
}
