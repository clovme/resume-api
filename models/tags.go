package models

import (
	"resume/types/enums"
)

// Tags 技能特长固定标签
type Tags struct {
	BaseModelWithRIDUID
	Type      enums.Tags `gorm:"column:type" json:"-"` // 标识
	Name      string     `json:"name"`                 // 标签名称
	IsChecked bool       `json:"isChecked"`            // 标签是否选中
}
