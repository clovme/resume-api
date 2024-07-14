package education

import (
	"github.com/gin-gonic/gin"
	"resume/libs"
	"resume/models"
)

func Get(c *gin.Context) {
	libs.QueryDataBaseArray[models.Education](c, []models.Education{})
}
