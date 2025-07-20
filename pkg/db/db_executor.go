package db

import "database/sql"

type DBExecutor interface {
	Query(query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) RowScanner
	Exec(query string, args ...any) (sql.Result, error)
}
