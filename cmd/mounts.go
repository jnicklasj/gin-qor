package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jnicklasj/gin-qor/models"
	"github.com/qor/admin"
)

func setupMounts(DB *gorm.DB, r *gin.Engine) {

	Admin := admin.New(&admin.AdminConfig{DB: DB})
	Admin.AddResource(models.Product{})

	mux := http.NewServeMux()
	qorPrefix := "/admin"
	Admin.MountTo(qorPrefix, mux)

	r.Any("/admin/*resources", gin.WrapH(mux))
}
