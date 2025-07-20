package db

// RowScanner abstracts scanning a single row result.
type RowScanner interface {
	Scan(dest ...any) error
}
