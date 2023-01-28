package main

import (
	"context"
	"log"

	"github.com/Bieler96/Portfolio/utils"
	"github.com/Bieler96/Portfolio/utils/web"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/*
main is the entry point of the whole project
*/
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

	createDB(*client)
	resetDB(ctx)
	initDB(ctx)

	// start webserver
	web.Webserver(ctx)
}

/*
createDB sets up the collections

@param client mongo.Client : connected to the MongoDB
*/
func createDB(client mongo.Client) {
	utils.DB_NavEntries = client.Database("db").Collection("nav")
	utils.DB_Projects = client.Database("db").Collection("projects")
	utils.DB_Services = client.Database("db").Collection("services")
}

/*
resetDB resets the collections

@param ctx context.Context : the context for the MongoDB operation
*/
func resetDB(ctx context.Context) {
	utils.ResetMongoNav(ctx)
	utils.ResetMongoProject(ctx)
	utils.ResetMongoService(ctx)
}

/*
initDB initializes the collections

@param ctx context.Context : the context for the MongoDB operation
*/
func initDB(ctx context.Context) {
	utils.InitNavMongo(ctx)
	utils.InitProjectMongo(ctx)
	utils.InitServiceMongo(ctx)
}
