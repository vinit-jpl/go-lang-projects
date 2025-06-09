package config

import (
	"context"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func ConnectMongo() (*MongoInstance, error) {

	uri := os.Getenv("MONGO_URI")

	if uri == "" {
		return nil, errors.New("MongoDB URI is not set in the environment variables")
	}

	dbName := os.Getenv("MONGO_DB_NAME")
	if dbName == "" {
		dbName = "blogDb" // default DB name
	}

	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancle()

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		// If there is an error connecting to MongoDB, return a wrapped error
		return nil, errors.New("failed to connect to MongoDB: " + err.Error())
	}

	// Ping the MongoDB server to verify the connection is alive
	err = client.Ping(ctx, nil)

	if err != nil {
		// If ping fails, return a wrapped error indicating MongoDB is unreachable
		return nil, errors.New("failed to ping MongoDB: " + err.Error())
	}

	// Store the connected client in the global MongoClient variable for reuse
	instance := &MongoInstance{
		Client: client,
		DB:     client.Database(dbName),
	}

	return instance, nil
}
