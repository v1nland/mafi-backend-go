package Models

import (
	"mafi-backend-go/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllItem(item *[]Item) (err error) {
	if err = Config.DB.Find(item).Error; err != nil {
		return err
	}
	return nil
}

func AddNewItem(item *Item) (err error) {
	if err = Config.DB.Create(item).Error; err != nil {
		return err
	}
	return nil
}

func GetOneItem(item *Item, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(item).Error; err != nil {
		return err
	}
	return nil
}

func PutOneItem(item *Item, id string) (err error) {
	fmt.Println(item)
	Config.DB.Save(item)
	return nil
}

func DeleteItem(item *Item, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(item)
	return nil
}
