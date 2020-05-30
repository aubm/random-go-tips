package handled

import (
	"encoding/json"
	"log"
	"net/http"

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

	ctx := r.Context()
	for _, imageUrl := range imagesUrl {
		select {
		case <-ctx.Done():
			log.Print(ctx.Err())
			return
		default:
			if _, err := img.ResizeFromUrl(imageUrl, 100, 0); err == nil {
				log.Printf("resized image, url: %v", imageUrl)
			} else {
				log.Printf("failed to resize image, url: %v, err: %v", imageUrl, err)
			}
		}
	}
}
