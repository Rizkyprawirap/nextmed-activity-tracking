package pkgpostgre

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)



type pkgPostgre struct {
	pool *sql.DB
}
type pkgPostgreTx struct {
	tx *sql.Tx
}


func New(cfg Config) (IPkgPostgreDB, error) {
	if cfg.SSLMode == "" {
		cfg.SSLMode = "disable"
	}
	fmt.Printf(
		"host=%s port=%d user=%s  dbname=%s sslmode=%s password=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.SSLMode, cfg.Password,
	)
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s  dbname=%s sslmode=%s password=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.SSLMode, cfg.Password,
	)
	if cfg.AppName != "" {
		dsn = dsn + fmt.Sprintf(" application_name=%s", cfg.AppName)
	}

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if cfg.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	} else {
		sqlDB.SetMaxOpenConns(10)
	}
	if cfg.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	} else {
		sqlDB.SetMaxIdleConns(5)
	}
	if cfg.ConnMaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	}
	if cfg.ConnMaxIdleTime > 0 {
		sqlDB.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		_ = sqlDB.Close()
		return nil, err
	}

	return &pkgPostgre{pool: sqlDB}, nil
}

func (p *pkgPostgre) GetConnection() *sql.DB {
	return p.pool
}
