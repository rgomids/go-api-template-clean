package app

import "testing"

func TestBuildContainer(t *testing.T) {
	c := BuildContainer("1.0.0")
	if c == nil || c.UserService == nil || c.UserHandler == nil || c.HealthHandler == nil {
		t.Fatal("container not built correctly")
	}
}

func TestConstructors(t *testing.T) {
	repo := NewUserRepository()
	if repo == nil {
		t.Fatal("repo nil")
	}
	svc := NewUserService(repo)
	if svc == nil {
		t.Fatal("service nil")
	}
	h := NewUserHandler(svc)
	if h == nil {
		t.Fatal("handler nil")
	}

	// cover dummy repository methods
	dr := repo.(*dummyUserRepository)
	dr.FindByID("1")
	dr.Save(nil)
	dr.Delete("1")
}
