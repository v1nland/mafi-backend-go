# Simple REST API

It is a just simple RESTful API with Go using:
1. **Gin Framework**
2. **Gorm**

## Installation & Run
```bash
# Download this project
$ go get github.com/dedidot/gorm-gin

# Download Gin Framework
$ go get github.com/gin-gonic/gin

# Download GORM
$ go get github.com/jinzhu/gorm
```

Setting DB in main.go
```go
Config.DB, err = gorm.Open("mysql", "DBUSERNAME:DBPASSWORD@tcp(127.0.0.1:3306)/DBNAME?charset=utf8&parseTime=True&loc=Local")
```

## Structure
```
├── Models
│   ├── Items.go // Items model
│   ├── Orders.go // Orders model
|	├── Scheme.go // Structures and tables
├── Config
│   └── Database.go // Global database variables
├── Controllers
│   └── Item.go // Item controller
│   └── Order.go // Order controller
├── ApiHelpers
│   └── Response.go // Response function
├── Routers
|   └── Routers.go // Routers
└── main.go
```

## API

Pending: OpenAPI 3.0 spec.

#### /item
* `GET` : Get all item
* `POST` : Create a new item

#### /item/:id
* `GET` : Get a item
* `PUT` : Update a item
* `DELETE` : Delete a item

#Item POST Params (JSON)
```
{
	"description": "Mochilita bonita jeje",
	"type": "KANK",
	"color": "ROJO"
	"purchase_price": "5000"
	"sell_price": "10000"
}
```
