package pkgredis

import (
	"github.com/redis/go-redis/v9"
)

type (
	pkgRedis struct {
		Client *redis.Client
	}
)

func New(address string, password string, db int) IRedis {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
	return &pkgRedis{
		Client: client,
	}
}
