package handlers

import (
	"go-remix-jokes/lib/utils"
	"net/http"

	"gorm.io/gorm"
)

var (
	homePage       = utils.NewPage("pages/index.go.html")
	jokesIndexPage = utils.NewPage("pages/jokes/index.go.html")
	loginPage      = utils.NewPage("pages/auth/login.go.html")
	registerPage   = utils.NewPage("pages/auth/register.go.html")
)

type Handler struct {
	DB *gorm.DB
}

func newRootHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

func Init(db *gorm.DB) {
	h := newRootHandler(db)

	// Index
	http.HandleFunc("/", h.Index)

	// Auth
	http.HandleFunc("/auth/login", h.Login)
	http.HandleFunc("/auth/register", h.Register)
	http.HandleFunc("/auth/logout", h.Logout)

	// Jokes
	http.HandleFunc("/jokes", h.JokesIndex)
}

func NotFoundResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	http.Error(w, "not found :(", http.StatusNotFound)
}
