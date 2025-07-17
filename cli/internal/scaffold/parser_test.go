package scaffold

import "testing"

func TestParseSuccess(t *testing.T) {
	args := []string{"Invoice", "number:string", "user:belongsTo"}
	spec, err := Parse(args)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if spec.EntityName != "Invoice" || len(spec.Fields) != 1 || len(spec.Relationships) != 1 {
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

func TestParseComplexSpec(t *testing.T) {
	args := []string{
		"Invoice",
		"number:string",
		"total:float",
		"status:enum[pending,paid]",
		"tags:array[string]",
		"metadata:json",
		"user:belongsTo",
	}
	spec, err := Parse(args)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(spec.Fields) != 5 {
		t.Fatalf("expected 5 fields, got %d", len(spec.Fields))
	}
	if len(spec.Relationships) != 1 {
		t.Fatalf("expected 1 relationship, got %d", len(spec.Relationships))
	}
	enumField := spec.Fields[2]
	if enumField.Type != "enum" || enumField.Subtype != "pending,paid" {
		t.Fatalf("unexpected enum field parsing: %+v", enumField)
	}
}

func TestParseArgFunction(t *testing.T) {
	f, r, err := parseArg("user:belongsTo")
	if err != nil || f != nil || r == nil {
		t.Fatalf("expected relationship parse, got f=%v r=%v err=%v", f, r, err)
	}

	f, r, err = parseArg("status:enum[pending,paid]")
	if err != nil || r != nil || f == nil {
		t.Fatalf("expected field parse, got f=%v r=%v err=%v", f, r, err)
	}
	if f.Type != "enum" || f.Subtype != "pending,paid" {
		t.Fatalf("unexpected field result: %+v", f)
	}

	if _, _, err = parseArg("badarg"); err == nil {
		t.Fatal("expected error for bad argument")
	}
}
