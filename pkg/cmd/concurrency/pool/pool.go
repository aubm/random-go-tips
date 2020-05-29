package pool

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

	nbJobs := len(imagesUrl)
	jobs := make(chan string, nbJobs)

	wg := sync.WaitGroup{}
	wg.Add(nbJobs)

	for w := 1; w <= 3; w++ {
		go resizeImagesWorker(jobs, &wg)
	}

	for _, imageUrl := range imagesUrl {
		jobs <- imageUrl
	}
	close(jobs)

	wg.Wait()
}

func resizeImagesWorker(imagesUrl chan string, wg *sync.WaitGroup) {
	for imageUrl := range imagesUrl {
		if _, err := img.ResizeFromUrl(imageUrl, 100, 0); err != nil {
			log.Printf("failed to resize image, url: %v, err: %v", imageUrl, err)
		}
		wg.Done()
	}
}
