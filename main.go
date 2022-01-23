package main

import (
	"fmt"
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
	fmt.Fprint(w, "Hello World")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
