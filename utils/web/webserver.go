package web

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/Bieler96/Portfolio/utils"
	"github.com/russross/blackfriday/v2"
)

// html files are in the directory template
var tmpl = template.Must(template.ParseGlob("template/*"))

var ctx context.Context

/*
Webserver starts a web server on http://localhost:8080. It sets up handlers for the following routes
"/" : index
"/show" : show
"/impressum" : impressum

It also sets up file servers for handling requests for files in the following directories
"./static/css" : access through route "/css/"
"./static/js" : access through route "/js/"
"./static/images" : access through route "/images/"

@param ctxx context.Context : the context for the web server
*/
func Webserver(ctxx context.Context) {
	ctx = ctxx
	log.Println("Server started on http://localhost:8080")

	http.HandleFunc("/", index)
	http.HandleFunc("/show", show)
	http.HandleFunc("/impressum", impressum)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./static/js"))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./static/images"))))

	go func() {
		<-time.After(100 * time.Millisecond)
		HandleOpenBrowser()
	}()

	http.ListenAndServe(":8080", nil)
}

func HandleOpenBrowser() {
	url := "http://localhost:8080"

	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "start", url)
		runCmd(cmd)
	} else if runtime.GOOS == "darwin" {
		cmd := exec.Command("open", url)
		runCmd(cmd)
	} else if runtime.GOOS == "linux" {
		cmd := exec.Command("xdg-open", url)
		runCmd(cmd)
	}
}

func runCmd(cmd *exec.Cmd) {
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

/*
index handles requests to the "/" route and renders the "Index" template

@param w http.ResponseWriter : the response writer to write the template
@param r *http.Request : the request object
*/
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

/*
show handles requests to the "/show" route and renders the "Project" template

@param w http.ResponseWriter : the response writer to write the template
@param r *http.Request : the request object
*/func show(w http.ResponseWriter, r *http.Request) {
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

/*
impressum handles requests to the "/impressum" route and renders the "Impressum" template

@param w http.ResponseWriter : the response writer to write the template
@param r *http.Request : the request object
*/
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
