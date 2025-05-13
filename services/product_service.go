package services

import (
	"multi_vendor/database"
	"multi_vendor/models"
)

func GetAllProductsWithCategories() ([]models.ProductWithCategory, error) {
	var result []models.ProductWithCategory
	err := database.DB.Table("products AS p").
		Select("p.id, p.name AS product_name, c.name AS category_name").
		Joins("JOIN categories c ON p.category_id = c.id").
		Scan(&result).Error
	return result, err
}

func GetProductsByCategoryID(categoryID int) ([]models.ProductBasic, error) {
	var products []models.ProductBasic
	err := database.DB.Table("products").
		Select("id, name").
		Where("category_id = ?", categoryID).
		Scan(&products).Error
	return products, err
}
