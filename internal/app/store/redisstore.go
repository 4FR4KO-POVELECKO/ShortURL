package store

import (
	"time"

	"github.com/go-redis/redis"
)

type redisStore struct {
	host    string
	db      int
	expires time.Duration
}

func NewStoreRedis(host string, db int, expires time.Duration) Store {
	return &redisStore{
		host:    host,
		db:      db,
		expires: expires,
	}
}

func (r *redisStore) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     r.host,
		Password: "",
		DB:       r.db,
	})
}

func (r *redisStore) Set(key string, value string) error {
	client := r.getClient()

	err := client.Set(key, value, r.expires*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *redisStore) Get(key string) (string, error) {
	client := r.getClient()

	val, err := client.Get(key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}
