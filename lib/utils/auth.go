package utils

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func IsSignedIn(r *http.Request) bool {
	_, err := r.Cookie(COOKIE_NAME)
	return err == nil
}

func CreateCookie() *http.Cookie {
	return &http.Cookie{
		Name:     COOKIE_NAME,
		Value:    "temp stuff",
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 30,
		HttpOnly: true,
	}
}

func DestoryCookie() *http.Cookie {
	return &http.Cookie{
		Name:     COOKIE_NAME,
		Value:    "temp stuff",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
}

func GenerateHashFromPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
}

func CompareHashAndPassword(passwordHash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}
