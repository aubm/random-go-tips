package img

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"

	"github.com/nfnt/resize"
)

func ResizeFromUrl(src string, width, height uint) ([]byte, error) {
	resp, err := http.Get(src)
	if err != nil {
		return nil, fmt.Errorf("failed to download %v: %v", src, err)
	}
	defer resp.Body.Close()

	srcImage, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image from response body of url %s: %v", src, err)
	}

	newImage := resize.Resize(width, height, srcImage, resize.Lanczos3)
	b := new(bytes.Buffer)
	if err := jpeg.Encode(b, newImage, nil); err != nil {
		return nil, fmt.Errorf("failed to encode resized image: %v", err)
	}

	return b.Bytes(), nil
}
