package handlers

import (
	"net/http"
)

func epsilon(w http.ResponseWriter, r *http.Request) {
	templates["epsilon"].ExecuteTemplate(w, "layout", nil)
}
