package routers

import (
	"resume/controllers/resume"
)

func (r *Routers) Resumes() {
	r.Protected.GET("/resumes", resume.Get)
	r.Protected.GET("/resumes/id", resume.GetResumeId)
	r.Protected.GET("/resumes/check", resume.CheckResume)

	r.Protected.PUT("/resumes", resume.Put)

	r.Protected.POST("/resumes", resume.Post)
	
	r.Protected.DELETE("/resumes", resume.Delete)
}
