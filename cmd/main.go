package main

import (
	"log"
	"net/http"
	"time"

	"github.com/blacknaml/hello-api/config"
	"github.com/blacknaml/hello-api/handlers"
	"github.com/blacknaml/hello-api/handlers/rest"
	"github.com/blacknaml/hello-api/translation"
)

func main() {

	cfg := config.LoadConfiguration()
	addr := cfg.Port

	var translationService rest.Translator
	translationService = translation.NewStaticService()
	if cfg.LegacyEndpoint != "" {
		log.Printf("creating external translation client: %s", cfg.LegacyEndpoint)
		client := translation.NewHelloClient(cfg.LegacyEndpoint)
		translationService = translation.NewRemoteService(client)
	}
	translateHandler := rest.NewTranslateHandler(translationService)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", translateHandler.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)
	mux.HandleFunc("/info", handlers.Info)

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
