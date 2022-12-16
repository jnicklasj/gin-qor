package model

import (
	modelTryOn "github.com/jnicklasj/gin-qor/models/try_on"

	"github.com/jinzhu/gorm"
)

type TryOnCmt struct {
	gorm.Model
	Tryons []modelTryOn.TryOn `gorm:"many2many:try_ons_cmts_tryon;"` // many 2 many
	Cmt    float32
}
