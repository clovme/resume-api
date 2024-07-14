package works

import (
	"github.com/gin-gonic/gin"
	"resume/libs"
	"resume/models"
)

func Post(c *gin.Context) {
	s := libs.Context(c)

	works := models.Works{BaseModelWithRIDUID: models.BaseModelWithRIDUID{RID: s.Resume.ID, UID: s.User.ID}, Content: "请输入新的工作内容"}

	libs.Create[models.Works](s, "工作经历", &works, models.Works{}, "id = ? AND uid = ? AND rid = ?", works.ID, s.User.ID, s.Resume.ID)
}
