package setting

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

	var setting []models.Setting

	if err := c.ShouldBind(&setting); err != nil {
		log.Println("简历标题数据反序列化失败！")
		s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
		return
	}

	libs.Update[models.Setting]("简历设置", setting, s, func(model *models.Setting) *gorm.DB {
		model.UID = s.User.ID
		model.RID = s.Resume.ID
		return s.Model(&models.Setting{}).Where("id = ? AND rid = ? AND uid = ?", model.ID, s.Resume.ID, s.User.ID)
	})
}
