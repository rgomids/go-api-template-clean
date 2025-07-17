package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/rgomids/go-api-template-clean/internal/app"
)

func TestRegisterRoutes(t *testing.T) {
	r := chi.NewRouter()
	c := app.BuildContainer()
	registerRoutes(r, c)
	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users", "application/json", nil)
	if err != nil || resp.StatusCode == http.StatusNotFound {
		t.Fatalf("route not registered: %v", err)
	}
}

func TestMainFunction(t *testing.T) {
	defer func() { listenAndServe = http.ListenAndServe }()
	listenAndServe = func(string, http.Handler) error { return nil }
	t.Setenv("DATABASE_URL", "db")
	t.Setenv("PORT", "0")
	main()
	listenAndServe = func(string, http.Handler) error { return errors.New("x") }
	main()
}

func TestMainFailure(t *testing.T) {
	defer func() { listenAndServe = http.ListenAndServe }()
	listenAndServe = func(string, http.Handler) error { return nil }
	os.Unsetenv("DATABASE_URL")
	main()
}
