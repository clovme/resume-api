package models

// Resumes 简历表
type Resumes struct {
	BaseModel
	UID  string `gorm:"column:uid;type:char(32);not null" json:"-"` // 用户ID
	Name string `gorm:"size:20;not null" json:"name"`               // 菜单名称
}
