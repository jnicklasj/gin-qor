package helper

import (
	"strings"

	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"github.com/qor/qor/utils"
	"github.com/qor/validations"
)

var Validators = []*resource.Validator{
	{
		Name: "check_has_name", // register another validator with same name will overwirte previous one
		Handler: func(record interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
			// Get meta's value from metaValues, metaValues is a struct that hold all post data
			if meta := metaValues.Get("Name"); meta != nil {
				if name := utils.ToString(meta.Value); strings.TrimSpace(name) == "" {
					return validations.NewError(record, "Name", "Name can't be blank")
				}
			}
			return nil
		},
	},
}
