package apimodelsclient

import (
	"context"
	"fmt"

	pkgpostgre "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/postgre"
	pkgredis "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/redis"
	"github.com/google/uuid"
)

type model struct {
	pkgPostgre pkgpostgre.IPkgPostgreDB
	pkgRedis   pkgredis.IRedis
}

func New(
	pkgPostgre pkgpostgre.IPkgPostgreDB,
	pkgRedis pkgredis.IRedis,
) IModelClient {
	return &model{
		pkgPostgre: pkgPostgre,
		pkgRedis:   pkgRedis,
	}
}

func (m *model) InsertClient(ctx context.Context, req InsertClientRequest) (*InsertClientResponse, error) {
	clientID := uuid.New()

	query := `
		INSERT INTO clients (
			client_id, name, email, api_key, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING client_id, name, email, api_key;
	`

	var result InsertClientResponse

	err := m.pkgPostgre.GetConnection().
		QueryRowContext(ctx, query,
			clientID,
			req.Name,
			req.Email,
			req.ApiKey,
		).
		Scan(
			&result.ClientID,
			&result.Name,
			&result.Email,
			&result.ApiKey,
		)

	if err != nil {
		return nil, fmt.Errorf("insert client failed: %w", err)
	}

	_ = m.pkgRedis.BumpUsageCacheVersion(ctx)

	_, _ = m.pkgRedis.Del(ctx, pkgredis.DelRequest{
		Key: "cache:usage:top:latest",
	})

	return &result, nil
}

func (m *model) GetClientByAPIKey(ctx context.Context, req GetClientByAPIKeyRequest) (*GetClientByAPIKeyResponse, error) {
	query := `
		SELECT client_id, name, email, api_key
		FROM clients
		WHERE api_key = $1
		LIMIT 1;
	`

	var c GetClientByAPIKeyResponse

	err := m.pkgPostgre.GetConnection().
		QueryRowContext(ctx, query, req.APIKey).
		Scan(&c.ClientID, &c.Name, &c.Email, &c.ApiKey)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (m *model) InsertLog(ctx context.Context, req InsertLogRequest) error {
	logID := uuid.New()

	query := `
		INSERT INTO logs (
			log_id, client_id, api_key, ip, endpoint, timestamp, created_at, updated_at
		)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW(), NOW());
	`

	_, err := m.pkgPostgre.GetConnection().ExecContext(ctx, query,
		logID,
		req.ClientID,
		req.APIKey,
		req.IP,
		req.Endpoint,
	)
	if err != nil {
		return fmt.Errorf("failed inserting log: %w", err)
	}

	_ = m.pkgRedis.BumpUsageCacheVersion(ctx)

	_, _ = m.pkgRedis.Del(ctx, pkgredis.DelRequest{
		Key: "cache:usage:top:latest",
	})

	return nil
}
