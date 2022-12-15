package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	internal "github.com/jnicklasj/gin-qor/internal/admin"
	"github.com/qor/auth"
	"github.com/qor/auth/providers/password"
	"github.com/qor/redirect_back"
	"github.com/qor/session/manager"
)

func setupMounts(DB *gorm.DB, r *gin.Engine) {
	var (
		qorPrefix    = "/admin"
		redirectBack = redirect_back.New(&redirect_back.Config{FallbackPath: "/admin", SessionManager: manager.SessionManager, IgnoredPrefixes: []string{"/auth"}})
		adminAuth    = auth.New(&auth.Config{ViewPaths: []string{"assets/views"}, DB: DB, Redirector: auth.Redirector{RedirectBack: redirectBack}})
	)

	mux := http.NewServeMux()

	adminAuth.RegisterProvider(password.New(&password.Config{}))
	internal.NewDummyAdmin(DB, true).MountTo(qorPrefix, mux)

	r.Any("/admin/*resources", gin.WrapH(mux))
	r.Any("/auth/*action", gin.WrapH(adminAuth.NewServeMux()))
}
