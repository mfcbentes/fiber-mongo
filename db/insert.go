package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Insert(colletion string, data any) (primitive.ObjectID, error) {
	client, ctx := GetConnection()
	defer client.Disconnect(ctx)

	c := client.Database(os.Getenv(TODO_DB_NAME)).Collection(colletion)

	resp, err := c.InsertOne(context.Background(), data)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return resp.InsertedID.(primitive.ObjectID), nil
}
