package web

import (
	"context"
	"html/template"
	"log"
	"net/http"

	"github.com/Bieler96/Portfolio/utils"
	"github.com/russross/blackfriday/v2"
)

// html files are in the directory template
var tmpl = template.Must(template.ParseGlob("template/*"))

var ctx context.Context

// starts the webserver
func Webserver(ctxx context.Context) {
	ctx = ctxx
	log.Println("Server started on http://localhost:8080")

	http.HandleFunc("/", index)
	http.HandleFunc("/show", show)
	http.HandleFunc("/impressum", impressum)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./static/js"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./static/images"))))

	http.ListenAndServe(":8080", nil)
}

// shows the index page
func index(w http.ResponseWriter, r *http.Request) {
	//projects
	projects, err := utils.GetProjects(ctx)

	if err != nil {
		log.Fatal(err)
	}

	//services
	services, err := utils.GetServices(ctx)

	if err != nil {
		log.Fatal(err)
	}

	//nav entries
	navEntries, err := utils.GetNavEntries(ctx)

	if err != nil {
		log.Fatal(err)
	}

	//data structure
	data := struct {
		Projects []utils.Project
		Services []utils.Service
		Nav      []utils.NavEntry
	}{projects, services, navEntries}

	tmpl.ExecuteTemplate(w, "Index", data)
}

// shows the selected project in detail
func show(w http.ResponseWriter, r *http.Request) {
	//project slug
	projectSlack := r.URL.Query().Get("slug")

	//project
	project, err := utils.GetProjectById(ctx, projectSlack)

	if err != nil {
		log.Println("bvla")
		tmpl.ExecuteTemplate(w, "404", nil)
		return
	}

	//nav entries
	navEntries, err := utils.GetNavEntries(ctx)

	if err != nil {
		log.Fatal(err)
	}

	var c template.HTML = template.HTML(blackfriday.Run(project.Description))

	//data structure
	data := struct {
		Project utils.Project
		Content template.HTML
		Nav     []utils.NavEntry
	}{project, c, navEntries}

	tmpl.ExecuteTemplate(w, "Project", data)
}

// shows the impressum page
func impressum(w http.ResponseWriter, r *http.Request) {
	//nav entries
	navEntries, err := utils.GetNavEntries(ctx)

	if err != nil {
		log.Fatal(err)
	}

	//data structure
	data := struct {
		Nav []utils.NavEntry
	}{navEntries}

	tmpl.ExecuteTemplate(w, "Impressum", data)
}
