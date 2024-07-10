package types

import (
	"database/sql/driver"
	"encoding/json"
)

// Value 实现 `gorm.Valuer` 接口，这样 GORM 才知道如何存储该类型
func (j JSON) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan 实现 `sql.Scanner` 接口，这样 GORM 才知道如何从数据库读取该类型
func (j *JSON) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), j)
}
