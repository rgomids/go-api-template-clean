package http

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestWriteJSONAndError(t *testing.T) {
	rr := httptest.NewRecorder()
	writeJSON(rr, 200, map[string]string{"foo": "bar"})
	if rr.Code != 200 {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	var data map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &data); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if data["foo"] != "bar" {
		t.Errorf("unexpected body: %v", data)
	}

	rr = httptest.NewRecorder()
	writeError(rr, 500, "boom")
	if rr.Code != 500 {
		t.Fatalf("expected 500, got %d", rr.Code)
	}
	if err := json.Unmarshal(rr.Body.Bytes(), &data); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if data["error"] != "boom" {
		t.Errorf("unexpected error message: %v", data)
	}
}
