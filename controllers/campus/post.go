package campus

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resume/libs"
	"resume/models"
)

func Post(c *gin.Context) {
	s := libs.Context(c)

	works := models.Campus{BaseModelWithRIDUID: models.BaseModelWithRIDUID{RID: s.Resume.ID, UID: s.User.ID}, Content: "请输入新的校园经历"}
	if libs.Create[models.Campus](s, "校园经历", &works, models.Campus{}, "id = ? AND uid = ? AND rid = ?", works.ID, s.User.ID, s.Resume.ID) {
		s.Json(http.StatusOK, "数据保存成功！", works)
		return
	}
	s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
}
