package model

import (
	lib "cinnabar/libs"

	"go.mongodb.org/mongo-driver/mongo"
)

var mongoDB *mongo.Client = lib.ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("cinnabar-rest").Collection(collectionName)
	return collection
}
