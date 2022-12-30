package main

import (
	"context"
	"log"

	"github.com/Bieler96/Portfolio/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// connection
	opt := options.Client().ApplyURI("mongodb://root:rootpassword@localhost:27017")
	client, err := mongo.Connect(ctx, opt)

	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
}

// create the collections for the database
func createDB(client mongo.Client) {
	utils.DB_NavEntries = client.Database("mydb").Collection("nav")
	utils.DB_Projects = client.Database("mydb").Collection("projects")
	utils.DB_Services = client.Database("mydb").Collection("services")
}

// reseted the database collections
func resetDB(ctx context.Context) {
	utils.ResetMongoNav(ctx)
	utils.ResetMongoProject(ctx)
	utils.ResetMongoService(ctx)
}

// fill the collections with initial data
func initDB(ctx context.Context) {
	utils.InitNavMongo(ctx)
	utils.InitProjectMongo(ctx)
	utils.InitServiceMongo(ctx)
}
