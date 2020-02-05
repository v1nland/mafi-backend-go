package Controllers

import (
	"mafi-backend-go/ApiHelpers"
	"mafi-backend-go/Models"
	"github.com/gin-gonic/gin"
)

func ListColor(c *gin.Context) {
	var color []Models.Color
	err := Models.GetAllColor(&color)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, color)
	} else {
		ApiHelpers.RespondJSON(c, 200, color)
	}
}

func AddNewColor(c *gin.Context) {
	var color Models.Color
	c.BindJSON(&color)
	err := Models.AddNewColor(&color)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, color)
	} else {
		ApiHelpers.RespondJSON(c, 200, color)
	}
}

func GetOneColor(c *gin.Context) {
	id := c.Params.ByName("id")
	var color Models.Color
	err := Models.GetOneColor(&color, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, color)
	} else {
		ApiHelpers.RespondJSON(c, 200, color)
	}
}

func PutOneColor(c *gin.Context) {
	var color Models.Color
	id := c.Params.ByName("id")
	err := Models.GetOneColor(&color, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, color)
	}
	c.BindJSON(&color)
	err = Models.PutOneColor(&color, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, color)
	} else {
		ApiHelpers.RespondJSON(c, 200, color)
	}
}

func DeleteColor(c *gin.Context) {
	var color Models.Color
	id := c.Params.ByName("id")
	err := Models.DeleteColor(&color, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, color)
	} else {
		ApiHelpers.RespondJSON(c, 200, color)
	}
}
