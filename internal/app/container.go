package app

// container.go centralizes dependency injection for the application. The
// concrete implementations are intentionally simple and work only as
// placeholders so the dependency graph can be composed without coupling the
// rest of the code to specific technologies.

// UserRepository defines the required persistence methods for users. The
// interface is intentionally empty for this template.
type UserRepository interface{}

// UserService exposes user related business operations. It is represented as an
// interface to allow multiple implementations.
type UserService interface{}

// userService is a minimal implementation of UserService used only for
// injection demonstration.
type userService struct {
	repo UserRepository
}

// UserHandler deals with HTTP requests related to users. Real handler logic is
// omitted in this template.
type UserHandler struct {
	service UserService
}

// AppContainer groups all dependencies that can be injected throughout the
// application.
type AppContainer struct {
	UserService UserService
	UserHandler *UserHandler
}

// NewUserRepository constructs a new UserRepository. At this stage it returns a
// nil implementation, serving only to demonstrate the dependency chain.
func NewUserRepository() UserRepository {
	return nil
}

// NewUserService receives a UserRepository and returns an instance of
// UserService. The concrete implementation is minimal, focusing on explicit
// injection of dependencies.
func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

// NewUserHandler builds a UserHandler using the provided UserService.
func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

// BuildContainer assembles all application dependencies and returns a fully
// populated AppContainer.
func BuildContainer() *AppContainer {
	repo := NewUserRepository()
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	return &AppContainer{
		UserService: service,
		UserHandler: handler,
	}
}
