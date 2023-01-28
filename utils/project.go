package utils

import (
	"context"
	"io/ioutil"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
Project represents a project
It contains: ID, slug, title, year, type, client, shortdescription, description, cover, images
*/
type Project struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Slug             string             `bson:"slug,omitempty" json:"slug,omitempty"`
	Title            string             `bson:"title,omitempty" json:"title,omitempty"`
	Year             string             `bson:"year,omitempty" json:"year,omitempty"`
	Type             string             `bson:"type,omitempty" json:"type,omitempty"`
	Client           string             `bson:"client,omitempty" json:"client,omitempty"`
	ShortDescription string             `bson:"shortdesc,omitempty" json:"shortdesc,omitempty"`
	Description      []byte             `bson:"desc,omitempty" json:"desc,omitempty"`
	Cover            string             `bson:"cover,omitempty" json:"cover,omitempty"`
	Images           []string           `bson:"images,omitempty" json:"images,omitempty"`
}

// projects in mongo collection
var DB_Projects *mongo.Collection

/*
InitProjectMongo initialized the project entries in MongoDB by inserting predefined entries

@param ctx context.Context : the context for the MongoDB operation
*/
func InitProjectMongo(ctx context.Context) {
	// bluetoothtrackerImages := []string{""}
	tankappImages := []string{"tankapp1.jpg"}
	zeitappImages := []string{"zeitapp1.jpg", "zeitapp2.jpg", "zeitapp3.jpg", "zeitapp4.jpg"}
	todoappImages := []string{"todoapp1.jpg", "todoapp2.jpg", "todoapp3.jpg"}
	bsvImages := []string{"bsv1.jpg", "bsv2.png"}
	modheadImages := []string{"modhead1.png", "modhead2.png", "modhead3.png", "modhead4.png", "modhead5.png"}
	modcharacterImages := []string{"modcharacter1.png", "modcharacter2.png", "modcharacter3.png", "modcharacter4.png"}
	plantaholicsImages := []string{"plantaholics1.jpg", "plantaholics2.jpg", "plantaholics3.jpg", "plantaholics4.jpg", "plantaholics5.jpg"}
	zsphereImages := []string{"zsphere1.png", "zsphere2.png", "zsphere3.png"}
	platformerengineImages := []string{"platformerengine1.png", "platformerengine2.png", "platformerengine3.png", "platformerengine4.png"}
	weingutImages := []string{"weingut1.jpg", "weingut2.png", "weingut3.png"}
	// arplakatImages := []string{""}

	entries := []interface{}{
		// bson.D{{"slug", "bluetoothtracker"}, {"title", "Obserwando Bluetooth Tracker"}, {"year", "2017"}, {"Type", "Entwicklung"}, {"client", "Rösler Software Technik"}, {"shortdesc", "Bluetooth Tracker"}, {"desc", getMarkdown("bluetoothtracker")}, {"cover", bluetoothtrackerImages[0]}, {"images", bluetoothtrackerImages}},
		bson.D{{"slug", "tankapp"}, {"title", "Obserwando Tank App"}, {"year", "2018"}, {"Type", "Entwicklung"}, {"client", "Rösler Software Technik"}, {"shortdesc", ""}, {"desc", getMarkdown("tankapp")}, {"cover", tankappImages[0]}, {"images", tankappImages}},
		bson.D{{"slug", "zeitapp"}, {"title", "Zeiterfassungs App"}, {"year", "2020"}, {"Type", "Entwicklung"}, {"client", "HS Harz"}, {"shortdesc", ""}, {"desc", getMarkdown("zeitapp")}, {"cover", zeitappImages[0]}, {"images", zeitappImages}},
		bson.D{{"slug", "todoapp"}, {"title", "Todo WebApp"}, {"year", "2018"}, {"Type", "Entwicklung"}, {"client", "Privat"}, {"shortdesc", ""}, {"desc", getMarkdown("todoapp")}, {"cover", todoappImages[0]}, {"images", todoappImages}},
		bson.D{{"slug", "bsv"}, {"title", "BSV"}, {"year", "2017"}, {"Type", "Entwicklung"}, {"client", "BBS 1 Delmenhorst"}, {"shortdesc", "Website für Sportverein mit weiteren Funktionen"}, {"desc", getMarkdown("bsv")}, {"cover", bsvImages[0]}, {"images", bsvImages}},
		bson.D{{"slug", "modhead"}, {"title", "ModHead"}, {"year", "2021"}, {"Type", "3D"}, {"client", "Privat"}, {"shortdesc", "Modulare Köpfe"}, {"desc", getMarkdown("modhead")}, {"cover", modheadImages[0]}, {"images", modheadImages}},
		bson.D{{"slug", "modcharacter"}, {"title", "ModCharacter"}, {"year", "2021"}, {"Type", "3D"}, {"client", "Privat"}, {"shortdesc", ""}, {"desc", getMarkdown("modcharacter")}, {"cover", modcharacterImages[1]}, {"images", modcharacterImages}},
		bson.D{{"slug", "plantaholics"}, {"title", "Plantaholics ePlant"}, {"year", "2022"}, {"Type", "Entwicklung"}, {"client", "HS Harz"}, {"shortdesc", ""}, {"desc", getMarkdown("plantaholics")}, {"cover", plantaholicsImages[0]}, {"images", plantaholicsImages}},
		bson.D{{"slug", "zsphere"}, {"title", "ZSphere"}, {"year", "2022"}, {"Type", "3D"}, {"client", "Privat"}, {"shortdesc", "ZSpeheres in Blender with Geo Nodes"}, {"desc", getMarkdown("zsphere")}, {"cover", zsphereImages[0]}, {"images", zsphereImages}},
		bson.D{{"slug", "platformerengine"}, {"title", "Platformer Engine"}, {"year", "2021"}, {"Type", "Game Dev"}, {"client", "Privat"}, {"shortdesc", ""}, {"desc", getMarkdown("platformerengine")}, {"cover", platformerengineImages[0]}, {"images", platformerengineImages}},
		bson.D{{"slug", "weingut"}, {"title", "Weingut"}, {"year", "2020"}, {"Type", "Gestaltung"}, {"client", "HS Harz"}, {"shortdesc", "Corporate Design für ein Weingut"}, {"desc", getMarkdown("weingut")}, {"cover", weingutImages[0]}, {"images", weingutImages}},
		// bson.D{{"slug", "arplakat"}, {"title", "AR Plakat Serie"}, {"year", "2022"}, {"Type", "Gestaltung"}, {"client", "HS Harz"}, {"shortdesc", ""}, {"desc", getMarkdown("arplakat")}, {"cover", arplakatImages[0]}, {"images", arplakatImages}},
	}

	_, err := DB_Projects.InsertMany(ctx, entries)

	if err != nil {
		log.Fatalf("Could not insert entries %v: %v", entries, err)
	}
}

