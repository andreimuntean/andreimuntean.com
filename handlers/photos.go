package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

type photo struct {
	ID          string
	Name        string
	Date        string
}

var model struct {
	Photos		[]photo
}

func init() {
	path, _ := filepath.Abs("assets/json/photos.json")
	file, e := ioutil.ReadFile(path)

	if e != nil {
		log.Fatal(e)
	}

	json.Unmarshal(file, &model.Photos)
}

func photos(w http.ResponseWriter, r *http.Request) {
	templates["photos"].ExecuteTemplate(w, "layout", &model)
}
