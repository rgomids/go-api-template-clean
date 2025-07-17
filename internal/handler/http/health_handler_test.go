package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthStatus(t *testing.T) {
	h := NewHealthHandler("1.0.0")
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()

	h.Status(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	expected := `{"status":"ok","version":"1.0.0"}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("unexpected body: %s", rr.Body.String())
	}
}
