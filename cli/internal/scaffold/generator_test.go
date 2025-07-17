package scaffold

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateNilSpec(t *testing.T) {
	if err := Generate(nil); err == nil {
		t.Fatal("expected error for nil spec")
	}
}

func TestGenerateOK(t *testing.T) {
	dir := t.TempDir()
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	os.Chdir(dir)

	spec := &ScaffoldSpec{EntityName: "Test"}
	if err := Generate(spec); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, err := os.Stat(filepath.Join("internal/domain/entity", "test.go")); err != nil {
		t.Fatalf("expected file not created: %v", err)
	}
}

func TestHelpers(t *testing.T) {
	if toPascal("user_profile") != "UserProfile" {
		t.Errorf("toPascal failed")
	}

	if goType(FieldSpec{Type: "int"}) != "int" {
		t.Errorf("goType int failed")
	}
	if sqlType(FieldSpec{Type: "json"}) != "jsonb" {
		t.Errorf("sqlType json failed")
	}
}
