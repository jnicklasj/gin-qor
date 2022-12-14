package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jnicklasj/gin-qor/handlers"
)

func setupRoutes(r *gin.Engine) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/pong", handlers.Pong)
}
