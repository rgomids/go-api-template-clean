package repository

import "github.com/seuusuario/go-api-template-clean/internal/domain/entity"

// UserRepository defines persistence operations for User entities.
type UserRepository interface {
	FindByID(id string) (*entity.User, error)
	Save(user *entity.User) error
	Delete(id string) error
}
