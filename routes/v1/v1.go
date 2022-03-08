package v1

import (
	"cinnabar-rest-api/controllers/api"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup){
	apiId := router.Group("/:apiId")
	{
		apiId.GET("/info", api.CheckApiInfo())
	}
}