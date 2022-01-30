package utils

import (
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func IsSignedIn(r *http.Request) bool {
	_, err := r.Cookie(COOKIE_NAME)
	return err == nil
}

func GetUserIdFromCookie(r *http.Request) (uint, error) {
	c, err := r.Cookie(COOKIE_NAME)
	if err != nil {
		return 0, err
	}
	id, err := strconv.ParseUint(c.Value, 10, 0)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func CreateCookie(uid string) *http.Cookie {
	return &http.Cookie{
		Name:     COOKIE_NAME,
		Value:    uid,
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
