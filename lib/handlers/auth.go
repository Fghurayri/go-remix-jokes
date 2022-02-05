package handlers

import (
	"log"
	"net/http"
	"strconv"

	"go-remix-jokes/lib/models"
	"go-remix-jokes/lib/utils"
)

var (
	authPage = utils.NewPage("html/auth.go.html")
)

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if utils.IsSignedIn(r) {
			http.Redirect(w, r, "/jokes", http.StatusFound)
			return
		}
		authPage.Render(w, r, nil)

	case http.MethodPost:
		username, password, authType := parseAuthForm(r)

		u := &models.User{
			Username: username,
		}

		if authType == "login" {
			err := u.VerifyCredentials(h.DB, password)
			if err != nil {
				log.Println("Error logging in the user", err.Error())
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
				return
			}
		} else {
			err := u.Create(h.DB, password)
			if err != nil {
				panic(err.Error())
			}
		}

		cookie := utils.CreateCookie(strconv.FormatUint(uint64(u.ID), 10))
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/jokes", http.StatusFound)

	default:
		NotFoundResponse(w)
	}
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := utils.DestoryCookie()
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/jokes", http.StatusFound)
}

func parseAuthForm(r *http.Request) (string, string, string) {
	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	authType := r.FormValue("auth-type")

	return username, password, authType
}
