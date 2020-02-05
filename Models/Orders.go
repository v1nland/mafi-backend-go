package Models

import (
	"mafi-backend-go/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllOrder(order *[]Order) (err error) {
	if err = Config.DB.Preload("Item").Find(order).Error; err != nil {
		return err
	}
	return nil
}

func GetPendingOrder(order *[]Order) (err error) {
	if err = Config.DB.Preload("Item").Where("finished = ?", 0).Find(order).Error; err != nil {
		return err
	}
	return nil
}

func AddNewOrder(order *Order) (err error) {
	if err = Config.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func GetOneOrder(order *Order, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}

func PutOneOrder(order *Order, id string) (err error) {
	fmt.Println(order)
	Config.DB.Save(order)
	return nil
}

func DeleteOrder(order *Order, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(order)
	return nil
}
