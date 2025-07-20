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
