package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./web/pub/assets/")
	router.StaticFile("/", "./web/pub/index.html")

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(307, "/")
	})

	// no login page exists yet
	router.GET("/login", func(c *gin.Context) {
		c.Redirect(307, "/")
	})

	// does not exists yet
	router.GET("/dashboard", func(c *gin.Context) {
		c.Redirect(307, "/")
	})

	router.Run(":8080")
}
