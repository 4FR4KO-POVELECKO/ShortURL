package store

import "time"

type Store interface {
	Set(key string, value string, exp time.Duration) error
	Get(key string) (string, error)
}
