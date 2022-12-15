package internal

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/jnicklasj/gin-qor/config"
	"github.com/jnicklasj/gin-qor/config/bindatafs"
	appFitTool "github.com/jnicklasj/gin-qor/internal/app/fit_tool"
	appFitToolFilter "github.com/jnicklasj/gin-qor/internal/app/fit_tool_filter"
	appGroup "github.com/jnicklasj/gin-qor/internal/app/group"
	appGroupItem "github.com/jnicklasj/gin-qor/internal/app/group_item"
	appGroupItemBom "github.com/jnicklasj/gin-qor/internal/app/group_item_bom"
	appGroupItemBomQty "github.com/jnicklasj/gin-qor/internal/app/group_item_bom_qty"
	appGroupItemDetail "github.com/jnicklasj/gin-qor/internal/app/group_item_detail"
	appGroupItemDetailDefault "github.com/jnicklasj/gin-qor/internal/app/group_item_detail_default"
	appGroupItemDetailFilter "github.com/jnicklasj/gin-qor/internal/app/group_item_detail_filter"
	appGroupItemYang "github.com/jnicklasj/gin-qor/internal/app/group_item_yang"
	appGroupItemYangDetail "github.com/jnicklasj/gin-qor/internal/app/group_item_yang_detail"
	appKind "github.com/jnicklasj/gin-qor/internal/app/kind"
	appNode "github.com/jnicklasj/gin-qor/internal/app/node"
	appNotification "github.com/jnicklasj/gin-qor/internal/app/notification"
	appSizeGroup "github.com/jnicklasj/gin-qor/internal/app/size_group"
	appSizeItem "github.com/jnicklasj/gin-qor/internal/app/size_item"
	appTryOn "github.com/jnicklasj/gin-qor/internal/app/try_on"
	appTryOnCmt "github.com/jnicklasj/gin-qor/internal/app/try_on_cmt"
	modelFitTool "github.com/jnicklasj/gin-qor/models/fit_tool"
	modelFitToolFitler "github.com/jnicklasj/gin-qor/models/fit_tool_filter"
	modelGroup "github.com/jnicklasj/gin-qor/models/group"
	modelGroupItem "github.com/jnicklasj/gin-qor/models/group_item"
	modelGroupItemBom "github.com/jnicklasj/gin-qor/models/group_item_bom"
	modelGroupItemBomQty "github.com/jnicklasj/gin-qor/models/group_item_bom_qty"
	modelGroupItemDetail "github.com/jnicklasj/gin-qor/models/group_item_detail"
	modelGroupItemDetailDefault "github.com/jnicklasj/gin-qor/models/group_item_detail_default"
	modelGroupItemDetailFilter "github.com/jnicklasj/gin-qor/models/group_item_detail_filter"
	modelGroupItemYang "github.com/jnicklasj/gin-qor/models/group_item_yang"
	modelGroupItemYangDetail "github.com/jnicklasj/gin-qor/models/group_item_yang_detail"
	modelKind "github.com/jnicklasj/gin-qor/models/kind"
	modelNode "github.com/jnicklasj/gin-qor/models/node"
	modelProduct "github.com/jnicklasj/gin-qor/models/product"
	modelSizeGroup "github.com/jnicklasj/gin-qor/models/size_group"
	modelSizeItem "github.com/jnicklasj/gin-qor/models/size_item"
	modelTryOn "github.com/jnicklasj/gin-qor/models/try_on"
	modelTryOnCmt "github.com/jnicklasj/gin-qor/models/try_on_cmt"
	modelUser "github.com/jnicklasj/gin-qor/models/user"
	"github.com/qor/admin"
	"github.com/qor/media"
	"github.com/qor/qor"
	"github.com/qor/roles"
)

func init() {
	roles.Register("administrator", func(req *http.Request, currentUser interface{}) bool {
		return currentUser != nil && currentUser.(*modelUser.User).Role == "Admin"
	})
}

// NewDummyAdmin generate admin for dummy app
func NewDummyAdmin(DB *gorm.DB, keepData ...bool) *admin.Admin {
	var (
		models = []interface{}{
			&modelProduct.Product{},
			&modelKind.Kind{},
			&modelNode.Node{},
			&modelGroup.Group{},
			&modelGroupItem.GroupsItem{},
			&modelGroupItemDetail.GroupItemDetail{},
			&modelGroupItemBom.GroupItemBom{},
			&modelGroupItemBomQty.GroupItemBomQty{},
			&modelGroupItemYang.GroupItemYang{},
			&modelGroupItemYangDetail.GroupItemYangDetail{},
			&modelGroupItemDetailFilter.GroupItemDetailFilter{},
			&modelGroupItemDetailDefault.GroupItemDetailDefault{},
			&modelSizeItem.SizeItem{},
			&modelSizeGroup.SizeGroup{},
			&modelFitTool.FitTool{},
			&modelFitToolFitler.FitToolFilter{},
			&modelTryOn.TryOn{},
			&modelTryOnCmt.TryOnCmt{},
		}
		Admin = admin.New(&admin.AdminConfig{DB: DB, Auth: AdminAuth{}, AssetFS: bindatafs.AssetFS.NameSpace("admin")})
	)

	// AssetFs RegisterPath
	Admin.AssetFS.RegisterPath("assets/views")
	// OSS
	media.RegisterCallbacks(DB)
	// AutoMigrate
	for _, model := range models {
		if len(keepData) == 0 {
			DB.DropTableIfExists(model)
		}
		// AutoMigrate 只能新增不存在的字段，不能修改已经存在的字段
		// will only add missing fields, won't delete/change current data
		DB.AutoMigrate(model)
	}

	// Register Apps
	appKind.Setup(DB, Admin)
	appNode.Setup(DB, Admin)
	appGroup.Setup(DB, Admin)
	appGroupItem.Setup(DB, Admin)
	appGroupItemDetail.Setup(DB, Admin)
	appGroupItemBom.Setup(DB, Admin)
	appGroupItemBomQty.Setup(DB, Admin)
	appGroupItemYang.Setup(DB, Admin)
	appGroupItemYangDetail.Setup(DB, Admin)
	appGroupItemDetailFilter.Setup(DB, Admin)
	appGroupItemDetailDefault.Setup(DB, Admin)
	appSizeItem.Setup(DB, Admin)
	appSizeGroup.Setup(DB, Admin)
	appFitTool.Setup(DB, Admin)
	appFitToolFilter.Setup(DB, Admin)
	appTryOn.Setup(DB, Admin)
	appTryOnCmt.Setup(DB, Admin)

	appNotification.Setup(DB, Admin)

	return Admin
}

type AdminAuth struct {
}

func (AdminAuth) LoginURL(ctx *admin.Context) string {
	return "/auth/login"
}

func (AdminAuth) LogoutURL(ctx *admin.Context) string {
	return "/auth/logout"
}

func (AdminAuth) GetCurrentUser(ctx *admin.Context) qor.CurrentUser {
	currentUser, _ := config.Auth.GetCurrentUser(ctx.Request).(qor.CurrentUser)
	fmt.Printf("1 The login user is %v\n", currentUser)

	matchedRoles := roles.MatchedRoles(ctx.Request, currentUser)
	fmt.Printf("2 The login role is %v\n", matchedRoles)
	return currentUser
}
