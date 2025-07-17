package http

import (
	"net/http"

	"github.com/rgomids/go-api-template-clean/internal/info"
)

// HealthHandler exposes a simple health check endpoint.
type HealthHandler struct{}

// NewHealthHandler returns a new HealthHandler instance.
func NewHealthHandler() *HealthHandler { return &HealthHandler{} }

// Check writes the API status and version.
func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	version, err := info.ReadVersion()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "unable to read version")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"version": version,
	})
}
