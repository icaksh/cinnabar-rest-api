package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TApiKey struct {
	ApiKey string `json:"api_key,omitempty" bson:"api_key"`
	Quota  int    `json:"quota,omitempty" bson:"quota"`
	Owner  string `json:"owner,omitempty" bson:"owner"`
}

const TABLES = "apikey"

var collection *mongo.Collection = GetCollection(mongoDB, TABLES)

func CreateApi(t TApiKey) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, t)
	if err != nil {
		return "0", err
	}

	return fmt.Sprintf("%v", result.InsertedID), err
}

func GetApi(id string) (TApiKey, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var t TApiKey

	if oid, err := primitive.ObjectIDFromHex(id); err != nil {
		return t, err
	} else {
		if err := collection.FindOne(ctx, bson.D{{Key: "_id", Value: oid}}).Decode(&t); err != nil {
			return t, err
		}
	}
	return t, nil
}

func GetApis() ([]TApiKey, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var api TApiKey
	var apis []TApiKey

	if cursor, err := collection.Find(ctx, bson.D{}); err != nil {
		defer cursor.Close(ctx)
		return apis, err
	} else {
		for cursor.Next(ctx) {
			if err := cursor.Decode(&api); err != nil {
				return apis, err
			}
			apis = append(apis, api)
		}
		return apis, nil
	}
}

func UpdateApi(id primitive.ObjectID, quota int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{"$set", bson.D{{Key: "quota", Value: quota}}}}
	_, err := collection.UpdateOne(
		ctx,
		filter,
		update,
	)
	return err
}

func DeleteApi(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return err
	}
	return nil
}

func ConsumeApiQuota(apiKey string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if res, err := GetApibyApiKey(apiKey); err != nil {
		return err
	} else {
		filter := bson.D{{Key: "api_key", Value: apiKey}}
		update := bson.D{{"$set", bson.D{{Key: "quota", Value: res.Quota - 1}}}}
		_, err := collection.UpdateOne(
			ctx,
			filter,
			update,
		)
		return err
	}
}

func GetApibyApiKey(apiKey string) (TApiKey, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var t TApiKey

	if err := collection.FindOne(ctx, bson.D{{Key: "api_key", Value: apiKey}}).Decode(&t); err != nil {
		return t, err
	}
	return t, nil
}

func ValidateApi(apiId string) (bool, error) {
	if isApiExist(apiId) {
		if ok, err := isQuotaExist(apiId); err != nil {
			return false, err
		} else {
			return ok, nil
		}
	} else {
		return false, errors.New("api_key invalid")
	}
}

func isQuotaExist(apiKey string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var t TApiKey

	if err := collection.FindOne(ctx, bson.D{{Key: "api_key", Value: apiKey}}).Decode(&t); err != nil {
		return false, err
	}
	if t.Quota > 1 {
		return true, nil
	} else {
		return false, errors.New("api_key limit reached")
	}
}

func isApiExist(apiId string) bool {
	_, err := GetApibyApiKey(apiId)
	return err == nil
}
