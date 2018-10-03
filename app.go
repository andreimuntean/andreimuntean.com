package main

import (
	"andreimuntean.com/handlers"
	"google.golang.org/appengine"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	initializeAssets()
	setUpRoutes()
	appengine.Main()
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