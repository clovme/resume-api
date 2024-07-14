package skills

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

	var skillsList []models.Skills

	if err := c.ShouldBind(&skillsList); err != nil {
		log.Println("技能特长数据反序列化失败！")
		s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
		return
	}

	libs.Update[models.Skills]("技能特长", skillsList, s, func(model *models.Skills) *gorm.DB {
		model.UID = s.User.ID
		model.RID = s.Resume.ID
		libs.SwitchTagsStatus(s.DB, s.Resume, model.TagsName, enums.Skills)
		return s.Model(&models.Skills{}).Where("id = ? AND rid = ? AND uid = ?", model.ID, s.Resume.ID, s.User.ID)
	})
}
