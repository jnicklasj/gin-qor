package model

import (
	"github.com/jinzhu/gorm"
	modelKind "github.com/jnicklasj/gin-qor/models/kind"
	modelNode "github.com/jnicklasj/gin-qor/models/node"
)

type FitTool struct {
	gorm.Model
	Name   string `gorm:"size:50"`
	Max    float32
	Min    float32
	Step   float32
	Active bool
	Sort   uint             `gorm:"default:1"`
	Kinds  []modelKind.Kind `gorm:"many2many:fit_tools_kind;"` // many 2 many
	Nodes  []modelNode.Node `gorm:"many2many:fit_tools_node;"` // many 2 many
}
