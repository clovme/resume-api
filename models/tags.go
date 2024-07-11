package models

import (
	"resume/types/enums"
	"time"
)

// Tags 技能特长固定标签
type Tags struct {
	ID        uint       `gorm:"primaryKey" json:"id"`         // ID
	RID       uint       `gorm:"column:rid;not null" json:"-"` // 简历ID
	UID       uint       `gorm:"column:uid;not null" json:"-"` // 用户ID
	Type      enums.Tags `gorm:"column:type" json:"-"`         // 标识
	Name      string     `json:"name"`                         // 标签名称
	IsChecked bool       `json:"isChecked"`                    // 标签是否选中
	CreatedAt time.Time  `json:"created_at"`                   // 创建时间
	UpdatedAt time.Time  `json:"updated_at"`                   // 更新时间
}
