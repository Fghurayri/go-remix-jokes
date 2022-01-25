package main

import (
	"fmt"
	"log"
	"net/http"

	"go-remix-jokes/lib/db"
	"go-remix-jokes/lib/env"
	"go-remix-jokes/lib/models"
	"go-remix-jokes/lib/page"
)

var (
	homePage       = page.NewPage("pages/index.go.html")
	jokesIndexPage = page.NewPage("pages/jokes/index.go.html")
	loginPage      = page.NewPage("pages/auth/login.go.html")
	registerPage   = page.NewPage("pages/auth/register.go.html")
)

func main() {
	log.Println("Connecting and auto migrating DB...")
	db.ConnectAndMigrateDB()

	log.Println("Setting up routes...")
	http.HandleFunc("/", Index)
	http.HandleFunc("/auth/login", Login)
	http.HandleFunc("/auth/register", Register)
	http.HandleFunc("/auth/logout", Logout)
	http.HandleFunc("/jokes", JokesIndex)

	p := env.GetEnv("PORT", "3000")
	log.Println("Started listening on port :" + p)
	log.Fatal(http.ListenAndServe(":"+p, nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
	homePage.Render(w, "Faisal")
}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		loginPage.Render(w, nil)

	case http.MethodPost:
		r.ParseForm()
		log.Println(r.PostForm)
		loginPage.Render(w, nil)

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found :("))
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		registerPage.Render(w, nil)

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			panic(err.Error())
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		u := &models.User{}
		err = u.CreateUser(email, password)
		if err != nil {
			panic(err.Error())
		}

		http.Redirect(w, r, "/jokes", http.StatusFound)

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found :("))
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "WIP")
}

func JokesIndex(w http.ResponseWriter, r *http.Request) {
	jokesIndexPage.Render(w, "Ali")
}
