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

type Order struct {
	gorm.Model
	Item_id			string	`json:"item_id"`
	Buyer			string	`json:"buyer"`
	Description		string	`json:"description"`
	Contact			string	`json:"contact"`
	Date			string	`json:"date"`
	Discount		string	`json:"discount"`
	Location		string	`json:"location"`
	Hour			string	`json:"hour"`
	Source			string	`json:"source"`
	Finished		string	`json:"finished"`
}

func (i *Order) TableName() string {
	return "orders"
}
