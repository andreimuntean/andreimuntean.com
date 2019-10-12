package imageService

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"net/http"
)

// CreateThumbnailFromUrl downloads the specified image and returns a thumbnail as a byte array.
func CreateThumbnailFromUrl(url string) []byte {
	return toBytes(createThumbnail(downloadImage(url)))
}

// Downloads an image from the specified URL.
func downloadImage(url string) image.Image {
	response, _ := http.Get(url)
	img, _, _ := image.Decode(response.Body)
	response.Body.Close()
	return img
}

// Creates a thumbnail from the specified image.
func createThumbnail(img image.Image) image.Image {
	return downscaleImage(img, 640)
}

// Shrinks the image, making its largest side equal to the specified size.
func downscaleImage(img image.Image, size uint) image.Image {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	if width > height {
		return resize.Resize(size, 0, img, resize.NearestNeighbor)
	}
	return resize.Resize(0, size, img, resize.NearestNeighbor)
}

// Converts the specified image to a byte array.
func toBytes(img image.Image) []byte {
	buffer := new(bytes.Buffer)
	jpeg.Encode(buffer, img, nil)
	return buffer.Bytes()
}
