package Models

import (
	"mafi-backend-go/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllType(_type *[]Type) (err error) {
	if err = Config.DB.Find(_type).Error; err != nil {
		return err
	}
	return nil
}

func AddNewType(_type *Type) (err error) {
	if err = Config.DB.Create(_type).Error; err != nil {
		return err
	}
	return nil
}

func GetOneType(_type *Type, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(_type).Error; err != nil {
		return err
	}
	return nil
}

func PutOneType(_type *Type, id string) (err error) {
	fmt.Println(_type)
	Config.DB.Save(_type)
	return nil
}

func DeleteType(_type *Type, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(_type)
	return nil
}
