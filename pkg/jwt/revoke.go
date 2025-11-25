package pkgjwt

import (
	"context"
	"errors"

	pkgredis "github.com/Rizkyprawirap/nextmed-activity-tracking/pkg/redis"
)

func (p *pkgJWT) Revoke(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("empty id")
	}
	_, err := p.Redis.Del(ctx, pkgredis.DelRequest{
		Key: id,
	})
	return err
}
