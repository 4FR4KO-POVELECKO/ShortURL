package grpcserver_test

import (
	"ShortURL/internal/app/store"
	"ShortURL/pkg/api"
	"ShortURL/pkg/grpcserver"
	"context"
	"log"
	"net"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func getConn() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	mr, err := miniredis.Run()
	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	store := store.NewStoreRedis(client)
	server := grpc.NewServer()
	srv := &grpcserver.GRPCServer{Store: store}
	api.RegisterShortlinkServer(server, srv)

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestShortlinkServer_Create(t *testing.T) {
	testCases := []struct {
		name    string
		payload *api.OriginUrl
		err     bool
	}{
		{
			name:    "valid",
			payload: &api.OriginUrl{Url: "google.com"},
			err:     false,
		},
		{
			name:    "invalid",
			payload: &api.OriginUrl{Url: "googlecom"},
			err:     true,
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(getConn()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := api.NewShortlinkClient(conn)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			payload := tc.payload

			_, err := client.Create(ctx, payload)
			eq := false

			if err != nil {
				eq = assert.Error(t, err)
			}

			assert.Equal(t, tc.err, eq)
		})
	}
}

func TestShortlinkServer_Get(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(getConn()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := api.NewShortlinkClient(conn)

	short, _ := client.Create(ctx, &api.OriginUrl{Url: "google.com"})
	origin, err := client.Get(ctx, short)
	assert.NoError(t, err)
	assert.Equal(t, "google.com", origin.Url)
}
