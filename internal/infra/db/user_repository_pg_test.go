package db

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/rgomids/go-api-template-clean/internal/domain/entity"
)

type stubRow struct {
	user *entity.User
	err  error
}

func (s stubRow) Scan(dest ...any) error {
	if s.err != nil {
		return s.err
	}
	if len(dest) != 4 {
		return errors.New("bad dest")
	}
	*dest[0].(*string) = s.user.ID
	*dest[1].(*string) = s.user.Name
	*dest[2].(*string) = s.user.Email
	*dest[3].(*time.Time) = s.user.CreatedAt
	return nil
}

type stubDB struct {
	row     rowScanner
	execErr error
}

func (s *stubDB) Exec(query string, args ...any) (sql.Result, error) {
	return nil, s.execErr
}

func (s *stubDB) QueryRow(query string, args ...any) rowScanner {
	return s.row
}

func TestPostgresUserRepositorySuccess(t *testing.T) {
	u := &entity.User{ID: "1", Name: "Jon", Email: "jon@example.com", CreatedAt: time.Now()}
	repo := NewPostgresUserRepository(&stubDB{row: stubRow{user: u}})
	if err := repo.Save(u); err != nil {
		t.Fatalf("save error: %v", err)
	}
	got, err := repo.FindByID("1")
	if err != nil {
		t.Fatalf("find error: %v", err)
	}
	if got.Email != u.Email {
		t.Errorf("unexpected user: %+v", got)
	}
	if err := repo.Delete("1"); err != nil {
		t.Fatalf("delete error: %v", err)
	}
}

func TestPostgresUserRepositoryErrors(t *testing.T) {
	repo := NewPostgresUserRepository(&stubDB{row: stubRow{err: errors.New("scan")}, execErr: errors.New("exec")})
	if _, err := repo.FindByID("1"); err == nil {
		t.Fatal("expected scan error")
	}
	if err := repo.Save(&entity.User{}); err == nil {
		t.Fatal("expected exec error")
	}
	if err := repo.Delete("1"); err == nil {
		t.Fatal("expected exec error")
	}
}
