package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	TODO_DB      = "TODO_DB"
	TODO_DB_NAME = "TODO_DB_NAME"
)

func GetConnection() (client *mongo.Client, ctx context.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv(TODO_DB)))
	if err != nil {
		log.Fatal(err)
	}

	return
}
