package models

type Unit struct {
	ID    int `gorm:"primaryKey"`
	Names string
}
