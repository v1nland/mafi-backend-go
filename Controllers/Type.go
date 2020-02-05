package Controllers

import (
	"mafi-backend-go/ApiHelpers"
	"mafi-backend-go/Models"
	"github.com/gin-gonic/gin"
)

func ListType(c *gin.Context) {
	var _type []Models.Type
	err := Models.GetAllType(&_type)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, _type)
	} else {
		ApiHelpers.RespondJSON(c, 200, _type)
	}
}

func AddNewType(c *gin.Context) {
	var _type Models.Type
	c.BindJSON(&_type)
	err := Models.AddNewType(&_type)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, _type)
	} else {
		ApiHelpers.RespondJSON(c, 200, _type)
	}
}

func GetOneType(c *gin.Context) {
	id := c.Params.ByName("id")
	var _type Models.Type
	err := Models.GetOneType(&_type, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, _type)
	} else {
		ApiHelpers.RespondJSON(c, 200, _type)
	}
}

func PutOneType(c *gin.Context) {
	var _type Models.Type
	id := c.Params.ByName("id")
	err := Models.GetOneType(&_type, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, _type)
	}
	c.BindJSON(&_type)
	err = Models.PutOneType(&_type, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, _type)
	} else {
		ApiHelpers.RespondJSON(c, 200, _type)
	}
}

func DeleteType(c *gin.Context) {
	var _type Models.Type
	id := c.Params.ByName("id")
	err := Models.DeleteType(&_type, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, _type)
	} else {
		ApiHelpers.RespondJSON(c, 200, _type)
	}
}
