package applicationinfo

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"resume/libs"
	"resume/models"
)

func GradePost(c *gin.Context) {
	s := libs.Context(c)

	var grade models.CourseGrade

	if err := c.ShouldBind(&grade); err != nil {
		log.Println("基础信息数据反序列化失败！")
		s.Msg(http.StatusBadRequest, "数据保存失败，请重试！")
		return
	}

	grade.UID = s.User.ID
	grade.RID = s.Resume.ID

	libs.Create[models.CourseGrade](s, "课程成绩", &grade, models.CourseGrade{}, "uid = ? AND rid = ? AND name = ?", s.User.ID, s.Resume.ID, grade.Name)
}
