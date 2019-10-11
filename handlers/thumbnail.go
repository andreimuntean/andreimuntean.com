package handlers

import (
	"../services/imageService"
	"image"
	"image/jpeg"
	"net/http"
	"strings"
)

var cachedThumbnails map[string]image.Image = make(map[string]image.Image)

func getThumbnailForId(id string) image.Image {
	if cachedThumbnails[id] == nil {
		cachedThumbnails[id] = imageService.CreateThumbnailFromUrl("https://storage.googleapis.com/andreimuntean/" + id)
	}

	return cachedThumbnails[id]
}

func thumbnail(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.URL.Path, "/")[2]
	jpeg.Encode(w, getThumbnailForId(id), nil)
}
