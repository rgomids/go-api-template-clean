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

	paths := []string{
		filepath.Join("internal/domain/entity", "test.go"),
		filepath.Join("internal/domain/usecase", "test_usecase_test.go"),
		filepath.Join("internal/handler/http", "test_handler_test.go"),
		filepath.Join("mocks", "test_repository_mock.go"),
	}

	for _, p := range paths {
		if _, err := os.Stat(p); err != nil {
			t.Fatalf("expected file %s not created: %v", p, err)
		}
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
