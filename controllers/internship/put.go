package internship

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

	var internship []models.Internship
	if err := c.ShouldBind(&internship); err != nil {
		log.Println("实习经验数据反序列化失败！", err)
		s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
		return
	}
	libs.Update[models.Internship]("实习经验", internship, s, func(model *models.Internship) *gorm.DB {
		model.UID = s.User.ID
		model.RID = s.Resume.ID
		return s.Model(&models.Internship{}).Where("id = ? AND rid = ? AND uid = ?", model.ID, s.Resume.ID, s.User.ID)
	})
}

func Sort(c *gin.Context) {
	libs.Sort[models.Internship]("实习经验", c, models.Internship{})
}
