package models

import (
	"gorm.io/datatypes"
)

// Honors 荣誉证书
type Honors struct {
	BaseModelWithRIDUID
	Content     string         `gorm:"type:text" json:"content"`     // 内容
	CheckedTags datatypes.JSON `gorm:"type:json" json:"checkedTags"` // 选中的标签
	TagsName    []string       `gorm:"-" json:"tagsName"`            // 表单标签
}
