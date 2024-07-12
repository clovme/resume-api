package models

// Evaluation 自我评价
type Evaluation struct {
	BaseModelWithRIDUID
	Content string `gorm:"type:text" json:"content"` // 自我评价内容
}
