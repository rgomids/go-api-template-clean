package scaffold

import "testing"

func TestUpdateNilSpec(t *testing.T) {
	if err := Update(nil); err == nil {
		t.Fatal("expected error for nil spec")
	}
}

func TestUpdateOK(t *testing.T) {
	spec := &ScaffoldSpec{Entity: "Test"}
	if err := Update(spec); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
