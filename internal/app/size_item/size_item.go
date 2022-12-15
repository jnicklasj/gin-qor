package size_item

import (
	"github.com/jinzhu/gorm"
	helpers "github.com/jnicklasj/gin-qor/internal/app/size_item/helper"
	models "github.com/jnicklasj/gin-qor/models/size_item"
	"github.com/qor/admin"
)

func Setup(Db *gorm.DB, Admin *admin.Admin) {

	var r = Admin.AddResource(
		&models.SizeItem{},
		&admin.Config{
			Menu: []string{"Settings 规格"},
			Name: "规格单项目",
		},
	)

	if len(helpers.NewAttrs) > 0 {
		r.NewAttrs(helpers.NewAttrs)
	}

	if len(helpers.EditAttrs) > 0 {
		r.EditAttrs(helpers.EditAttrs)
	}

	if len(helpers.IndexAttrs) > 0 {
		r.IndexAttrs(helpers.IndexAttrs)
	}

	if len(helpers.SearchAttrs) > 0 {
		r.SearchAttrs(helpers.SearchAttrs...)
	}

	// Search
	//oldOrderSearchHandler := orderResource.SearchHandler
	//orderResource.SearchHandler = func(keyword string, context *qor.Context) *gorm.DB {
	//	if strings.Contains(keyword, " ") {
	//		return oldOrderSearchHandler("", context).Where("id IN (?)", strings.Split(keyword, " "))
	//	}
	//
	//	return oldOrderSearchHandler(keyword, context)
	//}

	// Meta
	for _, meta := range helpers.Metas {
		r.Meta(meta)
	}

	// Filter
	for _, filter := range helpers.Filters {
		r.Filter(filter)
	}

	// Action
	//for _, action := range model.Actions {
	//	nodesResource.Action(action)
	//}

	// Validator
	for _, validator := range helpers.Validators {
		r.AddValidator(validator)
	}

}