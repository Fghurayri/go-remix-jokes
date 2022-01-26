package handlers

import (
	"go-remix-jokes/lib/utils"
	"net/http"
)

var (
	jokesIndexPage = utils.NewPage("html/jokes/index.go.html")
)

func (h *Handler) JokesIndex(w http.ResponseWriter, r *http.Request) {
	if !utils.IsSignedIn(r) {
		http.Redirect(w, r, "/auth/login", http.StatusFound)
	}
	jokesIndexPage.Render(w, "Ali")
}
