package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rgomids/go-api-template-clean/internal/app"
	"github.com/rgomids/go-api-template-clean/internal/config"
	httpmiddleware "github.com/rgomids/go-api-template-clean/internal/handler/http/middleware"
	httproutes "github.com/rgomids/go-api-template-clean/internal/handler/http/routes"
	"github.com/rgomids/go-api-template-clean/pkg/version"
)

var listenAndServe = http.ListenAndServe

// Entry point of the application.
// It loads the configuration, builds the dependency container, registers the
// routes and finally starts the HTTP server.
func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("failed to load configuration: %v", err)
		return
	}

	v, err := version.Load("VERSION")
	if err != nil {
		log.Printf("failed to load version: %v", err)
	}

	container := app.BuildContainer(v)

	router := chi.NewRouter()
	router.Use(httpmiddleware.LoggerMiddleware)
	registerRoutes(router, container)

	// Use configured port for the HTTP server.
	addr := ":" + cfg.Port

	log.Printf("starting HTTP server on %s", addr)
	if err := listenAndServe(addr, router); err != nil {
		log.Printf("server error: %v", err)
	}
}

// registerRoutes attaches handlers to the router using the container handlers.
func registerRoutes(r *chi.Mux, c *app.AppContainer) {
	httproutes.RegisterRoutes(r, c.UserHandler)
	r.Get("/health", c.HealthHandler.Status)
}
