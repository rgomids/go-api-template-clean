package util

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestJSONAndError(t *testing.T) {
	rr := httptest.NewRecorder()
	JSON(rr, 201, map[string]string{"ok": "yes"})
	if rr.Code != 201 {
		t.Fatalf("expected status 201, got %d", rr.Code)
	}
	var data map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &data); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if data["ok"] != "yes" {
		t.Errorf("unexpected body: %v", data)
	}

	rr = httptest.NewRecorder()
	Error(rr, 400, "bad")
	if rr.Code != 400 {
		t.Fatalf("expected status 400, got %d", rr.Code)
	}
	if err := json.Unmarshal(rr.Body.Bytes(), &data); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if data["error"] != "bad" {
		t.Errorf("unexpected error body: %v", data)
	}
}
