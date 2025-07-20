package usecase

import (
	"fmt"
	"time"

	"github.com/rgomids/go-api-template-clean/internal/domain/entity"
	"github.com/rgomids/go-api-template-clean/internal/domain/repository"
	"github.com/rgomids/go-api-template-clean/internal/domain/service"
	"github.com/rgomids/go-api-template-clean/pkg/util"
)

// UserUseCase implements UserService, orchestrating business logic.
type UserUseCase struct {
	repo repository.UserRepository
}

// NewUserUseCase creates a new UserUseCase with its required repository.
func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

// RegisterUser creates a new user after validating the email.
func (uc *UserUseCase) RegisterUser(name, email string) (*entity.User, error) {
	user := &entity.User{
		ID:        util.GenerateID(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}

	if !user.IsValidEmail() {
		return nil, fmt.Errorf("invalid email")
	}

	if err := uc.repo.Save(user); err != nil {
		return nil, err
	}
	return user, nil
}

// RemoveUser deletes the user by id.
func (uc *UserUseCase) RemoveUser(id string) error {
	return uc.repo.Delete(id)
}

// Ensure UserUseCase satisfies UserService at compile time.
var _ service.UserService = (*UserUseCase)(nil)
