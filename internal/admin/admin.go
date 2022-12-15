package internal

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jnicklasj/gin-qor/config"
	"github.com/jnicklasj/gin-qor/config/bindatafs"
	productModels "github.com/jnicklasj/gin-qor/models/product"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/roles"
)

// func init() {
// 	roles.Register("administrator", func(req *http.Request, currentUser interface{}) bool {
// 		return currentUser != nil && currentUser.(*userModel.User).Role == "Admin"
// 	})
// }

// NewDummyAdmin generate admin for dummy app
func NewDummyAdmin(Db *gorm.DB, keepData ...bool) *admin.Admin {
	var (
		models = []interface{}{
			&productModels.Product{},
		}
		Admin = admin.New(&admin.AdminConfig{DB: Db, Auth: AdminAuth{}, AssetFS: bindatafs.AssetFS.NameSpace("admin")})
	)

	// AssetFs RegisterPath
	Admin.AssetFS.RegisterPath("assets/views")
	// OSS
	// media.RegisterCallbacks(Db)
	// AutoMigrate
	for _, value := range models {
		if len(keepData) == 0 {
			Db.DropTableIfExists(value)
		}

		// AutoMigrate 只能新增不存在的字段，不能修改已经存在的字段
		// will only add missing fields, won't delete/change current data
		Db.AutoMigrate(value)
	}

	// Register Apps
	// kinds.Setup(Db, Admin)

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

	//currentUser := Auth.GetCurrentUser(ctx.Request)
	//if currentUser != nil {
	//	qorCurrentUser, ok := currentUser.(qor.CurrentUser)
	//	if !ok {
	//		fmt.Printf("User %#v haven't implement qor.CurrentUser interface\n", currentUser)
	//	}
	//
	//	fmt.Printf("1 The login user is %v\n", qorCurrentUser)
	//	fmt.Printf("2 The login name is %v\n", qorCurrentUser.DisplayName())
	//	return qorCurrentUser
	//}
	//return nil
}
