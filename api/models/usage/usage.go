package apimodelsusage

import (
	"context"
	"fmt"
	"time"

	apidto "github.com/Rizkyprawirap/nextmed-activity-tracking/api/dto"
	pkgpostgre "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/postgre"
	pkgredis "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/redis"
)

type model struct {
	pkgPostgre pkgpostgre.IPkgPostgreDB
	pkgRedis   pkgredis.IRedis
}

func New(
	pkgPostgre pkgpostgre.IPkgPostgreDB,
	pkgRedis pkgredis.IRedis,
) IModelUsage {
	return &model{
		pkgPostgre: pkgPostgre,
		pkgRedis:   pkgRedis,
	}
}

func (m *model) GetDailyUsage(ctx context.Context, req GetDailyUsageRequest) (*GetDailyUsageResponse, error) {

	version, err := m.pkgRedis.GetUsageCacheVersion(ctx)
	if err != nil {
		version = 1
	}

	cacheKey := fmt.Sprintf("cache:usage:daily:v%d:%s", version, req.APIKey)

	var cached []apidto.Usage
	found, _ := m.pkgRedis.GetJSON(ctx, cacheKey, &cached)
	if found {
		return (*GetDailyUsageResponse)(&cached), nil
	}

	query := `
		SELECT DATE(timestamp) AS day, COUNT(*) AS total
		FROM logs
		WHERE api_key = $1
		AND timestamp >= NOW() - INTERVAL '7 days'
		GROUP BY day
		ORDER BY day ASC;
	`

	rows, err := m.pkgPostgre.GetConnection().QueryContext(ctx, query, req.APIKey)
	if err != nil {
		return nil, fmt.Errorf("query daily usage failed: %w", err)
	}
	defer rows.Close()

	usageList := make([]apidto.Usage, 0)

	for rows.Next() {
		var (
			date  time.Time
			total int
		)

		if err := rows.Scan(&date, &total); err != nil {
			return nil, fmt.Errorf("scan daily usage failed: %w", err)
		}

		usageList = append(usageList, apidto.Usage{
			Date:  date.Format("2006-01-02"),
			Total: total,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	_ = m.pkgRedis.SetJSON(ctx, cacheKey, usageList, time.Hour)

	return (*GetDailyUsageResponse)(&usageList), nil
}

func (m *model) GetTopClientUsage(ctx context.Context) (*GetTopClientUsageResponse, error) {

	cacheKey := "cache:usage:top:latest"

	var cached []apidto.TopClientUsage
	found, _ := m.pkgRedis.GetJSON(ctx, cacheKey, &cached)
	if found {
		return (*GetTopClientUsageResponse)(&cached), nil
	}

	query := `
		SELECT c.client_id, c.name, c.email, c.api_key, COUNT(l.log_id) AS total
		FROM logs l
		JOIN clients c ON c.client_id = l.client_id
		WHERE l.timestamp >= NOW() - INTERVAL '24 hours'
		GROUP BY c.client_id, c.name, c.email, c.api_key
		ORDER BY total DESC
		LIMIT 3;
	`

	rows, err := m.pkgPostgre.GetConnection().QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []apidto.TopClientUsage{}

	for rows.Next() {
		var usage apidto.TopClientUsage
		if err := rows.Scan(
			&usage.ClientID,
			&usage.Name,
			&usage.Email,
			&usage.APIKey,
			&usage.Total,
		); err != nil {
			return nil, err
		}

		results = append(results, usage)
	}

	_ = m.pkgRedis.SetJSON(ctx, cacheKey, results, time.Hour)

	return (*GetTopClientUsageResponse)(&results), nil
}
