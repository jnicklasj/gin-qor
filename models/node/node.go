package model

import (
	"github.com/jinzhu/gorm"
	modelKind "github.com/jnicklasj/gin-qor/models/kind"
)

type Node struct {
	gorm.Model
	Name   string `gorm:"size:50"`
	Active bool
	Sort   uint             `gorm:"default:1"`
	Kinds  []modelKind.Kind `gorm:"many2many:node_kind;"` // many 2 many
}
