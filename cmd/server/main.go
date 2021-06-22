package main

import (
	"ShortURL/pkg/api"
	"ShortURL/pkg/grpcserver"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &grpcserver.GRPCServer{}

	api.RegisterShortlinkServer(s, srv)

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
