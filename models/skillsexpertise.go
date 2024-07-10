package models

import "time"

// SkillsExpertise 技能特长
type SkillsExpertise struct {
	ID        uint      `gorm:"primaryKey" json:"id"` // ID
	RID       uint      `gorm:"column:rid" json:"-"`  // 简历ID
	UID       uint      `gorm:"column:uid" json:"-"`  // 用户ID
	Content   string    `type:"TEXT" json:"content"`  // html内容
	CreatedAt time.Time `json:"created_at"`           // 创建时间
	UpdatedAt time.Time `json:"updated_at"`           // 更新时间
}

// SkillsExpertiseTags 技能特长固定标签
type SkillsExpertiseTags struct {
	ID        uint      `gorm:"primaryKey" json:"id"` // ID
	RID       uint      `gorm:"column:rid" json:"-"`  // 简历ID
	UID       uint      `gorm:"column:uid" json:"-"`  // 用户ID
	Name      string    `json:"name"`                 // 标签名称
	IsCheck   bool      `json:"isCheck"`              // 标签是否选中
	CreatedAt time.Time `json:"created_at"`           // 创建时间
	UpdatedAt time.Time `json:"updated_at"`           // 更新时间
}

// SkillsExpertiseCheckedTags 技能特长选中标签
type SkillsExpertiseCheckedTags struct {
	ID        uint      `json:"id"`                  // ID
	RID       uint      `gorm:"column:rid" json:"-"` // 简历ID
	UID       uint      `gorm:"column:uid" json:"-"` // 用户ID
	Name      string    `json:"name"`                // 标签名称
	Level     int       `json:"level"`               // 精通级别
	IsWord    bool      `json:"isWord"`              // 显示百分比或文字
	CreatedAt time.Time `json:"created_at"`          // 创建时间
	UpdatedAt time.Time `json:"updated_at"`          // 更新时间
}
