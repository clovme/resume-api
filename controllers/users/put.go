package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"resume/libs"
)

// Put 更新用户
func Put(c *gin.Context) {
	s := libs.Context(c)

	fmt.Println(s.User)
	s.Json(http.StatusOK, "用户", s.User)
}
