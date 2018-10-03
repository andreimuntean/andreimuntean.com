package main

import (
	"./handlers"
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
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
