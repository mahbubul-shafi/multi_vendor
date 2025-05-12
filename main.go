package main

import (
	"github.com/kataras/iris/v12"
	"multi_vendor/database"
)

func main() {
	// Connect to DB
	database.Connect()

	// Create new Iris app
	app := iris.New()

	// Start server
	app.Listen(":8080")
}
