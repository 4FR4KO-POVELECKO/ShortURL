package store

import (
	"time"

	"github.com/go-redis/redis"
)

type redisStore struct {
	client redis.Cmdable
}

func NewStoreRedis(client redis.Cmdable) Store {
	return &redisStore{client: client}
}

func (r *redisStore) Set(key string, value string, exp time.Duration) error {
	return r.client.Set(key, value, exp).Err()
}

func (r *redisStore) Get(key string) (string, error) {
	result := r.client.Get(key)

	return result.Result()
}

func (r *redisStore) Incr(key string) (int64, error) {
	result := r.client.Incr(key)
	return result.Result()
}
