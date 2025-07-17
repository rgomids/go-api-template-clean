package usecase

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/seuusuario/go-api-template-clean/internal/domain/entity"
	"github.com/seuusuario/go-api-template-clean/internal/domain/repository"
	"github.com/seuusuario/go-api-template-clean/internal/domain/service"
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
		ID:        generateID(),
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

func generateID() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	return hex.EncodeToString(b)
}
