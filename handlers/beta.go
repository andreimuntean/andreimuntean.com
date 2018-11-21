package handlers

import (
	"net/http"
)

func beta(w http.ResponseWriter, r *http.Request) {
	templates["beta"].ExecuteTemplate(w, "layout", nil)
}
