package education

import (
	"github.com/gin-gonic/gin"
	"resume/libs"
	"resume/models"
)

func Post(c *gin.Context) {
	s := libs.Context(c)

	education := models.Education{BaseModelWithRIDUID: models.BaseModelWithRIDUID{RID: s.Resume.ID, UID: s.User.ID}, Content: "请输入新的教育成绩"}

	libs.Create[models.Education](s, "教育经历", &education, models.Education{}, "id = ? AND uid = ? AND rid = ?", education.ID, s.User.ID, s.Resume.ID)
}
