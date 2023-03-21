package main

import (
	"net/http"
	"os"
	"time"

	//"cinnabar/controllers/private"
	"cinnabar/controllers/private"
	public "cinnabar/controllers/public"
	_ "cinnabar/docs"

	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mvrilo/go-redoc"
	ginredoc "github.com/mvrilo/go-redoc/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "OPTIONS", "PUT", "GET", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// },
		MaxAge: 12 * time.Hour,
	}))
	
	publicRedoc := redoc.Redoc{
		Title:       "Cinnabar Public REST API",
		Description: "Cinnabar Public REST API (Free to Use)",
		SpecFile:    "./docs/public_swagger.json",
		SpecPath:    "/docs/public_swagger.yaml",
		DocsPath:    "/public/docs",
	}

	privateRedoc := redoc.Redoc{
		Title:       "Cinnabar Private REST API",
		Description: "Cinnabar Private REST API (Paid to Use)",
		SpecFile:    "./docs/private_swagger.json",
		SpecPath:    "/docs/private_swagger.yaml",
		DocsPath:    "/private/docs",
	}
	r.Static("/view", "./view")
	//r.GET("/swagger/public/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.InstanceName("public"),ginSwagger.DefaultModelsExpandDepth(-1)))
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/public/docs")
	})
	r.Use(ginredoc.New(publicRedoc))
	r.Use(ginredoc.New(privateRedoc))
	public.Register(r)
	private.Register(r)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":"+port)
}
