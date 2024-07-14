package routers

import (
	"resume/controllers/evaluation"
)

func (router *Routers) EvaluationGET() {
	router.Protected.GET("/evaluation", evaluation.Get)
}

func (router *Routers) EvaluationPUT() {
	router.Protected.PUT("/evaluation", evaluation.Put)
}
