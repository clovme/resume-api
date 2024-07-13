package basicinfo

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"resume/libs"
	"resume/models"
)

func Put(c *gin.Context) {
	s := libs.Context(c)

	var basic models.Basicinfo
	basic.UID = s.User.ID
	basic.RID = s.Resume.ID

	if err := c.ShouldBind(&basic); err != nil {
		log.Println("基础信息数据反序列化失败！")
		s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
		return
	}

	if libs.UpdateNotCreatedAt[models.Basicinfo](&basic, s) {
		s.Msg(http.StatusOK, "数据保存成功！")
		return
	}
	s.Msg(http.StatusNotFound, "数据保存失败，请重试！")
}
