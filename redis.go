package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// Service wraps a redis client.
type Service struct {
	LockExpireUnit time.Duration

	*redis.Client
}

// Dial connects grpc server and call Register method.
func (s *Service) Dial(ctx context.Context, cfg Config) error {
	s.LockExpireUnit = time.Duration(cfg.LockExpireUnit) * time.Millisecond

	s.Client = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	_, err := s.Client.Ping(ctx).Result()

	return err
}

func (s *Service) Close(ctx context.Context) error {
	if s.Client != nil {
		s.Client.Close()
	}

	return nil
}
