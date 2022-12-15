package model

import (
	"github.com/jinzhu/gorm"
	modelGroupItem "github.com/jnicklasj/gin-qor/models/group_item"
	"github.com/qor/media/oss"
)

type GroupItemDetail struct {
	gorm.Model
	Code        string `gorm:"size:50"`
	Name        string `gorm:"size:50"`
	Sort        uint   `gorm:"default:1"`
	Default     bool
	Active      bool
	ExtPrice    float32
	GroupsItems []modelGroupItem.GroupsItem `gorm:"many2many:group_item_detail_item;"` // many 2 many

	Image oss.OSS
}

func (groupItemDetail GroupItemDetail) MainImageURL(styles ...string) string {
	style := "main"
	if len(styles) > 0 {
		style = styles[0]
	}

	return groupItemDetail.Image.URL(style)
}
