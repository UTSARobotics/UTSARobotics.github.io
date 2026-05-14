package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	_ "modernc.org/sqlite"
)

var db *sql.DB

func main() {
	db, err := sql.Open("sqlite", "users.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// create table if it does not exists
	table_create := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT
		username TEXT UNIQUE NOT NULL
		password TEXT NOT NULL
	)
	`

	if _, err := db.Exec(table_create); err != nil {
		log.Fatal(err)
	}

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

// HashPassword generates a bcrypt hash of the password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// RegisterUser hashes the password and saves the user to the DB
func RegisterUser(username, password string) error {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hashedPassword)
	return err
}

// AuthenticateUser checks if the credentials are valid
func AuthenticateUser(username, password string) bool {
	var hashedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&hashedPassword)
	if err != nil {
		return false // User not found
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
