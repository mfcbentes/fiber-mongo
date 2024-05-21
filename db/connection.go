package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	TODO_DB = "TODO_DB"
)

func GetConnection() (client *mongo.Client, ctx context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("TODO_DB")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 10)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return
}
