package routers

import (
	"resume/controllers/resume"
)

func (router *Routers) ResumesGET() {
	router.Protected.GET("/resumes", resume.Get)
	router.Protected.GET("/resumes/id", resume.GetResumeId)
	router.Protected.GET("/resumes/check", resume.CheckResume)
}

func (router *Routers) ResumesPUT() {
	router.Protected.PUT("/resumes", resume.Put)
}

func (router *Routers) ResumesPOST() {
	router.Protected.POST("/resumes", resume.Post)
	router.Protected.POST("/resumes/copy", resume.Copy)
}

func (router *Routers) ResumesDELETE() {
	router.Protected.DELETE("/resumes", resume.Delete)
}
