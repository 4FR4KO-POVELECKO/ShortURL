package store

type Store interface {
	Set(key string) string
	Get(key string) string
}
