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
