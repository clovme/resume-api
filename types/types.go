package types

import "gorm.io/gorm"

type Callback func(tx *gorm.DB)

// JSON 定义一个自定义类型，用于 JSON 字段
type JSON map[string]interface{}
