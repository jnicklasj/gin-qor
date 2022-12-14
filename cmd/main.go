package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jnicklasj/gin-qor/config/db"
)

func main() {

	r := gin.Default()

	setupMounts(db.DB, r)
	setupRoutes(r)

	r.Run(":3001")
}
