package models

type Product struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Description string
	CategoryID  int
	Category    Category `gorm:"foreignKey:CategoryID"`
	VendorID    int
	Vendor      Vendor `gorm:"foreignKey:VendorID"`
}
