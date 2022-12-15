package model

import (
	modelGroupItem "github.com/jnicklasj/gin-qor/models/group_item"
	modelGroupItemDetail "github.com/jnicklasj/gin-qor/models/group_item_detail"

	"github.com/jinzhu/gorm"
)

type GroupItemDetailFilter struct {
	gorm.Model
	FiltersAt    []modelGroupItemDetail.GroupItemDetail `gorm:"many2many:group_item_detail_filter_at;"`    // many 2 many
	FiltersGroup []modelGroupItem.GroupsItem            `gorm:"many2many:group_item_detail_filter_group;"` // many 2 many
	FiltersFor   []modelGroupItemDetail.GroupItemDetail `gorm:"many2many:group_item_detail_filter_for;"`   // many 2 many
	Active       bool
	Sort         uint `gorm:"default:1"`
}
