package handlers

import (
	"net/http"
)

func alpha(w http.ResponseWriter, r *http.Request) {
	templates["alpha"].ExecuteTemplate(w, "layout", nil)
}
