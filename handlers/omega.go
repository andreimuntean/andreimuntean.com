package handlers

import (
	"net/http"
)

func omega(w http.ResponseWriter, r *http.Request) {
	templates["omega"].ExecuteTemplate(w, "layout", nil)
}
