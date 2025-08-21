package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/blacknaml/hello-api/handlers"
	"github.com/blacknaml/hello-api/handlers/rest"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = ":80"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", rest.TranslateHandler)
	mux.HandleFunc("/helth", handlers.HealthCheck)

	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
