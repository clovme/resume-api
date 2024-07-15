package models

// Resumes 简历
type Resumes struct {
	BaseModel
	UID  string `gorm:"column:uid;type:char(32);not null" json:"-"` // 用户ID
	Name string `gorm:"size:20;not null" json:"name"`               // 简历名称
}
