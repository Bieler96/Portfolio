package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
Service represent activities I like to do
It contains: ID, title, description,icon
*/
type Service struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title       string             `bson:"title,omitempty" json:"title,omitempty"`
	Description string             `bson:"desc,omitempty" json:"desc,omitempty"`
	Icon        string             `bson:"icon,omitempty" json:"icon,omitempty"`
}

// services in mongo collection
var DB_Services *mongo.Collection

/*
InitServiceMongo initialized the service entries in MongoDB by inserting predefined entries

@param ctx context.Context : the context for the MongoDB operation
*/
func InitServiceMongo(ctx context.Context) {
	entries := []interface{}{
		bson.D{{"title", "UI / UX"}, {"desc", "Nicht nur das gestalten von User Interfaces sondern auch das Testen der Useability."}, {"icon", ""}},
		bson.D{{"title", "Branding"}, {"desc", "Visuelle Identität, Marketingmaterialien."}, {"icon", ""}},
		bson.D{{"title", "Entwicklung"}, {"desc", "Entwickeln von Apps, Webseiten und auch Spiele"}, {"icon", ""}},
	}

	_, err := DB_Services.InsertMany(ctx, entries)

	if err != nil {
		log.Fatalf("Could not insert entries %v: %v", entries, err)
	}
}

/*
ResetMongoService reset the "Services" MongoDB collection by dropping it

@param ctx context.Context : the context for the MongoDB operation
*/
func ResetMongoService(ctx context.Context) {
	DB_Services.Drop(ctx)
}

/*
GetServices gets all the services from the MongoDB

@param ctx context.Context : the context for the MongoDB operation

@return []Service the services
@return error if some error occure
*/
func GetServices(ctx context.Context) ([]Service, error) {
	cursor, err := DB_Services.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	var services []Service

	if err = cursor.All(ctx, &services); err != nil {
		return nil, err
	}

	return services, err
}
