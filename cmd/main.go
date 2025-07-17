package main

import (
	"log"
	"net/http"

	"github.com/seuusuario/go-api-template-clean/internal/app"
	"github.com/seuusuario/go-api-template-clean/internal/config"
)

// Entry point of the application.
// It loads the configuration, builds the dependency container, registers the
// routes and finally starts the HTTP server.
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	container := app.BuildContainer()

	mux := http.NewServeMux()
	registerRoutes(mux, container)

	// Use configured address for the HTTP server.
	addr := cfg.ServerAddr

	log.Printf("starting HTTP server on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

// registerRoutes attaches handlers to the HTTP mux. The real routes should be
// defined here using the handlers provided by the container.
func registerRoutes(mux *http.ServeMux, c *app.AppContainer) {
	// TODO: register application routes using c.UserHandler
	_ = c
}
