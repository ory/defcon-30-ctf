package main

import (
	"log"
	"net/http"

	"github.com/ory/graceful"
)

func main() {
	config, err := newConfig()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := NewRepo(config)
	if err != nil {
		log.Fatal(err)
	}

	h, err := NewHandler(repo, config)
	if err != nil {
		log.Fatal(err)
	}

	server := graceful.WithDefaults(&http.Server{
		Addr:    config.ListenAddress,
		Handler: h,
	})

	if err := graceful.Graceful(func() error {
		log.Printf("Listening on http://%s", config.ListenAddress)
		return server.ListenAndServe()
	}, server.Shutdown); err != nil {
		log.Fatalf("Unable to gracefully shutdown HTTP(s) server because %v", err)
		return
	}
	log.Println("HTTP server was shutdown gracefully")
}
