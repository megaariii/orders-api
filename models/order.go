package models

import "time"

type Order struct {
	ID           string `gorm:"size:36;not null;uniqueIndex;primaryKey"`
	CustomerName string `gorm:"size:100;not null"`
	OrderedAt    time.Time
	Items []Item
}