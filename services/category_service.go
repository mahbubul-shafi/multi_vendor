package services

import (
	"multi_vendor/database"
	"multi_vendor/models"
)

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	result := database.DB.Find(&categories)
	return categories, result.Error
}
