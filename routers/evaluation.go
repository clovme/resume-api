package routers

import (
	"resume/controllers/evaluation"
)

func (r *Routers) Evaluation() {
	r.Protected.GET("/evaluation", evaluation.Get)
}
