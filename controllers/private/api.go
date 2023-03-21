package private

import (
	model "cinnabar/models"
	"math/rand"
	"net/http"
	"time"
	"unsafe"

	"github.com/gin-gonic/gin"
)

type TReqApi struct {
	// Api key
	Apikey string `json:"api_key"`
}

const letterBytes = "1234567890abcdefghijklmnopqrstuvwxyz"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

//	@Summary		REST - Response
//	@Description	Cinnabar REST API JSON Format
//	@Tags			basic
//	@Security		apikey
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.TSuccessResponse
//  @Failure		400 {object}	model.TErrorResponse
//  @Failure		500 {object}	model.TErrorResponse
//	@Router			/api/test [post]
func PostBasicResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		if apiKey := c.Request.Header.Get("apikey"); apiKey == "" {
			c.JSON(
				http.StatusUnauthorized,
				model.TErrorResponse{
					Error:    "Invalid Api Key",
				},
			)
		} else {
			c.JSON(
				http.StatusUnauthorized,
				model.TSuccessResponse{
					Data:    "Hi, your Api Key is"+c.Request.Header.Get("apikey"),
				},
			)
		}
	}
}

//	@Summary		REST - Api Information
//	@Description	Check your API Information
//	@Tags			basic
//	@Security		apikey
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.TApiKey
//	@Router			/api/info [post]
func PostCheckApiInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		if apiKey := c.Request.Header.Get("apikey"); apiKey == "" {
			c.JSON(
				http.StatusUnauthorized,
				model.TErrorResponse{
					Error:    "apikey not valid",
				},
			)
		} else {
			if result, err := model.GetApibyApiKey(c.Request.Header.Get("apikey")); err != nil {
				c.JSON(
					http.StatusInternalServerError,
					model.TErrorResponse{
						Error:    "apikey not found",
					},
				)
				return
			} else {
				c.JSON(
					http.StatusUnauthorized,
					model.TSuccessResponse{
						Data: result,
					},
				)
			}
		}
	}
}
