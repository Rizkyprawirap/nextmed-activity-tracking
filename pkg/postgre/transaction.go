package pkgpostgre

import (
	"context"
	"database/sql"
)

func (p *pkgPostgre) BeginTx(ctx context.Context, opts *sql.TxOptions) (IPkgPostgreTx, error) {
	tx, err := p.pool.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &pkgPostgreTx{tx: tx}, nil
}

func (t *pkgPostgreTx) Rollback() error {
	return t.tx.Rollback()
}

func (t *pkgPostgreTx) Commit() error {
	return t.tx.Commit()
}

func (t *pkgPostgreTx) GetConnection() *sql.Tx {
	return t.tx
}
