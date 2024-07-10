package database

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"resume/models"
	"resume/types"
	"time"
)

// 初始化基础信息表
func initBasicInfoData(db *gorm.DB) {
	basicinfos := []models.BasicInfo{
		{RID: 0, UID: 0, Name: "柳如烟", IsAge: true, Birthday: "2000-07", Gender: "女", MaritalStatus: "未婚", Height: "175", Weight: "50", EthnicGroup: "汉族", NativePlace: "京都", PoliticalStatus: "群众", WorkExperienceYears: "18年", DesiredPosition: "渣遍天下无敌手", DesiredCity: "京都", DesiredSalary: "300W", StartDate: "一周内到岗", PhoneNumber: "13333333333", EmailAddress: "ruyan@qq.com", Photo: "", IsShowPhoto: true, CreatedAt: time.Now(), UpdatedAt: time.Now(), CustomInfo: types.JSON{"毕业时间": "2003-07"}},
	}

	for _, basicinfo := range basicinfos {
		var data models.BasicInfo

		if err := db.Where("rid = ? and uid = ?", basicinfo.RID, basicinfo.UID).First(&data).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// 数据不存在，插入新记录
				if err := db.Create(&basicinfo).Error; err != nil {
					log.Println("[初始化]基础信息创建失败:", err)
				} else {
					log.Println("[初始化]基础信息创建创建成功:", data)
				}
			} else {
				log.Println("[初始化]基础信息查询失败:", data)
			}
		}
	}
}
