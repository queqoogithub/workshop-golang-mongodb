package baseDB

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gosoft.co.th/workshop-api/constants"
)

func GetMongoConnection() (*mongo.Client, context.Context, context.CancelFunc, error) {
	urlConn := os.Getenv(constants.MongoUrlEnv)

	client, clientErr := mongo.NewClient(options.Client().ApplyURI(urlConn))
	if clientErr != nil {
		return nil, nil, nil, clientErr
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	conErr := client.Connect(ctx)
	if conErr != nil {
		defer cancel()
		return nil, nil, nil, conErr
	}

	return client, ctx, cancel, nil
}

func OpenCollection(collectionName string) (*mongo.Collection, *mongo.Client, context.Context, context.CancelFunc, error) {
	client, ctx, cancel, connErr := GetMongoConnection()
	if connErr != nil {
		return nil, nil, nil, nil, connErr
	}
	MONGO_DB_NAME := os.Getenv(constants.MongoDbName)
	var collection *mongo.Collection = client.Database(MONGO_DB_NAME).Collection(collectionName)

	return collection, client, ctx, cancel, nil
}
