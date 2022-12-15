package model

import (
	"github.com/jinzhu/gorm"
	modelGroupItemDetail "github.com/jnicklasj/gin-qor/models/group_item_detail"
	modelKind "github.com/jnicklasj/gin-qor/models/kind"
	modelNode "github.com/jnicklasj/gin-qor/models/node"
)

type TryOn struct {
	gorm.Model
	Name   string `gorm:"size:50"`
	Active bool
	Sort   uint                                   `gorm:"default:1"`
	Kinds  []modelKind.Kind                       `gorm:"many2many:try_on_kind;"` // many 2 many
	Nodes  []modelNode.Node                       `gorm:"many2many:try_on_node;"` // many 2 many
	Items  []modelGroupItemDetail.GroupItemDetail `gorm:"many2many:try_on_item;"` // many 2 many
}
