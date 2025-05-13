package models

type Stock struct {
	ID                   int `gorm:"primaryKey"`
	WarehouseID          int
	Warehouse            Warehouse `gorm:"foreignKey:WarehouseID"`
	ProductCombinationID int
	ProductCombination   ProductCombination `gorm:"foreignKey:ProductCombinationID"`
	Quantity             int
}
