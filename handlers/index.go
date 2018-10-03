package handlers

import (
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// Redirect unknown routes to the 404 page.
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	templates["index"].ExecuteTemplate(w, "layout", nil)
}
