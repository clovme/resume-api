package applicationinfo

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"resume/libs"
	"resume/models"
)

func Put(c *gin.Context) {
	s := libs.Context(c)

	var applicationinfo []models.ApplicationInfo

	if err := c.ShouldBind(&applicationinfo); err != nil {
		log.Println("报考信息数据反序列化失败！")
		s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
		return
	}

	libs.Update[models.ApplicationInfo]("报考信息", applicationinfo, s, func(model *models.ApplicationInfo) *gorm.DB {
		model.UID = s.User.ID
		model.RID = s.Resume.ID
		return s.Model(&models.ApplicationInfo{}).Where("id = ? AND uid = ? AND rid = ?", model.ID, s.User.ID, s.Resume.ID)
	})
}

func GradePut(c *gin.Context) {
	s := libs.Context(c)

	var grade []models.CourseGrade

	if err := c.ShouldBind(&grade); err != nil {
		log.Println("课程分数数据反序列化失败！")
		s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
		return
	}

	libs.Update[models.CourseGrade]("课程分数", grade, s, func(model *models.CourseGrade) *gorm.DB {
		model.UID = s.User.ID
		model.RID = s.Resume.ID
		return s.Model(&models.CourseGrade{}).Where("id = ? AND uid = ? AND rid = ?", model.ID, s.User.ID, s.Resume.ID)
	})
}
