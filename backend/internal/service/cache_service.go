package service

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheService interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}

type redisCacheService struct {
	client *redis.Client
}

func NewCacheService(client *redis.Client) CacheService {
	return &redisCacheService{client: client}
}

func (s *redisCacheService) Get(ctx context.Context, key string) (string, error) {
	val, err := s.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // cache miss, return empty string with no error
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func (s *redisCacheService) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return s.client.Set(ctx, key, value, expiration).Err()
}
