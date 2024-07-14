package campus

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"resume/libs"
	"resume/models"
)

func Put(c *gin.Context) {
	s := libs.Context(c)

	var campus []models.Campus
	if err := c.ShouldBind(&campus); err != nil {
		log.Println("校园经历数据反序列化失败！", err)
		s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
		return
	}
	libs.Update[models.Campus]("校园经历", campus, s, func(model *models.Campus) *gorm.DB {
		model.UID = s.User.ID
		model.RID = s.Resume.ID
		return s.Model(&models.Campus{}).Where("id = ? AND rid = ? AND uid = ?", model.ID, s.Resume.ID, s.User.ID)
	})
}

func Sort(c *gin.Context) {
	libs.Sort[models.Campus]("校园经历", c, models.Campus{})
}
