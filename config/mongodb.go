package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DisconnectMongo func()

type MongoConn struct {
	Collection mongo.Collection
	Client     mongo.Client
	Disconnect DisconnectMongo
}

func ConnectMongo(collectionName string) MongoConn {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}

	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	collection := client.Database("themoviebakery").Collection(collectionName)

	newConnection := MongoConn{
		Collection: *collection,
		Client:     *client,
		Disconnect: func() {
			if err := client.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
		},
	}

	return newConnection
}
