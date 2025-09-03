package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/blacknaml/hello-api/handlers"
	"github.com/blacknaml/hello-api/handlers/rest"
	"github.com/blacknaml/hello-api/translation"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = ":8080"
	}

	mux := http.NewServeMux()

	translationService := translation.NewStaticService()
	translationHandler := rest.NewTranslateHandler(translationService)
	mux.HandleFunc("/translate/hello", translationHandler.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)

	server := &http.Server{
		Addr:              addr,
		ReadHeaderTimeout: 5 * time.Second,  // Timeout for reading request headers
		ReadTimeout:       10 * time.Second, // Timeout for reading the entire request (headers + body)
		WriteTimeout:      10 * time.Second, // Timeout for writing the response
		IdleTimeout:       60 * time.Second, // Timeout for keeping idle connections alive
		Handler:           mux,              // Use http.DefaultServeMux if nil
	}

	log.Printf("listening on %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}
