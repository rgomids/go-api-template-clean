package service

import "github.com/rgomids/go-api-template-clean/internal/domain/entity"

// UserService exposes user-related business operations.
type UserService interface {
	RegisterUser(name, email string) (*entity.User, error)
	RemoveUser(id string) error
}
