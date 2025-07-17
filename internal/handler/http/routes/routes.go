package routes

import (
	"github.com/go-chi/chi/v5"
	httphandler "github.com/rgomids/go-api-template-clean/internal/handler/http"
)

// RegisterRoutes maps API endpoints to handlers.
func RegisterRoutes(router *chi.Mux, handler *httphandler.UserHandler) {
	router.Route("/users", func(r chi.Router) {
		r.Post("/", handler.Register)
		r.Delete("/{id}", handler.Delete)
	})

	// [AUTO-GENERATED-ROUTES]
}
