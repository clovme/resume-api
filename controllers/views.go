package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetIndexView(c *gin.Context) {
	// 渲染HTML模板
	c.HTML(http.StatusOK, "index.html", nil)
}
