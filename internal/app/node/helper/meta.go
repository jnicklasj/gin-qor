package model

import (
	"github.com/qor/admin"
)

var Metas = []*admin.Meta{
	{
		Label: "编号",
		Name:  "ID",
	},
	{
		Label: "名称",
		Name:  "Name",
	},
	{
		Label: "状态",
		Name:  "Active",
	},
	{
		Label: "排序",
		Name:  "Sort",
	},
	{
		Label: "类别",
		Name:  "Kinds",
	},
}
