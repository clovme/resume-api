package models

// Setting 简历设置
type Setting struct {
	BaseModelWithRIDUID
	FontFamily string  `json:"fontFamily"` // 字体
	FontSize   int     `json:"fontSize"`   // 字体大小
	Module     int     `json:"module"`     // 模块间距
	Lines      float32 `json:"lines"`      // 行间距
	Page       int     `json:"page"`       // 页边距
}
