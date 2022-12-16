package model

import (
	"github.com/jinzhu/gorm"
	modelKind "github.com/jnicklasj/gin-qor/models/kind"
	modelNode "github.com/jnicklasj/gin-qor/models/node"
)

type GroupItemBom struct {
	gorm.Model
	Name   string `gorm:"size:50"`
	Active bool
	Must   bool
	Sort   uint             `gorm:"default:1"`
	Kinds  []modelKind.Kind `gorm:"many2many:group_item_boms_kind;"` // many 2 many
	Nodes  []modelNode.Node `gorm:"many2many:group_item_boms_node;"` // many 2 many
}
