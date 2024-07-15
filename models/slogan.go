package models

// Slogan 简历标题
type Slogan struct {
	BaseModelWithRIDUID
	Title  string `gorm:"not null" json:"title"`  // 简历名称
	Slogan string `gorm:"not null" json:"slogan"` // 简历口号
}
