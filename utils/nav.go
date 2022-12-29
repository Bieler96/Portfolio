package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// data structure for a nav entry
type NavEntry struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name string             `bson:"name,omitempty" json:"name,omitempty"`
	Link string             `bson:"link,omitempty" json:"link,omitempty"`
}

// nav entries in mongo collection
var DB_NavEntries *mongo.Collection

// initialize data for the nav into the mongodb
func InitNavMongo(ctx context.Context) {
	entries := []interface{}{
		bson.D{{"name", "Über Mich"}, {"link", "http://localhost:8080/#about"}},
		bson.D{{"name", "Projekte"}, {"link", "http://localhost:8080/#projects"}},
		bson.D{{"name", "Kontakt"}, {"link", "#contact"}},
	}

	_, err := DB_NavEntries.InsertMany(ctx, entries)

	if err != nil {
		log.Fatalf("Could not insert entries %v: %v", entries, err)
	}
}

// removes all nav entries in the mongodb
func ResetMongoNav(ctx context.Context) {
	DB_NavEntries.Drop(ctx)
}

// get all nav entries
func GetNavEntries(ctx context.Context) ([]NavEntry, error) {
	cursor, err := DB_NavEntries.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	var navEntries []NavEntry

	if err = cursor.All(ctx, &navEntries); err != nil {
		return nil, err
	}

	return navEntries, err
}
