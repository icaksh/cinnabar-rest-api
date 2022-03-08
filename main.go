package main

import (
	"cinnabar-rest-api/models"
	"cinnabar-rest-api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default();
	gin.SetMode(gin.ReleaseMode)
	models.ConnectDB()
	routes.Routes(r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":"+port)
}