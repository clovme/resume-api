package models

// Project 项目经验
type Project struct {
	BaseModelWithRIDUID
	StartAt string `json:"startAt"`                  // 项目经验开始时间
	EndAt   string `json:"endAt"`                    // 项目经验结束时间
	Name    string `json:"name"`                     // 在职公司名称
	Title   string `json:"title"`                    // 公司职位
	Content string `gorm:"type:text" json:"content"` // 工作内容
	ToNow   bool   `json:"toNow"`                    // 至今
	Sort    int    `json:"sort"`                     // 排序
}
