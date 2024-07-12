package models

// Menus 菜单
type Menus struct {
	BaseModelWithRIDUID
	Title      string `gorm:"size:10" json:"title"` // 菜单标题
	Name       string `gorm:"size:20" json:"name"`  // 菜单名称
	IsActivate bool   `json:"isActivate"`           // 激活选项
	IsOption   bool   `json:"isOption"`             // 是否可操作
	IsChecked  bool   `json:"isChecked"`            // 是否开启模块
	IsPage     bool   `json:"isPage"`               // 模块是否在页面显示
	Sort       int    `json:"sort"`                 // 排序
}
