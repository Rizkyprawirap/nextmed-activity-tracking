package pkgpostgre

import (
	"context"
	"database/sql"
)

func (p *pkgPostgre) QueryRowContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return p.pool.Query(query, args...)
}
