package main

import (
	"github.com/v1nland/mafi-backend-go/Config"
	"github.com/v1nland/mafi-backend-go/Models"
	"github.com/v1nland/mafi-backend-go/Routers"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
 	Config.DB, err = gorm.Open("mysql", "beab05944c9484:08c3abbf@(us-cdbr-iron-east-01.cleardb.net)/heroku_3ad9b9052d3cbc1?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.Item{})
	Config.DB.AutoMigrate(&Models.Order{})
	Config.DB.AutoMigrate(&Models.Purchase{})

	r := Routers.SetupRouter()

	r.Run()
}
