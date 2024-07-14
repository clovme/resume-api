package models

import (
	"gorm.io/datatypes"
)

// Hobbies 兴趣爱好
type Hobbies struct {
	BaseModelWithRIDUID
	Content     string         `gorm:"type:text" json:"content"`     // 内容
	CheckedTags datatypes.JSON `gorm:"type:json" json:"checkedTags"` // 选中的标签
	TagsName    []string       `gorm:"-" json:"tagsName"`            // 表单标签
}
