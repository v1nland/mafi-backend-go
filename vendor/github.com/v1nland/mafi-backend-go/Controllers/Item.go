package Controllers

import (
	"github.com/v1nland/mafi-backend-go/ApiHelpers"
	"github.com/v1nland/mafi-backend-go/Models"
	"github.com/gin-gonic/gin"
)

func ListItem(c *gin.Context) {
	var item []Models.Item
	err := Models.GetAllItem(&item)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, item)
	} else {
		ApiHelpers.RespondJSON(c, 200, item)
	}
}

func AddNewItem(c *gin.Context) {
	var item Models.Item
	c.BindJSON(&item)
	err := Models.AddNewItem(&item)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, item)
	} else {
		ApiHelpers.RespondJSON(c, 200, item)
	}
}

func GetOneItem(c *gin.Context) {
	id := c.Params.ByName("id")
	var item Models.Item
	err := Models.GetOneItem(&item, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, item)
	} else {
		ApiHelpers.RespondJSON(c, 200, item)
	}
}

func PutOneItem(c *gin.Context) {
	var item Models.Item
	id := c.Params.ByName("id")
	err := Models.GetOneItem(&item, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, item)
	}
	c.BindJSON(&item)
	err = Models.PutOneItem(&item, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, item)
	} else {
		ApiHelpers.RespondJSON(c, 200, item)
	}
}

func DeleteItem(c *gin.Context) {
	var item Models.Item
	id := c.Params.ByName("id")
	err := Models.DeleteItem(&item, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, item)
	} else {
		ApiHelpers.RespondJSON(c, 200, item)
	}
}
