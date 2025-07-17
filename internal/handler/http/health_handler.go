package http

import "net/http"

// HealthHandler provides a simple status endpoint.
type HealthHandler struct {
	Version string
}

// NewHealthHandler creates a HealthHandler with the given version.
func NewHealthHandler(v string) *HealthHandler {
	return &HealthHandler{Version: v}
}

// Status returns API status and version.
func (h *HealthHandler) Status(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"status":  "ok",
		"version": h.Version,
	})
}
