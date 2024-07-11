package models

import (
	"resume/types"
	"time"
)

// Skills 技能特长
type Skills struct {
	ID          uint       `gorm:"primaryKey" json:"id"`         // ID
	RID         uint       `gorm:"column:rid" json:"-"`          // 简历ID
	UID         uint       `gorm:"column:uid" json:"-"`          // 用户ID
	Content     string     `gorm:"type:text" json:"content"`     // html内容
	CreatedAt   time.Time  `json:"created_at"`                   // 创建时间
	UpdatedAt   time.Time  `json:"updated_at"`                   // 更新时间
	CheckedTags types.JSON `gorm:"type:json" json:"checkedTags"` // 选中的标签
}
