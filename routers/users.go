package routers

import (
	"resume/controllers/users"
)

func (r *Routers) Users() {
	r.Protected.PUT("/users", users.Put)
}
