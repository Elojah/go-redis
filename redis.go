package redis

import (
	"context"

	"github.com/redis/rueidis"
)

// Service wraps a redis client.
type Service struct {
	rueidis.Client
}

// Dial connects grpc server and call Register method.
func (s *Service) Dial(ctx context.Context, cfg Config) error {
	var err error

	s.Client, err = rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{cfg.Addr},
		Password:    cfg.Password,
		SelectDB:    cfg.DB,
	})
	if err != nil {
		return err
	}

	return s.Do(ctx, s.B().Ping().Build()).Error()
}

func (s *Service) Close(ctx context.Context) error {
	if s.Client != nil {
		s.Client.Close()
	}

	return nil
}
