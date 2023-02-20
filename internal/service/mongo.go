package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	connectUri = "mongodb://user:pass@mongodb:27017/"
	table      = "testing"
	collection = "users"
)

func MakeUser() error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectUri))
	if err != nil {
	    return err
	}
	usersCollection := client.Database(table).Collection(collection)

	var pow = []int64{1, 2, 4, 8, 16, 32, 64, 1, 4, 3, 5, 6}

	for _, age := range pow {
		user := bson.D{{"username", String(15)}, {"age", age}}

		_, err = usersCollection.InsertOne(context.TODO(), user)

		if err != nil {
		    return err
		}
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	return nil
}
