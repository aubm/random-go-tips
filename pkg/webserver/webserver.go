package webserver

import (
	"log"
	"net/http"
)

func Start(addr string, handler http.HandlerFunc) {
	http.HandleFunc("/", handler)

	log.Printf("Start listening on %v", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
