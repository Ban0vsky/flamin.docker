package main

import (
	"github.com/ban0vsky/flamindocker/worker/database"
	"github.com/ban0vsky/flamindocker/worker/router"
	"log"
	"net/http"
)

func main() {
	port := "3000"
	newRouter := router.NewRouter()

	err := database.Connect()
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}

	log.Print("\nServer started on port " + port)

	err = http.ListenAndServe(":"+port, newRouter)
	if err != nil {
		log.Fatalf("could not serve on port %s", port)
	}
}
