package handlers

import (
	"../services/imageService"
	"image/jpeg"
	"net/http"
	"strings"
)

func thumbnail(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.URL.Path, "/")[2]
	img := imageService.CreateThumbnailFromUrl("https://storage.googleapis.com/andreimuntean/" + id)
	jpeg.Encode(w, img, nil)
}
