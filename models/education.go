package models

import "time"

// Education 教育经历
type Education struct {
	ID        uint      `gorm:"primaryKey" json:"id"`         // ID
	RID       uint      `gorm:"column:rid;not null" json:"-"` // 简历ID
	UID       uint      `gorm:"column:uid;not null" json:"-"` // 用户ID
	StartAt   string    `json:"startAt"`                      // 教育经历开始时间
	EndAt     string    `json:"endAt"`                        // 教育经历结束时间
	Name      string    `json:"name"`                         // 学校名称
	Major     string    `json:"major"`                        // 所学专业
	Content   string    `gorm:"type:text" json:"content"`     // 工作内容
	ToNow     bool      `json:"toNow"`                        // 至今
	Degree    string    `json:"degree"`                       // 学历
	Sort      uint      `json:"sort"`                         // 排序
	CreatedAt time.Time `json:"createdAt"`                    // 创建时间
	UpdatedAt time.Time `json:"updatedAt"`                    // 更新时间
}
