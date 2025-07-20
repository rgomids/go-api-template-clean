package util

import (
	"testing"

	"github.com/google/uuid"
)

func TestGenerateID(t *testing.T) {
	id1 := GenerateID()
	if id1 == "" {
		t.Fatal("expected non-empty id")
	}
	if _, err := uuid.Parse(id1); err != nil {
		t.Fatalf("invalid uuid: %v", err)
	}
	if id2 := GenerateID(); id1 == id2 {
		t.Error("expected unique ids")
	}
}
