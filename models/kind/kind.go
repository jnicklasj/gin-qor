package models

import (
	"github.com/jinzhu/gorm"
)

type Kind struct {
	gorm.Model
	Name   string `gorm:"size:50"`
	Active bool
}
