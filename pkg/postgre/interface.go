package pkgpostgre

import (
	"context"
	"database/sql"
)

type IPkgPostgreDB interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (IPkgPostgreTx, error)
	GetConnection() *sql.DB
}

type IPkgPostgreTx interface {
	Commit() error
	Rollback() error
	GetConnection() *sql.Tx
}
