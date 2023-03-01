package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := struct {
			Message string `bson:"message"`
		}{}

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

		fmt.Println("READING MESSAGE")
		coll := client.Database("TP903").Collection("message")
		cur, err := coll.Find(context.Background(), bson.M{}, nil)
		if err != nil {
			log.Fatal(err)
		}

		for cur.Next(context.Background()) {
			err = cur.Decode(&response)
			fmt.Println(response.Message)
			if err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("MESSAGE READ")

		res, err := json.Marshal(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(res)

	})

	port := os.Getenv("APPPORT")
	fmt.Printf("running on http://127.0.0.1:%v\n", port)
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", port), nil); err != nil {
		log.Fatal(err)
	}
}
