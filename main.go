package main

import (
	"fmt"
	"io"
	// "net/http"

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
	router.StaticFile("/login", "./web/pub/login.html")

	// TODO: Check DB, Create Cookie, Send Cookie
	router.POST("login", func(c *gin.Context) {
		fmt.Println(io.ReadAll(c.Request.Body))

		c.JSON(200, gin.H{"yummy": "horse"})
	})

	// does not exists yet
	router.GET("/dashboard", func(c *gin.Context) {
		c.Redirect(307, "/")
	})

	fmt.Println("\nView the site at:\nhttp://localhost:8080\n")

	router.Run(":8080")
}
