package createUser

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(collection mongo.Collection) {
	res, err := collection.InsertOne(context.TODO(), bson.M{"hello": "bebie boe"})
	if err != nil {
		panic(err)
	}

	id := res.InsertedID
	fmt.Println("MY NEW ID", id)

	// return id
}
