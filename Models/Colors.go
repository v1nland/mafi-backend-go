package Models

import (
	"mafi-backend-go/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllColor(color *[]Color) (err error) {
	if err = Config.DB.Find(color).Error; err != nil {
		return err
	}
	return nil
}

func AddNewColor(color *Color) (err error) {
	if err = Config.DB.Create(color).Error; err != nil {
		return err
	}
	return nil
}

func GetOneColor(color *Color, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(color).Error; err != nil {
		return err
	}
	return nil
}

func PutOneColor(color *Color, id string) (err error) {
	fmt.Println(color)
	Config.DB.Save(color)
	return nil
}

func DeleteColor(color *Color, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(color)
	return nil
}
