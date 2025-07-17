package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rgomids/go-api-template-clean/internal/app"
	"github.com/rgomids/go-api-template-clean/internal/config"
	httpmiddleware "github.com/rgomids/go-api-template-clean/internal/handler/http/middleware"
	httproutes "github.com/rgomids/go-api-template-clean/internal/handler/http/routes"
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

	router := chi.NewRouter()
	router.Use(httpmiddleware.LoggerMiddleware)
	registerRoutes(router, container)

	// Use configured address for the HTTP server.
	addr := cfg.ServerAddr

	log.Printf("starting HTTP server on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

// registerRoutes attaches handlers to the router using the container handlers.
func registerRoutes(r *chi.Mux, c *app.AppContainer) {
	httproutes.RegisterRoutes(r, c.UserHandler)
}
