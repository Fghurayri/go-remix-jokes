package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	p := getEnv("PORT", "3000")
	fmt.Println("Started listening on port :" + p)
	http.HandleFunc("/", HandleFunc)
	log.Fatalln(http.ListenAndServe(":"+p, nil))
}

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("pages/layouts/root.go.html", "pages/index.go.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, "Faisal")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
