package routers

import (
	"resume/controllers/users"
)

func (router *Routers) Users() {
	router.Protected.PUT("/users", users.Put)
}
