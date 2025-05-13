package main

import (
	"github.com/kataras/iris/v12"
	"multi_vendor/database"
	"multi_vendor/routes"
)

func main() {
	database.Connect()

	app := iris.New()

	routes.Register(app)

	err := app.Listen(":8080")
	if err != nil {
		return
	}
}
