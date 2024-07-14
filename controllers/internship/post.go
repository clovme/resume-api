package internship

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resume/libs"
	"resume/models"
)

func Post(c *gin.Context) {
	s := libs.Context(c)

	internship := models.Internship{BaseModelWithRIDUID: models.BaseModelWithRIDUID{RID: s.Resume.ID, UID: s.User.ID}, Content: "请输入新的实习经验"}
	if libs.Create[models.Internship](s, "实习经验", &internship, models.Internship{}, "id = ? AND uid = ? AND rid = ?", internship.ID, s.User.ID, s.Resume.ID) {
		s.Json(http.StatusOK, "数据保存成功！", internship)
		return
	}
	s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
}
