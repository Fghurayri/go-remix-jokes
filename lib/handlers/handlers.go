package handlers

import (
	"net/http"

	"gorm.io/gorm"
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
	http.HandleFunc("/jokes", h.Jokes)
	http.HandleFunc("/jokes/new", h.NewJoke)
	http.HandleFunc("/jokes/", h.Joke)
}

func NotFoundResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	http.Error(w, "not found :(", http.StatusNotFound)
}
