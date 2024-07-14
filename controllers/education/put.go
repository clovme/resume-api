package education

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

	var education []models.Education
	if err := c.ShouldBind(&education); err != nil {
		log.Println("教育经历数据反序列化失败！", err)
		s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
		return
	}
	libs.Update[models.Education]("教育经历", education, s, func(model *models.Education) *gorm.DB {
		model.UID = s.User.ID
		model.RID = s.Resume.ID
		return s.Model(&models.Education{}).Where("id = ? AND rid = ? AND uid = ?", model.ID, s.Resume.ID, s.User.ID)
	})
}

func Sort(c *gin.Context) {
	libs.Sort[models.Works]("教育经历", c, models.Works{})
}
