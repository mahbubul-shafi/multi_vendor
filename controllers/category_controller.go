package controllers

import (
	"github.com/kataras/iris/v12"
	"multi_vendor/services"
)

func GetCategories(ctx iris.Context) {
	categories, err := services.GetAllCategories()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		err := ctx.JSON(iris.Map{"error": "Failed to fetch categories"})
		if err != nil {
			return
		}
		return
	}
	err = ctx.JSON(categories)
	if err != nil {
		return
	}
}
