package models

type AttributeValue struct {
	ID          int `gorm:"primaryKey"`
	AttributeID int
	Attribute   ProductAttribute `gorm:"foreignKey:AttributeID"`
	Value       string
}
