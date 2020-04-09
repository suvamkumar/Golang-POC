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
	Client *mongo.Client
)

func init() {
	var err error
	Client, err = mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal("error connecting to the database : ", err)
	}
	err = Client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("connection could not be established")
	}
	log.Println("connection established ...")
}
