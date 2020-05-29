package unbound

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/aubm/random-go-tips/pkg/config"
	"github.com/aubm/random-go-tips/pkg/fibonacci"
	"github.com/aubm/random-go-tips/pkg/webserver"
)

func Run(config config.Config) {
	webserver.Start(config.WebAppAddr, handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var userInput []uint64
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(len(userInput))

	for _, n := range userInput {
		go func(n uint64) {
			defer wg.Done()

			v := fibonacci.Fibonacci(n)
			_, _ = fmt.Fprintf(w, "Result for %v is %v\n", n, v)
		}(n)
	}

	wg.Wait()
}
