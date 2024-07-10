package libs

import (
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
	Code         int            `json:"code"`
	Message      string         `json:"message"`
	Data         interface{}    `json:"data"`
}

func (t *HttpStatus) Msg(code int, message string) {
	t.Code = code
	t.Message = message
	t.Context.JSON(http.StatusOK, t)
}

func (t *HttpStatus) Json(code int, message string, data interface{}) {
	t.Code = code
	t.Message = message
	t.Data = data
	t.Context.JSON(http.StatusOK, t)
}

func Context(c *gin.Context) HttpStatus {
	return c.MustGet("$").(HttpStatus)
}
