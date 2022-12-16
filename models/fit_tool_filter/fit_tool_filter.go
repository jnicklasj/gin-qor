package model

import (
	modelFitTool "github.com/jnicklasj/gin-qor/models/fit_tool"
	modelSizeItem "github.com/jnicklasj/gin-qor/models/size_item"

	"github.com/jinzhu/gorm"
)

type FitToolsFilter struct {
	gorm.Model
	FitToolsItems []modelFitTool.FitTool   `gorm:"many2many:fit_tools_filters_item;"` // many 2 many
	SizesItems    []modelSizeItem.SizeItem `gorm:"many2many:fit_tools_size_item;"`    // many 2 many
	Percentage    float32
	Active        bool
	Sort          uint `gorm:"default:1"`
}