/*
ResetMongoProject reset the "Projects" MongoDB collection by dropping it

@param ctx context.Context : the context for the MongoDB operation
*/
func ResetMongoProject(ctx context.Context) {
	DB_Projects.Drop(ctx)
}

/*
GetProjects gets all the projects from the MongoDB

@param ctx context.Context : the context for the MongoDB operation

@return []Project the projects
@return error if some error occure
*/
func GetProjects(ctx context.Context) ([]Project, error) {
	opts := options.Find().SetSort(bson.D{{"year", -1}})
	cursor, err := DB_Projects.Find(ctx, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var projects []Project

	if err = cursor.All(ctx, &projects); err != nil {
		return nil, err
	}

	//images add project slug
	for i, p := range projects {
		for j, v := range projects[i].Images {
			str := p.Slug + "/" + v
			projects[i].Images[j] = str
		}
	}

	//cover add project slug
	for i, p := range projects {
		str := p.Slug + "/" + p.Cover
		projects[i].Cover = str
	}

	return projects, err
}

/*
GetProjectById gets the project by the given slug

@param ctx context.Context : the context for the MongoDB operation
@param slug string: the slug of the searched entry

@return Project the searched project
@return error if some error occure
*/
func GetProjectById(ctx context.Context, slug string) (Project, error) {
	//project
	var project Project
	err := DB_Projects.FindOne(ctx, bson.M{"slug": slug}).Decode(&project)

	if err != nil {
		return project, err
	}

	//images add project slug
	for i, v := range project.Images {
		str := project.Slug + "/" + v
		project.Images[i] = str
	}

	//cover add project slug
	str := project.Slug + "/" + project.Cover
	project.Cover = str

	return project, err
}

/*
getMarkdown read in a markdown file with the given slug

@param slug string: the slug of the searched entry

@return []byte : the content of the specified markdown file
*/
func getMarkdown(slug string) []byte {
	b, errT := ioutil.ReadFile("./static/description/" + slug + ".md")
	if errT != nil {
		log.Fatal(errT)
	}

	return b
}
