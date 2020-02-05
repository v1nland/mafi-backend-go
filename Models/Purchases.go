package Models

import (
	"mafi-backend-go/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllPurchase(purchase *[]Purchase) (err error) {
	if err = Config.DB.Preload("Item").Find(purchase).Error; err != nil {
		return err
	}
	return nil
}

func AddNewPurchase(purchase *Purchase) (err error) {
	if err = Config.DB.Create(purchase).Error; err != nil {
		return err
	}
	return nil
}

func GetOnePurchase(purchase *Purchase, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(purchase).Error; err != nil {
		return err
	}
	return nil
}

func PutOnePurchase(purchase *Purchase, id string) (err error) {
	fmt.Println(purchase)
	Config.DB.Save(purchase)
	return nil
}

func DeletePurchase(purchase *Purchase, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(purchase)
	return nil
}
