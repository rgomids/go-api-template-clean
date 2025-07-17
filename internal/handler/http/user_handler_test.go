package http

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/rgomids/go-api-template-clean/internal/domain/entity"
)

type stubService struct {
	regUser func(name, email string) (*entity.User, error)
	remUser func(id string) error
}

func (s *stubService) RegisterUser(name, email string) (*entity.User, error) {
	if s.regUser != nil {
		return s.regUser(name, email)
	}
	return &entity.User{ID: "1", Name: name, Email: email}, nil
}

func (s *stubService) RemoveUser(id string) error {
	if s.remUser != nil {
		return s.remUser(id)
	}
	return nil
}

func TestUserHandlerRegisterSuccess(t *testing.T) {
	h := NewUserHandler(&stubService{})
	body := bytes.NewBufferString(`{"name":"Jon","email":"jon@example.com"}`)
	req := httptest.NewRequest(http.MethodPost, "/", body)
	rr := httptest.NewRecorder()
	h.Register(rr, req)
	if rr.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rr.Code)
	}
}

func TestUserHandlerRegisterBadBody(t *testing.T) {
	h := NewUserHandler(&stubService{})
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString("{"))
	rr := httptest.NewRecorder()
	h.Register(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}

func TestUserHandlerRegisterServiceError(t *testing.T) {
	h := NewUserHandler(&stubService{regUser: func(name, email string) (*entity.User, error) {
		return nil, errors.New("fail")
	}})
	body := bytes.NewBufferString(`{"name":"Jon","email":"jon@example.com"}`)
	req := httptest.NewRequest(http.MethodPost, "/", body)
	rr := httptest.NewRecorder()
	h.Register(rr, req)
	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", rr.Code)
	}
}

func TestUserHandlerDelete(t *testing.T) {
	h := NewUserHandler(&stubService{})
	req := httptest.NewRequest(http.MethodDelete, "/123", nil)
	rr := httptest.NewRecorder()
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "123")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	h.Delete(rr, req)
	if rr.Code != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", rr.Code)
	}
}

func TestUserHandlerDeleteMissingID(t *testing.T) {
	h := NewUserHandler(&stubService{})
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rr := httptest.NewRecorder()
	h.Delete(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}

func TestUserHandlerDeleteServiceError(t *testing.T) {
	h := NewUserHandler(&stubService{remUser: func(id string) error { return errors.New("fail") }})
	req := httptest.NewRequest(http.MethodDelete, "/123", nil)
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "123")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	rr := httptest.NewRecorder()
	h.Delete(rr, req)
	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", rr.Code)
	}
}
