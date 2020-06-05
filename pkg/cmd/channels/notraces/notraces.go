package notraces

import (
	"bytes"
	"encoding/json"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nfnt/resize"

	"github.com/aubm/random-go-tips/pkg/config"
	"github.com/aubm/random-go-tips/pkg/webserver"
)

func Run(config config.Config) {
	urlChan := make(chan string)
	imageChan := make(chan []byte)
	resizedDataChan := make(chan []byte)

	go downloadImages(urlChan, imageChan)
	go resizeImages(imageChan, resizedDataChan)
	go readResizedImages(resizedDataChan)

	webserver.Start(config.WebAppAddr, func(w http.ResponseWriter, r *http.Request) {
		var imagesUrl []string
		if err := json.NewDecoder(r.Body).Decode(&imagesUrl); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for _, imageUrl := range imagesUrl {
			urlChan <- imageUrl
		}
	})
}

func downloadImages(urls chan string, out chan []byte) {
	for {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				log.Printf("failed to download %v: %v", url, err)
				return
			}
			defer resp.Body.Close()

			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("failed to read response body for %v: %v", url, err)
				return
			}

			out <- b
		}(<-urls)
	}
}

func resizeImages(in chan []byte, out chan []byte) {
	for {
		go func(data []byte) {
			img, _, err := image.Decode(bytes.NewBuffer(data))
			if err != nil {
				log.Printf("failed to decode image from data: %v", err)
				return
			}

			newImage := resize.Resize(50, 0, img, resize.Lanczos3)
			b := new(bytes.Buffer)
			if err := jpeg.Encode(b, newImage, nil); err != nil {
				log.Printf("failed to encode resized image: %v", err)
				return
			}

			out <- b.Bytes()
		}(<-in)
	}
}

func readResizedImages(imgs chan []byte) {
	for {
		go func(img []byte) {
			log.Print("read resized image data")
		}(<-imgs)
	}
}
