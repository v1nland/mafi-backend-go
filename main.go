package main

import (
	"mafi-backend-go/Config"
	// "mafi-backend-go/Models"
	"mafi-backend-go/Routers"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
 	Config.DB, err = gorm.Open("mysql", "bb734e7268e276:7a973a38@(us-cdbr-iron-east-04.cleardb.net)/heroku_3f9539851a7ef01?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config.DB.Close()

	// Config.DB.AutoMigrate(&Models.Item{})
	// Config.DB.AutoMigrate(&Models.Order{})
	// Config.DB.AutoMigrate(&Models.Purchase{})
	// Config.DB.AutoMigrate(&Models.Color{})
	// Config.DB.AutoMigrate(&Models.Type{})

	r := Routers.SetupRouter()

	r.Run()
}
