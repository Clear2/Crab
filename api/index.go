package api

import (
	"bee.com/conf"
	"bee.com/models"
	"bee.com/utils"
	"github.com/gin-gonic/gin"
)

func Index(c * gin.Context)  {

	var product models.Product
	result := conf.DB.First(&product)
	utils.ResSuccess(c, result)
}
