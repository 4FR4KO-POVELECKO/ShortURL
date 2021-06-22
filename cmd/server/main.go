package main

import (
	"ShortURL/internal/app/store"
	"ShortURL/pkg/api"
	"ShortURL/pkg/grpcserver"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db := store.NewStoreRedis("localhost:6379", 0, 10)

	s := grpc.NewServer()
	srv := &grpcserver.GRPCServer{}
	srv.Store = db

	api.RegisterShortlinkServer(s, srv)

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
