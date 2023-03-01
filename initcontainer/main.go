package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("STARTING")
	message := struct {
		Message string `bson:"message"`
	}{
		Message: string(os.Getenv("MESSAGE")),
	}

	fmt.Println("CONNECTING TO DB")
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CONNECTED")

	fmt.Println("WRITING IN DB")
	coll := client.Database("TP903").Collection("message")
	_, err = coll.InsertOne(context.Background(), message, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ENDED")
}
