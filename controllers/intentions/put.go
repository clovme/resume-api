package intentions

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

	var intentions []models.Intentions
	if err := c.ShouldBind(&intentions); err != nil {
		log.Println("求职意向数据反序列化失败！", err)
		s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
		return
	}
	libs.Update[models.Intentions]("求职意向", intentions, s, func(model *models.Intentions) *gorm.DB {
		model.UID = s.User.ID
		model.RID = s.Resume.ID
		return s.Model(&models.Intentions{}).Where("id = ? AND rid = ? AND uid = ?", model.ID, s.Resume.ID, s.User.ID)
	})
}
