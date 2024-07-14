package skills

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"resume/libs"
	"resume/models"
	"resume/types/enums"
)

func Get(c *gin.Context) {
	s := libs.Context(c)

	var skills models.Skills

	if err := s.Where("rid = ? AND uid = ?", s.Resume.ID, s.User.ID).First(&skills).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := s.Where("rid = ? AND uid = ?", enums.Vx32, enums.Vx32).First(&skills).Error; err != nil {
				log.Println("没有查询到数据，并且查询异常:", err.Error())
				s.Msg(http.StatusNotFound, "没有查询到数据！")
				return
			}
			s.Json(http.StatusOK, "数据查询成功", skills)
			return
		}
		log.Println("没有查询到数据，并且查询异常:", err.Error())
		s.Msg(http.StatusNotFound, "没有查询到数据！")
		return
	}

	s.Json(http.StatusOK, "数据查询成功", skills)
}
