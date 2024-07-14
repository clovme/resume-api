package education

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resume/libs"
	"resume/models"
)

func Post(c *gin.Context) {
	s := libs.Context(c)

	education := models.Education{BaseModelWithRIDUID: models.BaseModelWithRIDUID{RID: s.Resume.ID, UID: s.User.ID}, Content: "请输入新的教育成绩"}
	if libs.Create[models.Education](s, "教育经历", &education, models.Education{}, "id = ? AND uid = ? AND rid = ?", education.ID, s.User.ID, s.Resume.ID) {
		s.Json(http.StatusOK, "数据保存成功！", education)
		return
	}
	s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
}
