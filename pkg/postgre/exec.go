package pkgpostgre

import (
	"context"
	"database/sql"
)

func (p *pkgPostgre) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return p.pool.ExecContext(ctx, query, args...)
}
