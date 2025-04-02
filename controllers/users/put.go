package users

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"resume/libs"
)

// Put 更新用户
func Put(c *gin.Context) {
	s := libs.Context(c)

	log.Println(s.User)
	s.Json(http.StatusOK, "用户", s.User)
}
