package model

import (
	modelGroupItemDetail "github.com/jnicklasj/gin-qor/models/group_item_detail"
	modelGroupItemYang "github.com/jnicklasj/gin-qor/models/group_item_yang"

	"github.com/jinzhu/gorm"
)

type GroupItemYangDetail struct {
	gorm.Model

	Details []modelGroupItemDetail.GroupItemDetail `gorm:"many2many:group_item_yang_detail_detail;"` // many 2 many
	Yangs   []modelGroupItemYang.GroupItemYang     `gorm:"many2many:group_item_yang_detail_yang;"`   // many 2 many
	Value   string                                 `gorm:"size:50"`
}
