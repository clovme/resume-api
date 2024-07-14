package honors

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"resume/libs"
	"resume/models"
	"resume/types/enums"
)

func Put(c *gin.Context) {
	s := libs.Context(c)

	var honors []models.Honors

	if err := c.ShouldBind(&honors); err != nil {
		log.Println("荣誉证书数据反序列化失败！")
		s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
		return
	}

	libs.Update[models.Honors]("荣誉证书", honors, s, func(model *models.Honors) *gorm.DB {
		model.UID = s.User.ID
		model.RID = s.Resume.ID
		libs.SwitchTagsStatus(s.DB, s.Resume, model.TagsName, enums.Honors)
		return s.Model(&models.Honors{}).Where("id = ? AND rid = ? AND uid = ?", model.ID, s.Resume.ID, s.User.ID)
	})
}
