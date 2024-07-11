package tags

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resume/libs"
	"resume/models"
)

func Get(c *gin.Context) {
	s := libs.Context(c)

	_type := c.Query("type")

	var tags []models.Tags
	if err := s.Find(&tags, "rid = ? and uid = ? and type = ?", s.Resume.ID, s.User.ID, _type).Limit(12).Error; err != nil {
		s.Msg(http.StatusNotFound, "没有查询到基础信息！")
		return
	}
	if len(tags) <= 0 {
		if err := s.Find(&tags, "rid = 0 and uid = 0 and type = ?", _type).Limit(12).Error; err != nil {
			s.Msg(http.StatusNotFound, "没有查询到基础信息！")
			return
		}
	}
	s.Json(http.StatusOK, "数据查询成功", tags)
}
