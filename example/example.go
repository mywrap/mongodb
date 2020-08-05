package main

import (
	"context"
	"log"

	"github.com/mywrap/mongodb"
)

func main() {
	client, err := mongodb.Connect(mongodb.LoadEnvConfig())
	if err != nil {
		log.Fatalf("error Connect: %v", err)
	}
	_, err = client.Database("database0").Collection("testColl0").
		InsertOne(context.Background(), map[string]interface{}{"key0": "val0"})
	log.Printf("err InsertOne: %v", err)
}
