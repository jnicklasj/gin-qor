package model

import (
	modelKind "github.com/jnicklasj/gin-qor/models/kind"
	modelNode "github.com/jnicklasj/gin-qor/models/node"
	modelSizeItem "github.com/jnicklasj/gin-qor/models/size_item"

	"github.com/jinzhu/gorm"
)

type SizeGroup struct {
	gorm.Model
	Name   string `gorm:"size:50"`
	Active bool
	Sort   uint                     `gorm:"default:1"`
	Kinds  []modelKind.Kind         `gorm:"many2many:size_groups_kind;"` // many 2 many
	Nodes  []modelNode.Node         `gorm:"many2many:size_groups_node;"` // many 2 many
	Items  []modelSizeItem.SizeItem `gorm:"many2many:size_groups_item;"` // many 2 many
}
