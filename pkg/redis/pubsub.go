package pkgredis

import (
	"context"
)

func (r *pkgRedis) Publish(ctx context.Context, channel string, message string) error {
	return r.Client.Publish(ctx, channel, message).Err()
}

func (r *pkgRedis) Subscribe(ctx context.Context, channel string) (<-chan string, error) {
	sub := r.Client.Subscribe(ctx, channel)

	ch := make(chan string)

	go func() {
		for msg := range sub.Channel() {
			ch <- msg.Payload
		}
	}()

	return ch, nil
}
