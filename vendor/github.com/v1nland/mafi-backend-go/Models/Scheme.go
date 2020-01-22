package Models

import (
	"github.com/jinzhu/gorm"
)

type Item struct {
	gorm.Model
	Description		string	`json:"description"`
	Type			string	`json:"type"`
	Color			string	`json:"color"`
	Purchase_price	string	`json:"purchase_price"`
	Sell_price		string	`json:"sell_price"`
}

func (i *Item) TableName() string {
	return "items"
}
