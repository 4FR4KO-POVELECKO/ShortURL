package main

import (
	"ShortURL/internal/app/shorten"
	"ShortURL/pkg/api"
	"ShortURL/pkg/grpcserver"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &grpcserver.GRPCServer{}
	api.RegisterShortlinkServer(s, srv)

	url := shorten.Shorten()

	fmt.Println(url)

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
