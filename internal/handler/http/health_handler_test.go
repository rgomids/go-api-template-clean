package http

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/rgomids/go-api-template-clean/internal/info"
)

func TestHealthHandlerCheck(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "VERSION")
	os.WriteFile(path, []byte("0.0.1"), 0o600)
	old := info.FilePath
	info.FilePath = path
	defer func() { info.FilePath = old }()

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)

	NewHealthHandler().Check(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	var resp map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if resp["status"] != "ok" || resp["version"] != "0.0.1" {
		t.Errorf("unexpected response: %#v", resp)
	}
}

func TestHealthHandlerVersionError(t *testing.T) {
	dir := t.TempDir()
	old := info.FilePath
	info.FilePath = filepath.Join(dir, "missing")
	defer func() { info.FilePath = old }()

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)

	NewHealthHandler().Check(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", rr.Code)
	}
}
