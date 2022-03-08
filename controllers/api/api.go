package api

import (
	"cinnabar-rest-api/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
type Response struct {
	Status	int	`json:"status"`
	Message	string `json:"message"`
	Data	map[string]interface{} `json:"data"`
}

type ApiResponse struct {
	Status	int	`json:"status"`
	Message	string `json:"message"`
	Data	map[string]interface{} `json:"data"`
}

func CheckApi(apiId string) bool{
	var api models.Apikey
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err := apiCollection.FindOne(ctx, bson.M{"apikey": apiId}).Decode(&api)
	defer cancel()
	return err == nil
}

var apiCollection *mongo.Collection = models.GetCollection(models.DB, "apiKey")
var validate = validator.New()

func CreateApi() gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var api models.Apikey
		defer cancel()

		if err := c.BindJSON(&api); err != nil{
			c.JSON(
				http.StatusBadRequest, 
				ApiResponse{
					Status: http.StatusBadRequest, 
					Message: "error", 
					Data: map[string]interface{}{"data":err.Error()},
				},
			)
			return
		}

		if validationErr := validate.Struct(&api); validationErr != nil{
			c.JSON(
				http.StatusBadRequest,
				ApiResponse{
					Status: http.StatusBadRequest, 
					Message: "error", 
					Data: map[string]interface{}{"data":validationErr.Error()},
				},
			)
			return
		}

		newApi := models.Apikey{
			ApiKey: api.ApiKey,
			Quota: api.Quota,
			Owner: api.Owner,
		}
		
		result, err := apiCollection.InsertOne(ctx, newApi)
		if err != nil{
			c.JSON(
				http.StatusInternalServerError,
				ApiResponse{
					Status: http.StatusInternalServerError,
					Message: "error",
					Data: map[string]interface{}{"data":err.Error()},
				},
			)
			return
		}

		c.JSON(
			http.StatusCreated,
			ApiResponse{
				Status: http.StatusCreated,
				Message: "success",
				Data: map[string]interface{}{"data":result},
			},
		)
	}
}

func CheckApiInfo() gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		apiId := c.Param("apiId")
		var api models.Apikey
		defer cancel()

		err := apiCollection.FindOne(ctx, bson.M{"apikey": apiId}).Decode(&api)
		if err != nil{
			errType := err.Error()
			
			if err == mongo.ErrNoDocuments {
				errType = "apikey not found"
			}
			c.JSON(
				http.StatusInternalServerError,
				ApiResponse{
					Status: http.StatusInternalServerError,
					Message: "error",
					Data: map[string]interface{}{"data": errType},
				},
			)
			return
		}
		c.JSON(
			http.StatusOK,
			ApiResponse{
				Status: http.StatusOK,
				Message: "success",
				Data: map[string]interface{}{"data": api},
			},
		)
	}
}