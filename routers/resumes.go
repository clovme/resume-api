package routers

import (
	"resume/controllers/resume"
)

func (router *Routers) Resumes() {
	router.Protected.GET("/resumes", resume.Get)
	router.Protected.GET("/resumes/id", resume.GetResumeId)
	router.Protected.GET("/resumes/check", resume.CheckResume)

	router.Protected.PUT("/resumes", resume.Put)

	router.Protected.POST("/resumes", resume.Post)

	router.Protected.DELETE("/resumes", resume.Delete)
}
