package main

import (
	"fmt"
	"log"
	"net/http"

	"go-remix-jokes/lib/env"
	"go-remix-jokes/lib/page"
)

var (
	homePage       = page.NewPage("pages/index.go.html")
	jokesIndexPage = page.NewPage("pages/jokes/index.go.html")
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/jokes", JokesIndex)

	p := env.GetEnv("PORT", "3000")
	fmt.Println("Started listening on port :" + p)
	log.Fatalln(http.ListenAndServe(":"+p, nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
	homePage.Render(w, "Faisal")
}

func JokesIndex(w http.ResponseWriter, r *http.Request) {
	jokesIndexPage.Render(w, "Ali")
}
