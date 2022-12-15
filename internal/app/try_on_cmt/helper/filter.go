package helper

import (
	"github.com/jnicklasj/gin-qor/internal/library/qor/metax"

	"github.com/qor/admin"
)

var Filters = []*admin.Filter{
	// 选项匹配
	{

		Name:       "Groups",
		Label:      "分组",
		Operations: []string{"contains"},
		Type:       metax.SelectOne.String(),
		Config:     &admin.SelectOneConfig{Collection: []string{"款式设计", "工艺细节"}},
	},
	{

		Name:       "Active",
		Label:      "状态",
		Operations: []string{"contains"},
		Type:       metax.SelectOne.String(),
		Config:     &admin.SelectOneConfig{Collection: []string{"True", "False"}},
	},
	// 内容匹配
	{
		Name:  "Name",
		Label: "名称",
		Type:  "string",
		Visible: func(context *admin.Context) bool {
			return true
		},
	},
}
