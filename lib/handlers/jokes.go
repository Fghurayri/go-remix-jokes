package handlers

import (
	"go-remix-jokes/lib/models"
	"go-remix-jokes/lib/utils"
	"log"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

var (
	jokesIndexPage = utils.NewPage("html/jokes/index.go.html")
	jokePage       = utils.NewPage("html/jokes/joke.go.html")
)

func (h *Handler) Jokes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if !utils.IsSignedIn(r) {
			http.Redirect(w, r, "/auth/login", http.StatusFound)
		}

		j := &models.Joke{}
		var js []models.Joke
		jd := make(map[string]interface{})
		jd["Jokes"] = &js

		err := j.GetAll(h.DB, &js)
		if err != nil {
			jokesIndexPage.Render(w, r, jd)
		}

		jokesIndexPage.Render(w, r, jd)

	case http.MethodPost:
		if !utils.IsSignedIn(r) {
			http.Redirect(w, r, "/auth/login", http.StatusFound)
		}

		uid, err := utils.GetUserIdFromCookie(r)
		if err != nil {
			http.Redirect(w, r, "/auth/login", http.StatusFound)
		}

		name, content := parseAddJokeForm(r)

		j := &models.Joke{
			Name:    name,
			Content: content,
			UserID:  uid,
		}

		err = j.Create(h.DB)
		if err != nil {
			log.Println("Error creating the joke", err.Error())
			http.Error(w, "Unable to create the joke", http.StatusInternalServerError)
		}
		http.Redirect(w, r, "/jokes/"+strconv.FormatUint(uint64(j.ID), 10), http.StatusFound)

	default:
		NotFoundResponse(w)
	}
}

func (h *Handler) Joke(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		s := strings.Split(r.RequestURI, "/jokes/")
		jID, err := strconv.ParseUint(s[1], 10, 0)
		if err != nil {
			http.Redirect(w, r, "/jokes", http.StatusFound)
		}

		j := &models.Joke{
			Model: gorm.Model{ID: uint(jID)},
		}

		err = j.GetById(h.DB)
		if err != nil {
			http.Redirect(w, r, "/jokes", http.StatusFound)
		}

		d := make(map[string]interface{})
		d["Joke"] = j

		jokePage.Render(w, r, d)

	case http.MethodPut:
		if !utils.IsSignedIn(r) {
			http.Redirect(w, r, "/auth/login", http.StatusFound)
		}

	case http.MethodDelete:
		if !utils.IsSignedIn(r) {
			http.Redirect(w, r, "/auth/login", http.StatusFound)
		}

		http.Redirect(w, r, "/jokes", http.StatusFound)
	default:
		NotFoundResponse(w)
	}
}

func parseAddJokeForm(r *http.Request) (string, string) {
	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
	}

	name := r.FormValue("name")
	content := r.FormValue("content")

	return name, content
}
