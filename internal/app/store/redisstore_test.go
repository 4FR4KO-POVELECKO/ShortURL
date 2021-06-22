package store_test

import (
	"ShortURL/internal/app/store"
	"log"
	"os"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

var client *redis.Client

var (
	key = "key"
	val = "val"
)

func TestMain(m *testing.M) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatal(err)
	}

	client = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	code := m.Run()
	os.Exit(code)
}

func TestStore_Set(t *testing.T) {
	exp := time.Duration(0)

	mock := redismock.NewNiceMock(client)
	mock.On("Set", key, val, exp).Return(redis.NewStatusResult("", nil))

	r := store.NewStoreRedis(mock)
	err := r.Set(key, val, exp)
	assert.NoError(t, err)

}
func TestStore_Get(t *testing.T) {
	mock := redismock.NewNiceMock(client)
	mock.On("Get", key).Return(redis.NewStringResult(val, nil))

	r := store.NewStoreRedis(mock)
	res, err := r.Get(key)

	assert.NoError(t, err)
	assert.Equal(t, val, res)
}
