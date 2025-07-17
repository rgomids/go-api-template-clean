package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/rgomids/go-api-template-clean/internal/domain/entity"
	httphandler "github.com/rgomids/go-api-template-clean/internal/handler/http"
)

type fakeService struct{}

func (fakeService) RegisterUser(name, email string) (*entity.User, error) { return &entity.User{}, nil }
func (fakeService) RemoveUser(id string) error                            { return nil }

func TestRegisterRoutes(t *testing.T) {
	h := httphandler.NewUserHandler(new(fakeService))
	r := chi.NewRouter()
	RegisterRoutes(r, h)

	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users", "application/json", nil)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	if resp.StatusCode == http.StatusNotFound {
		t.Fatal("route not registered")
	}
}
