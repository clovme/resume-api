package routers

import (
	"resume/controllers/evaluation"
)

func (router *Routers) Evaluation() {
	router.Protected.GET("/evaluation", evaluation.Get)
}
