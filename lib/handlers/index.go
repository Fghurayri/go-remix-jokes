package handlers

import "net/http"

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	homePage.Render(w, "Faisal")
}
