package Models

import (
	"github.com/v1nland/mafi-backend-go/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllItem(i *[]Item) (err error) {
	if err = Config.DB.Find(i).Error; err != nil {
		return err
	}
	return nil
}

func AddNewItem(i *Item) (err error) {
	if err = Config.DB.Create(i).Error; err != nil {
		return err
	}
	return nil
}

func GetOneItem(i *Item, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(i).Error; err != nil {
		return err
	}
	return nil
}

func PutOneItem(i *Item, id string) (err error) {
	fmt.Println(i)
	Config.DB.Save(i)
	return nil
}

func DeleteItem(i *Item, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(i)
	return nil
}
