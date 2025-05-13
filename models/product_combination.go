package models

type ProductCombination struct {
	ID                 int `gorm:"primaryKey"`
	ProductID          int
	Product            Product `gorm:"foreignKey:ProductID"`
	HashedCombinations string
	UnitID             int
	Unit               Unit `gorm:"foreignKey:UnitID"`
	Prices             int
	Approved           bool
}
