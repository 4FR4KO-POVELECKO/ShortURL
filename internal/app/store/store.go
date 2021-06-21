package store

type Store interface {
	Set(key string)
	Get(key string) string
}
