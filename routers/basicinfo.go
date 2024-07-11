package routers

import "resume/controllers/basicinfo"

func (r *Routers) BasicInfo() {
	r.Protected.GET("/basicinfo", basicinfo.Get)
}
