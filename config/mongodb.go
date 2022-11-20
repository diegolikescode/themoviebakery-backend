package config

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DisconnectMongo func()

type mongoConn struct {
	collection *mongo.Collection
	client     *mongo.Client
	disconnect DisconnectMongo
}

func ConnectMongo() *mongoConn {
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
	collection := client.Database("themoviebakery").Collection("users")

	newMongoConnection := new(mongoConn)

	newMongoConnection.collection = collection
	newMongoConnection.client = client
	newMongoConnection.disconnect = func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}

	// newMongoConnection := mongoConn{
	// 	collection,
	// 	client,
	// 	func() {
	// 		if err := client.Disconnect(context.TODO()); err != nil {
	// 			panic(err)
	// 		}
	// 	},
	// }
	res2B, _ := json.Marshal(newMongoConnection)
	fmt.Println(string(res2B))

	return newMongoConnection
}
