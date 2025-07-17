package scaffold

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUpdateNilSpec(t *testing.T) {
	if err := Update(nil); err == nil {
		t.Fatal("expected error for nil spec")
	}
}

func TestUpdateOK(t *testing.T) {
	dir := t.TempDir()
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	os.Chdir(dir)

	os.MkdirAll(filepath.Join("internal/handler/http/routes"), 0o755)
	os.MkdirAll(filepath.Join("internal/app"), 0o755)

	routes := `package routes

import (
    "github.com/go-chi/chi/v5"
    httphandler "github.com/rgomids/go-api-template-clean/internal/handler/http"
)

func RegisterRoutes(router *chi.Mux, handler *httphandler.UserHandler) {
    router.Route("/users", func(r chi.Router) {})
    // [AUTO-GENERATED-ROUTES]
}
`
	os.WriteFile(filepath.Join("internal/handler/http/routes", "routes.go"), []byte(routes), 0o600)

	container := `package app

import (
    domainrepo "github.com/rgomids/go-api-template-clean/internal/domain/repository"
    "github.com/rgomids/go-api-template-clean/internal/domain/service"
    httphandler "github.com/rgomids/go-api-template-clean/internal/handler/http"
)

type AppContainer struct {
    UserService service.UserService
    UserHandler *httphandler.UserHandler
    // [AUTO-GENERATED-CONTAINER]
}

func BuildContainer(version string) *AppContainer {
    repo := NewUserRepository()
    svc := NewUserService(repo)
    handler := NewUserHandler(svc)
    // [AUTO-GENERATED-CONTAINER]
    return &AppContainer{
        UserService: svc,
        UserHandler: handler,
        // [AUTO-GENERATED-CONTAINER]
    }
}
`
	os.WriteFile(filepath.Join("internal/app", "container.go"), []byte(container), 0o600)

	spec := &ScaffoldSpec{EntityName: "Invoice"}
	if err := Update(spec); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	b, _ := os.ReadFile(filepath.Join("internal/handler/http/routes", "routes.go"))
	content := string(b)
	if !strings.Contains(content, "/invoices") {
		t.Fatalf("route not inserted")
	}
	if strings.Count(content, "/invoices") != 1 {
		t.Fatalf("route duplicated")
	}

	b, _ = os.ReadFile(filepath.Join("internal/app", "container.go"))
	c := string(b)
	if !strings.Contains(c, "InvoiceHandler") {
		t.Fatalf("container not updated")
	}
	if strings.Count(c, "InvoiceHandler") != 4 {
		t.Fatalf("container content unexpected")
	}

	if err := Update(spec); err != nil {
		t.Fatalf("unexpected error on second update: %v", err)
	}

	b, _ = os.ReadFile(filepath.Join("internal/handler/http/routes", "routes.go"))
	if strings.Count(string(b), "/invoices") != 1 {
		t.Fatalf("idempotence failed for routes")
	}
	b, _ = os.ReadFile(filepath.Join("internal/app", "container.go"))
	if strings.Count(string(b), "InvoiceHandler") != 4 {
		t.Fatalf("idempotence failed for container")
	}
}
