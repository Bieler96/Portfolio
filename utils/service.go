package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// data structure of a service
type Service struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title       string             `bson:"title,omitempty" json:"title,omitempty"`
	Description string             `bson:"desc,omitempty" json:"desc,omitempty"`
	Icon        string             `bson:"icon,omitempty" json:"icon,omitempty"`
}

// services in mongo collection
var DB_Services *mongo.Collection

// initialize data for the services into the mongodb
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

// removes all services in the mongodb
func ResetMongoService(ctx context.Context) {
	DB_Services.Drop(ctx)
}

// get all services
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
