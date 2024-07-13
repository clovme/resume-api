package models

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        string    `gorm:"primaryKey;type:char(32)" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	hash := md5.Sum([]byte(uuid.New().String()))
	base.ID = hex.EncodeToString(hash[:])
	return
}

type BaseModelWithRIDUID struct {
	BaseModel
	RID string `gorm:"column:rid;type:char(32);not null" json:"-"` // 简历ID
	UID string `gorm:"column:uid;type:char(32);not null" json:"-"` // 用户ID
}
