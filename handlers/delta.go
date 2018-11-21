package handlers

import (
	"net/http"
)

func delta(w http.ResponseWriter, r *http.Request) {
	templates["delta"].ExecuteTemplate(w, "layout", nil)
}
