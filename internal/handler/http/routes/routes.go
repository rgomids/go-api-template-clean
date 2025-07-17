package routes

import (
	"github.com/go-chi/chi/v5"
	httphandler "github.com/rgomids/go-api-template-clean/internal/handler/http"
)

// RegisterRoutes maps API endpoints to handlers.
func RegisterRoutes(router *chi.Mux, userHandler *httphandler.UserHandler, healthHandler *httphandler.HealthHandler) {
	router.Get("/health", healthHandler.Check)

	router.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.Register)
		r.Delete("/{id}", userHandler.Delete)
	})
}
