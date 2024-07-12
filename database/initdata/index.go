package initdata

import (
	"gorm.io/gorm"
	"resume/models"
	"resume/types"
)

type InitData struct {
	Db *gorm.DB
}

func (_ *InitData) ridUID() models.BaseModelWithRIDUID {
	return models.BaseModelWithRIDUID{RID: types.Ox320, UID: types.Ox320}
}
