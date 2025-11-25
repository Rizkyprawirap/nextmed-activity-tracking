package pkgjwt

import pkgredis "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/redis"

type pkgJWT struct {
	Secret string
	Redis  pkgredis.IRedis
}

func New(
	secret string,
	redis pkgredis.IRedis,
) IPkgJWT {
	return &pkgJWT{
		Secret: secret,
		Redis:  redis,
	}
}
