package node

import (
	helpers "github.com/jnicklasj/gin-qor/internal/app/node/helper"
	models "github.com/jnicklasj/gin-qor/models/node"

	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
)

func Setup(Db *gorm.DB, Admin *admin.Admin) {

	var r = Admin.AddResource(
		&models.Node{},
		&admin.Config{
			Menu: []string{"Settings 结构"},
			Name: "品类",
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
	//	r.Action(action)
	//}

	// Validator
	for _, validator := range helpers.Validators {
		r.AddValidator(validator)
	}

	//order.SaveHandler = func(result interface{}, cont *qor.Context) error {
	//	trx := apm.DefaultTracer.StartTransaction(fmt.Sprintf("%v %v", cont.Request.Method, cont.Request.RequestURI), "request")
	//	trxCtx := apm.ContextWithTransaction(cont.Request.Context(), trx)
	//	defer trx.End()
	//	if (cont.GetDB().NewScope(result).PrimaryKeyZero() &&
	//		order.HasPermission(roles.Create, cont)) || // has create permission
	//		order.HasPermission(roles.Update, cont) { // has update permission
	//
	//		mediaFile, _ := result.(*model.Order)
	//		if httpreq.HasMediaFile(cont.Request) {
	//			_, err := o.minioClient.UploadFileToMinioFromRequest(trxCtx, cont.Request, mediaFile.File)
	//			if err != nil {
	//				return err
	//			}
	//		}
	//
	//		if err := cont.GetDB().Save(result).Error; err != nil {
	//			return err
	//		}
	//		return nil
	//	}
	//	return roles.ErrPermissionDenied
	//}
	//
	//order.DeleteHandler = func(result interface{}, context *qor.Context) error {
	//	transaction := apm.DefaultTracer.StartTransaction(fmt.Sprintf("%v %v", context.Request.Method, context.Request.RequestURI), "request")
	//	defer transaction.End()
	//	if order.HasPermission(roles.Delete, context) {
	//		if primaryQuerySQL, primaryParams := order.ToPrimaryQueryParams(context.ResourceID, context); primaryQuerySQL != "" {
	//			if !context.GetDB().First(result, append([]interface{}{primaryQuerySQL}, primaryParams...)...).RecordNotFound() {
	//				err := context.GetDB().Delete(result).Error
	//				return err
	//			}
	//		}
	//		return gorm.ErrRecordNotFound
	//	}
	//	return roles.ErrPermissionDenied
	//}
	//
	//order.FindOneHandler = func(result interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
	//	if order.HasPermission(roles.Read, context) {
	//		var (
	//			primaryQuerySQL string
	//			primaryParams   []interface{}
	//		)
	//		if metaValues == nil {
	//			primaryQuerySQL, primaryParams = order.ToPrimaryQueryParams(context.ResourceID, context)
	//		} else {
	//			primaryQuerySQL, primaryParams = order.ToPrimaryQueryParamsFromMetaValue(metaValues, context)
	//		}
	//
	//		if primaryQuerySQL != "" {
	//			if metaValues != nil {
	//				if destroy := metaValues.Get("_destroy"); destroy != nil {
	//					if fmt.Sprint(destroy.Value) != "0" && order.HasPermission(roles.Delete, context) {
	//						context.GetDB().Delete(result, append([]interface{}{primaryQuerySQL}, primaryParams...)...)
	//						return resource.ErrProcessorSkipLeft
	//					}
	//				}
	//			}
	//			err := context.GetDB().First(result, append([]interface{}{primaryQuerySQL}, primaryParams...)...).Error
	//			return err
	//		}
	//		return errors.New("failed to find")
	//	}
	//	return roles.ErrPermissionDenied
	//}
	//
	//order.FindManyHandler = func(result interface{}, context *qor.Context) error {
	//	if order.HasPermission(roles.Read, context) {
	//		db := context.GetDB()
	//		if _, ok := db.Get("qorm:getting_total_count"); ok {
	//			if err := context.GetDB().Count(result).Error; err != nil {
	//				return err
	//			}
	//			return nil
	//		}
	//		if err := context.GetDB().Set("gorm:order_by_primary_key", "DESC").Find(result).Error; err != nil {
	//			return err
	//		}
	//		return nil
	//	}
	//	return roles.ErrPermissionDenied
	//}

}
