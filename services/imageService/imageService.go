package imageService

import (
	"github.com/nfnt/resize"
	"image"
	"net/http"
)

// CreateThumbnailFromUrl downloads the specified image and returns a thumbnail.
func CreateThumbnailFromUrl(url string) image.Image {
	return CreateThumbnail(DownloadImage(url))
}

// DownloadImage downloads an image from the specified URL.
func DownloadImage(url string) image.Image {
	response, _ := http.Get(url)
	img, _, _ := image.Decode(response.Body)
	response.Body.Close()
	return img
}

// CreateThumbnail creates a thumbnail from the specified image.
func CreateThumbnail(img image.Image) image.Image {
	return DownscaleImage(img, 640)
}

// DownscaleImage shrinks the image, making its largest side equal to the specified size.
func DownscaleImage(img image.Image, size uint) image.Image {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	if width > height {
		return resize.Resize(size, 0, img, resize.NearestNeighbor)
	}
	return resize.Resize(0, size, img, resize.NearestNeighbor)
}
