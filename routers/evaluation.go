package routers

import (
	"resume/controllers/evaluation"
)

func (router *Routers) EvaluationGET() {
	router.Protected.GET("/evaluation", evaluation.Get)
}
