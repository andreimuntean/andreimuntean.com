package handlers

import (
	"andreimuntean.com/services/cacheService"
	"andreimuntean.com/services/imageService"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"net/http"
	"strings"
)

func getThumbnailForId(c context.Context, id string) []byte {
	img := cacheService.Get(c, id)
	if img == nil {
		img = imageService.CreateThumbnailFromUrl("https://storage.googleapis.com/andreimuntean/" + id)
		cacheService.Add(c, id, img)
	}
	return img
}

func thumbnail(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	id := strings.Split(r.URL.Path, "/")[2]
	w.Write(getThumbnailForId(c, id))
}
