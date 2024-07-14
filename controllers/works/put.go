package works

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

	var works []models.Works
	if err := c.ShouldBind(&works); err != nil {
		log.Println("工作经历数据反序列化失败！", err)
		s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
		return
	}
	libs.Update[models.Works]("工作经历", works, s, func(model *models.Works) *gorm.DB {
		model.UID = s.User.ID
		model.RID = s.Resume.ID
		return s.Model(&models.Works{}).Where("id = ? AND rid = ? AND uid = ?", model.ID, s.Resume.ID, s.User.ID)
	})
}

func Sort(c *gin.Context) {
	libs.Sort[models.Works]("工作经历", c, models.Works{})
}
