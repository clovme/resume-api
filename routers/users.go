package routers

import (
	"resume/controllers/users"
)

func (router *Routers) UsersPUT() {
	router.Protected.PUT("/users", users.Put)
}
