package testing

import (
	"fmt"
	"testing"
)

func TestStubServiceExecute(t *testing.T) {
	s := StubService[int]{Response: 42}
	res, err := s.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != 42 {
		t.Fatalf("expected 42, got %d", res)
	}
}

func TestStubServiceExecuteError(t *testing.T) {
	expectedErr := fmt.Errorf("fail")
	s := StubService[int]{Error: expectedErr}
	_, err := s.Execute()
	if err != expectedErr {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
}

func TestCRUDStubServiceDefaults(t *testing.T) {
	s := CRUDStubService[int]{}
	if err := s.Create(nil); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if _, err := s.GetByID("1"); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if _, err := s.List(map[string]interface{}{}); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if err := s.Update("1", nil); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if err := s.Delete("1"); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

func TestCRUDStubServiceOverrides(t *testing.T) {
	called := false
	s := CRUDStubService[int]{
		CreateFn: func(*int) error { called = true; return fmt.Errorf("fail") },
	}
	if err := s.Create(nil); err == nil {
		t.Fatal("expected error")
	}
	if !called {
		t.Fatal("override not called")
	}
}
