package models

type ProductCombinationValue struct {
	ID                   int `gorm:"primaryKey"`
	ProductCombinationID int
	ProductCombination   ProductCombination `gorm:"foreignKey:ProductCombinationID"`
	AttributeValueID     int
	AttributeValue       AttributeValue `gorm:"foreignKey:AttributeValueID"`
}
