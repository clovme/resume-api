package models

import "time"

// Menus 菜单
type Menus struct {
	ID         uint      `gorm:"primaryKey" json:"id"`         // 数据ID
	UID        uint      `gorm:"column:uid;not null" json:"-"` // 用户ID
	RID        uint      `gorm:"column:rid;not null" json:"-"` // 简历ID
	Title      string    `gorm:"size:10" json:"title"`         // 菜单标题
	Name       string    `gorm:"size:20" json:"name"`          // 菜单名称
	IsActivate bool      `json:"isActivate"`                   // 激活选项
	IsOption   bool      `json:"isOption"`                     // 是否可操作
	IsChecked  bool      `json:"isChecked"`                    // 是否开启模块
	IsPage     bool      `json:"isPage"`                       // 模块是否在页面显示
	Sort       int       `json:"sort"`                         // 排序
	CreatedAt  time.Time `json:"created_at"`                   // 创建时间
	UpdatedAt  time.Time `json:"updated_at"`                   // 更新时间
}
