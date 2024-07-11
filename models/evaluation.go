package models

import "time"

// Evaluation 自我评价
type Evaluation struct {
	ID        uint      `gorm:"primaryKey" json:"id"`         // ID
	RID       uint      `gorm:"column:rid;not null" json:"-"` // 简历ID
	UID       uint      `gorm:"column:uid;not null" json:"-"` // 用户ID
	Content   string    `gorm:"type:text" json:"content"`     // 自我评价内容
	CreatedAt time.Time `json:"createdAt"`                    // 创建时间
	UpdatedAt time.Time `json:"updatedAt"`                    // 更新时间
}
