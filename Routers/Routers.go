package Routers

import (
	"github.com/v1nland/mafi-backend-go/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// segment API by version
	v1 := r.Group("/v1")
	{
		v1.GET("item", Controllers.ListItem)
		v1.POST("item", Controllers.AddNewItem)
		v1.GET("item/:id", Controllers.GetOneItem)
		v1.PUT("item/:id", Controllers.PutOneItem)
		v1.DELETE("item/:id", Controllers.DeleteItem)

		v1.GET("order", Controllers.ListOrder)
		v1.POST("order", Controllers.AddNewOrder)
		v1.GET("order/:id", Controllers.GetOneOrder)
		v1.PUT("order/:id", Controllers.PutOneOrder)
		v1.DELETE("order/:id", Controllers.DeleteOrder)

		v1.GET("purchase", Controllers.ListPurchase)
		v1.POST("purchase", Controllers.AddNewPurchase)
		v1.GET("purchase/:id", Controllers.GetOnePurchase)
		v1.PUT("purchase/:id", Controllers.PutOnePurchase)
		v1.DELETE("purchase/:id", Controllers.DeletePurchase)
	}

	return r
}
