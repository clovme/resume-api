package basicinfo

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

	var basicList []models.Basicinfo

	if err := c.ShouldBind(&basicList); err != nil {
		log.Println("基础信息数据反序列化失败！")
		s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
		return
	}

	libs.Update[models.Basicinfo]("基础信息", basicList, s, func(model *models.Basicinfo) *gorm.DB {
		model.UID = s.User.ID
		model.RID = s.Resume.ID
		return s.Model(&models.Basicinfo{}).Where("id = ? AND uid = ? AND rid = ?", model.ID, s.User.ID, s.Resume.ID)
	})
}
