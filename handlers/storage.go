package handlers

import (
	"net/http"
	"strings"
)

func storage(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.URL.Path, "/")[2]
	http.Redirect(w, r, "https://storage.googleapis.com/andreimuntean/" + id, http.StatusMovedPermanently)
}
