package models

// Education 教育经历
type Education struct {
	BaseModelWithRIDUID
	StartAt string `json:"startAt"`                  // 教育经历开始时间
	EndAt   string `json:"endAt"`                    // 教育经历结束时间
	Name    string `json:"name"`                     // 学校名称
	Major   string `json:"major"`                    // 所学专业
	Content string `gorm:"type:text" json:"content"` // 工作内容
	ToNow   bool   `json:"toNow"`                    // 至今
	Degree  string `json:"degree"`                   // 学历
	Sort    int    `json:"sort"`                     // 排序
}
