package initdata

import (
	"gorm.io/gorm"
	"resume/models"
	"resume/types/enums"
)

// Tags 初始化标签
func (d *InitData) Tags() {
	var modelList []models.Tags

	for _, name := range []string{"Office软件", "沟通能力", "口才", "文案编辑", "数据分析", "推广运营", "产品设计", "JavaScript", "Python", "PHP", "NodeJs", "英语"} {
		modelList = append(modelList, models.Tags{BaseModelWithRIDUID: ridUID(), Type: enums.Skills, Name: name, IsChecked: false})
	}
	for _, name := range []string{"英语四级", "英语六级", "计算机一级", "计算机二级", "计算机三级", "计算机四级", "普通话一级", "普通话二级", "王者荣耀春季赛10强", "10大青年领军人才"} {
		modelList = append(modelList, models.Tags{BaseModelWithRIDUID: ridUID(), Type: enums.Honors, Name: name, IsChecked: false})
	}
	for _, name := range []string{"篮球", "跑步", "唱歌", "跳舞", "摄影", "旅行", "健身", "养宠物", "园艺", "桌游", "王者荣耀", "吃鸡"} {
		modelList = append(modelList, models.Tags{BaseModelWithRIDUID: ridUID(), Type: enums.Hobbies, Name: name, IsChecked: false})
	}

	insertRecord[models.Tags]("技能特长", modelList, func(model models.Tags) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ? and type = ? and name = ?", model.RID, model.UID, model.Type, model.Name)
	})
}
