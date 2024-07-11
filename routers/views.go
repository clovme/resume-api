package routers

import "resume/controllers"

func (r *Routers) Views() {
	r.Router.GET("", controllers.GetIndexView)
}
