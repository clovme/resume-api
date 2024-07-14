package internship

import (
	"github.com/gin-gonic/gin"
	"resume/libs"
	"resume/models"
)

func Post(c *gin.Context) {
	s := libs.Context(c)

	internship := models.Internship{BaseModelWithRIDUID: models.BaseModelWithRIDUID{RID: s.Resume.ID, UID: s.User.ID}, Content: "请输入新的实习经验"}

	libs.Create[models.Internship](s, "实习经验", &internship, models.Internship{}, "id = ? AND uid = ? AND rid = ?", internship.ID, s.User.ID, s.Resume.ID)
}
