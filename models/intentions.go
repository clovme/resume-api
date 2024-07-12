package models

// Intentions 求职意向
type Intentions struct {
	BaseModelWithRIDUID
	Content string `gorm:"type:text" json:"content"` // 技能特长内容
}
