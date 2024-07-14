package hobbies

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

	var hobbies []models.Hobbies

	if err := c.ShouldBind(&hobbies); err != nil {
		log.Println("荣誉证书数据反序列化失败！")
		s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
		return
	}

	libs.Update[models.Hobbies]("荣誉证书", hobbies, s, func(model *models.Hobbies) *gorm.DB {
		model.UID = s.User.ID
		model.RID = s.Resume.ID
		libs.SwitchTagsStatus(s.DB, s.Resume, model.TagsName, enums.Hobbies)
		return s.Model(&models.Hobbies{}).Where("id = ? AND rid = ? AND uid = ?", model.ID, s.Resume.ID, s.User.ID)
	})
}
