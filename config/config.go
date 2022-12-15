package config

import (
	"github.com/jnicklasj/gin-qor/config/db"
	userModels "github.com/jnicklasj/gin-qor/models/user"

	"github.com/qor/auth"
	"github.com/qor/auth_themes/clean"
)

var (
	Auth = clean.New(&auth.Config{})
)

func init() {

	Auth = clean.New(&auth.Config{
		DB:        db.DB,
		UserModel: &userModels.User{},
	})

}
