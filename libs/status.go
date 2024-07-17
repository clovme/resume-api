package libs

import (
	"embed"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"resume/models"
)

type HttpStatus struct {
	*gorm.DB     `json:"-"`
	*gin.Context `json:"-"`
	User         models.Users   `json:"-"`
	Resume       models.Resumes `json:"-"`
	Embed        *embed.FS      `json:"-"`
	Code         int            `json:"code"`
	Message      string         `json:"message"`
	Data         interface{}    `json:"data"`
}

func (s *HttpStatus) Msg(code int, message string) {
	s.Code = code
	s.Message = message
	s.Context.JSON(http.StatusOK, s)
}

func (s *HttpStatus) Json(code int, message string, data interface{}) {
	s.Code = code
	s.Message = message
	s.Data = data
	s.Context.JSON(http.StatusOK, s)
}

func Context(c *gin.Context) HttpStatus {
	return c.MustGet("$").(HttpStatus)
}
