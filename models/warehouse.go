package models

type Warehouse struct {
	ID       int `gorm:"primaryKey"`
	VendorID int
	Vendor   Vendor `gorm:"foreignKey:VendorID"`
	Name     string
}
