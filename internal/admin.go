package dummy

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jnicklasj/gin-qor/models"
	"github.com/qor/admin"
)

func NewDummyAdmin(Db *gorm.DB) *admin.Admin {
	var Admin = admin.New(&admin.AdminConfig{DB: Db})

	Admin.AddResource(&models.Product{})

	fmt.Println("NewDummyAdmin Mounted...")

	return Admin
}
