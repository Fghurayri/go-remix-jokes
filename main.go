package main

import (
	"log"
	"net/http"

	"go-remix-jokes/lib/db"
	"go-remix-jokes/lib/env"
	"go-remix-jokes/lib/page"
)

var (
	homePage       = page.NewPage("pages/index.go.html")
	jokesIndexPage = page.NewPage("pages/jokes/index.go.html")
)

func main() {
	log.Println("Connecting and auto migrating DB...")
	db.ConnectAndMigrateDB()

	log.Println("Setting up routes...")
	http.HandleFunc("/", Index)
	http.HandleFunc("/jokes", JokesIndex)

	p := env.GetEnv("PORT", "3000")
	log.Println("Started listening on port :" + p)
	log.Fatal(http.ListenAndServe(":"+p, nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
	homePage.Render(w, "Faisal")
}

func JokesIndex(w http.ResponseWriter, r *http.Request) {
	jokesIndexPage.Render(w, "Ali")
}
