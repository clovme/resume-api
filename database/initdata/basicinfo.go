package initdata

import (
	"errors"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"log"
	"resume/models"
)

// BasicInfo 初始化基础信息表
func (d *InitData) BasicInfo() {
	modelList := []models.Basicinfo{
		{BaseModelWithRIDUID: d.ridUID(), Name: "柳如烟", IsAge: true, Birthday: "2000-07", Gender: "女", MaritalStatus: "未婚", Height: "175", Weight: "50", EthnicGroup: "汉族", NativePlace: "京都", PoliticalStatus: "群众", WorkExperienceYears: "18年", DesiredPosition: "渣遍天下无敌手", DesiredCity: "京都", DesiredSalary: "300W", StartDate: "一周内到岗", PhoneNumber: "13333333333", EmailAddress: "ruyan@qq.com", Photo: "", IsShowPhoto: true, CustomInfo: datatypes.JSON([]byte(`{"毕业时间": "2003-07"}`))},
	}

	for _, model := range modelList {
		if err := d.Db.Where("rid = ? and uid = ?", model.RID, model.UID).First(&models.Basicinfo{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 数据不存在，插入新记录
				if err := d.Db.Create(&model).Error; err != nil {
					log.Fatalln("[初始化]基础信息创建失败:", err)
				} else {
					log.Println("[初始化]基础信息创建创建成功:", model)
				}
			} else {
				log.Fatalln("[初始化]基础信息查询失败:", err)
			}
		}
	}
}
