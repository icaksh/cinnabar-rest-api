package routes

import (
	"cinnabar-rest-api/routes/public"
	v1 "cinnabar-rest-api/routes/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Routes(router *gin.Engine){
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			nil,
		)
	})
	api := router.Group("api")
	{
		publicGroup := api.Group("/public")
		{
			public.Routes(publicGroup)
		}
		first := api.Group("/v1")
		{
			v1.Routes(first)
		}
	}
}