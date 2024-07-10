package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resume/libs"
	"resume/models"
)

func GetMenus(c *gin.Context) {
	s := libs.Context(c)
	rid := c.Query("rid")

	var resume models.Resumes
	if err := s.First(&resume, "id = ? and uid = ?", rid, s.User.ID).Error; err != nil {
		rid = "0"
		s.User.ID = 0
	}

	var menus []models.Menus
	if err := s.Find(&menus, "rid = ? and uid = ?", rid, s.User.ID).Error; err != nil {
		s.Msg(http.StatusNotFound, "没有查询到数据")
		return
	}
	s.Json(http.StatusOK, "数据查询成功", menus)
}
