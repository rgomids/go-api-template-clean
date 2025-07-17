package usecase

import (
	"crypto/rand"
	"errors"
	"testing"
)

type errorReader struct{}

func (errorReader) Read(p []byte) (int, error) { return 0, errors.New("fail") }

func TestGenerateID(t *testing.T) {
	id := generateID()
	if len(id) == 0 {
		t.Fatal("expected id")
	}
}

func TestGenerateIDError(t *testing.T) {
	orig := rand.Reader
	rand.Reader = errorReader{}
	id := generateID()
	rand.Reader = orig
	if id == "" {
		t.Fatal("fallback id not generated")
	}
}
