package routers

import (
	"resume/controllers/applicationinfo"
)

func (r *Routers) ApplicationInfo() {
	r.Protected.GET("/applicationinfo", applicationinfo.Get)
}
