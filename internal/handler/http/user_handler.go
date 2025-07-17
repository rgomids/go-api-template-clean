package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rgomids/go-api-template-clean/internal/domain/service"
)

// UserHandler adapts UserService to HTTP transport.
type UserHandler struct {
	service service.UserService
}

// NewUserHandler creates a new UserHandler with its dependencies.
func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// Register handles user registration requests.
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	user, err := h.service.RegisterUser(req.Name, req.Email)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, user)
}

// Delete handles removal of a user by id.
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		writeError(w, http.StatusBadRequest, "missing id")
		return
	}

	if err := h.service.RemoveUser(id); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
