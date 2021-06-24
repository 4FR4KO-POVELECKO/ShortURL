package main

import (
	"ShortURL/internal/app/store"
	"ShortURL/pkg/api"
	"ShortURL/pkg/grpcserver"
	"log"
	"net"

	"github.com/go-redis/redis"
	"google.golang.org/grpc"
)

func main() {
	// Connect Redis
	client := redis.NewClient(&redis.Options{Addr: ":6379", Password: "", DB: 0})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	store := store.NewStoreRedis(client)

	// GRPC server
	s := grpc.NewServer()
	srv := &grpcserver.GRPCServer{Store: store}

	api.RegisterShortlinkServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(s.Serve(l))
}
