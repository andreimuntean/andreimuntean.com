package handlers

import (
	"net/http"
)

func gamma(w http.ResponseWriter, r *http.Request) {
	templates["gamma"].ExecuteTemplate(w, "layout", nil)
}
