package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

var templates *template.Template

func landingHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "landing", nil)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "dashboard", nil)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create", nil)
}

func projectHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "project", nil)
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "projects", nil)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "profile", nil)
}

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))

	r := mux.NewRouter()
	r.HandleFunc("/", landingHandler)
	r.HandleFunc("/dashboard", dashboardHandler)
	r.HandleFunc("/create", createHandler)
	r.HandleFunc("/project", projectHandler)
	r.HandleFunc("/projects", projectsHandler)
	r.HandleFunc("/profile", profileHandler)

	n := negroni.Classic()
	n.UseHandler(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	n.Run(":"+port)
}

/*
import (
	"flag"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "text/html")
	http.ServeFile(w, r, "public/main.html")
}

func main() {
	flag.Parse()

	store.Init()

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/login", p("/login", login))
	router.HandleFunc("/verify", p("/verify", verify))

	apiRouter := mux.NewRouter()
	apiRouter.HandleFunc("/logout", logout)
	apiRouter.HandleFunc("/user", user)
	apiRouter.HandleFunc("/user/project", userProject)
	apiRouter.HandleFunc("/project", project)
	apiRouter.HandleFunc("/project/member", member)
	apiRouter.HandleFunc("/ws", ws)

	go h.run()

	n := negroni.Classic()
	n.UseHandler(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	n.Run(":"+port)
}
*/
