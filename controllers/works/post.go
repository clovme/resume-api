package works

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resume/libs"
	"resume/models"
)

func Post(c *gin.Context) {
	s := libs.Context(c)

	works := models.Works{BaseModelWithRIDUID: models.BaseModelWithRIDUID{RID: s.Resume.ID, UID: s.User.ID}, Content: "请输入新的工作内容"}
	if libs.Create[models.Works](s, "工作经历", &works, models.Works{}, "id = ? AND uid = ? AND rid = ?", works.ID, s.User.ID, s.Resume.ID) {
		s.Json(http.StatusOK, "数据保存成功！", works)
		return
	}
	s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
}
