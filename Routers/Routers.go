package Routers

import (
	"mafi-backend-go/Controllers"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())

	// segment API by version
	v1 := r.Group("/v1")
	{
		item := v1.Group("/item")
		{
			item.GET("", Controllers.ListItem)
			item.POST("", Controllers.AddNewItem)
			item.GET("id/:id", Controllers.GetOneItem)
			item.PUT(":id", Controllers.PutOneItem)
			item.DELETE(":id", Controllers.DeleteItem)
		}

		order := v1.Group("/order")
		{
			order.GET("", Controllers.ListOrder)
			order.GET("pending", Controllers.ListPendingOrder)
			order.POST("", Controllers.AddNewOrder)
			order.GET("id/:id", Controllers.GetOneOrder)
			order.PUT(":id", Controllers.PutOneOrder)
			order.DELETE(":id", Controllers.DeleteOrder)
		}

		purchase := v1.Group("/purchase")
		{
			purchase.GET("", Controllers.ListPurchase)
			purchase.POST("", Controllers.AddNewPurchase)
			purchase.GET("id/:id", Controllers.GetOnePurchase)
			purchase.PUT(":id", Controllers.PutOnePurchase)
			purchase.DELETE(":id", Controllers.DeletePurchase)
		}

		color := v1.Group("/color")
		{
			color.GET("", Controllers.ListColor)
			color.POST("", Controllers.AddNewColor)
			color.GET("id/:id", Controllers.GetOneColor)
			color.PUT(":id", Controllers.PutOneColor)
			color.DELETE(":id", Controllers.DeleteColor)
		}

		_type := v1.Group("/type")
		{
			_type.GET("", Controllers.ListType)
			_type.POST("", Controllers.AddNewType)
			_type.GET("id/:id", Controllers.GetOneType)
			_type.PUT(":id", Controllers.PutOneType)
			_type.DELETE(":id", Controllers.DeleteType)
		}

		stat := v1.Group("/stat")
		{
			stat.GET("box", Controllers.GetCurrency)
			stat.GET("sources", Controllers.GetSources)
			stat.GET("mostSoldItem", Controllers.GetMostSold)
			stat.GET("itemStock", Controllers.GetStock)
		}
	}

	return r
}
