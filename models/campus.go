package models

// Campus 校园经历
type Campus struct {
	BaseModelWithRIDUID
	StartAt string `json:"startAt"`                  // 开始时间
	EndAt   string `json:"endAt"`                    // 结束时间
	Name    string `json:"name"`                     // 学校名称
	Title   string `json:"title"`                    // 角色
	Content string `gorm:"type:text" json:"content"` // 校园描述
	ToNow   bool   `json:"toNow"`                    // 至今
	Sort    uint   `json:"sort"`                     // 排序
}
