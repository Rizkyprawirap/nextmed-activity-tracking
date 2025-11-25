package pkg

import (
	pkgpostgre "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/postgre"
	pkgredis "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/redis"
)

type (
	PkgDatabase struct {
		PkgRedis   pkgredis.IRedis
		PkgPostgre pkgpostgre.IPkgPostgreDB
	}
)
