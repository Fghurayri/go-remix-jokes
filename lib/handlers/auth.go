package handlers

import (
	"log"
	"net/http"

	"go-remix-jokes/lib/models"
	"go-remix-jokes/lib/utils"
)

var (
	loginPage    = utils.NewPage("html/auth/login.go.html")
	registerPage = utils.NewPage("html/auth/register.go.html")
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if utils.IsSignedIn(r) {
			http.Redirect(w, r, "/jokes", http.StatusFound)
			return
		}
		loginPage.Render(w, nil)

	case http.MethodPost:
		username, password := parseAuthForm(r)

		var user models.User
		err := user.GetUserByUsername(h.DB, username)
		if err != nil {
			log.Println("Error fetching the user", err.Error())
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		err = utils.CompareHashAndPassword(user.PasswordHash, password)
		if err != nil {
			log.Println("Error comparing the hash", err.Error())
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		cookie := utils.CreateCookie()
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/jokes", http.StatusFound)

	default:
		NotFoundResponse(w)
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if utils.IsSignedIn(r) {
			http.Redirect(w, r, "/jokes", http.StatusFound)
			return
		}
		registerPage.Render(w, nil)

	case http.MethodPost:
		username, password := parseAuthForm(r)

		hashedPassword, err := utils.GenerateHashFromPassword(password)
		if err != nil {
			panic(err.Error())
		}

		user := &models.User{
			Username:     username,
			PasswordHash: string(hashedPassword),
		}
		err = user.CreateUser(h.DB)
		if err != nil {
			panic(err.Error())
		}

		cookie := utils.CreateCookie()
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

func parseAuthForm(r *http.Request) (string, string) {
	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	return username, password
}
