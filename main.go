package main

import (
	"log"
	"net/http"

	"go-remix-jokes/lib/db"
	"go-remix-jokes/lib/env"
	"go-remix-jokes/lib/models"
	"go-remix-jokes/lib/page"

	"golang.org/x/crypto/bcrypt"
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
		_, err := r.Cookie("RJ_session")
		if err == nil {
			http.Redirect(w, r, "/jokes", http.StatusFound)
			return
		}
		loginPage.Render(w, nil)

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			panic(err.Error())
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		var user models.User
		err = db.DB.Where("username = ?", username).First(&user).Error
		if err != nil {
			log.Println("Error fetching the user", err.Error())
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
		if err != nil {
			log.Println("Error comparing the hash", err.Error())
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		cookie := &http.Cookie{
			Name:     "RJ_session",
			Value:    "temp stuff",
			Path:     "/",
			MaxAge:   60 * 60 * 24 * 30,
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/jokes", http.StatusFound)

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found :("))
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		_, err := r.Cookie("RJ_session")
		if err == nil {
			http.Redirect(w, r, "/jokes", http.StatusFound)
			return
		}
		registerPage.Render(w, nil)

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			panic(err.Error())
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			panic(err.Error())
		}

		err = db.DB.Create(&models.User{
			Username:     username,
			PasswordHash: string(hashedPassword),
		}).Error

		if err != nil {
			panic(err.Error())
		}

		cookie := &http.Cookie{
			Name:     "RJ_session",
			Value:    "temp stuff",
			Path:     "/",
			MaxAge:   60 * 60 * 24 * 30,
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/jokes", http.StatusFound)

	default:
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "not found :(", http.StatusNotFound)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "RJ_session",
		Value:    "temp stuff",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/jokes", http.StatusFound)
}

func JokesIndex(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("RJ_session")
	if err != nil {
		http.Redirect(w, r, "/auth/login", http.StatusFound)
	}
	jokesIndexPage.Render(w, "Ali")
}
