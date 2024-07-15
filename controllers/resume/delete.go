package resume

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"resume/libs"
	"resume/models"
	"resume/types"
	"strings"
)

func Delete(c *gin.Context) {
	s := libs.Context(c)
	rid := c.Query("rid")
	var resume models.Resumes

	if len(strings.Trim(rid, "")) <= 0 {
		s.Msg(http.StatusNotFound, "简历标识不能为空！")
		return
	}

	if err := s.First(&resume, "uid = ? AND id = ?", s.User.ID, rid).Error; err != nil {
		s.Msg(http.StatusNotFound, "删除失败，简历名称不存在！")
		return
	}

	libs.Transaction(s, func(tx *gorm.DB) {
		// 删除简历
		if !libs.DeleteData(s, tx, "简历", models.Resumes{}, "id = ? AND uid = ?", rid, s.User.ID) {
			return
		}

		var modelList []types.ModelItem
		modelList = append(modelList, types.ModelItem{Name: "菜单", Model: models.Menus{}})
		modelList = append(modelList, types.ModelItem{Name: "标签", Model: models.Tags{}})
		modelList = append(modelList, types.ModelItem{Name: "基础信息", Model: models.Basicinfo{}})
		modelList = append(modelList, types.ModelItem{Name: "技能特长", Model: models.Skills{}})
		modelList = append(modelList, types.ModelItem{Name: "工作经历", Model: models.Works{}})
		modelList = append(modelList, types.ModelItem{Name: "项目经验", Model: models.Project{}})
		modelList = append(modelList, types.ModelItem{Name: "教育经历", Model: models.Education{}})
		modelList = append(modelList, types.ModelItem{Name: "自我评价", Model: models.Evaluation{}})
		modelList = append(modelList, types.ModelItem{Name: "实习经验", Model: models.Internship{}})
		modelList = append(modelList, types.ModelItem{Name: "报考信息", Model: models.ApplicationInfo{}})
		modelList = append(modelList, types.ModelItem{Name: "校园经历", Model: models.Campus{}})
		modelList = append(modelList, types.ModelItem{Name: "荣誉证书", Model: models.Honors{}})
		modelList = append(modelList, types.ModelItem{Name: "兴趣爱好", Model: models.Hobbies{}})
		modelList = append(modelList, types.ModelItem{Name: "求职意向", Model: models.Intentions{}})
		modelList = append(modelList, types.ModelItem{Name: "课程成绩", Model: models.CourseGrade{}})
		modelList = append(modelList, types.ModelItem{Name: "简历标题", Model: models.Slogan{}})
		modelList = append(modelList, types.ModelItem{Name: "简历设置", Model: models.Setting{}})

		for _, model := range modelList {
			if !libs.DeleteData(s, tx, model.Name, model.Model, "uid = ? AND rid = ?", resume.UID, resume.ID) {
				return
			}
		}

		var resumes []models.Resumes
		if err := s.Find(&resumes, "uid = ?", s.User.ID).Error; err != nil {
			s.Msg(http.StatusNotFound, "暂无数据")
			return
		}

		s.Json(http.StatusOK, "数据删除成功", resumes)
	})
}
