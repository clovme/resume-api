package evaluation

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

	var evaluation []models.Evaluation
	if err := c.ShouldBind(&evaluation); err != nil {
		log.Println("自我教育数据反序列化失败！", err)
		s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
		return
	}
	libs.Update[models.Evaluation]("自我教育", evaluation, s, func(model *models.Evaluation) *gorm.DB {
		model.UID = s.User.ID
		model.RID = s.Resume.ID
		return s.Model(&models.Evaluation{}).Where("id = ? AND rid = ? AND uid = ?", model.ID, s.Resume.ID, s.User.ID)
	})
}
