package helper

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	clientOnce sync.Once
)

const (
	uri    = "mongodb://mongodb:27017"
	dbname = "smartm2m"
)

func ConnectDB() (*mongo.Client, error) {
	var err error

	// Use sync.Once to ensure the connection is established only once
	clientOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(uri)

		// Connect to MongoDB
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		c, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatalf("Error connecting to MongoDB: %v", err)
		}

		// Check the connection
		err = c.Ping(ctx, nil)
		if err != nil {
			log.Fatalf("Error pinging MongoDB: %v", err)
		}

		client = c
	})

	return client, err
}

func GetDB() *mongo.Database {
	return client.Database(dbname)
}
