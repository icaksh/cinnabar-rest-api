package public

import (
	"cinnabar-rest-api/controllers/daily"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup){
	router.GET("/daily/wiki", daily.TahukahAnda())
}