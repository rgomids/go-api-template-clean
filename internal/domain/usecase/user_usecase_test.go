package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/rgomids/go-api-template-clean/internal/domain/entity"
)

type stubRepo struct {
	saved     *entity.User
	deleted   string
	saveErr   error
	deleteErr error
}

func (s *stubRepo) FindByID(id string) (*entity.User, error) { return nil, nil }
func (s *stubRepo) Save(u *entity.User) error {
	s.saved = u
	return s.saveErr
}
func (s *stubRepo) Delete(id string) error {
	s.deleted = id
	return s.deleteErr
}

func TestRegisterUserSuccess(t *testing.T) {
	repo := &stubRepo{}
	uc := NewUserUseCase(repo)

	u, err := uc.RegisterUser("Jon", "jon@example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u == nil || u.ID == "" {
		t.Fatal("user not created")
	}
	if repo.saved == nil || repo.saved.Email != "jon@example.com" {
		t.Fatal("user was not saved")
	}
	if time.Since(repo.saved.CreatedAt) > time.Second {
		t.Error("created time not set correctly")
	}
}

func TestRegisterUserInvalidEmail(t *testing.T) {
	repo := &stubRepo{}
	uc := NewUserUseCase(repo)

	if _, err := uc.RegisterUser("Jon", "invalid"); err == nil {
		t.Fatal("expected error")
	}
	if repo.saved != nil {
		t.Error("user should not be saved")
	}
}

func TestRegisterUserRepoError(t *testing.T) {
	repo := &stubRepo{saveErr: errors.New("save fail")}
	uc := NewUserUseCase(repo)

	if _, err := uc.RegisterUser("Jon", "jon@example.com"); err == nil {
		t.Fatal("expected repo error")
	}
}

func TestRemoveUser(t *testing.T) {
	repo := &stubRepo{}
	uc := NewUserUseCase(repo)

	if err := uc.RemoveUser("123"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if repo.deleted != "123" {
		t.Errorf("expected delete id '123', got %s", repo.deleted)
	}
}

func TestRemoveUserRepoError(t *testing.T) {
	repo := &stubRepo{deleteErr: errors.New("delete fail")}
	uc := NewUserUseCase(repo)

	if err := uc.RemoveUser("123"); err == nil {
		t.Fatal("expected repo error")
	}
}
