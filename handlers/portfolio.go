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
	path, _ := filepath.Abs("assets/json/portfolio.json")
	file, e := ioutil.ReadFile(path)

	if e != nil {
		log.Fatal(e)
	}

	json.Unmarshal(file, &model.Photos)
}

func portfolio(w http.ResponseWriter, r *http.Request) {
	templates["portfolio"].ExecuteTemplate(w, "layout", &model)
}
