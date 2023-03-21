package public

import (
	model "cinnabar/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Summary		REST - Response
//	@Description	Cinnabar REST API JSON Format
//	@Tags			basic
//	@Produce		json
//	@Success		200	{object}	model.TSuccessResponse
//  @Failure		400 {object}	model.TErrorResponse
//  @Failure		500 {object}	model.TErrorResponse
//	@Router			/test [get]
func GetBasicResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			model.TSuccessResponse{
				Data:    "Hi",
			},
		)
	}
}
