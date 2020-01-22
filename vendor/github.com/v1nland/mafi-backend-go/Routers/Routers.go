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
	}

	return r
}
