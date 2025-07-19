package formatter

import "testing"

func TestToSnake(t *testing.T) {
	if ToSnake("CamelCase") != "camel_case" {
		t.Errorf("unexpected result")
	}
	if ToSnake("lower") != "lower" {
		t.Errorf("unexpected result for lower")
	}
}

func TestBuildEntityMeta(t *testing.T) {
	meta := BuildEntityMeta("library")
	if meta.PluralPascal != "Libraries" {
		t.Errorf("expected Libraries, got %s", meta.PluralPascal)
	}
	if meta.PluralSnake != "libraries" {
		t.Errorf("expected libraries, got %s", meta.PluralSnake)
	}
	if meta.PluralKebab != "libraries" {
		t.Errorf("expected libraries, got %s", meta.PluralKebab)
	}

	meta = BuildEntityMeta("person")
	if meta.PluralPascal != "People" {
		t.Errorf("expected People, got %s", meta.PluralPascal)
	}
	if meta.PluralSnake != "people" {
		t.Errorf("expected people, got %s", meta.PluralSnake)
	}
	if meta.EntityName != "Person" {
		t.Errorf("expected Person, got %s", meta.EntityName)
	}
}
