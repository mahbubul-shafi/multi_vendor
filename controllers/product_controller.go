package controllers

import (
	"github.com/kataras/iris/v12"
	"multi_vendor/services"
)

func GetProductsWithCategories(ctx iris.Context) {
	data, err := services.GetAllProductsWithCategories()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		err := ctx.JSON(iris.Map{"error": "Failed to fetch data"})
		if err != nil {
			return
		}
		return
	}
	err = ctx.JSON(data)
	if err != nil {
		return
	}
}

func GetProductsByCategory(ctx iris.Context) {
	categoryID, err := ctx.Params().GetInt("id")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		err := ctx.JSON(iris.Map{"error": "Invalid category ID"})
		if err != nil {
			return
		}
		return
	}

	products, err := services.GetProductsByCategoryID(categoryID)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		err := ctx.JSON(iris.Map{"error": "Failed to fetch products"})
		if err != nil {
			return
		}
		return
	}

	err = ctx.JSON(products)
	if err != nil {
		return
	}
}
