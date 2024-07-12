package models

import (
	"gorm.io/datatypes"
)

// Basicinfo 基础信息
type Basicinfo struct {
	BaseModelWithRIDUID
	Name                string         `gorm:"size:20" json:"name"`         // 姓名
	IsAge               bool           `json:"isAge"`                       // 转年龄
	Birthday            string         `json:"birthday"`                    // 生日
	Gender              string         `json:"gender"`                      // 性别
	MaritalStatus       string         `json:"maritalStatus"`               // 婚姻状况
	Height              string         `json:"height"`                      // 身高
	Weight              string         `json:"weight"`                      // 体重
	EthnicGroup         string         `json:"ethnicGroup"`                 // 民族
	NativePlace         string         `json:"nativePlace"`                 // 籍贯
	PoliticalStatus     string         `json:"politicalStatus"`             // 政治面貌
	WorkExperienceYears string         `json:"workExperienceYears"`         // 工作年限
	DesiredPosition     string         `json:"desiredPosition"`             // 求职岗位
	DesiredCity         string         `json:"desiredCity"`                 // 意向城市
	DesiredSalary       string         `json:"desiredSalary"`               // 期望薪资
	StartDate           string         `json:"startDate"`                   // 入职时间
	PhoneNumber         string         `json:"phoneNumber"`                 // 电话
	EmailAddress        string         `json:"emailAddress"`                // 邮箱
	Photo               string         `json:"photo"`                       // 照片设置
	IsShowPhoto         bool           `json:"iShowPhoto"`                  // 展示照片
	CustomInfo          datatypes.JSON `gorm:"type:json" json:"customInfo"` // 自定义表单
}
