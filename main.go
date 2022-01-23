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
	log.Fatalln(http.ListenAndServe(":"+p, nil))
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
