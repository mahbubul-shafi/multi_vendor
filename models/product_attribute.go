package models

type ProductAttribute struct {
	ID         int `gorm:"primaryKey"`
	CategoryID int
	Category   Category `gorm:"foreignKey:CategoryID"`
	Name       string
}
