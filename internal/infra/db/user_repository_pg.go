package db

import (
	"database/sql"

	"github.com/rgomids/go-api-template-clean/internal/domain/entity"
	"github.com/rgomids/go-api-template-clean/internal/domain/repository"
)

// PostgresUserRepository implements UserRepository using a PostgreSQL database.
type dbExecutor interface {
	Exec(query string, args ...any) (sql.Result, error)
	QueryRow(query string, args ...any) rowScanner
}

type rowScanner interface {
	Scan(dest ...any) error
}

type PostgresUserRepository struct {
	db dbExecutor
}

// NewPostgresUserRepository creates a repository with the given DB connection.
func NewPostgresUserRepository(db dbExecutor) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

// Ensure the implementation satisfies the interface at compile time.
var _ repository.UserRepository = (*PostgresUserRepository)(nil)

// FindByID retrieves a user by id.
func (r *PostgresUserRepository) FindByID(id string) (*entity.User, error) {
	const query = `SELECT id, name, email, created_at FROM users WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var u entity.User
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}

// Save inserts a new user into the database.
func (r *PostgresUserRepository) Save(user *entity.User) error {
	const query = `INSERT INTO users (id, name, email, created_at) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, user.ID, user.Name, user.Email, user.CreatedAt)
	return err
}

// Delete removes the user with the given id.
func (r *PostgresUserRepository) Delete(id string) error {
	const query = `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
