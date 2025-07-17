package scaffold

import "testing"

func TestGenerateNilSpec(t *testing.T) {
	if err := Generate(nil); err == nil {
		t.Fatal("expected error for nil spec")
	}
}

func TestGenerateOK(t *testing.T) {
	spec := &ScaffoldSpec{EntityName: "Test"}
	if err := Generate(spec); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
