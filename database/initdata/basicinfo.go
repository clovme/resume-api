package initdata

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"resume/models"
)

// BasicInfo 初始化基础信息表
func (d *InitData) BasicInfo() {
	modelList := []models.Basicinfo{
		{BaseModelWithRIDUID: ridUID(), Name: "柳如烟", IsAge: true, Birthday: "2000-07", Gender: "女", MaritalStatus: "未婚", Height: "175", Weight: "50", EthnicGroup: "汉族", NativePlace: "京都", PoliticalStatus: "群众", WorkExperienceYears: "18年", DesiredPosition: "渣遍天下无敌手", DesiredCity: "京都", DesiredSalary: "300W", StartDate: "一周内到岗", PhoneNumber: "13333333333", EmailAddress: "ruyan@qq.com", Photo: "", IsShowPhoto: true, CustomInfo: datatypes.JSON([]byte(`{"毕业时间": "2003-07"}`))},
	}
	insertRecord[models.Basicinfo]("基础信息", modelList, func(model models.Basicinfo) (db, where *gorm.DB) {
		return d.Db, d.Db.Where("rid = ? and uid = ?", model.RID, model.UID)
	})
}
