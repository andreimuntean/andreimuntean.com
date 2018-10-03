package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type messageModel struct {
	Title   string
	Message string
}

var templates map[string]*template.Template

// Execute loads templates from the specified directory and configures routes.
func Execute(templateDirectory string) error {
	loadTemplates(templateDirectory)
	setUpRoutes()

	return nil
}

func loadTemplates(templateDirectory string) error {
	if _, err := os.Stat(templateDirectory); err != nil {
		return fmt.Errorf("Could not find template directory '%s'.", templateDirectory)
	}

	// Loads template paths.
	templatePaths, _ := filepath.Glob(filepath.Join(templateDirectory, "*.tmpl"))
	sharedPaths, _ := filepath.Glob(filepath.Join(templateDirectory, "shared/*.tmpl"))

	// Loads the templates.
	templates = make(map[string]*template.Template)

	for _, templatePath := range templatePaths {
		t, err := template.ParseFiles(append(sharedPaths, templatePath)...)

		if err != nil {
			return err
		}

		name := strings.Split(filepath.Base(templatePath), ".")[0]
		templates[name] = t
	}

	return nil
}

func setUpRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/robots.txt", robots)
}
