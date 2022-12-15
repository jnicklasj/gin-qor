package model

import (
	"github.com/jinzhu/gorm"
)

type Group struct {
	gorm.Model
	Name   string `gorm:"size:50"`
	Sort   uint   `gorm:"default:1"`
	Active bool
}
