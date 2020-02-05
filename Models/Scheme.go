package Models

import (
	"github.com/jinzhu/gorm"
)

type Item struct {
	gorm.Model
	Description		string	`json:"description"`
	Type			string	`json:"type"`
	Color			string	`json:"color"`
	PurchasePrice	uint	`json:"purchase_price"`
	SellPrice		uint	`json:"sell_price"`
}

func (i *Item) TableName() string {
	return "items"
}

type Order struct {
	gorm.Model
	ItemID			uint	`json:"item_id"`
	Item			Item	`json:"item"`
	Buyer			string	`json:"buyer"`
	Description		string	`json:"description"`
	Contact			uint	`json:"contact"`
	Date			string	`json:"date"`
	Discount		float32	`json:"discount"`
	Location		string	`json:"location"`
	Hour			string	`json:"hour"`
	Source			string	`json:"source"`
	Finished		uint	`json:"finished"`
}

func (o *Order) TableName() string {
	return "orders"
}

type Purchase struct {
	gorm.Model
	ItemID			uint	`json:"item_id"`
	Item			Item	`json:"item"`
	Description		string	`json:"description"`
	Date			string	`json:"date"`
	ItemQty			string	`json:"item_qty"`
}

func (p *Purchase) TableName() string {
	return "purchases"
}

type Color struct {
	gorm.Model
	Color			string	`json:"color"`
}

func (p *Color) TableName() string {
	return "colors"
}

type Type struct {
	gorm.Model
	Type			string	`json:"type"`
}

func (p *Type) TableName() string {
	return "types"
}
