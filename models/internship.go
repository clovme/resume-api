package models

// Internship 实习经验
type Internship struct {
	BaseModelWithRIDUID
	StartAt string `json:"startAt"`                  // 实习经历开始时间
	EndAt   string `json:"endAt"`                    // 实习经历结束时间
	Name    string `json:"name"`                     // 在职公司名称
	Title   string `json:"title"`                    // 公司职位
	Content string `gorm:"type:text" json:"content"` // 工作内容
	ToNow   bool   `json:"toNow"`                    // 至今
	Sort    int    `json:"sort"`                     // 排序
}
