package project

import (
	"github.com/gin-gonic/gin"
	"resume/libs"
	"resume/models"
)

func Post(c *gin.Context) {
	s := libs.Context(c)

	project := models.Project{BaseModelWithRIDUID: models.BaseModelWithRIDUID{RID: s.Resume.ID, UID: s.User.ID}, Content: "请输入新的项目经验"}

	libs.Create[models.Project](s, "项目经验", &project, models.Project{}, "id = ? AND uid = ? AND rid = ?", project.ID, s.User.ID, s.Resume.ID)
}
