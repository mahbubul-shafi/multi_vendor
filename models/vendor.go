package models

type Vendor struct {
	ID   int `gorm:"primaryKey"`
	Name string
}
