package usecase

import (
	"testing"

	"github.com/rgomids/go-api-template-clean/pkg/util"
)

func TestGenerateID(t *testing.T) {
	id1 := util.GenerateID()
	if id1 == "" {
		t.Fatal("expected id")
	}
	id2 := util.GenerateID()
	if id1 == id2 {
		t.Fatal("ids should be unique")
	}
}
