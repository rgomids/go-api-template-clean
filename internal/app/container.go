package app

import (
	"github.com/rgomids/go-api-template-clean/internal/domain/entity"
	domainrepo "github.com/rgomids/go-api-template-clean/internal/domain/repository"
	"github.com/rgomids/go-api-template-clean/internal/domain/service"
	"github.com/rgomids/go-api-template-clean/internal/domain/usecase"
	httphandler "github.com/rgomids/go-api-template-clean/internal/handler/http"
)

// AppContainer groups dependencies for injection across the application.
type AppContainer struct {
	UserService   service.UserService
	UserHandler   *httphandler.UserHandler
	HealthHandler *httphandler.HealthHandler
}

// dummyUserRepository is a minimal in-memory implementation of UserRepository.
type dummyUserRepository struct{}

func (d *dummyUserRepository) FindByID(id string) (*entity.User, error) { return nil, nil }
func (d *dummyUserRepository) Save(u *entity.User) error                { return nil }
func (d *dummyUserRepository) Delete(id string) error                   { return nil }

// NewUserRepository constructs a repository instance.
func NewUserRepository() domainrepo.UserRepository {
	return &dummyUserRepository{}
}

// NewUserService builds the user service using the provided repository.
func NewUserService(repo domainrepo.UserRepository) service.UserService {
	return usecase.NewUserUseCase(repo)
}

// NewUserHandler builds an HTTP handler for users.
func NewUserHandler(svc service.UserService) *httphandler.UserHandler {
	return httphandler.NewUserHandler(svc)
}

// NewHealthHandler builds an HTTP handler for health checks.
func NewHealthHandler() *httphandler.HealthHandler {
	return httphandler.NewHealthHandler()
}

// BuildContainer assembles all dependencies of the application.
func BuildContainer() *AppContainer {
	repo := NewUserRepository()
	svc := NewUserService(repo)
	userHandler := NewUserHandler(svc)
	healthHandler := NewHealthHandler()

	return &AppContainer{
		UserService:   svc,
		UserHandler:   userHandler,
		HealthHandler: healthHandler,
	}
}
