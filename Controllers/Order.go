package Controllers

import (
	"mafi-backend-go/ApiHelpers"
	"mafi-backend-go/Models"
	"github.com/gin-gonic/gin"
)

func ListOrder(c *gin.Context) {
	var order []Models.Order
	err := Models.GetAllOrder(&order)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, order)
	} else {
		ApiHelpers.RespondJSON(c, 200, order)
	}
}

func AddNewOrder(c *gin.Context) {
	var order Models.Order
	c.BindJSON(&order)
	err := Models.AddNewOrder(&order)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, order)
	} else {
		ApiHelpers.RespondJSON(c, 200, order)
	}
}

func GetOneOrder(c *gin.Context) {
	id := c.Params.ByName("id")
	var order Models.Order
	err := Models.GetOneOrder(&order, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, order)
	} else {
		ApiHelpers.RespondJSON(c, 200, order)
	}
}

func PutOneOrder(c *gin.Context) {
	var order Models.Order
	id := c.Params.ByName("id")
	err := Models.GetOneOrder(&order, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, order)
	}
	c.BindJSON(&order)
	err = Models.PutOneOrder(&order, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, order)
	} else {
		ApiHelpers.RespondJSON(c, 200, order)
	}
}

func DeleteOrder(c *gin.Context) {
	var order Models.Order
	id := c.Params.ByName("id")
	err := Models.DeleteOrder(&order, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, order)
	} else {
		ApiHelpers.RespondJSON(c, 200, order)
	}
}
