package models

import "time"

// Works 工作经历
type Works struct {
	ID        uint      `gorm:"primaryKey" json:"id"`         // ID
	RID       uint      `gorm:"column:rid;not null" json:"-"` // 简历ID
	UID       uint      `gorm:"column:uid;not null" json:"-"` // 用户ID
	StartAt   string    `json:"startAt"`                      // 工作经历开始时间
	EndAt     string    `json:"endAt"`                        // 工作经历结束时间
	Name      string    `json:"name"`                         // 在职公司名称
	Title     string    `json:"title"`                        // 公司职位
	Content   string    `gorm:"type:text" json:"content"`     // 工作内容
	ToNow     bool      `json:"toNow"`                        // 至今
	Sort      uint      `json:"sort"`                         // 排序
	CreatedAt time.Time `json:"createdAt"`                    // 创建时间
	UpdatedAt time.Time `json:"updatedAt"`                    // 更新时间
}
