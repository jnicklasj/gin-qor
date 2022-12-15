package internal

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/jnicklasj/gin-qor/config"
	"github.com/jnicklasj/gin-qor/config/bindatafs"
	appKind "github.com/jnicklasj/gin-qor/internal/app/kind"
	modelKind "github.com/jnicklasj/gin-qor/models/kind"
	modelProduct "github.com/jnicklasj/gin-qor/models/product"
	modelUser "github.com/jnicklasj/gin-qor/models/user"
	"github.com/qor/admin"
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
		}
		Admin = admin.New(&admin.AdminConfig{DB: DB, Auth: AdminAuth{}, AssetFS: bindatafs.AssetFS.NameSpace("admin")})
	)

	// AssetFs RegisterPath
	Admin.AssetFS.RegisterPath("assets/views")
	// OSS
	// media.RegisterCallbacks(Db)
	// AutoMigrate
	for _, value := range models {
		if len(keepData) == 0 {
			DB.DropTableIfExists(value)
		}

		// AutoMigrate 只能新增不存在的字段，不能修改已经存在的字段
		// will only add missing fields, won't delete/change current data
		DB.AutoMigrate(value)
	}

	// Register Apps
	appKind.Setup(DB, Admin)

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
