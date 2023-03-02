package service

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	username string
	age      int64
}

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

	_, err = usersCollection.InsertMany(context.TODO(), []interface{}{
		User{username: String(15), age: 5},
		User{username: String(15), age: 2},
		User{username: String(15), age: 7},
		User{username: String(15), age: 9},
		User{username: String(15), age: 14},
		User{username: String(15), age: 18},
		User{username: String(15), age: 22},
	})

	if err != nil {
	    return err
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	return nil
}

func FetchFirstFiveUser() (string, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectUri))
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	usersCollection := client.Database(table).Collection(collection)

    filter := bson.D{}
	opts := options.Find().SetLimit(5)
	cursor, err := usersCollection.Find(context.TODO(), filter, opts)

	var results []User
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	result, err := json.Marshal(results)
	if err != nil {
		fmt.Println(err.Error())

		return "", err
	}

	return string(result), nil
}
