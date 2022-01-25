package main

import (
	"log"
	"net/http"

	"go-remix-jokes/lib/db"
	"go-remix-jokes/lib/handlers"
	"go-remix-jokes/lib/utils"
)

func main() {
	log.Println("Connecting and auto migrating DB...")
	db := db.Init()

	log.Println("Setting up handlers...")
	handlers.Init(db)

	p := utils.GetEnv("PORT", "3000")
	log.Println("Started listening on port :" + p)
	log.Fatal(http.ListenAndServe(":"+p, nil))
}
