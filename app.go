package main

import (
	"andreimuntean.com/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	initializeAssets()
	setUpRoutes()
	startServer()
}

func initializeAssets() {
	// Make the "assets" folder public.
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
}

func setUpRoutes() {
	dir, _ := os.Getwd()
	templateDirectory := filepath.Join(dir, "templates")

	if err := handlers.Execute(templateDirectory); err != nil {
		log.Fatal(err)
	}
}

func startServer() {
	port := os.Getenv("PORT")
	if port == "" {
			port = "8080"
			log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
