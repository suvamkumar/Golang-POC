package users_db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	connectionString = "mongodb://localhost:27017"
)

var (
	//Client accessible
	client *mongo.Client
)

func init() {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("error connecting to the database : ", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("connection could not be established")
	}
	log.Println("connection established ...")
}

//GetMongoInstance ...
func GetMongoInstance() *mongo.Client {
	return client
}
