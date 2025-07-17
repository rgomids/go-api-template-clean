package scaffold

import "testing"

func TestParseSuccess(t *testing.T) {
	args := []string{"Invoice", "number:string", "user:belongsTo"}
	spec, err := Parse(args)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if spec.Entity != "Invoice" || len(spec.Fields) != 1 || len(spec.Relationships) != 1 {
		t.Fatalf("parsed spec not as expected: %+v", spec)
	}
}

func TestParseInvalidArg(t *testing.T) {
	if _, err := Parse([]string{"Invoice", "badarg"}); err == nil {
		t.Fatal("expected error for invalid arg")
	}
}

func TestParseNoEntity(t *testing.T) {
	if _, err := Parse([]string{}); err == nil {
		t.Fatal("expected error for missing entity")
	}
}
