package models

import "time"

// Users 用户表
type Users struct {
	ID        uint      `gorm:"primaryKey" json:"id"`         // ID
	Nickname  string    `gorm:"not null" json:"nickname"`     // 昵称
	Username  string    `gorm:"not null" json:"username"`     // 用户名
	Password  string    `gorm:"not null" json:"password"`     // 密码
	Token     string    `gorm:"unique;not null" json:"token"` // Token
	ExpiresAt time.Time `json:"expires_at"`                   // 过期时间
	CreatedAt time.Time `json:"created_at"`                   // 创建时间
	UpdatedAt time.Time `json:"updated_at"`                   // 更新时间
}
