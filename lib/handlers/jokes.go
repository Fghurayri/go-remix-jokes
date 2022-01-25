package handlers

import (
	"go-remix-jokes/lib/utils"
	"net/http"
)

func (h *Handler) JokesIndex(w http.ResponseWriter, r *http.Request) {
	if !utils.IsSignedIn(r) {
		http.Redirect(w, r, "/auth/login", http.StatusFound)
	}
	jokesIndexPage.Render(w, "Ali")
}
