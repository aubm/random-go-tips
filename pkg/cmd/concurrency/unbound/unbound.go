package unbound

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/aubm/random-go-tips/pkg/config"
	"github.com/aubm/random-go-tips/pkg/img"
	"github.com/aubm/random-go-tips/pkg/webserver"
)

func Run(config config.Config) {
	webserver.Start(config.WebAppAddr, handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var imagesUrl []string
	if err := json.NewDecoder(r.Body).Decode(&imagesUrl); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(len(imagesUrl))

	for _, imageUrl := range imagesUrl {
		go func(imageUrl string) {
			defer wg.Done()
			if _, err := img.ResizeFromUrl(imageUrl, 100, 0); err != nil {
				log.Printf("failed to resize image, url: %v, err: %v", imageUrl, err)
			}
		}(imageUrl)
	}

	wg.Wait()
}
