package models

import "time"

// Resumes 简历表
type Resumes struct {
	ID        uint      `gorm:"primaryKey" json:"id"`         // 简历ID
	UID       uint      `gorm:"column:uid;not null" json:"-"` // 用户ID
	Name      string    `gorm:"size:20;not null" json:"name"` // 菜单名称
	CreatedAt time.Time `json:"created_at"`                   // 创建时间
	UpdatedAt time.Time `json:"updated_at"`                   // 更新时间
}
