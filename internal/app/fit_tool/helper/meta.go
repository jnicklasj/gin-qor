package helper

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
	{
		Label: "分组",
		Name:  "Nodes",
	},
	{
		Label: "最大调整值",
		Name:  "Max",
	},
	{
		Label: "最小调整值",
		Name:  "Min",
	},
	{
		Label: "调整步长",
		Name:  "Step",
	},
}
