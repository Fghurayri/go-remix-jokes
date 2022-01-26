package handlers

import (
	"go-remix-jokes/lib/utils"
	"net/http"
)

var (
	homePage = utils.NewPage("html/index.go.html")
)

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	homePage.Render(w, "Faisal")
}
