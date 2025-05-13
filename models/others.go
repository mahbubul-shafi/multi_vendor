package models

type ProductWithCategory struct {
	ID           int    `json:"id"`
	ProductName  string `json:"product_name"`
	CategoryName string `json:"category_name"`
}
type ProductBasic struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
