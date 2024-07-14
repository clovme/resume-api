package campus

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resume/libs"
	"resume/models"
)

func Delete(c *gin.Context) {
	s := libs.Context(c)

	if result := s.Delete(&models.Campus{}, "id = ? AND uid = ? AND rid = ?", c.Query("id"), s.User.ID, s.Resume.ID); result.Error != nil {
		s.Msg(http.StatusBadRequest, "删除错误，请重试！")
		return
	}
	s.Msg(http.StatusOK, "删除成功！")
}
