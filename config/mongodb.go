package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongodb+srv://night_driver:<password>@themoviebakery.p2jppue.mongodb.net/?retryWrites=true&w=majority

func ConnectMongo() mongo.Collection {
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

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("themoviebakery").Collection("users")

	// res, err := collection.InsertOne(context.TODO(), bson.M{"hello": "bebie boe"})

	// if err != nil {
	// 	panic(err)
	// }

	// id := res.InsertedID
	// fmt.Println("MY NEW ID", id)

	return *collection
}
