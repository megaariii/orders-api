package models

type Item struct {
	ID          string `gorm:"size:36;not null;uniqueIndex;primaryKey"`
	ItemCode    string `gorm:"size:36;not null"`
	Description string `gorm:"size:255;not null"`
	Quantity    int    `gorm:"not null"`
	OrderID     string `gorm:"size:36;index"`
}