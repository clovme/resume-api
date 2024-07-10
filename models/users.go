package models

import "time"

type Users struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Nickname  string    `json:"nickname"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Password  string    `json:"password"`
	Token     string    `gorm:"unique;not null" json:"token"` // Token
	ExpiresAt time.Time `json:"expires_at"`                   // 过期时间
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
