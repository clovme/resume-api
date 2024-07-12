package models

import (
	"resume/types"
)

// Skills 技能特长
type Skills struct {
	BaseModelWithRIDUID
	Content     string     `gorm:"type:text" json:"content"`     // 技能特长内容
	CheckedTags types.JSON `gorm:"type:json" json:"checkedTags"` // 选中的标签
}
