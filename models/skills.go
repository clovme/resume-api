package models

import (
	"gorm.io/datatypes"
)

// Skills 技能特长
type Skills struct {
	BaseModelWithRIDUID
	Content     string         `gorm:"type:text" json:"content"`     // 技能特长内容
	CheckedTags datatypes.JSON `gorm:"type:json" json:"checkedTags"` // 选中的标签
	TagsName    []string       `gorm:"-" json:"tagsName"`            // 表单标签
}
