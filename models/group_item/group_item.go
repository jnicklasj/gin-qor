package model

import (
	"github.com/jinzhu/gorm"
	modelGroup "github.com/jnicklasj/gin-qor/models/group"
	modelKind "github.com/jnicklasj/gin-qor/models/kind"
	modelNode "github.com/jnicklasj/gin-qor/models/node"
)

type GroupsItem struct {
	gorm.Model
	Name   string `gorm:"size:50"`
	Active bool
	Bom    bool
	Index  uint               `gorm:"default:1"`
	Sort   uint               `gorm:"default:1"`
	Kinds  []modelKind.Kind   `gorm:"many2many:group_item_kind;"`  // many 2 many
	Nodes  []modelNode.Node   `gorm:"many2many:group_item_node;"`  // many 2 many
	Groups []modelGroup.Group `gorm:"many2many:group_item_group;"` // many 2 many
}
