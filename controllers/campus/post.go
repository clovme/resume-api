package campus

import (
	"github.com/gin-gonic/gin"
	"resume/libs"
	"resume/models"
)

func Post(c *gin.Context) {
	s := libs.Context(c)

	works := models.Campus{BaseModelWithRIDUID: models.BaseModelWithRIDUID{RID: s.Resume.ID, UID: s.User.ID}, Content: "请输入新的校园经历"}

	libs.Create[models.Campus](s, "校园经历", &works, models.Campus{}, "id = ? AND uid = ? AND rid = ?", works.ID, s.User.ID, s.Resume.ID)
}
