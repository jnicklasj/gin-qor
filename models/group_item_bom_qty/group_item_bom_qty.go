package model

import (
	modelGroupItemBom "github.com/jnicklasj/gin-qor/models/group_item_bom"
	modelGroupItemDetail "github.com/jnicklasj/gin-qor/models/group_item_detail"

	"github.com/jinzhu/gorm"
)

type GroupItemBomQty struct {
	gorm.Model

	Details []modelGroupItemDetail.GroupItemDetail `gorm:"many2many:group_item_boms_qty_detail;"` // many 2 many
	Boms    []modelGroupItemBom.GroupItemBom       `gorm:"many2many:group_item_boms_qty_bom;"`    // many 2 many
	Qty     float32
}
