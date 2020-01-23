package Controllers

import (
	"mafi-backend-go/ApiHelpers"
	"mafi-backend-go/Models"
	"github.com/gin-gonic/gin"
)

func ListPurchase(c *gin.Context) {
	var purchase []Models.Purchase
	err := Models.GetAllPurchase(&purchase)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, purchase)
	} else {
		ApiHelpers.RespondJSON(c, 200, purchase)
	}
}

func AddNewPurchase(c *gin.Context) {
	var purchase Models.Purchase
	c.BindJSON(&purchase)
	err := Models.AddNewPurchase(&purchase)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, purchase)
	} else {
		ApiHelpers.RespondJSON(c, 200, purchase)
	}
}

func GetOnePurchase(c *gin.Context) {
	id := c.Params.ByName("id")
	var purchase Models.Purchase
	err := Models.GetOnePurchase(&purchase, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, purchase)
	} else {
		ApiHelpers.RespondJSON(c, 200, purchase)
	}
}

func PutOnePurchase(c *gin.Context) {
	var purchase Models.Purchase
	id := c.Params.ByName("id")
	err := Models.GetOnePurchase(&purchase, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, purchase)
	}
	c.BindJSON(&purchase)
	err = Models.PutOnePurchase(&purchase, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, purchase)
	} else {
		ApiHelpers.RespondJSON(c, 200, purchase)
	}
}

func DeletePurchase(c *gin.Context) {
	var purchase Models.Purchase
	id := c.Params.ByName("id")
	err := Models.DeletePurchase(&purchase, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, purchase)
	} else {
		ApiHelpers.RespondJSON(c, 200, purchase)
	}
}
