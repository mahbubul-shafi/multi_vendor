package routes

import (
	"github.com/kataras/iris/v12"
	"multi_vendor/controllers"
)

func Register(app *iris.Application) {
	app.Get("/categories", controllers.GetCategories)                           //getting all category names
	app.Get("/products-with-categories", controllers.GetProductsWithCategories) //getting all products with their category names
	app.Get("/products/category/{id:int}", controllers.GetProductsByCategory)   //getting all products by their category id

}
