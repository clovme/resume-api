package project

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

	var project []models.Project
	if err := c.ShouldBind(&project); err != nil {
		log.Println("项目经验数据反序列化失败！", err)
		s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
		return
	}
	libs.Update[models.Project]("项目经验", project, s, func(model *models.Project) *gorm.DB {
		model.UID = s.User.ID
		model.RID = s.Resume.ID
		return s.Model(&models.Project{}).Where("id = ? AND rid = ? AND uid = ?", model.ID, s.Resume.ID, s.User.ID)
	})
}

func Sort(c *gin.Context) {
	libs.Sort[models.Project]("项目经验", c, models.Project{})
}
