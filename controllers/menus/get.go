package menus

import (
	"github.com/gin-gonic/gin"
	"resume/libs"
	"resume/models"
)

func Get(c *gin.Context) {
	libs.QueryDataBaseArray[models.Menus](c, []models.Menus{})
}
