package handlers

import (
	"fmt"
	"net/http"
)

func ads(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "google.com, pub-5089386830685380, DIRECT, f08c47fec0942fa0")
}

func robots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "User-agent: *")
	fmt.Fprintln(w, "Allow: /assets/json/")
	fmt.Fprintln(w, "Disallow: /assets/")
}
